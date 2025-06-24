package models

import (
	"time"

	"github.com/google/uuid"
)

type NoteModel struct {
	Id        uuid.UUID `gorm:"type:uuid; primary_key; default:gen_random_uuid()"`
	Title     string    `gorm:"type:text; uniqueIndex; not null"`
	Content   string    `gorm:"type:text; not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`

	// TODO: put in many-to-many relationship for tags
}

func (NoteModel) TableName() string {
	return "notes"
}
