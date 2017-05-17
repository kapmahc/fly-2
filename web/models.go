package web

import (
	"time"

	"github.com/gin-gonic/gin"
)

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

// NewDropdown new dropdown
func NewDropdown(label, icon string, items ...Link) Dropdown {
	return Dropdown{"label": label, "icon": icon, "items": items}
}

// Dropdown dropdown
type Dropdown gin.H

// Append append items
func (p Dropdown) Append(items ...Link) {
	p["items"] = append(p["items"].([]Link), items...)
}

// Link link
type Link gin.H

// NewLink new link
func NewLink(label, href string) Link {
	return Link{"label": label, "href": href}
}
