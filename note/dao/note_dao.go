package dao

import (
	"github.com/jinzhu/gorm"
	"note/dao/model"
)

type NoteDAO struct {
	DB *gorm.DB
}

func NewNoteDAO(db *gorm.DB) *NoteDAO {
	return &NoteDAO{DB: db}
}

// Create a new note
func (dao *NoteDAO) CreateNote(note *model.Note) error {
	//解决只传入id时name会被覆盖为空的bug
	if note.Tags != nil {
		var tags []*model.Tag
		for _, t := range note.Tags {
			var tag model.Tag
			if err := dao.DB.First(&tag, t.ID).Error; err != nil {
				return err
			}
			tags = append(tags, &tag)
		}
		note.Tags = tags
	}
	return dao.DB.Create(note).Error
}

// Get a note by ID
func (dao *NoteDAO) GetNoteByID(id uint) (*model.Note, error) {
	var note model.Note
	err := dao.DB.Preload("Tags").First(&note, id).Error
	return &note, err
}

// Update a note
func (dao *NoteDAO) UpdateNote(note *model.Note) error {
	return dao.DB.Save(note).Error
}

// Delete a note
func (dao *NoteDAO) DeleteNote(id uint) error {
	return dao.DB.Unscoped().Delete(&model.Note{}, id).Error
}

// List all notes
func (dao *NoteDAO) ListNote() ([]model.Note, error) {
	var notes []model.Note
	err := dao.DB.Preload("Tags").Find(&notes).Error
	return notes, err
}
