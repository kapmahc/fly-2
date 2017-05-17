package web

import "time"

const (
	// TypeMARKDOWN markdown format
	TypeMARKDOWN = "markdown"
	// TypeHTML html format
	TypeHTML = "html"
)

// Timestamp timestamp
type Timestamp struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

//Model base model
type Model struct {
	Timestamp
	UpdatedAt time.Time `json:"updatedAt"`
}

// Media media
type Media struct {
	Model
	Body string `json:"body"`
	Type string `json:"type"`
}

// Dropdown dropdown
type Dropdown struct {
	Label string
	Items []*Link
}

// Link link
type Link struct {
	Label string
	Href  string
}
