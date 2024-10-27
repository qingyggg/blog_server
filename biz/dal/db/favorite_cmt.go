package db

import (
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	"github.com/qingyggg/blog_server/biz/model/query"
	"github.com/qingyggg/blog_server/pkg/utils"
)

func CmtFavoriteAction(UHashId string, CHashId string, AHashId string, actionTp int32) error {
	cf := query.CommentFavorite
	err := cf.Create(&orm_gen.CommentFavorite{
		ArticleID: utils.ConvertStringHashToByte(AHashId),
		UserID:    utils.ConvertStringHashToByte(UHashId),
		CommentID: utils.ConvertStringHashToByte(CHashId),
		Status:    actionTp,
	})
	return err
}

func CmtFavoriteStatusFlush(UHashId string, CHashId string) error {
	cf := query.CommentFavorite
	_, err := cf.Where(cf.UserID.Eq(utils.ConvertStringHashToByte(UHashId)), cf.CommentID.Eq(utils.ConvertStringHashToByte(CHashId))).Delete()
	return err
}

//status:1 like,2dislike
func CmtFavoriteStatus(UHashId string, CHashId string) (error, int32) {
	cf := query.CommentFavorite
	cmt, err := cf.Where(cf.UserID.Eq(utils.ConvertStringHashToByte(UHashId)), cf.CommentID.Eq(utils.ConvertStringHashToByte(CHashId))).Take()
	if err != nil {
		return err, 0
	}
	return nil, cmt.Status //1:like,2:dislike
}

func CmtFavorieExist(UHashId string, CHashId string) (error, bool) {
	cf := query.CommentFavorite
	count, err := cf.Where(cf.UserID.Eq(utils.ConvertStringHashToByte(UHashId)), cf.CommentID.Eq(utils.ConvertStringHashToByte(CHashId))).Count()
	if err != nil {
		return err, false
	}
	if count == 0 {
		return nil, false
	}
	return nil, true
}

func GetCmtFavoriteStatusMap(CHashIds []string, UHashId string) (error, map[string]int32) {
	cf := query.CommentFavorite
	var CByteHashIDs [][]byte
	statusMap := make(map[string]int32)
	//如果该文章没有评论
	if len(CHashIds) == 0 {
		return nil, statusMap
	}
	for _, v := range CHashIds {
		CByteHashIDs = append(CByteHashIDs, utils.ConvertStringHashToByte(v))
	}
	cfs, err := cf.Where(cf.UserID.Eq(utils.ConvertStringHashToByte(UHashId)), cf.CommentID.In(CByteHashIDs...)).Find()
	if err != nil {
		return err, nil
	}

	var curStatus int32 //1:like,2:dislike,3:none status
	for _, v := range CHashIds {
		curStatus = 3
		for _, v2 := range cfs {
			if utils.ConvertByteHashToString(v2.CommentID) == v {
				curStatus = v2.Status //1 or 2
				break
			}
		}
		statusMap[v] = curStatus
	}

	return nil, statusMap
}

// CmtFavoriteFlushByCmtId 删除评论的时候，调用该函数
func CmtFavoriteFlushByCmtId(CHashId string) error {
	cf := query.CommentFavorite
	_, err := cf.Where(cf.CommentID.Eq(utils.ConvertStringHashToByte(CHashId))).Delete()
	return err
}

// CmtFavoriteFlushByArticleId 删除文章的时候，调用该函数
func CmtFavoriteFlushByArticleId(AHashId string) error {
	cf := query.CommentFavorite
	_, err := cf.Where(cf.ArticleID.Eq(utils.ConvertStringHashToByte(AHashId))).Delete()
	return err
}

func GetCmtFavoriteCtMap(cids []string) (error, map[string]int64) {
	// 查询MySQL
	var cfa = query.CommentFavorite
	var resMap = make(map[string]int64)
	var results []struct {
		CommentID []byte
		Count     int64
	}
	var cByteIds [][]byte
	for _, v := range cids {
		cByteIds = append(cByteIds, utils.ConvertStringHashToByte(v))
	}
	// 使用聚合查询来统计每个comment的点赞数
	err := cfa.Select(cfa.CommentID, cfa.UserID.Count().As("Count")).
		Where(cfa.CommentID.In(cByteIds...), cfa.Status.Eq(1)).
		Group(cfa.CommentID).
		Scan(&results)
	if err != nil {
		return err, nil
	}
	var curCt int64
	for _, cid := range cids {
		curCt = 0
		for _, result := range results {
			if utils.ConvertByteHashToString(result.CommentID) == cid {
				curCt = result.Count
				break
			}
		}
		resMap[cid] = curCt
	}
	return nil, resMap
}
