package db

import (
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	"github.com/qingyggg/blog_server/biz/model/query"
	"github.com/qingyggg/blog_server/pkg/errno"
)

var c = query.Comment

// AddNewComment add a comment
func AddNewComment(comment *orm_gen.Comment) error {
	if ok, _ := CheckUserExistById(comment.UserID); !ok {
		return errno.UserIsNotExistErr
	}
	if ok, _ := CheckArticleExistById(comment.ArticleID); !ok {
		return errno.ArticleIsNotExistErr
	}
	err := c.Create(comment)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCommentById delete comment by comment id
func DeleteCommentById(comment_id int64) error {
	if ok, _ := CheckCommentExist(comment_id); !ok {
		return errno.CommentIsNotExistErr
	}
	_, err := c.Where(c.ID.Eq(comment_id)).Delete()
	if err != nil {
		return err
	}
	return nil
}

func CheckCommentExist(comment_id int64) (bool, error) {
	count, err := c.Where(c.ID.Eq(comment_id)).Count()
	if err != nil {
		return false, err
	} else if count != 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func GetCommentListByArticleID(aid int64) (comments []*orm_gen.Comment, err error) {
	if ok, _ := CheckArticleExistById(aid); !ok {
		return comments, errno.ArticleIsNotExistErr
	}
	comments, err = c.Where(c.ArticleID.Eq(aid), c.Floor.Eq(1)).Find()
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func CheckCommentExistById(cid int64) (bool, error) {
	count, err := a.Where(c.ID.Eq(cid)).Count()
	if err != nil {
		return false, err
	} else if count != 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func GetCommentListByTopCommentID(aid int64, cid int64) (comments []*orm_gen.Comment, err error) {
	if ok, _ := CheckCommentExistById(cid); !ok {
		return comments, errno.CommentIsNotExistErr
	}
	comments, err = c.Where(c.ArticleID.Eq(aid), c.RepliedCommentID.Eq(cid)).Find()
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func GetCommentCountByArticleID(aid int64) (int64, error) {
	count, err := c.Where(c.ArticleID.Eq(aid)).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetCommentCountByTopCommentID(aid int64, comment_id int64) (int64, error) {
	count, err := c.Where(c.RepliedCommentID.Eq(comment_id), c.ArticleID.Eq(aid)).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}
