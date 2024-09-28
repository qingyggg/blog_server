// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm_gen

import (
	"time"

	"gorm.io/gorm"
)

const TableNameFollow = "follows"

// Follow 关注表
type Follow struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增主键" json:"id"`                            // 自增主键
	UserID     int64          `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`                                       // 用户ID
	FollowerID int64          `gorm:"column:follower_id;not null;comment:粉丝ID" json:"follower_id"`                               // 粉丝ID
	CreatedAt  time.Time      `gorm:"column:created_at;not null;default:current_timestamp();comment:关注关系创建时间" json:"created_at"` // 关注关系创建时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;comment:关注关系删除时间" json:"deleted_at"`                                      // 关注关系删除时间
}

// TableName Follow's table name
func (*Follow) TableName() string {
	return TableNameFollow
}
