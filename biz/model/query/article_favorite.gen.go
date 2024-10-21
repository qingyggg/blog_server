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

func newArticleFavorite(db *gorm.DB, opts ...gen.DOOption) articleFavorite {
	_articleFavorite := articleFavorite{}

	_articleFavorite.articleFavoriteDo.UseDB(db, opts...)
	_articleFavorite.articleFavoriteDo.UseModel(&orm_gen.ArticleFavorite{})

	tableName := _articleFavorite.articleFavoriteDo.TableName()
	_articleFavorite.ALL = field.NewAsterisk(tableName)
	_articleFavorite.ID = field.NewInt64(tableName, "id")
	_articleFavorite.ArticleID = field.NewBytes(tableName, "article_id")
	_articleFavorite.UserID = field.NewBytes(tableName, "user_id")
	_articleFavorite.Status = field.NewInt32(tableName, "status")

	_articleFavorite.fillFieldMap()

	return _articleFavorite
}

// articleFavorite 文章点赞表
type articleFavorite struct {
	articleFavoriteDo

	ALL       field.Asterisk
	ID        field.Int64 // 主键
	ArticleID field.Bytes // 评论文章ID
	UserID    field.Bytes // 用户ID
	Status    field.Int32 // 2：踩, 1：点赞

	fieldMap map[string]field.Expr
}

func (a articleFavorite) Table(newTableName string) *articleFavorite {
	a.articleFavoriteDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a articleFavorite) As(alias string) *articleFavorite {
	a.articleFavoriteDo.DO = *(a.articleFavoriteDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *articleFavorite) updateTableName(table string) *articleFavorite {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.ArticleID = field.NewBytes(table, "article_id")
	a.UserID = field.NewBytes(table, "user_id")
	a.Status = field.NewInt32(table, "status")

	a.fillFieldMap()

	return a
}

func (a *articleFavorite) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *articleFavorite) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 4)
	a.fieldMap["id"] = a.ID
	a.fieldMap["article_id"] = a.ArticleID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["status"] = a.Status
}

func (a articleFavorite) clone(db *gorm.DB) articleFavorite {
	a.articleFavoriteDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a articleFavorite) replaceDB(db *gorm.DB) articleFavorite {
	a.articleFavoriteDo.ReplaceDB(db)
	return a
}

type articleFavoriteDo struct{ gen.DO }

type IArticleFavoriteDo interface {
	gen.SubQuery
	Debug() IArticleFavoriteDo
	WithContext(ctx context.Context) IArticleFavoriteDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IArticleFavoriteDo
	WriteDB() IArticleFavoriteDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IArticleFavoriteDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IArticleFavoriteDo
	Not(conds ...gen.Condition) IArticleFavoriteDo
	Or(conds ...gen.Condition) IArticleFavoriteDo
	Select(conds ...field.Expr) IArticleFavoriteDo
	Where(conds ...gen.Condition) IArticleFavoriteDo
	Order(conds ...field.Expr) IArticleFavoriteDo
	Distinct(cols ...field.Expr) IArticleFavoriteDo
	Omit(cols ...field.Expr) IArticleFavoriteDo
	Join(table schema.Tabler, on ...field.Expr) IArticleFavoriteDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IArticleFavoriteDo
	RightJoin(table schema.Tabler, on ...field.Expr) IArticleFavoriteDo
	Group(cols ...field.Expr) IArticleFavoriteDo
	Having(conds ...gen.Condition) IArticleFavoriteDo
	Limit(limit int) IArticleFavoriteDo
	Offset(offset int) IArticleFavoriteDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleFavoriteDo
	Unscoped() IArticleFavoriteDo
	Create(values ...*orm_gen.ArticleFavorite) error
	CreateInBatches(values []*orm_gen.ArticleFavorite, batchSize int) error
	Save(values ...*orm_gen.ArticleFavorite) error
	First() (*orm_gen.ArticleFavorite, error)
	Take() (*orm_gen.ArticleFavorite, error)
	Last() (*orm_gen.ArticleFavorite, error)
	Find() ([]*orm_gen.ArticleFavorite, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.ArticleFavorite, err error)
	FindInBatches(result *[]*orm_gen.ArticleFavorite, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*orm_gen.ArticleFavorite) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IArticleFavoriteDo
	Assign(attrs ...field.AssignExpr) IArticleFavoriteDo
	Joins(fields ...field.RelationField) IArticleFavoriteDo
	Preload(fields ...field.RelationField) IArticleFavoriteDo
	FirstOrInit() (*orm_gen.ArticleFavorite, error)
	FirstOrCreate() (*orm_gen.ArticleFavorite, error)
	FindByPage(offset int, limit int) (result []*orm_gen.ArticleFavorite, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IArticleFavoriteDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a articleFavoriteDo) Debug() IArticleFavoriteDo {
	return a.withDO(a.DO.Debug())
}

func (a articleFavoriteDo) WithContext(ctx context.Context) IArticleFavoriteDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articleFavoriteDo) ReadDB() IArticleFavoriteDo {
	return a.Clauses(dbresolver.Read)
}

