package controller

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/yakarim/website-walid/cfg"
	"github.com/yakarim/website-walid/database"
	"github.com/yakarim/website-walid/model"
)

// User ...
type User struct {
	model.Mdl
	cfg.Cfg
	H cfg.H
}

// QueryOne User...
func (c *User) QueryOne(ctx *atreugo.RequestCtx) error {
	key := ctx.UserValue("key")
	data, _ := c.UserM.QueryOne(key.(string))
	return c.JSON(ctx, cfg.H{"user": data}, 200)
}

// QueryAll User...
func (c *User) QueryAll(ctx *atreugo.RequestCtx) error {
	data, _ := c.UserM.Query()
	return c.JSON(ctx, cfg.H{"users": data}, 200)
}

// Create User ...
func (c *User) Create(ctx *atreugo.RequestCtx) error {
	var userl database.User
	json.Unmarshal(ctx.PostBody(), &userl)
	if string(userl.Username) == "" || string(userl.Email) == "" || string(userl.Password) == "" {
		c.JSON(ctx, cfg.H{"msg": "form harus disi"}, 404)
	} else {
		if c.UserM.Create(userl) != nil {
			c.JSON(ctx, cfg.H{"msg": "Email Found"}, 404)
		} else {
			c.JSON(ctx, cfg.H{"msg": "sukses"}, 200)
		}
	}
	return nil
}

// Update user ...
func (c *User) Update(ctx *atreugo.RequestCtx) error {
	var user database.User
	json.Unmarshal(ctx.PostBody(), &user)
	err := c.UserM.Update(user)
	if c.UserM.Update(user) != nil {
		c.JSON(ctx, cfg.H{"msg": err.Error()}, 404)
	} else {
		c.JSON(ctx, cfg.H{"msg": "sukses"}, 200)
	}
	return nil
}

// Delete user.
func (c *User) Delete(ctx *atreugo.RequestCtx) error {
	key := ctx.UserValue("key")
	err := c.UserM.Delete(key.(string))
	if err != nil {
		c.JSON(ctx, cfg.H{"msg": err.Error()}, 404)
	}
	return c.JSON(ctx, cfg.H{"msg": "sukses"}, 200)
}
