package controller

import "github.com/savsgio/atreugo/v11"

// Router ...
func (c *Ctrl) Router(ctx *atreugo.Atreugo) {
	ctx.GET("/", c.Index)
	ctx.GET("/login", c.Login)
	ctx.POST("/login__jwt", c.LoginJwt)
	ctx.GET("/admin", c.Admin).UseBefore(c.AuthMiddleware)
	ctx.GET("/useradmin", c.UserHTML).UseBefore(c.AuthMiddleware)

	user := ctx.NewGroupPath("/userapi")
	user.UseBefore(c.AuthMiddleware)
	user.GET("/query", c.UserC.QueryAll)
	user.POST("/", c.UserC.Create)
	user.DELETE("/{key}", c.UserC.Delete)
}
