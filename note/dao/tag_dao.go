package dao

import (
	"github.com/jinzhu/gorm"
	"note/dao/model"
)

type TagDAO struct {
	DB *gorm.DB
}

func NewTagDAO(db *gorm.DB) *TagDAO {
	return &TagDAO{DB: db}
}

// Create a new tag
func (dao *TagDAO) CreateTag(tag *model.Tag) error {
	return dao.DB.Create(tag).Error
}

// Get a tag by ID
func (dao *TagDAO) GetTagByID(id uint) (*model.Tag, error) {
	var tag model.Tag
	err := dao.DB.Preload("Notes").First(&tag, id).Error
	return &tag, err
}

// Update a tag
func (dao *TagDAO) UpdateTag(tag *model.Tag) error {
	return dao.DB.Save(tag).Error
}

// Delete a tag
func (dao *TagDAO) DeleteTag(id uint) error {
	return dao.DB.Delete(&model.Tag{}, id).Error
}

// List all tags
func (dao *TagDAO) ListTag() ([]model.Tag, error) {
	var tags []model.Tag
	err := dao.DB.Preload("Notes").Find(&tags).Error
	return tags, err
}
