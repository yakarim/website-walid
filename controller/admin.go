package controller

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/yakarim/website-walid/cfg"
)

// Admin ...
func (c Ctrl) Admin(ctx *atreugo.RequestCtx) error {

	return c.HTML(ctx, 200, "admin/index", cfg.H{
		"title":  "admin",
		"user":   c.SessionUsername(ctx),
		"signIn": c.SessionID(ctx),
	})
}

/*
// WasmRust ...
func WasmRust() {
	wasmBytes, _ := ioutil.ReadFile("wasm_bg.wasm")

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	module, _ := wasmer.NewModule(store, wasmBytes)

	importObject := wasmer.NewImportObject()
	instance, _ := wasmer.NewInstance(module, importObject)

	sum, _ := instance.Exports.GetFunction("sum")

	result, _ := sum(5, 37)

	fmt.Println(result) // 42!
}
*/
