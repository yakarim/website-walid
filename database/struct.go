package database

import (
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Base ...
type Base struct {
	ID        string `gorm:"primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid.String())
}

// User struct.
type User struct {
	Base
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Auth ...
type Auth struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID   string `gorm:";not null;" json:"user_id"`
	AuthUUID string `gorm:"size:255;not null;" json:"auth_uuid"`
}

// Post struct.
type Post struct {
	Base
	Title        string  `json:"title"`
	Slug         string  `json:"slug"`
	Content      string  `json:"content"`
	AuthorID     string  `json:"author_id"`
	Status       int     `json:"status"`
	Type         int     `json:"type"`
	CommentCount big.Int `json:"comment_count"`
}
