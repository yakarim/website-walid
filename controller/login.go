package controller

import (
	"fmt"

	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
	"github.com/yakarim/website-walid/cfg"
	"github.com/yakarim/website-walid/database"
)

// Login pages.
func (c Ctrl) Login(ctx *atreugo.RequestCtx) error {
	c.SessionDestroy(ctx)
	return c.HTML(ctx, 200, "pages/login", cfg.H{
		"title":  "Login",
		"user":   database.GetSession(ctx, "Username"),
		"signIn": database.SessionAuth(ctx),
	})
}

// LoginJwt ...
func (c Ctrl) LoginJwt(ctx *atreugo.RequestCtx) error {

	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	if len(email) == 0 || len(password) == 0 {
		ctx.RedirectResponse("/login", ctx.Response.StatusCode())
	}
	u, b, _ := c.ValidateUser(string(email), string(password))
	if b == false {
		ctx.RedirectResponse("/login", ctx.Response.StatusCode())
	} else {
		/*
			session := cfg.Session.StartFasthttp(ctx.RequestCtx)
			session.Set("ID", u.ID)
			session.Set("Username", u.Username)
		*/
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
		fmt.Println(store.Get("Username"))

		ctx.RedirectResponse("/admin", ctx.Response.StatusCode())
	}
	return nil
}

func (c Ctrl) deleteSession(ctx *atreugo.RequestCtx) {

}
