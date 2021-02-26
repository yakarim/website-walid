/* eslint no-unused-vars: "off"*/
if (!WebAssembly.instantiateStreaming) { // polyfill
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
    };
}
const go = new Go();
let mod, inst;
WebAssembly.instantiateStreaming(fetch("./wasm/calc.wasm"), go.importObject).then((result) => {
    mod = result.module;
    inst = result.instance;
    go.run(inst);
    WebAssembly.instantiate(mod, go.importObject); 
    async function change(val, val2) {
        return waAdd(...Array(val, val2))
     }// reset instance
});