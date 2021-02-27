'use strict';

const WASM_URL = '/wasm/button.wasm';
var wasm;

function init() {
    const go = new Go();

    if ('instantiateStreaming' in WebAssembly) {
        WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(function (obj) {
            wasm = obj.instance;
            go.run(wasm);
        })
    }
}

init();

console.log(wasm.exports.add(a, b));