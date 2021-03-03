package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/atreugo/cors"
	"github.com/joho/godotenv"
	"github.com/savsgio/atreugo/v11"
	fastprefork "github.com/valyala/fasthttp/prefork"
	"github.com/yakarim/website-walid/cfg"
	"github.com/yakarim/website-walid/cfg/cache"
	"github.com/yakarim/website-walid/controller"
	"github.com/yakarim/website-walid/templates"
)

var (
	c       cfg.Cfg
	ctrl    controller.Ctrl
	storage cache.Storage
	t       templates.Tmpl
)

func init() {
	var err error

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	environmentPath := filepath.Join(dir, ".env")

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
		err = godotenv.Load(environmentPath)
		if err != nil {
			log.Fatalf("Error getting env, %v", err)
		}
	}
	storage = cache.NewStorage()
}

func main() {
	config, port := c.Port()
	ctx := atreugo.New(config)

	cors := cors.New(cors.Config{
		AllowedOrigins: []string{"http://localhost:3000/", "http://website-walid.herokuapp.com/", "https://website-walid.herokuapp.com/"},
		//	AllowedHeaders:   []string{"Content-Type", "X-Custom"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		ExposedHeaders:   []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		AllowMaxAge:      5600,
	})

	ctx.UseBefore(cors)
	//StaticPath(ctx)
	c.Static(ctx)
	ctrl.Router(ctx)
	t.RouteQuick(ctx)
	preforkServer := &fastprefork.Prefork{
		RecoverThreshold: runtime.GOMAXPROCS(0) / 4,
		ServeFunc:        ctx.Serve,
	}

	if err := preforkServer.ListenAndServe(port); err != nil {
		panic(err)
	}
}
