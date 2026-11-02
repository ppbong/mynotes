# electron-vite

The Build Tool for Electron based on Vite. 

## Install

```bash
# equires Node.js version 20.19+, 22.12+ and Vite version 5.0+
npm i electron-vite -D
```

## Command line interface

- npx

```bash
npx electron-vite
# show cli options like '--outDir'
npx electron-vite -h
```

- package.json / scripts

```json
{
  "scripts": {
    // start electron app to preview production build
    "start": "electron-vite preview",
    // start dev server and electron app
    "dev": "electron-vite dev",
    // build for production
    "prebuild": "electron-vite build"
  }
}
```

## Conigure

When running `electron-vite` will automatically try to resolve a config file named `electron.vite.config.js` inside project root.

```js
export default {
  main: {
    // vite config options
  },
  preload: {
    // vite config options
  },
  renderer: {
    // vite config options
  }
}
```

## Electron entry point

- default: --outDir = out
- package.json

```json
{
  "name": "electron-app",
  "version": "1.0.0",
  // electron main process entry file
  "main": "./out/main/index.js"
}
```

## Scaffolding electron-vite project

- Project name: `electron-app`
- Select a framework: `vue` / `vue-ts` / `react` / `react-ts` / ..
- Add TypeScript: `No` / `Yes`
- Add Electron updater plugin: `No` / `Yes`
- Enable Electron download mirror proxy: `No` / `Yes`

```bash
npm create @quick-start/electron@latest
npm create @quick-start/electron@latest my-app -- --template vue
```

## Electron mirror

- .npmrc

```bash
electron_mirror=https://npmmirror.com/mirrors/electron/
electron_builder_binaries_mirror=https://npmmirror.com/mirrors/electron-builder-binaries/
```

## Project structure

### Conventions

It is recommended to use the following project structure:

```text
.
├──src
│  ├──main
│  │  ├──index.ts
│  │  └──...
│  ├──preload
│  │  ├──index.ts
│  │  └──...
│  └──renderer    # with vue, react, etc.
│     ├──src
│     ├──index.html
│     └──...
├──electron.vite.config.ts
├──package.json
└──...
```

With this convention, electron-vite can work with minimal configuration, and it will automatically find the default entry points：

- Main process: `<root>/src/main/{index|main}.{js|ts|mjs|cjs}`
- Preload script: `<root>/src/preload/{index|preload}.{js|ts|mjs|cjs}`
- Renderer: `<root>/src/renderer/index.html`

### Customizing

It is recommended to use the following project structure:

```text
.
├──electron
│  ├──main
│  │  ├──index.ts
│  │  └──...
│  └──preload
│     ├──index.ts
│     └──...
├──src   # with vue, react, etc.
├──index.html
├──electron.vite.config.ts
├──package.json
└──...
```

The `electron.vite.config.ts` should be:

```ts
import { defineConfig } from 'electron-vite'
import { resolve } from 'path'

export default defineConfig({
  main: {
    build: {
      rollupOptions: {
        input: {
          index: resolve(__dirname, 'electron/main/index.ts')
        }
      }
    }
  },
  preload: {
    build: {
      rollupOptions: {
        input: {
          index: resolve(__dirname, 'electron/preload/index.ts')
        }
      }
    }
  },
  renderer: {
    root: '.',
    build: {
      rollupOptions: {
        input: {
          index: resolve(__dirname, 'index.html')
        }
      }
    }
  }
})
```

## Using preload scripts

### Example

- preload.js

```js
import { contextBridge, ipcRenderer } from 'electron'

contextBridge.exposeInMainWorld('electron', {
  ping: () => ipcRenderer.invoke('ping')
})
```

- main.js

```js
import { ipcMain } from 'electron'

ipcMain.handle('ping', () => 'pong')
```

- renderer.js

```js
const func = async () => {
  const response = await window.electron.ping()
  console.log(response) // pong
}

func()
```

### Limitations of sandboxing

From Electron 20 onwards, preload scripts are sandboxed by default and no longer have access to a full Node.js environment. 

In sandbox mode, if you `require` or `import` other modules in a preload script, you may encounter the following error:

```text
Unable to load preload scripts -> Error: module not found: 'foo'
```

To handle dependencies correctly, you have two options:

- Disable sandbox

```js
const win = new BrowserWindow({
  webPreferences: {
    sandbox: false
  }
})
```

- Fully bundle dependencies with electron-vite

### Toolkit

For development efficiency, it is recommended to use `@electron-toolkit/preload`. It provides an easy way to expose Electron APIs (`ipcRenderer`, `webFrame`, `process`) to the renderer process.

