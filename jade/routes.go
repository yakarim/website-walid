package jade

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/yakarim/website-walid/database"
)

// Jade ...
type Jade struct{}

// Routes ...
func (j Jade) Routes(ctx *atreugo.Atreugo) {
	ctx.GET("/home", Home)
}

// Home ...
func Home(ctx *atreugo.RequestCtx) error {

	Index("Jade.go", database.SessionAuth(ctx), ctx)
	ctx.SetContentType("text/html; charset=utf-8")

	return nil
}