func (a articleFavoriteDo) WriteDB() IArticleFavoriteDo {
	return a.Clauses(dbresolver.Write)
}

func (a articleFavoriteDo) Session(config *gorm.Session) IArticleFavoriteDo {
	return a.withDO(a.DO.Session(config))
}

func (a articleFavoriteDo) Clauses(conds ...clause.Expression) IArticleFavoriteDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articleFavoriteDo) Returning(value interface{}, columns ...string) IArticleFavoriteDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articleFavoriteDo) Not(conds ...gen.Condition) IArticleFavoriteDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articleFavoriteDo) Or(conds ...gen.Condition) IArticleFavoriteDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articleFavoriteDo) Select(conds ...field.Expr) IArticleFavoriteDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articleFavoriteDo) Where(conds ...gen.Condition) IArticleFavoriteDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articleFavoriteDo) Order(conds ...field.Expr) IArticleFavoriteDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articleFavoriteDo) Distinct(cols ...field.Expr) IArticleFavoriteDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articleFavoriteDo) Omit(cols ...field.Expr) IArticleFavoriteDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articleFavoriteDo) Join(table schema.Tabler, on ...field.Expr) IArticleFavoriteDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articleFavoriteDo) LeftJoin(table schema.Tabler, on ...field.Expr) IArticleFavoriteDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articleFavoriteDo) RightJoin(table schema.Tabler, on ...field.Expr) IArticleFavoriteDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articleFavoriteDo) Group(cols ...field.Expr) IArticleFavoriteDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articleFavoriteDo) Having(conds ...gen.Condition) IArticleFavoriteDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articleFavoriteDo) Limit(limit int) IArticleFavoriteDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articleFavoriteDo) Offset(offset int) IArticleFavoriteDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articleFavoriteDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleFavoriteDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articleFavoriteDo) Unscoped() IArticleFavoriteDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articleFavoriteDo) Create(values ...*orm_gen.ArticleFavorite) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articleFavoriteDo) CreateInBatches(values []*orm_gen.ArticleFavorite, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articleFavoriteDo) Save(values ...*orm_gen.ArticleFavorite) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articleFavoriteDo) First() (*orm_gen.ArticleFavorite, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.ArticleFavorite), nil
	}
}

func (a articleFavoriteDo) Take() (*orm_gen.ArticleFavorite, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.ArticleFavorite), nil
	}
}

func (a articleFavoriteDo) Last() (*orm_gen.ArticleFavorite, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.ArticleFavorite), nil
	}
}

func (a articleFavoriteDo) Find() ([]*orm_gen.ArticleFavorite, error) {
	result, err := a.DO.Find()
	return result.([]*orm_gen.ArticleFavorite), err
}

func (a articleFavoriteDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.ArticleFavorite, err error) {
	buf := make([]*orm_gen.ArticleFavorite, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articleFavoriteDo) FindInBatches(result *[]*orm_gen.ArticleFavorite, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articleFavoriteDo) Attrs(attrs ...field.AssignExpr) IArticleFavoriteDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articleFavoriteDo) Assign(attrs ...field.AssignExpr) IArticleFavoriteDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articleFavoriteDo) Joins(fields ...field.RelationField) IArticleFavoriteDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articleFavoriteDo) Preload(fields ...field.RelationField) IArticleFavoriteDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articleFavoriteDo) FirstOrInit() (*orm_gen.ArticleFavorite, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.ArticleFavorite), nil
	}
}

func (a articleFavoriteDo) FirstOrCreate() (*orm_gen.ArticleFavorite, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.ArticleFavorite), nil
	}
}

func (a articleFavoriteDo) FindByPage(offset int, limit int) (result []*orm_gen.ArticleFavorite, count int64, err error) {
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

func (a articleFavoriteDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articleFavoriteDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articleFavoriteDo) Delete(models ...*orm_gen.ArticleFavorite) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articleFavoriteDo) withDO(do gen.Dao) *articleFavoriteDo {
	a.DO = *do.(*gen.DO)
	return a
}