- preload.js

```js
import { contextBridge } from 'electron'
import { electronAPI } from '@electron-toolkit/preload'

if (process.contextIsolated) {
  try {
    contextBridge.exposeInMainWorld('electron', electronAPI)
  } catch (error) {
    console.error(error)
  }
} else {
  window.electron = electronAPI
}
```

- renderer.js

```js
// Send a message to the main process with no response
window.electron.ipcRenderer.send('electron:say', 'hello')

// Send a message to the main process with the response asynchronously
window.electron.ipcRenderer.invoke('electron:doSomeThing', '').then(res => {
  console.log(res)
})

// Receive messages from the main process
window.electron.ipcRenderer.on('electron:reply', (_, args) => {
  console.log(args)
})
```

### Webview

```js
mainWindow.webContents.on('will-attach-webview', (e, webPreferences) => {
  webPreferences.preload = join(import.meta.dirname, '../preload/index.js')
})
```

## Multiple Windows App

- electron.vite.config.js

```js
export default {
  main: {},
  preload: {
    build: {
      rollupOptions: {
        input: {
          browser: resolve(__dirname, 'src/preload/browser.js'),
          webview: resolve(__dirname, 'src/preload/webview.js')
        }
      }
    }
  },
  renderer: {
    build: {
      rollupOptions: {
        input: {
          browser: resolve(__dirname, 'src/renderer/browser.html'),
          webview: resolve(__dirname, 'src/renderer/webview.html')
        }
      }
    }
  }
}
```

## ESM Support in Electron

- package.json

```json
{
    "type": "module"
}
```

- electron.vite.config.js

```js
export default defineConfig({
  main: {
    build: {
      rollupOptions: {
        output: {
          format: 'es' // end with the .mjs extension
        }
      }
    }
  },
  preload: {
    build: {
      rollupOptions: {
        output: {
          format: 'es' // end with the .mjs extension
        }
      }
    }
  }
  // ...
})
```

- main.js

```js
import { fileURLToPath } from 'url'
import path from 'path'

const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)

const preload = path.join(__dirname, '../preload/index.mjs')
preload = fileURLToPath(new URL('../preload/index.mjs', import.meta.url))
```

## Public directory

- main process and preload scripts: `<root>/resources`
  - main.publicDir
  - preload.publicDir
- renderer process: `src/renderer/public`
  - renderer.publicDir

## TypeScript

- *.d.ts

```ts
/// <reference types="electron-vite/node" />
```

- tsconfig.node.json

```json
{
  "compilerOptions": {
    "types": ["electron-vite/node"]
  }
}
```

## Building for production

- By default output dir is `<ROOT>/out`:

```text
├──out/
│  ├──main
│  ├──preload
│  └──renderer
└──...
```

- specify it via a command line flag `--outDir`:

```bash
npx electron-vite dev --outDir=dist
npx electron-vite build --outDir=dist
npx electron-vite preview --outDir=dist
```

- use `build.outDir` option in `electron.vite.config.js`:

```js
export default defineConfig({
  main: {
    build: {
      outDir: 'dist/main'
    }
  },
  preload: {
    build: {
      outDir: 'dist/preload'
    }
  },
  renderer: {
    build: {
      outDir: 'dist/renderer'
    }
  }
})
```

## Chunking Strategy

- electron.vite.config.ts

```js
import { defineConfig } from 'electron-vite'

export default defineConfig({
  main: {
    build: {
      minify: true,
      rollupOptions: {
        output: {
          manualChunks(id): string | void {
            if (id.includes('foo')) {
              return 'foo'
            }
          }
        }
      }
    }
  },
  // ...
})
```

## Distribution

### Distributing app with electron builder

- electron-builder.yml

