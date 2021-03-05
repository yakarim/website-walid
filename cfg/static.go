package cfg

import (
	"os"
	"strings"

	"github.com/savsgio/atreugo/v11"
)

// Port ...
func (c *Cfg) Port() (atreugo.Config, string) {
	port := os.Getenv("PORT")
	var config atreugo.Config

	if len(port) == 0 {
		port = ":3000"
		config = atreugo.Config{
			Addr:              "0.0.0.0:" + port,
			Name:              "Kreasindo Pratama",
			ReduceMemoryUsage: true,
			Compress:          true,
			//Concurrency:       100,
			//GracefulShutdown: true,
			Debug: true,
		}
	} else if !strings.HasPrefix(":", port) {
		port = ":" + port
		config = atreugo.Config{
			Addr:              "0.0.0.0:" + port,
			Name:              "Kreasindo Pratama",
			ReduceMemoryUsage: true,
			Compress:          true,
			Concurrency:       100,
			//GracefulShutdown:  true,
		}
	}

	return config, port
}

// Static ...
func (c *Cfg) Static(ctx *atreugo.Atreugo) {
	costumStatic(ctx, "css")
	costumStatic(ctx, "js")
	costumStatic(ctx, "images")
	costumStatic(ctx, "wasm")
}

func costumStatic(ctx *atreugo.Atreugo, name string) {
	rootFS := &atreugo.StaticFS{
		Root:               "./static/" + name,
		AcceptByteRange:    true,
		GenerateIndexPages: true,
		//	CacheDuration:      15 * time.Hour,
		//Compress:           true,
	}
	ctx.StaticCustom("/"+name, rootFS)
}
