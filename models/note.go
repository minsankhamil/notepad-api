package models

import (
	"time"
)

type Note struct {
	IdNote      uint   `gorm:"primaryKey" json:"id_note"`
	NoteHeading string `json:"note_heading"`
	NoteBody    string `json:"note_body"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
