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

func homefunc(ctx *atreugo.RequestCtx) error {
	ctx.Response.Header.Set("Content-Type", "text/html; charset=UTF-8")
	home("home", database.GetSession(ctx, "Username").(string), database.SessionAuth(ctx), ctx)
	return nil
}
