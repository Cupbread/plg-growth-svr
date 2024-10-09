package service

import (
	"note/dao"
	"note/dao/model"
)

type TagService struct {
	TagDAO *dao.TagDAO
}

func NewTagService(dao *dao.TagDAO) *TagService {
	return &TagService{TagDAO: dao}
}

func (s *TagService) CreateTag(tag *model.Tag) error {
	return s.TagDAO.CreateTag(tag)
}

func (s *TagService) GetTagByID(id uint) (*model.Tag, error) {
	return s.TagDAO.GetTagByID(id)
}

func (s *TagService) UpdateTag(tag *model.Tag) error {
	return s.TagDAO.UpdateTag(tag)
}

func (s *TagService) DeleteTag(id uint) error {
	return s.TagDAO.DeleteTag(id)
}

func (s *TagService) ListTags() ([]model.Tag, error) {
	return s.TagDAO.ListTag()
}
