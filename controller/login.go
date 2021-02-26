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
		"user":   c.SessionUsername(ctx),
		"signIn": c.SessionID(ctx),
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

		session := cfg.Session.StartFasthttp(ctx.RequestCtx)
		session.Set("ID", u.ID)
		session.Set("Username", u.Username)

		store, err := database.ServerSession.Get(ctx.RequestCtx)
		if err != nil {
			return err
		}

		defer func() {
			err = database.ServerSession.Save(ctx.RequestCtx, store)
		}()

		store.Set("foo", "bar")
		fmt.Println(store.Get("foo"))
		ctx.RedirectResponse("/admin", ctx.Response.StatusCode())
	}
	return nil
}

func (c Ctrl) deleteSession(ctx *atreugo.RequestCtx) {
	err := session.Destroy(ctx.RequestCtx)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

}
