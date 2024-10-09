package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"note/dao/model"
	"note/service"
	"strconv"
)

type TagsController struct {
	TagService *service.TagService
}

func NewTagsController(tagService *service.TagService) *TagsController {
	return &TagsController{TagService: tagService}
}

func (tc *TagsController) CreateTag(c *gin.Context) {
	var tag model.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.TagService.CreateTag(&tag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tag"})
		return
	}

	c.JSON(http.StatusCreated, tag)
}

func (tc *TagsController) GetTags(c *gin.Context) {
	tags, err := tc.TagService.ListTags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tags"})
		return
	}

	c.JSON(http.StatusOK, tags)
}

func (tc *TagsController) GetTag(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	tagID := uint(id)

	tag, err := tc.TagService.GetTagByID(tagID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	c.JSON(http.StatusOK, tag)
}

func (tc *TagsController) UpdateTag(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	tagID := uint(id)

	var tag model.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag.ID = tagID

	if err := tc.TagService.UpdateTag(&tag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tag"})
		return
	}

	c.JSON(http.StatusOK, tag)
}

func (tc *TagsController) DeleteTag(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	tagID := uint(id)

	if err := tc.TagService.DeleteTag(tagID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted successfully"})
}
