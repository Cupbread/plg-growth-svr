package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"note/dao/model"
	"note/service"
	"strconv"
)

type NotesController struct {
	NoteService *service.NoteService
}

func NewNotesController(noteService *service.NoteService) *NotesController {
	return &NotesController{NoteService: noteService}
}

func (nc *NotesController) CreateNote(c *gin.Context) {
	var note model.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := nc.NoteService.CreateNote(&note); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	c.JSON(http.StatusCreated, note)
}

func (nc *NotesController) GetNotes(c *gin.Context) {
	notes, err := nc.NoteService.ListNotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
		return
	}

	c.JSON(http.StatusOK, notes)
}

func (nc *NotesController) GetNote(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	noteID := uint(id)

	note, err := nc.NoteService.GetNoteByID(noteID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(http.StatusOK, note)
}

func (nc *NotesController) UpdateNote(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	noteID := uint(id)

	var note model.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note.ID = noteID

	if err := nc.NoteService.UpdateNote(&note); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}

	c.JSON(http.StatusOK, note)
}

// 置status为deleted
func (nc *NotesController) DeleteNote(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	noteID := uint(id)

	if err := nc.NoteService.DeleteNote(noteID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}

// 执行删除逻辑
func (nc *NotesController) DeleteRecycleNote(c *gin.Context) {

}
