// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePlanView = "plan_view"

// PlanView mapped from table <plan_view>
type PlanView struct {
	ID        int64     `gorm:"column:id;not null" json:"id"`
	TeacherID int64     `gorm:"column:teacher_id;not null" json:"teacher_id"`
	StartTime time.Time `gorm:"column:start_time;not null" json:"start_time"`
	Duration  int64     `gorm:"column:duration;not null" json:"duration"`
	Place     string    `gorm:"column:place;not null" json:"place"`
	Quota     int8      `gorm:"column:quota;not null" json:"quota"`
	Message   *string   `gorm:"column:message" json:"message"`
	IsValid   bool      `gorm:"column:is_valid;not null" json:"is_valid"`
	QuotaLeft int8      `gorm:"column:quota_left;not null" json:"quota_left"`
}

// TableName PlanView's table name
func (*PlanView) TableName() string {
	return TableNamePlanView
}
