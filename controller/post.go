package controller

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/yakarim/website-walid/cfg"
	"github.com/yakarim/website-walid/database"
	"github.com/yakarim/website-walid/model"
)

// Post type
type Post struct {
	p model.Post
	u model.User
	cfg.Cfg
}

// PostHTML ...
func (c Ctrl) PostHTML(ctx *atreugo.RequestCtx) error {
	return c.HTML(ctx, 200, "admin/post", cfg.H{
		"title":  "user",
		"user":   database.GetSession(ctx, "Username"),
		"signIn": database.SessionAuth(ctx),
	})
}

// QueryOne User...
func (c Post) QueryOne(ctx *atreugo.RequestCtx) error {
	key := ctx.UserValue("key")
	data, _ := c.p.QueryOne(key.(string))
	return c.JSON(ctx, cfg.H{"post": data}, 200)
}

// QueryAll User...
func (c Post) QueryAll(ctx *atreugo.RequestCtx) error {
	data, _ := c.p.QueryAll()
	for i := 0; i < len(data); i++ {
		user, _ := c.u.QueryOne(data[i].AuthorID)
		data[i].AuthorID = user.Username
	}
	return c.JSON(ctx, cfg.H{"posts": data}, 200)
}

// Create User ...
func (c Post) Create(ctx *atreugo.RequestCtx) error {
	var postl database.Post
	json.Unmarshal(ctx.PostBody(), &postl)

	if string(postl.Title) == "" {
		return c.JSON(ctx, cfg.H{"msg": "form harus disi"}, 404)
	}
	postl.AuthorID = database.GetSession(ctx, "ID").(string)
	if c.p.Create(postl) != nil {
		return c.JSON(ctx, cfg.H{"msg": "Title Found"}, 404)
	}
	return c.JSON(ctx, cfg.H{"msg": "sukses"}, 200)
}

// Update post controller...
func (c Post) Update(ctx *atreugo.RequestCtx) error {
	key := ctx.UserValue("key")

	var postl database.Post
	json.Unmarshal(ctx.PostBody(), &postl)

	if string(postl.Title) == "" || string(postl.Content) == "" {
		return c.JSON(ctx, cfg.H{"msg": "form harus disi"}, 404)
	}
	postl.ID = key.(string)

	//postl.AuthorID = database.GetSession(ctx, "ID").(string)
	if c.p.Update(postl) != nil {
		return c.JSON(ctx, cfg.H{"msg": "Title Found"}, 404)
	}
	return c.JSON(ctx, cfg.H{"msg": "sukses"}, 200)
}

// Delete user.
func (c Post) Delete(ctx *atreugo.RequestCtx) error {
	key := ctx.UserValue("key")
	err := c.p.Delete(key.(string))
	if err != nil {
		return c.JSON(ctx, cfg.H{"msg": "gagal"}, 404)
	}
	return c.JSON(ctx, cfg.H{"msg": "sukses"}, 200)
}
