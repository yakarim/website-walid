package main

import (
	"net"
	"testing"

	"github.com/CloudyKit/jet/v3"
	"github.com/savsgio/go-logger"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
	"github.com/yakarim/website-walid/jade"
)

var ()

func serve(handler fasthttp.RequestHandler, req *fasthttp.Request, res *fasthttp.Response) error {
	ln := fasthttputil.NewInmemoryListener()
	defer ln.Close()

	go func() {
		err := fasthttp.Serve(ln, handler)
		if err != nil {
			logger.Error(err)
			panic(err)
		}
	}()

	client := fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return ln.Dial()
		},
	}

	return client.Do(req, res)
}

func TestListAllSubscription(t *testing.T) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("/uri") // task URI
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()
	err := serve(handler, req, resp)
	if err != nil {

	}
	logger.Info("resp from Postman.Post: ", resp)
	logger.Info("resp status code", resp.StatusCode())
	logger.Info("resp body", string(resp.Body()))
}
func handler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "text/html; charset=UTF-8")
	jade.Index("home", "string", true, ctx)
	ctx.SetStatusCode(fasthttp.StatusOK)
}

var jetSet = jet.NewHTMLSet("./template")

func handlerjet(ctx *fasthttp.RequestCtx) {
	tmpl, _ := jetSet.GetTemplate("pages/index.html")
	err := tmpl.Execute(ctx, nil, nil)
	if err != nil {
	}
}
