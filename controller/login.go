package controller

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
	"github.com/yakarim/website-walid/cfg"
)

// Login pages.
func (c Ctrl) Login(ctx *atreugo.RequestCtx) error {
	c.deleteSession(ctx)
	return c.HTML(ctx, 200, "pages/login", cfg.H{
		"title": "Login",
	})
}

// LoginJwt ...
func (c Ctrl) LoginJwt(ctx *atreugo.RequestCtx) error {

	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	if len(email) == 0 || len(password) == 0 {
		return ctx.RedirectResponse("/login", ctx.Response.StatusCode())
	}
	u, b, _ := c.ValidateUser(string(email), string(password))
	if b == false {
		return ctx.RedirectResponse("/login", ctx.Response.StatusCode())
	}
	store, err := session.Get(ctx.RequestCtx)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
	defer func() {
		if err := session.Save(ctx.RequestCtx, store); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}
	}()
	store.Set("ID", u.ID)
	store.Set("Username", u.Username)
	return ctx.RedirectResponse("/admin", ctx.Response.StatusCode())
}

func (c Ctrl) deleteSession(ctx *atreugo.RequestCtx) {
	session.Destroy(ctx.RequestCtx)
}
