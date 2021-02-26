package controller

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/yakarim/website-walid/cfg"
	"github.com/yakarim/website-walid/database"
)

// Index ...
func (c Ctrl) Index(ctx *atreugo.RequestCtx) error {

	return c.HTML(ctx, 200, "pages/index", cfg.H{
		"title":  "Home",
		"user":   database.GetSession(ctx, "Username"),
		"signIn": database.SessionAuth(ctx),
	})
}
