package controllers

import (
	"ginapi/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NoteController struct {
	notesService services.NotesService
}

func (n *NoteController) InitController(noteService services.NotesService) *NoteController {
	n.notesService = noteService
	return n
}

func (n *NoteController) InitRoute(router *gin.Engine) {
	notes := router.Group("/notes")

	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreateNotes())
	notes.PUT("/", n.CreateNotes())
}

func (n *NoteController) GetNotes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status := ctx.Query("status")
		actualStatus, err := strconv.ParseBool(status)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		notes, err := n.notesService.GetNotesSrvices(actualStatus)

		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
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

func (n *NoteController) UpdateNotes() gin.HandlerFunc {
	type NoteBody struct {
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status"`
		Id     int    `json:"id" binding required`
	}
	return func(ctx *gin.Context) {
		var noteBody NoteBody
		if err := ctx.BindJSON(&noteBody); err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		note, err := n.notesService.UpdateNotesService(noteBody.Title, noteBody.Status, noteBody.Id)
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
