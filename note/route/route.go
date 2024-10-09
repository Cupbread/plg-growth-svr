package route

import (
	"github.com/gin-gonic/gin"
	"note/controller"
	"note/dao"
	"note/service"
	"note/storge/mysql"
)

func RegisterWebRoutes(router *gin.Engine) {

	db := mysql.DBGorm
	noteDAO := dao.NewNoteDAO(db)
	tagDAO := dao.NewTagDAO(db)

	noteService := service.NewNoteService(noteDAO)
	tagService := service.NewTagService(tagDAO)

	nc := controller.NewNotesController(noteService)
	tc := controller.NewTagsController(tagService)

	// Note Routes
	router.POST("/notes", nc.CreateNote)
	router.GET("/notes", nc.GetNotes)
	router.GET("/notes/:id", nc.GetNote)
	router.PUT("/notes/:id", nc.UpdateNote)
	router.DELETE("/notes/:id", nc.DeleteNote)
	router.DELETE("/notes/recycle/:id", nc.DeleteRecycleNote)

	// Tag Routes
	router.POST("/tags", tc.CreateTag)
	router.GET("/tags", tc.GetTags)
	router.GET("/tags/:id", tc.GetTag)
	router.PUT("/tags/:id", tc.UpdateTag)
	router.DELETE("/tags/:id", tc.DeleteTag)
}
