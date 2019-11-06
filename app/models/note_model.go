package models

import (
	"github.com/jinzhu/gorm"
)

type NoteModel struct {
	DB *gorm.DB
}

func NewNoteModel(DB *gorm.DB) *NoteModel {
	model := NoteModel{
		DB: DB,
	}
	note := Note{}
	DB.AutoMigrate(&note)

	return &model
}

func (m *NoteModel) GetNote(noteID string) *Note {
	fetchNote := Note{}
	notFound := m.DB.Where("ID = ?", noteID).First(&fetchNote).RecordNotFound()

	if notFound {
		return nil
	}

	return &fetchNote
}

func (m *NoteModel) Upsert(noteID, title, text string) (*Note, error) {
	newNote := Note{
		ID:    noteID,
		Title: title,
		Text:  text,
	}

	fetchNote := Note{}

	err := m.DB.Where("ID = ?", noteID).Assign(newNote).FirstOrCreate(&fetchNote).Error
	if err != nil {
		return nil, err
	}

	return &fetchNote, err
}

func (m *NoteModel) GetNotes() ([]*Note, error) {
	notes := make([]*Note, 0)
	err := m.DB.Find(&notes).Error

	return notes, err
}

func (m *NoteModel) RemoveNote(noteID string) error {
	note := Note{
		ID: noteID,
	}
	return m.DB.Unscoped().Delete(&note).Error
}
