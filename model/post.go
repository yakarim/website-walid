package model

import (
	"errors"

	"github.com/gosimple/slug"
	"github.com/yakarim/website-walid/cfg"
	"github.com/yakarim/website-walid/database"
)

// Post model ...
type Post struct {
	database.Post
	cfg.Cfg
}

// QueryOne User ...
func (m Post) QueryOne(ID string) (database.Post, error) {
	var post database.Post
	if err := db.Where("id = ?", ID).Find(&post).Error; err != nil {
		return post, err
	}
	return post, nil
}

// QueryAll User ...
func (m Post) QueryAll() ([]database.Post, error) {
	var post []database.Post
	if err := db.Order("created_at desc").Find(&post).Error; err != nil {
		return post, err
	}
	return post, nil
}

// Create user Model ...
func (m Post) Create(post database.Post) error {
	var wpost database.Post
	if !db.Where("title = ?", post.Title).First(&wpost).RecordNotFound() {
		return errors.New("POST_FOUND")
	}
	var cpost database.Post
	cpost.Title = post.Title
	cpost.Slug = slug.Make(post.Title)
	cpost.AuthorID = post.AuthorID
	cpost.Status = post.Status
	cpost.Type = post.Type
	cpost.Content = post.Content
	if err := db.Model(&m.Post).Create(&cpost).Error; err != nil {
		return err
	}
	return nil
}

// Update post model.
func (m Post) Update(post database.Post) error {
	var wpost database.Post
	if db.Where("id = ?", post.ID).First(&wpost).RecordNotFound() {
		return errors.New("ID_NOT_FOUND")
	}
	var cpost database.Post
	cpost.ID = post.ID
	cpost.Title = post.Title
	cpost.Slug = slug.Make(post.Title)
	cpost.AuthorID = post.AuthorID
	cpost.Status = post.Status
	cpost.Type = post.Type
	cpost.Content = post.Content
	if err := db.Model(&m.Post).Update(&cpost).Error; err != nil {
		return err
	}
	return nil
}

// Delete ...
func (m Post) Delete(id string) error {
	if err := db.Delete(&database.Post{}, database.Post{Base: database.Base{ID: id}}).Error; err != nil {
		return err
	}
	return nil

}
