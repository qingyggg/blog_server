// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm_gen

import (
	"time"

	"gorm.io/gorm"
)

const TableNameArticleCollect = "article_collect"

// ArticleCollect 收藏表
type ArticleCollect struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增主键" json:"id"`                        // 自增主键
	UserID      int64          `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`                                   // 用户ID
	ArticleID   int64          `gorm:"column:article_id;not null;comment:文章ID" json:"article_id"`                             // 文章ID
	CollectName string         `gorm:"column:collect_name;default:default;comment:收藏的类型" json:"collect_name"`                 // 收藏的类型
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:收藏创建时间" json:"created_at"` // 收藏创建时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:收藏删除时间" json:"deleted_at"`                                    // 收藏删除时间
}

// TableName ArticleCollect's table name
func (*ArticleCollect) TableName() string {
	return TableNameArticleCollect
}
