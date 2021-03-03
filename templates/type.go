package templates

import (
	"github.com/savsgio/atreugo/v11"
)

// Tmpl ...
type Tmpl struct {
}

// RouteQuick ...
func (t Tmpl) RouteQuick(ctx *atreugo.Atreugo) {
	ctx.GET("/Home", Home)
}

// Home ...
func Home(ctx *atreugo.RequestCtx) error {
	p := &MainPage{
		CTX: ctx.RequestCtx,
	}
	ctx.SetContentType("text/html; charset=utf-8")

	WritePageTemplate(ctx.RequestCtx, p)
	return nil
}
