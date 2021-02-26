package model

import (
	"github.com/yakarim/website-walid/database"
)

// Mdl ...
type Mdl struct {
	UserM User
}

var db = database.DB
