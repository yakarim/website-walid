import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { store } from './stores'
import VueFinalModal from 'vue-final-modal'
import 'bootstrap'
import axios from 'axios'
import Go from './wasm_exec'


const app = createApp(App)
/* eslint no-undef: "off"*/
app.config.globalProperties.$http = () => {
    axios
}
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
     app.provide('$go', { 'satu': 'dia', 'add': change })

});


app.use(router)
app.use(store)
app.use(VueFinalModal())
app.mount('#app')