package cfg

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
	"github.com/yakarim/website-walid/database"
)

// Auth login
func (c *Cfg) Auth(ctx *atreugo.RequestCtx) (string, bool) {

	return "", false
}

// SecurityTime ...
func (c *Cfg) SecurityTime(ctx *atreugo.RequestCtx) error {
	ctx.Response.Header.Set("X-Frame-Options", "SAMEORIGIN")
	ctx.Response.Header.Set("X-XSS-Protection", "1; mode=block")
	ctx.Response.Header.Set("Strict-Transport-Security", "max-age= max-age=30672000")
	ctx.Response.Header.Set("X-Content-Type-Options", "nosniff")
	return ctx.Next()
}

// AuthMiddleware ...
func (c *Cfg) AuthMiddleware(ctx *atreugo.RequestCtx) error {
	if string(ctx.Path()) == "/login" {
		return ctx.Next()
	}

	if database.SessionAuth(ctx) == false {
		return ctx.RedirectResponse("/login", fasthttp.StatusUnauthorized)
	}

	return ctx.Next()
}
