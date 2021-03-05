package model

import (
	"github.com/yakarim/website-walid/database"
)

// Img type model
type Img struct {
	database.Images
}

// Create user Model ...
func (m Img) Create(post database.Images) error {
	if err := db.Model(&m.Images).Create(&post).Error; err != nil {
		return err
	}
	return nil
}

// QueryOne User ...
func (m Img) QueryOne(ID string) (database.Images, error) {
	var post database.Images
	if err := db.Where("uid = ?", ID).Find(&post).Error; err != nil {
		return post, err
	}
	return post, nil
}
