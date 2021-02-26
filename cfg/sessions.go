package cfg

import (
	"os"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/kataras/go-sessions/v3"
	"github.com/savsgio/atreugo/v11"
)

// Session ...
var Session *sessions.Sessions

// Session ...
func init() {

	cookieName := os.Getenv("SESSION_SECRET")
	// AES only supports key sizes of 16, 24 or 32 bytes.
	// You either need to provide exactly that amount or you derive the key from what you type in.
	hashKey := []byte("the-big-and-secret-fash-key-here")
	blockKey := []byte("lot-secret-of-characters-big-too")
	secureCookie := securecookie.New(hashKey, blockKey)

	mySessions := sessions.New(sessions.Config{
		Cookie:  cookieName,
		Encode:  secureCookie.Encode,
		Decode:  secureCookie.Decode,
		Expires: 24 * 30 * time.Hour,
	})

	Session = mySessions
}

// SessionUsername ...
func (c Cfg) SessionUsername(ctx *atreugo.RequestCtx) interface{} {
	session := Session.StartFasthttp(ctx.RequestCtx)
	if session != nil {
		return session.Get("Username")
	}
	return nil
}

// SessionID ...
func (c Cfg) SessionID(ctx *atreugo.RequestCtx) bool {
	session := Session.StartFasthttp(ctx.RequestCtx)
	if session.Get("ID") != nil {
		return true
	}
	return false
}

// SessionDestroy ...
func (c Cfg) SessionDestroy(ctx *atreugo.RequestCtx) error {
	session := Session.StartFasthttp(ctx.RequestCtx)
	session.Destroy()
	return nil
}
