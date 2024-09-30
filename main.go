package main

import (
	"fmt"
	"ginapi/controllers"
	"ginapi/services"

	internal "ginapi/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := internal.InitDb()

	if db == nil {
		fmt.Println("Database is not init properly")
	}

	notesService := &services.NotesService{}
	notesService.InitService(db)

	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"mesage": "Hello gin yy",
	// 	})
	// })

	// router.GET("/me/:id", func(ctx *gin.Context) {
	// 	var id = ctx.Param("id")
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"name": "Aziz",
	// 		"id":   id,
	// 	})
	// })

	// router.POST("/me", func(ctx *gin.Context) {
	// 	type MeRequest struct {
	// 		Email    string `json:"email"`
	// 		Password string `json:"password"`
	// 	}

	// 	var meRequest MeRequest
	// 	ctx.BindJSON(&meRequest)

	// 	ctx.JSON(http.StatusCreated, gin.H{
	// 		"message":  "User created",
	// 		"email":    meRequest.Email,
	// 		"password": meRequest.Password,
	// 	})
	// })

	noteControllers := &controllers.NoteController{}
	noteControllers.InitNotesController(router, *notesService)
	router.Run(":8000") //listen and server on 0.0.0.0:8080
}
