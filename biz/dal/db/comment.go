package db

import (
	"context"
	mymongo "github.com/qingyggg/blog_server/biz/mw/mongo"
	"github.com/qingyggg/blog_server/biz/mw/redis"
	"github.com/qingyggg/blog_server/pkg/errno"
	"go.mongodb.org/mongo-driver/v2/bson"
	"sync"
	"time"
)

var rdComment = redis.Comment{}

// AddNewComment add a comment
func AddNewComment(ctx context.Context, comment *mymongo.Comment) error {
	var err error
	//初始化创建时间
	comment.CreateTime = bson.NewDateTimeFromTime(time.Now())
	//顶级评论
	if comment.Degree == 1 {
		comment.ParentID = comment.HashID
		rootCmtL := comment
		rootClosure := &mymongo.CommentClosure{AncestorID: comment.HashID, DescendantID: comment.HashID, Depth: 0, ArticleID: comment.ArticleID}
		_, err = mymongo.CommentCol.InsertOne(ctx, rootCmtL)
		_, err = mymongo.CmtClosureCol.InsertOne(ctx, rootClosure)
		if err != nil {
			return err
		}
	} else if comment.Degree == 2 { //子集评论
		//这里判断parent id 是否存在
		count, err := mymongo.CommentCol.CountDocuments(ctx, bson.M{"hash_id": comment.ParentID})
		if err != nil {
			return err
		}
		if count == 0 {
			return errno.ServiceErr.WithMessage("该评论的parentID不存在")
		}
		_, err = mymongo.CommentCol.InsertOne(ctx, comment) //插入评论
		if err != nil {
			return err
		}
		cursor, err := mymongo.CmtClosureCol.Find(ctx, bson.M{"descendant": comment.ParentID})
		if err != nil {
			return err
		}
		defer cursor.Close(ctx)

		// 遍历祖先评论并生成新的闭包表记录
		var closures []interface{}
		for cursor.Next(ctx) {
			closure := new(mymongo.CommentClosure)
			if err := cursor.Decode(&closure); err != nil {
				return err
			}
			newClosure := mymongo.CommentClosure{
				DescendantID: comment.HashID,
				AncestorID:   closure.AncestorID,
				Depth:        closure.Depth + 1,
				ArticleID:    comment.ArticleID,
			}
			closures = append(closures, newClosure)
		}
		//增添自己的闭包表
		closures = append(closures, mymongo.CommentClosure{
			AncestorID:   comment.HashID,
			Depth:        0,
			DescendantID: comment.HashID,
			ArticleID:    comment.ArticleID,
		})
		// 插入所有闭包表记录
		if len(closures) > 0 {
			_, err = mymongo.CmtClosureCol.InsertMany(ctx, closures)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// DelCommentByHashID DeleteCommentById delete comment by comment id
func DelCommentByHashID(ctx context.Context, cHashId string, aHashId string) error {
	var childIds []string //该评论的子评论数组
	var wg sync.WaitGroup
	cfrdb := redis.Comment{}
	errChan := make(chan error, 4) // 处理错误的channel
	defer close(errChan)
	//查找该评论以及该评论的子评论的id
	cursor, err := mymongo.CmtClosureCol.Find(ctx, bson.M{"ancestor": cHashId})
	if err != nil {
		return err
	}
	curClo := new(mymongo.CommentClosure)
	for cursor.Next(ctx) {
		err = cursor.Decode(curClo)
		if err != nil {
			return err
		}
		childIds = append(childIds, curClo.DescendantID)
	}
	//获取评论数量在删除之前
	cmtCount, err := mymongo.CommentCol.CountDocuments(ctx, bson.M{"article_id": aHashId})
	if err != nil {
		return err
	}
	wg.Add(4)
	//删掉评论以及子评论
	go func() {
		defer wg.Done()
		_, err := mymongo.CommentCol.DeleteMany(ctx, bson.M{"hash_id": bson.M{"$in": childIds}})
		if err != nil {
			errChan <- err
			return
		}
	}()
	// 删除closure依赖
	go func() {
		defer wg.Done()
		_, err := mymongo.CmtClosureCol.DeleteMany(ctx, bson.M{
			"$or": []bson.M{
				{"ancestor": bson.M{"$in": childIds}},
				{"descendant": bson.M{"$in": childIds}},
			},
		})
		if err != nil {
			errChan <- err
			return
		}
	}()
	//删除掉该评论的favorite以及该评论的子评论的favorite
	go func() {
		defer wg.Done()
		err := CmtFavoriteFlushByCmtId(cHashId)
		if err != nil {
			errChan <- err
			return
		}
	}()
	//redis comment count减小
	go func() {
		defer wg.Done()
		err := cfrdb.CommentCtAssign(aHashId, cmtCount-int64(len(childIds)))
		if err != nil {
			errChan <- err
			return
		}
	}()
	wg.Wait()

	select {
	case err := <-errChan:
		if err != nil {
			return err
		}
	default:
	}
	return nil
}

// GetCommentListByArticleID 通过文章id获取所有的一级评论
func GetCommentListByArticleID(ctx context.Context, aHashId string) ([]*mymongo.CommentItem, error) {
	// 查询所有顶级评论
	var err error
	var firstList []mymongo.Comment
	//初始化cmtList
	var finalList []*mymongo.CommentItem
	cursor, err := mymongo.CommentCol.Find(ctx, bson.M{
		"article_id": aHashId,
		"degree":     1, // 查找顶级评论
	})
	defer cursor.Close(ctx)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &firstList); err != nil {
		return nil, err
	}
	if len(firstList) != 0 {
		//获取所有的hashid
		var hashIDs []string
		for _, v := range firstList {
			hashIDs = append(hashIDs, v.HashID)
		}
		pipeline := []bson.M{
			{
				"$match": bson.M{"ancestor": bson.M{"$in": hashIDs}}, // 匹配祖先ID在给定列表中的记录
			},
			{
				"$group": bson.M{
					"_id":   "$ancestor",       // 按祖先ID分组
					"count": bson.M{"$sum": 1}, // 统计每组的数量
				},
			},
		}

		// 执行聚合查询
		cursor, err = mymongo.CmtClosureCol.Aggregate(ctx, pipeline)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)

		// map[评论id]回复总数
		sumMaps := make(map[string]int64)

		// 解析聚合结果
		for cursor.Next(ctx) {
			var item struct {
				ID    string `bson:"_id"`
				Count int64  `bson:"count"`
			}
			if err := cursor.Decode(&item); err != nil {
				return nil, err
			}
			sumMaps[item.ID] = item.Count // 将结果存入映射
		}
		//处理结果
		for _, v := range firstList {
			finalList = append(finalList, &mymongo.CommentItem{
				ArticleID:  v.ArticleID,
				Content:    v.Content,
				Degree:     v.Degree,
				UserID:     v.UserID,
				HashID:     v.HashID,
				CreateTime: v.CreateTime,
				ChildNum:   sumMaps[v.HashID] - 1, //减掉1的原因：closure table(has)-->ancestor:hashid ,descendant:hashid
			})
		}
	}
	return finalList, nil
}

func GetCommentListByTopCommentID(ctx context.Context, aHashId string, cHashId string) (cmtList []*mymongo.CommentItem, err error) {
	//通过closure获取所有要请求的评论ID
	cursor, err := mymongo.CmtClosureCol.Find(ctx, &bson.M{"ancestor": cHashId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	//将结果推入cids数组
	var cids []string
	for cursor.Next(ctx) {
		item := new(mymongo.CommentClosure)
		if err := cursor.Decode(item); err != nil {
			return nil, err
		}
		if item.DescendantID != cHashId {
			cids = append(cids, item.DescendantID)
		}

	}
	//根据cids去请求评论列表
	var firstList []*mymongo.Comment
	cursor, err = mymongo.CommentCol.Find(ctx, &bson.M{"hash_id": bson.M{"$in": cids}, "article_id": aHashId})
	if err != nil {
		return nil, err
	}
	//map[comment id]user id
	uMaps := make(map[string]string)
	for cursor.Next(ctx) {
		item := new(mymongo.Comment)
		if err := cursor.Decode(item); err != nil {
			return nil, err
		}
		uMaps[item.HashID] = item.UserID
		firstList = append(firstList, item)
	}
	for _, v := range firstList {
		cmtList = append(cmtList, &mymongo.CommentItem{
			ArticleID:  v.ArticleID,
			Content:    v.Content,
			Degree:     v.Degree,
			UserID:     v.UserID,
			HashID:     v.HashID,
			ParentUID:  uMaps[v.ParentID],
			CreateTime: v.CreateTime,
		})
	}
	return cmtList, nil
}

// GetCommentByTopCmtID 根据comment hashid 获取评论
func GetCommentByCmtID(ctx context.Context, cHashId string) (cmt *mymongo.Comment, err error) {
	cmt = new(mymongo.Comment)
	//挑选某一个顶级评论
	err = mymongo.CommentCol.FindOne(ctx, bson.M{
		"hash_id": cHashId,
	}).Decode(cmt)
	if err != nil {
		return nil, err
	}
	return cmt, nil
}

func CheckCmtExistById(ctx context.Context, cHashId string) (exist bool, err error) {
	count, err := mymongo.CommentCol.CountDocuments(ctx, bson.M{
		"hash_id": cHashId,
	})
	if err != nil {
		return false, err
	}
	if count != 0 {
		return true, nil
	} else {
		return false, nil
	}
}

// GetCmtCtByAids 通过文章id获取该文章的评论数
func GetCmtCtByAids(ctx context.Context, aHashIds []string) (error, map[string]int64) {
	// 将所有的aid从Redis里获取，获取不到的aid收集起来，一起从数据库中获取，获取后将这些键值对存储进Redis
	var aidInCache []string // 存在于缓存里的aHashId:count
	var aidInNoCache []string
	var resMap = make(map[string]int64) // 初始化结果map
	var wg sync.WaitGroup
	var mu sync.Mutex              // 用于保护resMap的并发写入
	errChan := make(chan error, 2) // 处理错误的channel
	defer close(errChan)
	// 遍历输入的aHashIds，检查每个ID是否存在于缓存中
	for _, v := range aHashIds {
		err, exist := rdComment.CheckCommentCt(v)
		if err != nil {
			return err, nil
		}
		if exist {
			aidInCache = append(aidInCache, v)
		} else {
			aidInNoCache = append(aidInNoCache, v)
		}
	}

	wg.Add(2) // 启动两个goroutine

	// 从缓存中获取点赞数的goroutine
	go func() {
		defer wg.Done()
		for _, v := range aidInCache {
			err, count := rdComment.CountComment(v)
			if err != nil {
				errChan <- err
				return
			}
			mu.Lock()
			resMap[v] = count
			mu.Unlock()
		}
	}()
	// 查询MongoDB数据库并将结果同步到Redis的goroutine
	go func() {
		defer wg.Done()
		if len(aidInNoCache) > 0 {
			pipeline := []bson.M{
				{
					"$match": bson.M{"article_id": bson.M{"$in": aidInNoCache}}, // 匹配祖先ID在给定列表中的记录
				},
				{
					"$group": bson.M{
						"_id":   "$article_id",     // 按祖先ID分组
						"count": bson.M{"$sum": 1}, // 统计每组的数量
					},
				},
				{
					"$addFields": bson.M{
						//添加两个新字段,它的值分别为分组字段的值
						"article_id": "$article_id",
					},
				},
			}

			// 执行聚合查询
			cursor, err := mymongo.CommentCol.Aggregate(ctx, pipeline)
			if err != nil {
				errChan <- err
				return
			}
			defer cursor.Close(ctx)

			// 解析聚合结果,缓存进redis
			fpipe := rdComment.GetCommentClient().TxPipeline()
			for cursor.Next(ctx) {
				// 查询Mongo
				var item struct {
					_id       string
					Count     int64
					ArticleId string
				}
				if err := cursor.Decode(&item); err != nil {
					errChan <- err
					return
				}
				mu.Lock()
				resMap[item._id] = item.Count
				mu.Unlock()
				fpipe.Set(item._id+redis.CommentCountSuffix, item.Count, redis.ExpireTime)
			}
			// 执行事务
			_, err = fpipe.Exec()
			if err != nil {
				errChan <- err
				return
			}
		}
	}()
	wg.Wait()
	// 检查是否有错误发生
	select {
	case err := <-errChan:
		if err != nil {
			return err, nil
		}
	default:
	}
	//这里如果缓存和MongoDB都没有查找到comment count，
	//说明这篇文章并没有被评论过，所以默认赋值为0
	for _, v := range aHashIds {
		if _, exist := resMap[v]; exist == false {
			resMap[v] = 0
		}
	}
	return nil, resMap
}

// DelCommentByAid 删除文章前，调用该函数
func DelCommentByAid(ctx context.Context, aHashId string) error {
	var wg sync.WaitGroup
	cmtRdb := redis.Comment{}
	errChan := make(chan error, 4)
	wg.Add(4)
	defer close(errChan)
	//根据文章id删除评论
	go func() {
		defer wg.Done()
		_, err := mymongo.CommentCol.DeleteMany(ctx, bson.M{"article_id": aHashId})
		if err != nil {
			errChan <- err
			return
		}
	}()
	//根据文章id删除评论closure
	go func() {
		defer wg.Done()
		_, err := mymongo.CmtClosureCol.DeleteMany(ctx, bson.M{"article_id": aHashId})
		if err != nil {
			errChan <- err
			return
		}
	}()
	//清除redis缓存
	go func() {
		defer wg.Done()
		err, exist := cmtRdb.CheckCommentCt(aHashId)
		if err != nil {
			errChan <- err
			return
		}
		if exist {
			err := cmtRdb.DelCommentCt(aHashId)
			if err != nil {
				errChan <- err
				return
			}
		}
	}()
	go func() {
		defer wg.Done()
		//清除掉cmt favorite
		err := CmtFavoriteFlushByArticleId(aHashId)
		if err != nil {
			errChan <- err
			return
		}
	}()
	wg.Wait()

	select {
	case err := <-errChan:
		if err != nil {
			return err
		}
	default:
	}
	return nil
}

func GetCmtCountByAid(ctx context.Context, aHashId string) (error, int64) {
	err, exist := rdComment.CheckCommentCt(aHashId)
	if err != nil {
		return err, 0
	}
	if exist {
		return rdComment.CountComment(aHashId)
	} else {
		count, err := mymongo.CommentCol.CountDocuments(ctx, &bson.M{"article_id": aHashId})
		if err != nil {
			return err, 0
		}
		return nil, count
	}
}
