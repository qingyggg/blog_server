// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/qingyggg/blog_server/biz/model/orm_gen"
)

func newCommentFavorite(db *gorm.DB, opts ...gen.DOOption) commentFavorite {
	_commentFavorite := commentFavorite{}

	_commentFavorite.commentFavoriteDo.UseDB(db, opts...)
	_commentFavorite.commentFavoriteDo.UseModel(&orm_gen.CommentFavorite{})

	tableName := _commentFavorite.commentFavoriteDo.TableName()
	_commentFavorite.ALL = field.NewAsterisk(tableName)
	_commentFavorite.ID = field.NewInt64(tableName, "id")
	_commentFavorite.ArticleID = field.NewBytes(tableName, "article_id")
	_commentFavorite.CommentID = field.NewBytes(tableName, "comment_id")
	_commentFavorite.UserID = field.NewBytes(tableName, "user_id")
	_commentFavorite.Status = field.NewInt32(tableName, "status")

	_commentFavorite.fillFieldMap()

	return _commentFavorite
}

// commentFavorite 用户对评论的点赞或者踩
type commentFavorite struct {
	commentFavoriteDo

	ALL       field.Asterisk
	ID        field.Int64 // 主键
	ArticleID field.Bytes // 评论文章ID
	CommentID field.Bytes // 被点赞或踩的评论 ID
	UserID    field.Bytes // 用户ID
	Status    field.Int32 // 2：踩, 1：点赞

	fieldMap map[string]field.Expr
}

func (c commentFavorite) Table(newTableName string) *commentFavorite {
	c.commentFavoriteDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c commentFavorite) As(alias string) *commentFavorite {
	c.commentFavoriteDo.DO = *(c.commentFavoriteDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *commentFavorite) updateTableName(table string) *commentFavorite {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.ArticleID = field.NewBytes(table, "article_id")
	c.CommentID = field.NewBytes(table, "comment_id")
	c.UserID = field.NewBytes(table, "user_id")
	c.Status = field.NewInt32(table, "status")

	c.fillFieldMap()

	return c
}

func (c *commentFavorite) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *commentFavorite) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 5)
	c.fieldMap["id"] = c.ID
	c.fieldMap["article_id"] = c.ArticleID
	c.fieldMap["comment_id"] = c.CommentID
	c.fieldMap["user_id"] = c.UserID
	c.fieldMap["status"] = c.Status
}

func (c commentFavorite) clone(db *gorm.DB) commentFavorite {
	c.commentFavoriteDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c commentFavorite) replaceDB(db *gorm.DB) commentFavorite {
	c.commentFavoriteDo.ReplaceDB(db)
	return c
}

type commentFavoriteDo struct{ gen.DO }

type ICommentFavoriteDo interface {
	gen.SubQuery
	Debug() ICommentFavoriteDo
	WithContext(ctx context.Context) ICommentFavoriteDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICommentFavoriteDo
	WriteDB() ICommentFavoriteDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICommentFavoriteDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICommentFavoriteDo
	Not(conds ...gen.Condition) ICommentFavoriteDo
	Or(conds ...gen.Condition) ICommentFavoriteDo
	Select(conds ...field.Expr) ICommentFavoriteDo
	Where(conds ...gen.Condition) ICommentFavoriteDo
	Order(conds ...field.Expr) ICommentFavoriteDo
	Distinct(cols ...field.Expr) ICommentFavoriteDo
	Omit(cols ...field.Expr) ICommentFavoriteDo
	Join(table schema.Tabler, on ...field.Expr) ICommentFavoriteDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICommentFavoriteDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICommentFavoriteDo
	Group(cols ...field.Expr) ICommentFavoriteDo
	Having(conds ...gen.Condition) ICommentFavoriteDo
	Limit(limit int) ICommentFavoriteDo
	Offset(offset int) ICommentFavoriteDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICommentFavoriteDo
	Unscoped() ICommentFavoriteDo
	Create(values ...*orm_gen.CommentFavorite) error
	CreateInBatches(values []*orm_gen.CommentFavorite, batchSize int) error
	Save(values ...*orm_gen.CommentFavorite) error
	First() (*orm_gen.CommentFavorite, error)
	Take() (*orm_gen.CommentFavorite, error)
	Last() (*orm_gen.CommentFavorite, error)
	Find() ([]*orm_gen.CommentFavorite, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.CommentFavorite, err error)
	FindInBatches(result *[]*orm_gen.CommentFavorite, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*orm_gen.CommentFavorite) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICommentFavoriteDo
	Assign(attrs ...field.AssignExpr) ICommentFavoriteDo
	Joins(fields ...field.RelationField) ICommentFavoriteDo
	Preload(fields ...field.RelationField) ICommentFavoriteDo
	FirstOrInit() (*orm_gen.CommentFavorite, error)
	FirstOrCreate() (*orm_gen.CommentFavorite, error)
	FindByPage(offset int, limit int) (result []*orm_gen.CommentFavorite, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICommentFavoriteDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c commentFavoriteDo) Debug() ICommentFavoriteDo {
	return c.withDO(c.DO.Debug())
}

func (c commentFavoriteDo) WithContext(ctx context.Context) ICommentFavoriteDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c commentFavoriteDo) ReadDB() ICommentFavoriteDo {
	return c.Clauses(dbresolver.Read)
}

