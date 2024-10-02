package db

import (
	"context"
	mymongo "github.com/qingyggg/blog_server/biz/mw/mongo"
	"github.com/qingyggg/blog_server/pkg/errno"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

// AddNewComment add a comment
func AddNewComment(ctx context.Context, comment *mymongo.Comment) error {
	var err error
	//初始化创建时间
	comment.CreateTime = bson.NewDateTimeFromTime(time.Now())
	//顶级评论
	if comment.Degree == 1 {
		comment.ParentID = comment.HashID
		rootCmtL := comment
		rootClosure := &mymongo.CommentClosure{AncestorID: comment.HashID, DescendantID: comment.HashID, Depth: 0}
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
			}
			closures = append(closures, newClosure)
		}
		//增添自己的闭包表
		closures = append(closures, mymongo.CommentClosure{
			AncestorID:   comment.HashID,
			Depth:        0,
			DescendantID: comment.HashID,
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
func DelCommentByHashID(ctx context.Context, cHashId string) error {
	var childIds []string //该评论的子评论数组
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
	//删掉他们
	_, err = mymongo.CommentCol.DeleteMany(ctx, bson.M{"hash_id": bson.M{"$in": childIds}})
	if err != nil {
		return err
	}
	// 删除closure依赖
	_, err = mymongo.CmtClosureCol.DeleteMany(ctx, bson.M{
		"$or": []bson.M{
			{"ancestor": bson.M{"$in": childIds}},
			{"descendant": bson.M{"$in": childIds}},
		},
	})
	return err
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
	return finalList, nil
}

func GetCommentListByTopCommentID(ctx context.Context, aHashId string, cHashId string) (cmtList []*mymongo.CommentItemForSub, err error) {
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
		cmtList = append(cmtList, &mymongo.CommentItemForSub{
			ArticleID:  v.ArticleID,
			Content:    v.Content,
			Degree:     v.Degree,
			UserID:     v.UserID,
			HashID:     v.HashID,
			ParentUID:  uMaps[v.HashID],
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
