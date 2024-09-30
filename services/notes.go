package services

import (
	"fmt"

	"gorm.io/gorm"

	internal "ginapi/internal/models"
)

type NotesService struct {
	db *gorm.DB
}

func (n *NotesService) InitService(database *gorm.DB) {
	n.db = database
	n.db.AutoMigrate(&internal.Notes{})
}

type Note struct {
	Id   int
	Name string
}

func (n *NotesService) GetNoteService() []Note {
	data := []Note{
		{Id: 1, Name: "Note 1"},
		{Id: 2, Name: "Note 2"},
	}

	err := n.db.Create(&internal.Notes{
		Id:     2,
		Title:  "Aziz",
		Status: true,
	})

	if err != nil {
		fmt.Println("DB created failed", err)
	}
	return data
}
