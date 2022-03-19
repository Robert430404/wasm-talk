// Setup event name
const eventName = 'wasm:module:ready';
const threshold = 1000000;

// Setup the listener
document.addEventListener(eventName, () => {
  // Add listener for native js implementation
  const nativeJSButton = document.querySelector('.Demo__javascriptButton');
  const nativeJSOutput = document.querySelector('.Demo__javascriptOutput');

  nativeJSButton.addEventListener('click', (event) => {
    computed = 0;

    for (let i = 0; i < threshold; i += 1) {
      computed = Math.ceil(Math.random() * 3617296361323339000);
    }

    nativeJSOutput.value = computed;
  });

  // Add listener for wasm implementation
  const wasmButton = document.querySelector('.Demo__webassemblyButton');

  wasmButton.addEventListener('click', (event) => {
    generateNumber(threshold);
  });
});

// Load in our web assembly module
const go = new Go();

WebAssembly.instantiateStreaming(
  fetch('./build/print.wasm'),
  go.importObject,
).then((result) => {
  go.run(result.instance);

  document.dispatchEvent(new Event(eventName));
});
