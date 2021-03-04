package cfg

import "github.com/savsgio/atreugo/v11"

// Username ...
func (c Cfg) Username(ctx *atreugo.RequestCtx) string {
	store, err := session.Get(ctx.RequestCtx)
	if err != nil {
		return ""
	}
	str := store.Get("Username")
	if str != nil {
		return str.(string)
	}
	return ""
}

// SessionAuth ...
func (c Cfg) SessionAuth(ctx *atreugo.RequestCtx) bool {
	store, err := session.Get(ctx.RequestCtx)
	if err != nil {
		return false
	}
	id := store.Get("ID")
	if id != nil {
		return true
	}
	return false
}
