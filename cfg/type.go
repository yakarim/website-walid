package cfg

import (
	"github.com/CloudyKit/jet/v3"
	"github.com/yakarim/website-walid/database"
)

// Cfg ...
type Cfg struct{}

var (
	views = jet.NewHTMLSet("./template")

	db      = database.DB
	session = database.ServerSession
)
