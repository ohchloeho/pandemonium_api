package models

import "time"

type Project struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"createdBy"`
	Markdown    string    `json:"markdown"`
	VoiceNotes  []VoiceNote
}
