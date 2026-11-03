# [rolldown](https://rolldown.rs/)

Blazing Fast Rust-based `bundler` for `JavaScript` with Rollup-compatible API and esbuild feature parity.

## Installation

```bash
npm install -D rolldown
```

## Using the CLI

```bash
# ./node_modules/.bin/rolldown --version
npx rolldown --version
npx rolldown --help
```

## Example Code

- src/main.js

```js
import { hello } from './hello.js';

hello();
```

- src/hello.js

```js
export function hello() {
  console.log('Hello Rolldown!');
}
```

## Bundle

- cli

```bash
npx rolldown src/main.js --file bundle.js
node bundle.js
```

- package.json

```json
{
  "scripts": {
    "build": "rolldown src/main.js --file bundle.js"
  }
}
```

```bash
npm run build
```

## Using the config file

- rolldown.config.js

```js
import { defineConfig } from 'rolldown';

export default defineConfig({
  input: 'src/main.js',
  output: {
    file: 'bundle.js',
  },
});
```

```json
{
  "scripts": {
    "build": "rolldown -c" // --config
  }
}
```

## Multiple builds in the same config

```js
import { defineConfig } from 'rolldown';

export default defineConfig([
  {
    input: 'src/main.js',
    output: {
      format: 'esm',
    },
  },
  {
    input: 'src/worker.js',
    output: {
      format: 'iife',
      dir: 'dist/worker',
    },
  },
]);
```
