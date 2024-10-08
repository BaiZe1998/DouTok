// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameFollow = "follow"

// Follow mapped from table <follow>
type Follow struct {
	ID           int64     `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	UserID       int64     `gorm:"column:user_id;type:bigint;not null;index:user_id_idx,priority:1" json:"user_id"`
	TargetUserID int64     `gorm:"column:target_user_id;type:bigint;not null;index:user_id_idx,priority:2;comment:被关注的用户id" json:"target_user_id"` // 被关注的用户id
	IsDeleted    bool      `gorm:"column:is_deleted;type:tinyint(1);not null;index:user_id_idx,priority:3" json:"is_deleted"`
	CreateTime   time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName Follow's table name
func (*Follow) TableName() string {
	return TableNameFollow
}
