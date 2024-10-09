package service

import (
	"note/dao"
	"note/dao/model"
)

type NoteService struct {
	NoteDAO *dao.NoteDAO
}

func NewNoteService(dao *dao.NoteDAO) *NoteService {
	return &NoteService{NoteDAO: dao}
}

func (s *NoteService) CreateNote(note *model.Note) error {
	return s.NoteDAO.CreateNote(note)
}

func (s *NoteService) GetNoteByID(id uint) (*model.Note, error) {
	return s.NoteDAO.GetNoteByID(id)
}

func (s *NoteService) UpdateNote(note *model.Note) error {
	return s.NoteDAO.UpdateNote(note)
}

func (s *NoteService) DeleteNote(id uint) error {
	return s.NoteDAO.DeleteNote(id)
}

func (s *NoteService) ListNotes() ([]model.Note, error) {
	return s.NoteDAO.ListNote()
}
