package main

/*
import (
	"testing"

	"github.com/atreugo/httptest"
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
)

func Test_View(t *testing.T) {
	t.Parallel()

	body := "TEST"

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.SetRequestURI("/test")

	handler := func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse(body)
	}

	httptest.View(t, req, handler, func(resp *fasthttp.Response) {
		if string(resp.Body()) != body {
			t.Errorf("Response.Body == %s, want %s", resp.Body(), body)
		}
	})
}
*/
