package model

import (
	"github.com/twinj/uuid"
	"github.com/yakarim/website-walid/cfg"
	"github.com/yakarim/website-walid/database"
)

// FetchAuth ...
func (m *Mdl) FetchAuth(authD *cfg.AuthDetails) (*database.Auth, error) {

	au := &database.Auth{}

	err := db.Debug().Where("user_id = ? AND auth_uuid = ?", authD.UserID, authD.AuthUUID).Take(&au).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}

// DeleteAuth ...
func (m *Mdl) DeleteAuth(authD *cfg.AuthDetails) error {
	au := &database.Auth{}

	if authD != nil {
		db := db.Where("user_id = ? AND auth_uuid = ?", authD.UserID, authD.AuthUUID).Take(&au).Delete(&au)
		if db.Error != nil {
			return db.Error
		}
	}

	return nil
}

// CreateAuth Once the user signup/login, create a row in the auth table, with a new uuid
func (m *Mdl) CreateAuth(userID string) (*database.Auth, error) {
	au := &database.Auth{}
	au.AuthUUID = uuid.NewV4().String() //generate a new UUID each time
	au.UserID = userID
	err := db.Debug().Create(&au).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}
