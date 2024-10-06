package services

import (
	"fmt"
	internal "ginapi/internal/models"

	"gorm.io/gorm"
)

type NotesService struct {
	db *gorm.DB
}

func (n *NotesService) InitService(database *gorm.DB) {
	n.db = database
	//n.db.AutoMigrate(&internal.Notes{})
}

type Note struct {
	Id   int
	Name string
}

func (n *NotesService) GetNotesSrvices(status bool) ([]*internal.Notes, error) {
	var notes []*internal.Notes
	if err := n.db.Where("status=?", status).Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func (n *NotesService) CreateNotesService(title string, status bool) (*internal.Notes, error) {
	note := &internal.Notes{
		Title:  title,
		Status: status,
	}

	if err := n.db.Create(note).Error; err != nil {
		fmt.Print(err)
		return nil, err
	}
	return note, nil
}
