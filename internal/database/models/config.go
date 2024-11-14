package models

import "time"

// Note represents the general structure of a note for various content types.
type Note struct {
	ID          string         `json:"id" bson:"_id"`               // Unique identifier for the note
	Title       string         `json:"title" bson:"title"`          // Title of the note
	ContentType ContentType    `json:"content_type" bson:"content_type"` // Type of content (e.g., Recipe, StudyNote)
	Content     Content        `json:"content" bson:"content"`      // Embedded struct for content-specific data
	Tags        []string       `json:"tags" bson:"tags"`            // Tags for categorizing the note
	CreatedAt   time.Time      `json:"created_at" bson:"created_at"` // Creation timestamp
	UpdatedAt   time.Time      `json:"updated_at" bson:"updated_at"` // Last updated timestamp
	Metadata    Metadata       `json:"metadata" bson:"metadata"`    // Additional metadata (e.g., priority, category)
}

// ContentType defines the allowed types for content, which can expand as needed.
type ContentType string

const (
	Recipe     ContentType = "Recipe"
	StudyNote  ContentType = "StudyNote"
	Journal    ContentType = "Journal"
	CheatSheet ContentType = "CheatSheet"
	BusinessIdea ContentType = "BusinessIdea"
	TravelPlan ContentType = "TravelPlan"
	Expense    ContentType = "Expense"
	// Add more content types as needed
)

// Content is a flexible struct that contains fields specific to each content type.
// Fields are optional depending on the ContentType of the Note.
type Content struct {
	Text           string     `json:"text,omitempty" bson:"text,omitempty"`           // General text (for notes or ideas)
	Ingredients    []string   `json:"ingredients,omitempty" bson:"ingredients,omitempty"` // List of ingredients (for recipes)
	Steps          []string   `json:"steps,omitempty" bson:"steps,omitempty"`         // Steps (for recipes)
	Subject        string     `json:"subject,omitempty" bson:"subject,omitempty"`     // Subject (for study notes)
	Source         string     `json:"source,omitempty" bson:"source,omitempty"`       // Source of information (e.g., URL, book)
	ProgrammingLang string    `json:"programming_lang,omitempty" bson:"programming_lang,omitempty"` // Language (for cheat sheets)
	Amount         float64    `json:"amount,omitempty" bson:"amount,omitempty"`       // Amount spent (for expenses)
	Location       string     `json:"location,omitempty" bson:"location,omitempty"`   // Location (for travel plans)
	Date           time.Time  `json:"date,omitempty" bson:"date,omitempty"`           // Specific date (e.g., for journals or travel plans)
	// Add more fields as needed for specific content types
}

// Metadata contains additional information about the note for sorting or prioritization.
type Metadata struct {
	Priority    int      `json:"priority,omitempty" bson:"priority,omitempty"` // Priority level (e.g., 1 = high, 5 = low)
	Categories  []string `json:"categories,omitempty" bson:"categories,omitempty"` // Categories to group notes
	RelatedTags []string `json:"related_tags,omitempty" bson:"related_tags,omitempty"` // Additional tags for linking
}
