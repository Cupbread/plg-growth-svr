package model

import (
	"github.com/jinzhu/gorm"
)

//type Note struct {
//	NoteID     uint64    `gorm:"primaryKey;autoIncrement;comment:'记事本主键'" json:"note_id"`
//	CreateTime time.Time `gorm:"default:NULL;comment:'记事本创建时间'" json:"create_time"`
//	UpdateTime time.Time `gorm:"default:NULL;comment:'记事本更新时间'" json:"update_time"`
//	DeleteTime time.Time `gorm:"default:NULL;comment:'记事本删除时间'" json:"delete_time"`
//	UserID     uint64    `gorm:"not null;index:idx_user_id;comment:'创建记事本用户的id'" json:"user_id"`
//	Title      string    `gorm:"size:255;default:NULL;comment:'记事本标题'" json:"title"`
//	Content    string    `gorm:"type:text;default:NULL;comment:'记事本内容'" json:"content"`
//	NoteStatus int8      `gorm:"default:0;comment:'记事本状态，0表示未删除，1表示回收站，2表示彻底删除'" json:"note_status"`
//	Tags       []Tag     `gorm:"many2many:note_tag" json:"tags,omitempty"`
//}

type Note struct {
	gorm.Model
	UserID     uint64
	Title      string
	Content    string
	NoteStatus int8
	Tags       []Tag `gorm:"many2many:note_tags;"`
}
