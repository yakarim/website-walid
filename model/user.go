package model

import (
	"errors"

	"github.com/yakarim/website-walid/cfg"
	"github.com/yakarim/website-walid/database"
)

// User model ...
type User struct {
	database.User
	cfg.Cfg
}

// QueryOne User ...
func (m *User) QueryOne(ID string) (database.User, error) {
	var user database.User
	if err := db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// QueryAll User ...
func (m *User) QueryAll() ([]database.User, error) {
	var user []database.User
	if err := db.Order("created_at desc").Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Create user Model ...
func (m *User) Create(user database.User) error {
	var wuser database.User
	if !db.Where("email = ?", user.Email).First(&wuser).RecordNotFound() {
		return errors.New("EMAIL_FOUND")
	}
	var cuser database.User
	cuser.Username = user.Username
	cuser.Email = user.Email
	cuser.Password = m.HashAndSalt(m.GetPwd(user.Password))
	if err := db.Model(&m.User).Create(&cuser).Error; err != nil {
		return err
	}
	return nil
}

// Update user.
func (m *User) Update(user database.User) error {
	if !db.Where("email = ?", user.Email).First(&user).RecordNotFound() {
		return errors.New("EMAIL_FOUND")
	}
	if user.Password == "" {
		if err := db.Model(&m.User).Update(&user).Error; err != nil {
			return err
		}
	} else {
		user.Password = m.HashAndSalt(m.GetPwd(user.Password))
		if err := db.Model(&m.User).Update(&user).Error; err != nil {
			return err
		}
	}
	return nil
}

// Delete ...
func (m *User) Delete(id string) error {
	if err := db.Delete(&database.User{}, database.User{Base: database.Base{ID: id}}).Error; err != nil {
		return err
	}
	return nil

}
