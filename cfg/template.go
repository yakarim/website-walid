package cfg

import (
	"log"
	"os"
	"time"

	"github.com/CloudyKit/jet/v3"
	"github.com/savsgio/atreugo/v11"
)

// HTML ...
func (c *Cfg) HTML(ctx *atreugo.RequestCtx, code int, page string, data H) error {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		views.SetDevelopmentMode(true)
	} else {
		views.SetDevelopmentMode(false)
	}
	t, vars := c.template(ctx, code, page)
	return t.Execute(ctx.RequestCtx, vars, data)
}

func (c *Cfg) template(ctx *atreugo.RequestCtx, code int, page string) (*jet.Template, jet.VarMap) {
	views.Delims("[%", "%]")
	t, err := views.GetTemplate(page + ".html")
	if err != nil {
		log.Println("Unexpected template err:", err.Error())
	}
	c.globalFunc(views)
	vars := make(jet.VarMap)
	ctx.SetStatusCode(code)
	ctx.Response.Header.Set("Content-Type", "text/html; charset=UTF-8")
	return t, vars
}

func (c *Cfg) globalFunc(view *jet.Set) {
	tm, _ := c.TimeIn(time.Now(), "Local")
	view.AddGlobal("time", tm.Format("01/02/2006"))
}
