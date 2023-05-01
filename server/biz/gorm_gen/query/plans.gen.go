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

	"meetplan/biz/gorm_gen"
)

func newPlan(db *gorm.DB, opts ...gen.DOOption) plan {
	_plan := plan{}

	_plan.planDo.UseDB(db, opts...)
	_plan.planDo.UseModel(&gorm_gen.Plan{})

	tableName := _plan.planDo.TableName()
	_plan.ALL = field.NewAsterisk(tableName)
	_plan.ID = field.NewInt64(tableName, "id")
	_plan.TeacherID = field.NewInt64(tableName, "teacher_id")
	_plan.StartTime = field.NewTime(tableName, "start_time")
	_plan.Duration = field.NewInt64(tableName, "duration")
	_plan.Place = field.NewString(tableName, "place")
	_plan.Quota = field.NewInt8(tableName, "quota")
	_plan.Message = field.NewString(tableName, "message")
	_plan.Orders = planHasManyOrders{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Orders", "gorm_gen.Order"),
	}

	_plan.fillFieldMap()

	return _plan
}

type plan struct {
	planDo

	ALL       field.Asterisk
	ID        field.Int64
	TeacherID field.Int64
	StartTime field.Time
	Duration  field.Int64
	Place     field.String
	Quota     field.Int8
	Message   field.String
	Orders    planHasManyOrders

	fieldMap map[string]field.Expr
}

func (p plan) Table(newTableName string) *plan {
	p.planDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p plan) As(alias string) *plan {
	p.planDo.DO = *(p.planDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *plan) updateTableName(table string) *plan {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.TeacherID = field.NewInt64(table, "teacher_id")
	p.StartTime = field.NewTime(table, "start_time")
	p.Duration = field.NewInt64(table, "duration")
	p.Place = field.NewString(table, "place")
	p.Quota = field.NewInt8(table, "quota")
	p.Message = field.NewString(table, "message")

	p.fillFieldMap()

	return p
}

func (p *plan) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *plan) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 8)
	p.fieldMap["id"] = p.ID
	p.fieldMap["teacher_id"] = p.TeacherID
	p.fieldMap["start_time"] = p.StartTime
	p.fieldMap["duration"] = p.Duration
	p.fieldMap["place"] = p.Place
	p.fieldMap["quota"] = p.Quota
	p.fieldMap["message"] = p.Message

}

func (p plan) clone(db *gorm.DB) plan {
	p.planDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p plan) replaceDB(db *gorm.DB) plan {
	p.planDo.ReplaceDB(db)
	return p
}

type planHasManyOrders struct {
	db *gorm.DB

	field.RelationField
}

func (a planHasManyOrders) Where(conds ...field.Expr) *planHasManyOrders {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a planHasManyOrders) WithContext(ctx context.Context) *planHasManyOrders {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a planHasManyOrders) Session(session *gorm.Session) *planHasManyOrders {
	a.db = a.db.Session(session)
	return &a
}

func (a planHasManyOrders) Model(m *gorm_gen.Plan) *planHasManyOrdersTx {
	return &planHasManyOrdersTx{a.db.Model(m).Association(a.Name())}
}

type planHasManyOrdersTx struct{ tx *gorm.Association }

func (a planHasManyOrdersTx) Find() (result []*gorm_gen.Order, err error) {
	return result, a.tx.Find(&result)
}

func (a planHasManyOrdersTx) Append(values ...*gorm_gen.Order) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a planHasManyOrdersTx) Replace(values ...*gorm_gen.Order) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a planHasManyOrdersTx) Delete(values ...*gorm_gen.Order) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a planHasManyOrdersTx) Clear() error {
	return a.tx.Clear()
}

func (a planHasManyOrdersTx) Count() int64 {
	return a.tx.Count()
}

type planDo struct{ gen.DO }

