// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"github.com/pkuphysu/meetplan/gorm_gen/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.PlanView{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.PlanView{}) fail: %s", err)
	}
}

func Test_planViewQuery(t *testing.T) {
	planView := newPlanView(db)
	planView = *planView.As(planView.TableName())
	_do := planView.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(planView.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <plan_view> fail:", err)
		return
	}

	_, ok := planView.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from planView success")
	}

	err = _do.Create(&model.PlanView{})
	if err != nil {
		t.Error("create item in table <plan_view> fail:", err)
	}

	err = _do.Save(&model.PlanView{})
	if err != nil {
		t.Error("create item in table <plan_view> fail:", err)
	}

	err = _do.CreateInBatches([]*model.PlanView{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <plan_view> fail:", err)
	}

	_, err = _do.Select(planView.ALL).Take()
	if err != nil {
		t.Error("Take() on table <plan_view> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <plan_view> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <plan_view> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <plan_view> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.PlanView{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <plan_view> fail:", err)
	}

	_, err = _do.Select(planView.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <plan_view> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <plan_view> fail:", err)
	}

	_, err = _do.Select(planView.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <plan_view> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <plan_view> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <plan_view> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <plan_view> fail:", err)
	}

	_, err = _do.ScanByPage(&model.PlanView{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <plan_view> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <plan_view> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <plan_view> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <plan_view> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <plan_view> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <plan_view> fail:", err)
	}
}
