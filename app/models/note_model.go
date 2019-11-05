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
    refNote := Note{}
    DB.AutoMigrate(&refNote)

    return &model
}

func (m *NoteModel) GetNote(ID int) (*Note, error) {
    var note *Note
    err := m.DB.Where("id = ?", ID).Find(note).Error

    return note, err
}

func (m *NoteModel) CreateNote(title, text string) (*Note, error) {
    var note = &Note{
        Title:     title,
        Text:      text,
    }

    err := m.DB.Create(&note).Error

    return note, err
}

func (m *NoteModel) GetNotes() ([]*Note, error) {
    var notes = make([]*Note, 0)
    err := m.DB.Find(&notes).Error

    return notes, err
}