func (c commentFavoriteDo) WriteDB() ICommentFavoriteDo {
	return c.Clauses(dbresolver.Write)
}

func (c commentFavoriteDo) Session(config *gorm.Session) ICommentFavoriteDo {
	return c.withDO(c.DO.Session(config))
}

func (c commentFavoriteDo) Clauses(conds ...clause.Expression) ICommentFavoriteDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c commentFavoriteDo) Returning(value interface{}, columns ...string) ICommentFavoriteDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c commentFavoriteDo) Not(conds ...gen.Condition) ICommentFavoriteDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c commentFavoriteDo) Or(conds ...gen.Condition) ICommentFavoriteDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c commentFavoriteDo) Select(conds ...field.Expr) ICommentFavoriteDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c commentFavoriteDo) Where(conds ...gen.Condition) ICommentFavoriteDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c commentFavoriteDo) Order(conds ...field.Expr) ICommentFavoriteDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c commentFavoriteDo) Distinct(cols ...field.Expr) ICommentFavoriteDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c commentFavoriteDo) Omit(cols ...field.Expr) ICommentFavoriteDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c commentFavoriteDo) Join(table schema.Tabler, on ...field.Expr) ICommentFavoriteDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c commentFavoriteDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICommentFavoriteDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c commentFavoriteDo) RightJoin(table schema.Tabler, on ...field.Expr) ICommentFavoriteDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c commentFavoriteDo) Group(cols ...field.Expr) ICommentFavoriteDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c commentFavoriteDo) Having(conds ...gen.Condition) ICommentFavoriteDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c commentFavoriteDo) Limit(limit int) ICommentFavoriteDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c commentFavoriteDo) Offset(offset int) ICommentFavoriteDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c commentFavoriteDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICommentFavoriteDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c commentFavoriteDo) Unscoped() ICommentFavoriteDo {
	return c.withDO(c.DO.Unscoped())
}

func (c commentFavoriteDo) Create(values ...*orm_gen.CommentFavorite) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c commentFavoriteDo) CreateInBatches(values []*orm_gen.CommentFavorite, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c commentFavoriteDo) Save(values ...*orm_gen.CommentFavorite) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c commentFavoriteDo) First() (*orm_gen.CommentFavorite, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.CommentFavorite), nil
	}
}

func (c commentFavoriteDo) Take() (*orm_gen.CommentFavorite, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.CommentFavorite), nil
	}
}

func (c commentFavoriteDo) Last() (*orm_gen.CommentFavorite, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.CommentFavorite), nil
	}
}

func (c commentFavoriteDo) Find() ([]*orm_gen.CommentFavorite, error) {
	result, err := c.DO.Find()
	return result.([]*orm_gen.CommentFavorite), err
}

func (c commentFavoriteDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.CommentFavorite, err error) {
	buf := make([]*orm_gen.CommentFavorite, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c commentFavoriteDo) FindInBatches(result *[]*orm_gen.CommentFavorite, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c commentFavoriteDo) Attrs(attrs ...field.AssignExpr) ICommentFavoriteDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c commentFavoriteDo) Assign(attrs ...field.AssignExpr) ICommentFavoriteDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c commentFavoriteDo) Joins(fields ...field.RelationField) ICommentFavoriteDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c commentFavoriteDo) Preload(fields ...field.RelationField) ICommentFavoriteDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c commentFavoriteDo) FirstOrInit() (*orm_gen.CommentFavorite, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.CommentFavorite), nil
	}
}

func (c commentFavoriteDo) FirstOrCreate() (*orm_gen.CommentFavorite, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.CommentFavorite), nil
	}
}

func (c commentFavoriteDo) FindByPage(offset int, limit int) (result []*orm_gen.CommentFavorite, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c commentFavoriteDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c commentFavoriteDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c commentFavoriteDo) Delete(models ...*orm_gen.CommentFavorite) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *commentFavoriteDo) withDO(do gen.Dao) *commentFavoriteDo {
	c.DO = *do.(*gen.DO)
	return c
}
