// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"SimpleDouYin/app/dao/model"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newPublish(db *gorm.DB) publish {
	_publish := publish{}

	_publish.publishDo.UseDB(db)
	_publish.publishDo.UseModel(&model.Publish{})

	tableName := _publish.publishDo.TableName()
	_publish.ALL = field.NewAsterisk(tableName)
	_publish.Title = field.NewString(tableName, "title")
	_publish.PublishTime = field.NewInt64(tableName, "publish_time")
	_publish.UserID = field.NewInt64(tableName, "user_id")
	_publish.VideoID = field.NewInt64(tableName, "video_id")

	_publish.fillFieldMap()

	return _publish
}

type publish struct {
	publishDo publishDo

	ALL         field.Asterisk
	Title       field.String // 标题
	PublishTime field.Int64  // 时间戳 unix
	UserID      field.Int64  // 用户id
	VideoID     field.Int64  // 视频id

	fieldMap map[string]field.Expr
}

func (p publish) Table(newTableName string) *publish {
	p.publishDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p publish) As(alias string) *publish {
	p.publishDo.DO = *(p.publishDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *publish) updateTableName(table string) *publish {
	p.ALL = field.NewAsterisk(table)
	p.Title = field.NewString(table, "title")
	p.PublishTime = field.NewInt64(table, "publish_time")
	p.UserID = field.NewInt64(table, "user_id")
	p.VideoID = field.NewInt64(table, "video_id")

	p.fillFieldMap()

	return p
}

func (p *publish) WithContext(ctx context.Context) *publishDo { return p.publishDo.WithContext(ctx) }

func (p publish) TableName() string { return p.publishDo.TableName() }

func (p publish) Alias() string { return p.publishDo.Alias() }

func (p *publish) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *publish) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["title"] = p.Title
	p.fieldMap["publish_time"] = p.PublishTime
	p.fieldMap["user_id"] = p.UserID
	p.fieldMap["video_id"] = p.VideoID
}

func (p publish) clone(db *gorm.DB) publish {
	p.publishDo.ReplaceDB(db)
	return p
}

type publishDo struct{ gen.DO }

func (p publishDo) Debug() *publishDo {
	return p.withDO(p.DO.Debug())
}

func (p publishDo) WithContext(ctx context.Context) *publishDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p publishDo) ReadDB() *publishDo {
	return p.Clauses(dbresolver.Read)
}

func (p publishDo) WriteDB() *publishDo {
	return p.Clauses(dbresolver.Write)
}

func (p publishDo) Clauses(conds ...clause.Expression) *publishDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p publishDo) Returning(value interface{}, columns ...string) *publishDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p publishDo) Not(conds ...gen.Condition) *publishDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p publishDo) Or(conds ...gen.Condition) *publishDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p publishDo) Select(conds ...field.Expr) *publishDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p publishDo) Where(conds ...gen.Condition) *publishDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p publishDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *publishDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p publishDo) Order(conds ...field.Expr) *publishDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p publishDo) Distinct(cols ...field.Expr) *publishDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p publishDo) Omit(cols ...field.Expr) *publishDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p publishDo) Join(table schema.Tabler, on ...field.Expr) *publishDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p publishDo) LeftJoin(table schema.Tabler, on ...field.Expr) *publishDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p publishDo) RightJoin(table schema.Tabler, on ...field.Expr) *publishDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p publishDo) Group(cols ...field.Expr) *publishDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p publishDo) Having(conds ...gen.Condition) *publishDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p publishDo) Limit(limit int) *publishDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p publishDo) Offset(offset int) *publishDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p publishDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *publishDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p publishDo) Unscoped() *publishDo {
	return p.withDO(p.DO.Unscoped())
}

func (p publishDo) Create(values ...*model.Publish) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p publishDo) CreateInBatches(values []*model.Publish, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p publishDo) Save(values ...*model.Publish) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p publishDo) First() (*model.Publish, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Publish), nil
	}
}

func (p publishDo) Take() (*model.Publish, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Publish), nil
	}
}

func (p publishDo) Last() (*model.Publish, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Publish), nil
	}
}

func (p publishDo) Find() ([]*model.Publish, error) {
	result, err := p.DO.Find()
	return result.([]*model.Publish), err
}

func (p publishDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Publish, err error) {
	buf := make([]*model.Publish, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p publishDo) FindInBatches(result *[]*model.Publish, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p publishDo) Attrs(attrs ...field.AssignExpr) *publishDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p publishDo) Assign(attrs ...field.AssignExpr) *publishDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p publishDo) Joins(fields ...field.RelationField) *publishDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p publishDo) Preload(fields ...field.RelationField) *publishDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p publishDo) FirstOrInit() (*model.Publish, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Publish), nil
	}
}

func (p publishDo) FirstOrCreate() (*model.Publish, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Publish), nil
	}
}

func (p publishDo) FindByPage(offset int, limit int) (result []*model.Publish, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p publishDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p publishDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p publishDo) Delete(models ...*model.Publish) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *publishDo) withDO(do gen.Dao) *publishDo {
	p.DO = *do.(*gen.DO)
	return p
}