package model

import "github.com/jinzhu/gorm"

type Status string

const (
	Deleted   Status = "deleted"
	Undeleted Status = "undeleted"
)

type Note struct {
	gorm.Model
	UserID  uint
	Title   string
	Content string
	Status  Status `gorm:"type:enum('deleted','undeleted');default:undeleted"`
	Tags    []*Tag `gorm:"many2many:note_tags;"`
}