```yml
appId: com.electron.app
productName: electron-app
directories:
  buildResources: build
files:
  - '!**/.vscode/*'
  - '!src/*'
  - '!electron.vite.config.{js,ts,mjs,cjs}'
  - '!{.eslintignore,.eslintrc.cjs,.prettierignore,.prettierrc.yaml,dev-app-update.yml,CHANGELOG.md,README.md}'
  - '!{.env,.env.*,.npmrc,pnpm-lock.yaml}'
  - '!{tsconfig.json,tsconfig.node.json,tsconfig.web.json}'
asarUnpack:
  - resources/**
afterSign: build/notarize.js
win:
  executableName: electron-app
nsis:
  artifactName: ${name}-${version}-setup.${ext}
  shortcutName: ${productName}
  uninstallDisplayName: ${productName}
  createDesktopShortcut: always
mac:
  entitlementsInherit: build/entitlements.mac.plist
  extendInfo:
    - NSCameraUsageDescription: Application requests access to the device's camera.
    - NSMicrophoneUsageDescription: Application requests access to the device's microphone.
    - NSDocumentsFolderUsageDescription: Application requests access to the user's Documents folder.
    - NSDownloadsFolderUsageDescription: Application requests access to the user's Downloads folder.
dmg:
  artifactName: ${name}-${version}.${ext}
linux:
  target:
    - AppImage
    - snap
    - deb
  maintainer: electronjs.org
  category: Utility
appImage:
  artifactName: ${name}-${version}.${ext}
npmRebuild: false
publish:
  provider: generic
  url: https://example.com/auto-updates
```

- package.json

```json
{
  "scripts": {
    "build": "electron-vite build",
    "build:win": "npm run build && electron-builder --win --config",
    "build:mac": "npm run build && electron-builder --mac --config",
    "build:linux": "npm run build && electron-builder --linux --config"
  }
}
```

### Github Action CI/CD

- .github/workflows/release.yml

```yml
name: Build/release Electron app

on:
  push:
    tags:
      - v*.*.*

jobs:
  release:
    runs-on: ${{ matrix.os }}

    strategy:
      matrix:
        os: [ubuntu-latest, macos-latest, windows-latest]

    steps:
      - name: Check out Git repository
        uses: actions/checkout@v3

      - name: Install Node.js
        uses: actions/setup-node@v3
        with:
          node-version: 16

      - name: Install Dependencies
        run: npm install

      - name: build-linux
        if: matrix.os == 'ubuntu-latest'
        run: npm run build:linux

      - name: build-mac
        if: matrix.os == 'macos-latest'
        run: npm run build:mac

      - name: build-win
        if: matrix.os == 'windows-latest'
        run: npm run build:win

      - name: release
        uses: softprops/action-gh-release@v1
        with:
          draft: true
          files: |
            dist/*.exe
            dist/*.zip
            dist/*.dmg
            dist/*.AppImage
            dist/*.snap
            dist/*.deb
            dist/*.rpm
            dist/*.tar.gz
            dist/*.yml
            dist/*.blockmap
        env:
          GITHUB_TOKEN: ${{ secrets.ACCESS_TOKEN }}
```

## Global env variables

By default, variables prefixed with `MAIN_VITE_` are exposed to the main process, `PRELOAD_VITE_` to preload scripts, `RENDERER_VITE_` to renderers and `VITE_` to all.

```ini
# .env
KEY=123                # not available
MAIN_VITE_KEY=123      # only main process available
PRELOAD_VITE_KEY=123   # only preload scripts available
RENDERER_VITE_KEY=123  # only renderers available
VITE_KEY=123           # all available
```

## Debugging

electron-vite supports debugging both the main process and the renderer process code.

- .vscode/launch.json

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Debug Main Process",
      "type": "node",
      "request": "launch",
      "cwd": "${workspaceRoot}",
      "runtimeExecutable": "${workspaceRoot}/node_modules/.bin/electron-vite",
      "windows": {
        "runtimeExecutable": "${workspaceRoot}/node_modules/.bin/electron-vite.cmd"
      },
      "runtimeArgs": ["--sourcemap"],
      "env": {
        "REMOTE_DEBUGGING_PORT": "9222"
      }
    },
    {
      "name": "Debug Renderer Process",
      "port": 9222,
      "request": "attach",
      "type": "chrome",
      "webRoot": "${workspaceFolder}/src/renderer",
      "timeout": 60000,
      "presentation": {
        "hidden": true
      }
    }
  ],
  "compounds": [
    {
      "name": "Debug All",
      "configurations": ["Debug Main Process", "Debug Renderer Process"],
      "presentation": {
        "order": 1
      }
    }
  ]
}
```

## Troubleshooting

### `vue-router` or `react-router-dom` works in development but not production

Electron does not manage browser history and instead relies on a synchronized URL. Therefore, only a hash-based router will work properly in production.

- For `vue-router`, use `createWebHashHistory` instead of `createWebHistory`.
- For `react-router-dom`, use `HashRouter` instead of `BrowserRouter`.

When using a hash router, you can specify the hash value in the second argument of `BrowserWindow.loadFile` to load the corresponding page.

```js
win.loadFile(path.join(import.meta.dirname, '../renderer/index.html'), { hash: 'home' })
```
