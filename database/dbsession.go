package database

import (
	"fmt"
	"log"

	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
	psgsession "github.com/yakarim/psgsession"
	"github.com/yakarim/psgsession/providers/postgre"
)

// Session ...
var Session *psgsession.Session

func init() {

	encoder := psgsession.Base64Encode
	decoder := psgsession.Base64Decode

	var provider psgsession.Provider
	var err error

	//portln := os.Getenv("PORT")
	//if len(portln) != 0 {
	//	cfg := postgre.NewConfigWith("127.0.0.1", 5432, "postgres", "1234", "walid", "session")
	//	provider, err = postgre.New(cfg)

	//} else if !strings.HasPrefix(":", portln) {
	cfg := postgre.NewConfigWith("ec2-34-235-240-133.compute-1.amazonaws.com", 5432, "qrkpjazlukiadp", "129bc7576d446c3e85369edbd5563edd18d9be2f44521e89dca5956bd5ee0df0", "dfingkdn5kgjg", "session")
	provider, err = postgre.New(cfg)
	fmt.Println(err)
	//}

	//provider, err = memory.New(memory.Config{})
	cfg1 := psgsession.NewDefaultConfig()
	cfg1.EncodeFunc = encoder
	cfg1.DecodeFunc = decoder
	Session = psgsession.New(cfg1)

	if err = Session.SetProvider(provider); err != nil {
		log.Fatal(err)
	}

}

// GetSession ...
func GetSession(ctx *atreugo.RequestCtx, str string) interface{} {
	store, err := Session.Get(ctx.RequestCtx)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
	return store.Get(str)
}

// SessionAuth ...
func SessionAuth(ctx *atreugo.RequestCtx) bool {
	store, err := Session.Get(ctx.RequestCtx)
	if err != nil {
		return false
	}
	id := store.Get("ID")
	if id != nil {
		return true
	}
	return false
}
