package model

import "github.com/jinzhu/gorm"

type NoteTag struct {
	gorm.Model
	NoteID uint
	TagID  uint
}
