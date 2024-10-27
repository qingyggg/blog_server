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

func newArticle(db *gorm.DB, opts ...gen.DOOption) article {
	_article := article{}

	_article.articleDo.UseDB(db, opts...)
	_article.articleDo.UseModel(&orm_gen.Article{})

	tableName := _article.articleDo.TableName()
	_article.ALL = field.NewAsterisk(tableName)
	_article.ID = field.NewInt64(tableName, "id")
	_article.UserID = field.NewBytes(tableName, "user_id")
	_article.Title = field.NewString(tableName, "title")
	_article.Note = field.NewString(tableName, "note")
	_article.CoverURL = field.NewString(tableName, "cover_url")
	_article.PublishTime = field.NewTime(tableName, "publish_time")
	_article.HashID = field.NewBytes(tableName, "hash_id")

	_article.fillFieldMap()

	return _article
}

// article 文章表
type article struct {
	articleDo

	ALL         field.Asterisk
	ID          field.Int64  // 文章ID
	UserID      field.Bytes  // 作者ID
	Title       field.String // 文章标题
	Note        field.String // 文章小记
	CoverURL    field.String // 背景图URL
	PublishTime field.Time   // 发布时间戳
	HashID      field.Bytes  // 文章的hashID值

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
	a.UserID = field.NewBytes(table, "user_id")
	a.Title = field.NewString(table, "title")
	a.Note = field.NewString(table, "note")
	a.CoverURL = field.NewString(table, "cover_url")
	a.PublishTime = field.NewTime(table, "publish_time")
	a.HashID = field.NewBytes(table, "hash_id")

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
	a.fieldMap = make(map[string]field.Expr, 7)
	a.fieldMap["id"] = a.ID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["title"] = a.Title
	a.fieldMap["note"] = a.Note
	a.fieldMap["cover_url"] = a.CoverURL
	a.fieldMap["publish_time"] = a.PublishTime
	a.fieldMap["hash_id"] = a.HashID
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
	Create(values ...*orm_gen.Article) error
	CreateInBatches(values []*orm_gen.Article, batchSize int) error
	Save(values ...*orm_gen.Article) error
	First() (*orm_gen.Article, error)
	Take() (*orm_gen.Article, error)
	Last() (*orm_gen.Article, error)
	Find() ([]*orm_gen.Article, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.Article, err error)
	FindInBatches(result *[]*orm_gen.Article, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*orm_gen.Article) (info gen.ResultInfo, err error)
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
	FirstOrInit() (*orm_gen.Article, error)
	FirstOrCreate() (*orm_gen.Article, error)
	FindByPage(offset int, limit int) (result []*orm_gen.Article, count int64, err error)
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

func (a articleDo) Create(values ...*orm_gen.Article) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articleDo) CreateInBatches(values []*orm_gen.Article, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articleDo) Save(values ...*orm_gen.Article) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articleDo) First() (*orm_gen.Article, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.Article), nil
	}
}

func (a articleDo) Take() (*orm_gen.Article, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.Article), nil
	}
}

func (a articleDo) Last() (*orm_gen.Article, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.Article), nil
	}
}

func (a articleDo) Find() ([]*orm_gen.Article, error) {
	result, err := a.DO.Find()
	return result.([]*orm_gen.Article), err
}

func (a articleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.Article, err error) {
	buf := make([]*orm_gen.Article, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articleDo) FindInBatches(result *[]*orm_gen.Article, batchSize int, fc func(tx gen.Dao, batch int) error) error {
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

func (a articleDo) FirstOrInit() (*orm_gen.Article, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.Article), nil
	}
}

func (a articleDo) FirstOrCreate() (*orm_gen.Article, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.Article), nil
	}
}

func (a articleDo) FindByPage(offset int, limit int) (result []*orm_gen.Article, count int64, err error) {
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

func (a articleDo) Delete(models ...*orm_gen.Article) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articleDo) withDO(do gen.Dao) *articleDo {
	a.DO = *do.(*gen.DO)
	return a
}
