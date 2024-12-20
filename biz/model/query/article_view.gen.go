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

func newArticleView(db *gorm.DB, opts ...gen.DOOption) articleView {
	_articleView := articleView{}

	_articleView.articleViewDo.UseDB(db, opts...)
	_articleView.articleViewDo.UseModel(&orm_gen.ArticleView{})

	tableName := _articleView.articleViewDo.TableName()
	_articleView.ALL = field.NewAsterisk(tableName)
	_articleView.ID = field.NewInt64(tableName, "id")
	_articleView.ArticleID = field.NewBytes(tableName, "article_id")
	_articleView.ViewCount = field.NewInt64(tableName, "view_count")

	_articleView.fillFieldMap()

	return _articleView
}

// articleView 统计文章阅读数的表
type articleView struct {
	articleViewDo

	ALL       field.Asterisk
	ID        field.Int64 // 主键
	ArticleID field.Bytes // 文章ID
	ViewCount field.Int64 // 文章阅读数

	fieldMap map[string]field.Expr
}

func (a articleView) Table(newTableName string) *articleView {
	a.articleViewDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a articleView) As(alias string) *articleView {
	a.articleViewDo.DO = *(a.articleViewDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *articleView) updateTableName(table string) *articleView {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.ArticleID = field.NewBytes(table, "article_id")
	a.ViewCount = field.NewInt64(table, "view_count")

	a.fillFieldMap()

	return a
}

func (a *articleView) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *articleView) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 3)
	a.fieldMap["id"] = a.ID
	a.fieldMap["article_id"] = a.ArticleID
	a.fieldMap["view_count"] = a.ViewCount
}

func (a articleView) clone(db *gorm.DB) articleView {
	a.articleViewDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a articleView) replaceDB(db *gorm.DB) articleView {
	a.articleViewDo.ReplaceDB(db)
	return a
}

type articleViewDo struct{ gen.DO }

type IArticleViewDo interface {
	gen.SubQuery
	Debug() IArticleViewDo
	WithContext(ctx context.Context) IArticleViewDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IArticleViewDo
	WriteDB() IArticleViewDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IArticleViewDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IArticleViewDo
	Not(conds ...gen.Condition) IArticleViewDo
	Or(conds ...gen.Condition) IArticleViewDo
	Select(conds ...field.Expr) IArticleViewDo
	Where(conds ...gen.Condition) IArticleViewDo
	Order(conds ...field.Expr) IArticleViewDo
	Distinct(cols ...field.Expr) IArticleViewDo
	Omit(cols ...field.Expr) IArticleViewDo
	Join(table schema.Tabler, on ...field.Expr) IArticleViewDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IArticleViewDo
	RightJoin(table schema.Tabler, on ...field.Expr) IArticleViewDo
	Group(cols ...field.Expr) IArticleViewDo
	Having(conds ...gen.Condition) IArticleViewDo
	Limit(limit int) IArticleViewDo
	Offset(offset int) IArticleViewDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleViewDo
	Unscoped() IArticleViewDo
	Create(values ...*orm_gen.ArticleView) error
	CreateInBatches(values []*orm_gen.ArticleView, batchSize int) error
	Save(values ...*orm_gen.ArticleView) error
	First() (*orm_gen.ArticleView, error)
	Take() (*orm_gen.ArticleView, error)
	Last() (*orm_gen.ArticleView, error)
	Find() ([]*orm_gen.ArticleView, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.ArticleView, err error)
	FindInBatches(result *[]*orm_gen.ArticleView, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*orm_gen.ArticleView) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IArticleViewDo
	Assign(attrs ...field.AssignExpr) IArticleViewDo
	Joins(fields ...field.RelationField) IArticleViewDo
	Preload(fields ...field.RelationField) IArticleViewDo
	FirstOrInit() (*orm_gen.ArticleView, error)
	FirstOrCreate() (*orm_gen.ArticleView, error)
	FindByPage(offset int, limit int) (result []*orm_gen.ArticleView, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IArticleViewDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a articleViewDo) Debug() IArticleViewDo {
	return a.withDO(a.DO.Debug())
}

func (a articleViewDo) WithContext(ctx context.Context) IArticleViewDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articleViewDo) ReadDB() IArticleViewDo {
	return a.Clauses(dbresolver.Read)
}

func (a articleViewDo) WriteDB() IArticleViewDo {
	return a.Clauses(dbresolver.Write)
}

func (a articleViewDo) Session(config *gorm.Session) IArticleViewDo {
	return a.withDO(a.DO.Session(config))
}

func (a articleViewDo) Clauses(conds ...clause.Expression) IArticleViewDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articleViewDo) Returning(value interface{}, columns ...string) IArticleViewDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articleViewDo) Not(conds ...gen.Condition) IArticleViewDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articleViewDo) Or(conds ...gen.Condition) IArticleViewDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articleViewDo) Select(conds ...field.Expr) IArticleViewDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articleViewDo) Where(conds ...gen.Condition) IArticleViewDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articleViewDo) Order(conds ...field.Expr) IArticleViewDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articleViewDo) Distinct(cols ...field.Expr) IArticleViewDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articleViewDo) Omit(cols ...field.Expr) IArticleViewDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articleViewDo) Join(table schema.Tabler, on ...field.Expr) IArticleViewDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articleViewDo) LeftJoin(table schema.Tabler, on ...field.Expr) IArticleViewDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articleViewDo) RightJoin(table schema.Tabler, on ...field.Expr) IArticleViewDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articleViewDo) Group(cols ...field.Expr) IArticleViewDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articleViewDo) Having(conds ...gen.Condition) IArticleViewDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articleViewDo) Limit(limit int) IArticleViewDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articleViewDo) Offset(offset int) IArticleViewDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articleViewDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleViewDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articleViewDo) Unscoped() IArticleViewDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articleViewDo) Create(values ...*orm_gen.ArticleView) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articleViewDo) CreateInBatches(values []*orm_gen.ArticleView, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articleViewDo) Save(values ...*orm_gen.ArticleView) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articleViewDo) First() (*orm_gen.ArticleView, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.ArticleView), nil
	}
}

func (a articleViewDo) Take() (*orm_gen.ArticleView, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.ArticleView), nil
	}
}

func (a articleViewDo) Last() (*orm_gen.ArticleView, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.ArticleView), nil
	}
}

func (a articleViewDo) Find() ([]*orm_gen.ArticleView, error) {
	result, err := a.DO.Find()
	return result.([]*orm_gen.ArticleView), err
}

func (a articleViewDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.ArticleView, err error) {
	buf := make([]*orm_gen.ArticleView, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articleViewDo) FindInBatches(result *[]*orm_gen.ArticleView, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articleViewDo) Attrs(attrs ...field.AssignExpr) IArticleViewDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articleViewDo) Assign(attrs ...field.AssignExpr) IArticleViewDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articleViewDo) Joins(fields ...field.RelationField) IArticleViewDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articleViewDo) Preload(fields ...field.RelationField) IArticleViewDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articleViewDo) FirstOrInit() (*orm_gen.ArticleView, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.ArticleView), nil
	}
}

func (a articleViewDo) FirstOrCreate() (*orm_gen.ArticleView, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.ArticleView), nil
	}
}

func (a articleViewDo) FindByPage(offset int, limit int) (result []*orm_gen.ArticleView, count int64, err error) {
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

func (a articleViewDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articleViewDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articleViewDo) Delete(models ...*orm_gen.ArticleView) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articleViewDo) withDO(do gen.Dao) *articleViewDo {
	a.DO = *do.(*gen.DO)
	return a
}
