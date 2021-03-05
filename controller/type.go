package controller

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/yakarim/website-walid/cfg"
	"github.com/yakarim/website-walid/database"
)

// Ctrl ...
type Ctrl struct {
	User User
	Post Post
	Img  Img
	cfg.Cfg
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var session = database.Session

/*
// RouterAPI ...
func (c *Ctrl) RouterAPI(ctx *atreugo.Atreugo) {

	ctx.POST("/user/create", c.UserC.Create)

	ctx.POST("/login", c.LoginPost)
	ctx.GET("/logoutdest", c.Logout)

	user := ctx.NewGroupPath("/user")
	user.UseBefore(c.AuthMiddleware)
	user.GET("/query", c.UserC.QueryAll)
	user.GET("/query/{key}", c.UserC.QueryOne)
	user.PUT("/update", c.UserC.Update)
	user.DELETE("/delete/{key}", c.UserC.Delete)

	admin := ctx.NewGroupPath("/admin")
	admin.UseBefore(c.AuthMiddleware)
	admin.GET("/user", c.Admin)
}
*/