type IPlanDo interface {
	gen.SubQuery
	Debug() IPlanDo
	WithContext(ctx context.Context) IPlanDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPlanDo
	WriteDB() IPlanDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPlanDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPlanDo
	Not(conds ...gen.Condition) IPlanDo
	Or(conds ...gen.Condition) IPlanDo
	Select(conds ...field.Expr) IPlanDo
	Where(conds ...gen.Condition) IPlanDo
	Order(conds ...field.Expr) IPlanDo
	Distinct(cols ...field.Expr) IPlanDo
	Omit(cols ...field.Expr) IPlanDo
	Join(table schema.Tabler, on ...field.Expr) IPlanDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPlanDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPlanDo
	Group(cols ...field.Expr) IPlanDo
	Having(conds ...gen.Condition) IPlanDo
	Limit(limit int) IPlanDo
	Offset(offset int) IPlanDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPlanDo
	Unscoped() IPlanDo
	Create(values ...*gorm_gen.Plan) error
	CreateInBatches(values []*gorm_gen.Plan, batchSize int) error
	Save(values ...*gorm_gen.Plan) error
	First() (*gorm_gen.Plan, error)
	Take() (*gorm_gen.Plan, error)
	Last() (*gorm_gen.Plan, error)
	Find() ([]*gorm_gen.Plan, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gorm_gen.Plan, err error)
	FindInBatches(result *[]*gorm_gen.Plan, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*gorm_gen.Plan) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPlanDo
	Assign(attrs ...field.AssignExpr) IPlanDo
	Joins(fields ...field.RelationField) IPlanDo
	Preload(fields ...field.RelationField) IPlanDo
	FirstOrInit() (*gorm_gen.Plan, error)
	FirstOrCreate() (*gorm_gen.Plan, error)
	FindByPage(offset int, limit int) (result []*gorm_gen.Plan, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPlanDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p planDo) Debug() IPlanDo {
	return p.withDO(p.DO.Debug())
}

func (p planDo) WithContext(ctx context.Context) IPlanDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p planDo) ReadDB() IPlanDo {
	return p.Clauses(dbresolver.Read)
}

func (p planDo) WriteDB() IPlanDo {
	return p.Clauses(dbresolver.Write)
}

func (p planDo) Session(config *gorm.Session) IPlanDo {
	return p.withDO(p.DO.Session(config))
}

func (p planDo) Clauses(conds ...clause.Expression) IPlanDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p planDo) Returning(value interface{}, columns ...string) IPlanDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p planDo) Not(conds ...gen.Condition) IPlanDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p planDo) Or(conds ...gen.Condition) IPlanDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p planDo) Select(conds ...field.Expr) IPlanDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p planDo) Where(conds ...gen.Condition) IPlanDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p planDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IPlanDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p planDo) Order(conds ...field.Expr) IPlanDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p planDo) Distinct(cols ...field.Expr) IPlanDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p planDo) Omit(cols ...field.Expr) IPlanDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p planDo) Join(table schema.Tabler, on ...field.Expr) IPlanDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p planDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPlanDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p planDo) RightJoin(table schema.Tabler, on ...field.Expr) IPlanDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p planDo) Group(cols ...field.Expr) IPlanDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p planDo) Having(conds ...gen.Condition) IPlanDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p planDo) Limit(limit int) IPlanDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p planDo) Offset(offset int) IPlanDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p planDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPlanDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p planDo) Unscoped() IPlanDo {
	return p.withDO(p.DO.Unscoped())
}

func (p planDo) Create(values ...*gorm_gen.Plan) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p planDo) CreateInBatches(values []*gorm_gen.Plan, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p planDo) Save(values ...*gorm_gen.Plan) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p planDo) First() (*gorm_gen.Plan, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen.Plan), nil
	}
}

func (p planDo) Take() (*gorm_gen.Plan, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen.Plan), nil
	}
}

func (p planDo) Last() (*gorm_gen.Plan, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen.Plan), nil
	}
}

func (p planDo) Find() ([]*gorm_gen.Plan, error) {
	result, err := p.DO.Find()
	return result.([]*gorm_gen.Plan), err
}

func (p planDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gorm_gen.Plan, err error) {
	buf := make([]*gorm_gen.Plan, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p planDo) FindInBatches(result *[]*gorm_gen.Plan, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p planDo) Attrs(attrs ...field.AssignExpr) IPlanDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p planDo) Assign(attrs ...field.AssignExpr) IPlanDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p planDo) Joins(fields ...field.RelationField) IPlanDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p planDo) Preload(fields ...field.RelationField) IPlanDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p planDo) FirstOrInit() (*gorm_gen.Plan, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen.Plan), nil
	}
}

func (p planDo) FirstOrCreate() (*gorm_gen.Plan, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen.Plan), nil
	}
}

func (p planDo) FindByPage(offset int, limit int) (result []*gorm_gen.Plan, count int64, err error) {
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

func (p planDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p planDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p planDo) Delete(models ...*gorm_gen.Plan) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *planDo) withDO(do gen.Dao) *planDo {
	p.DO = *do.(*gen.DO)
	return p
}
