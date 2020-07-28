const { performance } = require('perf_hooks');
let t0 = performance.now();
for (let i = 0; i < 10000000; i++) {
}
let t1 = performance.now();

console.log(t1 - t0);