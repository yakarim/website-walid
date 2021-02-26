package controller

import "github.com/savsgio/atreugo/v11"

// Router ...
func (c *Ctrl) Router(ctx *atreugo.Atreugo) {
	ctx.GET("/", c.Index)
	ctx.GET("/login", c.Login)
	ctx.POST("/login__jwt", c.LoginJwt)
	ctx.GET("/admin", c.Admin).UseBefore(c.AuthMiddleware)
}
