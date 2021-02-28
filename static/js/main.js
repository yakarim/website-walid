const go = new Go(); // Defined in wasm_exec.js
const WASM_URL = '/wasm/html.wasm';

var wasm;

if ('instantiateStreaming' in WebAssembly) {
	WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(function (obj) {
		wasm = obj.instance;
		go.run(wasm);
		async function MyFunc() {
			try {
				const response = await MyGoFunc('')
				const message = await response.json()
				console.log(message)
			} catch (err) {
				console.error('Caught exception', err)
			}
		}
	})
} else {
	fetch(WASM_URL).then(resp =>
		resp.arrayBuffer()
	).then(bytes =>
		WebAssembly.instantiate(bytes, go.importObject).then(function (obj) {
			wasm = obj.instance;
			go.run(wasm);
		})
	)
}

