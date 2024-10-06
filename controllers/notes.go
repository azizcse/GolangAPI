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
		notes, err := n.notesService.GetNotesSrvices()
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
		}
		ctx.JSON(200, gin.H{
			"message": "Success",
			"status":  "S200",
			"notes":   notes,
		})
	}
}

func (n *NoteController) CreateNotes() gin.HandlerFunc {
	type NoteBody struct {
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status"`
	}
	return func(ctx *gin.Context) {
		var noteBody NoteBody
		if err := ctx.BindJSON(&noteBody); err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		note, err := n.notesService.CreateNotesService(noteBody.Title, noteBody.Status)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err,
			})
			return
		}
		ctx.JSON(200, gin.H{
			"note": note,
		})

	}
}
