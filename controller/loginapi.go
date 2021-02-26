package controller

import (
	"log"

	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
	"github.com/yakarim/website-walid/cfg"
	"github.com/yakarim/website-walid/database"
)

// LoginPost ...
func (c *Ctrl) LoginPost(ctx *atreugo.RequestCtx) error {
	var userl database.User
	json.Unmarshal(ctx.PostBody(), &userl)
	if len(userl.Email) == 0 || len(userl.Password) == 0 {
		return c.JSON(ctx, cfg.H{"msg": "Isi Email & Password"}, 404)
	}
	u, b, msgerr := c.ValidateUser(string(userl.Email), string(userl.Password))
	if b == false {
		return c.JSON(ctx, cfg.H{"msg": msgerr.Error()}, 404)
	}
	uthData, err := c.CreateAuth(u.ID)
	if err != nil {
		return c.JSON(ctx, cfg.H{"msg": err.Error()}, fasthttp.StatusInternalServerError)
	}
	var authD cfg.AuthDetails
	authD.UserID = uthData.UserID
	authD.AuthUUID = uthData.AuthUUID
	token, loginErr := cfg.Authorize.SignIn(authD)
	if loginErr != nil {
		return c.JSON(ctx, cfg.H{"msg": "Please Try to login later"}, fasthttp.StatusForbidden)
	}
	return c.JSON(ctx, cfg.H{"msg": "sukses", "token": token, "id": authD.UserID, "expireAt": "expireAt"}, 200)
}

// Logout ...
func (c *Ctrl) Logout(ctx *atreugo.RequestCtx) error {
	au, err := cfg.ExtractTokenAuth(ctx)
	if err != nil {
		return c.JSON(ctx, cfg.H{"msg": "unauthorized"}, fasthttp.StatusUnauthorized)
	}
	delErr := c.DeleteAuth(au)
	if delErr != nil {
		log.Println(delErr)
		return c.JSON(ctx, cfg.H{"msg": "unauthorized"}, fasthttp.StatusUnauthorized)
	}
	return c.JSON(ctx, cfg.H{"msg": "succesfuly"}, fasthttp.StatusOK)
}
