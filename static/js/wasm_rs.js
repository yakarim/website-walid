WebAssembly.instantiateStreaming(fetch("/wasm/wasm_bg.wasm"))
    .then(wasmModule => {
        // this saves the exported function from WASM module for use in JS
        wasm_add = wasmModule.instance.exports.add;
    });