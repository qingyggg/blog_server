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

	"github.com/qingyggg/blog_server/biz/model/model"
)

func newArticle(db *gorm.DB, opts ...gen.DOOption) article {
	_article := article{}

	_article.articleDo.UseDB(db, opts...)
	_article.articleDo.UseModel(&model.Article{})

	tableName := _article.articleDo.TableName()
	_article.ALL = field.NewAsterisk(tableName)
	_article.ID = field.NewInt64(tableName, "id")
	_article.UserID = field.NewInt64(tableName, "user_id")
	_article.LikeCount = field.NewInt64(tableName, "like_count")
	_article.CommentCount = field.NewInt64(tableName, "comment_count")
	_article.CollectCount = field.NewInt64(tableName, "collect_count")
	_article.Title = field.NewString(tableName, "title")
	_article.Note = field.NewString(tableName, "note")
	_article.Content = field.NewString(tableName, "content")
	_article.CoverURL = field.NewString(tableName, "cover_url")

	_article.fillFieldMap()

	return _article
}

// article 文章表
type article struct {
	articleDo

	ALL          field.Asterisk
	ID           field.Int64  // 文章ID
	UserID       field.Int64  // 作者id
	LikeCount    field.Int64  // 点赞用户数目
	CommentCount field.Int64  // 评论数目
	CollectCount field.Int64  // 收藏数目
	Title        field.String // 文章标题
	Note         field.String // 文章小记
	Content      field.String // 文章内容
	CoverURL     field.String // 背景图url

	fieldMap map[string]field.Expr
}

func (a article) Table(newTableName string) *article {
	a.articleDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a article) As(alias string) *article {
	a.articleDo.DO = *(a.articleDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *article) updateTableName(table string) *article {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.UserID = field.NewInt64(table, "user_id")
	a.LikeCount = field.NewInt64(table, "like_count")
	a.CommentCount = field.NewInt64(table, "comment_count")
	a.CollectCount = field.NewInt64(table, "collect_count")
	a.Title = field.NewString(table, "title")
	a.Note = field.NewString(table, "note")
	a.Content = field.NewString(table, "content")
	a.CoverURL = field.NewString(table, "cover_url")

	a.fillFieldMap()

	return a
}

func (a *article) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *article) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["id"] = a.ID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["like_count"] = a.LikeCount
	a.fieldMap["comment_count"] = a.CommentCount
	a.fieldMap["collect_count"] = a.CollectCount
	a.fieldMap["title"] = a.Title
	a.fieldMap["note"] = a.Note
	a.fieldMap["content"] = a.Content
	a.fieldMap["cover_url"] = a.CoverURL
}

func (a article) clone(db *gorm.DB) article {
	a.articleDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a article) replaceDB(db *gorm.DB) article {
	a.articleDo.ReplaceDB(db)
	return a
}

type articleDo struct{ gen.DO }

type IArticleDo interface {
	gen.SubQuery
	Debug() IArticleDo
	WithContext(ctx context.Context) IArticleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IArticleDo
	WriteDB() IArticleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IArticleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IArticleDo
	Not(conds ...gen.Condition) IArticleDo
	Or(conds ...gen.Condition) IArticleDo
	Select(conds ...field.Expr) IArticleDo
	Where(conds ...gen.Condition) IArticleDo
	Order(conds ...field.Expr) IArticleDo
	Distinct(cols ...field.Expr) IArticleDo
	Omit(cols ...field.Expr) IArticleDo
	Join(table schema.Tabler, on ...field.Expr) IArticleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IArticleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IArticleDo
	Group(cols ...field.Expr) IArticleDo
	Having(conds ...gen.Condition) IArticleDo
	Limit(limit int) IArticleDo
	Offset(offset int) IArticleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleDo
	Unscoped() IArticleDo
	Create(values ...*model.Article) error
	CreateInBatches(values []*model.Article, batchSize int) error
	Save(values ...*model.Article) error
	First() (*model.Article, error)
	Take() (*model.Article, error)
	Last() (*model.Article, error)
	Find() ([]*model.Article, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Article, err error)
	FindInBatches(result *[]*model.Article, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Article) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IArticleDo
	Assign(attrs ...field.AssignExpr) IArticleDo
	Joins(fields ...field.RelationField) IArticleDo
	Preload(fields ...field.RelationField) IArticleDo
	FirstOrInit() (*model.Article, error)
	FirstOrCreate() (*model.Article, error)
	FindByPage(offset int, limit int) (result []*model.Article, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IArticleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a articleDo) Debug() IArticleDo {
	return a.withDO(a.DO.Debug())
}

func (a articleDo) WithContext(ctx context.Context) IArticleDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articleDo) ReadDB() IArticleDo {
	return a.Clauses(dbresolver.Read)
}

func (a articleDo) WriteDB() IArticleDo {
	return a.Clauses(dbresolver.Write)
}

func (a articleDo) Session(config *gorm.Session) IArticleDo {
	return a.withDO(a.DO.Session(config))
}

func (a articleDo) Clauses(conds ...clause.Expression) IArticleDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articleDo) Returning(value interface{}, columns ...string) IArticleDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articleDo) Not(conds ...gen.Condition) IArticleDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articleDo) Or(conds ...gen.Condition) IArticleDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articleDo) Select(conds ...field.Expr) IArticleDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articleDo) Where(conds ...gen.Condition) IArticleDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articleDo) Order(conds ...field.Expr) IArticleDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articleDo) Distinct(cols ...field.Expr) IArticleDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articleDo) Omit(cols ...field.Expr) IArticleDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articleDo) Join(table schema.Tabler, on ...field.Expr) IArticleDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IArticleDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articleDo) RightJoin(table schema.Tabler, on ...field.Expr) IArticleDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articleDo) Group(cols ...field.Expr) IArticleDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articleDo) Having(conds ...gen.Condition) IArticleDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articleDo) Limit(limit int) IArticleDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articleDo) Offset(offset int) IArticleDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articleDo) Unscoped() IArticleDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articleDo) Create(values ...*model.Article) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articleDo) CreateInBatches(values []*model.Article, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articleDo) Save(values ...*model.Article) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articleDo) First() (*model.Article, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Article), nil
	}
}

func (a articleDo) Take() (*model.Article, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Article), nil
	}
}

func (a articleDo) Last() (*model.Article, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Article), nil
	}
}

func (a articleDo) Find() ([]*model.Article, error) {
	result, err := a.DO.Find()
	return result.([]*model.Article), err
}

func (a articleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Article, err error) {
	buf := make([]*model.Article, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articleDo) FindInBatches(result *[]*model.Article, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articleDo) Attrs(attrs ...field.AssignExpr) IArticleDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articleDo) Assign(attrs ...field.AssignExpr) IArticleDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articleDo) Joins(fields ...field.RelationField) IArticleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articleDo) Preload(fields ...field.RelationField) IArticleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articleDo) FirstOrInit() (*model.Article, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Article), nil
	}
}

func (a articleDo) FirstOrCreate() (*model.Article, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Article), nil
	}
}

func (a articleDo) FindByPage(offset int, limit int) (result []*model.Article, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a articleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articleDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articleDo) Delete(models ...*model.Article) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articleDo) withDO(do gen.Dao) *articleDo {
	a.DO = *do.(*gen.DO)
	return a
}
