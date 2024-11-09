package models

import "time"

type VoiceNote struct {
    ID          int       `json:"id"`
    ProjectID   int       `json:"project_id"`
    Content     string    `json:"content"`      // Transcribed text
    AudioFile   string    `json:"audio_file"`   // Path or URL to the original audio
    CreatedAt   time.Time `json:"created_at"`
}
