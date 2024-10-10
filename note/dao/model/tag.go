package model

import "github.com/jinzhu/gorm"

//type Tag struct {
//	TagID      uint64    `gorm:"primaryKey;autoIncrement;comment:'标签主键'" json:"tag_id"`
//	Name       string    `gorm:"size:255;not null;comment:'标签名称'" json:"name"`
//	CreateTime time.Time `gorm:"default:NULL;comment:'标签创建时间'" json:"create_time"`
//	UpdateTime time.Time `gorm:"default:NULL;comment:'标签更新时间'" json:"update_time"`
//	DeleteTime time.Time `gorm:"default:NULL;comment:'标签删除时间'" json:"delete_time"`
//	UserID     uint64    `gorm:"not null;index:idx_user_id;comment:'创建标签用户的id'" json:"user_id"`
//	Notes      []Note    `gorm:"many2many:note_tags" json:"notes,omitempty"`
//}

type Tag struct {
	gorm.Model
	Name   string
	UserID uint64
	Notes  []Note `gorm:"many2many:note_tags;"`
}
