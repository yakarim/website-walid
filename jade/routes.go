package jade

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/yakarim/website-walid/database"
)

// Jade ...
type Jade struct{}

// Router ...
func (j Jade) Router(ctx *atreugo.Atreugo) {
	ctx.GET("/home", homefunc)
}

func username(ctx *atreugo.RequestCtx) string {
	store, err := database.Session.Get(ctx.RequestCtx)
	if err != nil {
		return ""
	}
	u := store.Get("Username")
	if u != nil {
		return u.(string)
	}
	return ""
}

func homefunc(ctx *atreugo.RequestCtx) error {
	ctx.Response.Header.Set("Content-Type", "text/html; charset=UTF-8")
	home("home", username(ctx), database.SessionAuth(ctx), ctx)
	return nil
}
