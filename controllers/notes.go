package controllers

import (
	"ginapi/services"

	"github.com/gin-gonic/gin"
)

type NoteController struct {
	notesService services.NotesService
}

func (n *NoteController) InitNotesController(router *gin.Engine, notesService services.NotesService) {
	notes := router.Group("/notes")

	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreateNotes())
	n.notesService = notesService
}

func (n *NoteController) GetNotes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"notes": n.notesService.GetNoteService(),
		})
	}
}

func (n *NoteController) CreateNotes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"notes": "Created",
		})
	}
}
