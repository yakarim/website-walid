package jade

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/yakarim/website-walid/database"
	"github.com/yakarim/website-walid/model"
)

// Jade ...
type Jade struct{}

var (
	m model.Img
)

// Router ...
func (j Jade) Router(ctx *atreugo.Atreugo) {
	ctx.GET("/home", homefunc)
	ctx.GET("/imgdelete/{key}", imgdelete)
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
	data, _ := m.QueryAll()
	ctx.Response.Header.Set("Content-Type", "text/html; charset=UTF-8")
	Index("home", username(ctx), database.SessionAuth(ctx), data, ctx)
	return nil
}

func imgdelete(ctx *atreugo.RequestCtx) error {
	key := ctx.UserValue("key")
	m.Delete(key.(string))
	return ctx.RedirectResponse("/home", 200)
}
