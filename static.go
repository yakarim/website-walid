package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path"
	"strings"
	"time"

	"github.com/savsgio/atreugo/v11"
)

// FilesServer ...
func FilesServer(ctx *atreugo.Atreugo) {
	//h := http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "dist", Fallback: "index.html"})
	h := http.FileServer(http.Dir("./public/dist"))
	ctx.NetHTTPPath("GET", "/", h)
	ctx.NetHTTPPath("GET", "/about", cached("25s", h))
	ctx.NetHTTPPath("GET", "/admin", cached("25s", h))

	ctx.NetHTTPPath("GET", "/login", cached("25s", h))
	ctx.NetHTTPPath("GET", "/signup", cached("25s", h))
	ctx.NetHTTPPath("GET", "/logout", cached("25s", h))
}

// StaticPath ..
func StaticPath(ctx *atreugo.Atreugo) {
	for _, v := range GzipAssetNames() {
		static(ctx, "css", v)
		static(ctx, "js", v)
		static(ctx, "wasm", v)
	}
}

func static(ctx *atreugo.Atreugo, str, v string) {
	if path.Ext(v) == "."+str {
		css := strings.TrimPrefix(v, "../static/"+str)
		ctx.GET("/"+str+css, func(ctx *atreugo.RequestCtx) error {

			data, err := GzipAsset("../static/" + str + css)
			if err != nil {
				println(err)
			}
			if path.Ext(v) == ".img" {
				key := "black"
				e := `"` + key + `"`

				if match := ctx.Request.Header.Peek("If-None-Match"); string(match) != "" {
					if strings.Contains(string(match), e) {
						ctx.SetStatusCode(http.StatusNotModified)
					}
				}
				ctx.Response.Header.Set("Etag", e)

			}

			ctx.Response.Header.Set("Vary", "Accept-Encoding")
			ctx.Response.Header.Set("Content-Type", "text/"+str)
			ctx.Response.Header.Set("Content-Encoding", "gzip")
			ctx.Response.Header.Set("Cache-Control", "max-age=2592000")

			ctx.SetBody(data)
			return nil
		})
	}
}
func baseName(s string) string {
	n := strings.LastIndexByte(s, '.')
	if n == -1 {
		return s
	}
	return s[:n]
}

func cached(duration string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content := storage.Get(r.RequestURI)
		if content != nil {
			fmt.Print("Cache Hit!\n")
			w.Write(content)
		} else {
			c := httptest.NewRecorder()
			handler.ServeHTTP(c, r)
			for k, v := range c.HeaderMap {
				w.Header()[k] = v
			}
			w.WriteHeader(c.Code)
			content := c.Body.Bytes()
			if d, err := time.ParseDuration(duration); err == nil {
				fmt.Printf("New page cached: %s for %s\n", r.RequestURI, duration)
				storage.Set(r.RequestURI, content, d)
			} else {
				fmt.Printf("Page not cached. err: %s\n", err)
			}
			w.Write(content)
		}
	})
}

// Index ..
func Index(ctx *atreugo.RequestCtx) error {
	//wasm()
	return ctx.JSONResponse("wasm", 200)
}

/*
func wasm() {
	wasmBytes, _ := ioutil.ReadFile("calc.wasm")

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, _ := wasmer.NewModule(store, wasmBytes)

	// Instantiates the module
	importObject := wasmer.NewImportObject()
	instance, _ := wasmer.NewInstance(module, importObject)

	// Gets the `sum` exported function from the WebAssembly instance.
	sum, _ := instance.Exports.GetFunction("sum")

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, _ := sum(5, 37)

	fmt.Println(result) // 42!
	//return result
}
*/
