package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/savsgio/atreugo/v11"
	"github.com/yakarim/website-walid/cfg"
	"github.com/yakarim/website-walid/database"
)

// Index ...
func (c Ctrl) Index(ctx *atreugo.RequestCtx) error {
	return c.HTML(ctx, 200, "pages/index", cfg.H{
		"title":  "Home walid",
		"user":   database.GetSession(ctx, "Username"),
		"signIn": database.SessionAuth(ctx),
		"data":   call("https://jsonplaceholder.typicode.com/posts"),
	})
}

func call(url string) interface{} {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err = res.Body.Close()
	if err != nil {
		panic(err)
	}
	str, err := prettyJSON(body)
	if err != nil {
		panic(err)
	}
	return string(str)
}

func prettyJSON(input []byte) (string, error) {
	var raw interface{}
	if err := json.Unmarshal(input, &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}
