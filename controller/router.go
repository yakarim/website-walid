package controller

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/yakarim/website-walid/jade"
)

// Router ...
func (c *Ctrl) Router(ctx *atreugo.Atreugo) {
	ctx.GET("/", c.Index)
	ctx.GET("/login", c.Login)
	ctx.POST("/loginjwt", c.LoginJwt)
	ctx.GET("/admin", c.Admin).UseBefore(c.AuthMiddleware)
	ctx.GET("/useradmin", c.UserHTML).UseBefore(c.AuthMiddleware)
	ctx.GET("/postadmin", c.PostHTML).UseBefore(c.AuthMiddleware)
	ctx.PUT("/media-save", c.Img.ImageSave).UseBefore(c.AuthMiddleware)
	ctx.GET("/media-get-{key}", c.Img.ImageGet).UseBefore(c.AuthMiddleware)
	ctx.GET("/media-get-all", c.Img.GetImages).UseBefore(c.AuthMiddleware)

	user := ctx.NewGroupPath("/userapi")
	user.UseBefore(c.AuthMiddleware)
	user.GET("/query", c.User.QueryAll)
	user.POST("/", c.User.Create)
	user.DELETE("/{key}", c.User.Delete)

	post := ctx.NewGroupPath("/postapi")
	post.UseBefore(c.AuthMiddleware)
	post.GET("/query", c.Post.QueryAll)
	post.POST("/", c.Post.Create)
	post.PUT("/{key}", c.Post.Update)
	post.DELETE("/{key}", c.Post.Delete)
}

// Home ...
func (c Ctrl) Home(ctx *atreugo.RequestCtx) error {
	ctx.Response.Header.Set("Content-Type", "text/html; charset=UTF-8")
	jade.Index("home", c.Username(ctx), c.SessionAuth(ctx), nil, ctx)
	return nil
}
