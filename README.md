# bun-promptx

bun-promptx is a terminal prompt library based on [bubbles](https://github.com/mritd/bubbles) via `bun:ffi`.

## Install

```bash
bun add bun-promptx
```

## Usage

### createSelection

The `createSelection` function lets you create a terminal single-selection list prompt. It provides the functions of page up and down and key movement, and supports custom rendering methods.

```js
import { createSelection } from 'bun-promptx'

const result = createSelection([
  { text: 'feat', description: 'Introducing new features' },
  { text: 'fix', description: 'Bug fix' },
  { text: 'docs', description: 'Writing docs' },
  { text: 'style', description: 'Improving structure/format of the code' },
  { text: 'refactor', description: 'Refactoring code' },
  { text: 'test', description: 'Refactoring code' },
  { text: 'chore', description: 'When adding missing tests' },
  { text: 'perf', description: 'Improving performance' }
], {
  headerText: 'Select Commit Type: ',
  perPage: 5,
  footerText: 'Footer here'
})

console.log(result)
// { selectedIndex: 2, error: null }
```

<img src="https://i.imgur.com/yE0qKyA.gif" alt="promptx demo" width="500" />

### createPrompt

The `createPrompt` function is a terminal input prompt library. It provides CJK character support and standard terminal shortcut keys (such as ctrl+a, ctrl+e), password input echo and other functions.

```js
import { createPrompt } from 'bun-promptx'

const username = createPrompt("Enter username: ")
// { value: "wobsoriano", error: null }

const password = createPrompt("Enter password: ", {
  echoMode: 'password'
})
// { value: "123456", error: null }
```

<img src="https://i.imgur.com/wx6BTUm.gif" alt="promptx demo" width="500" />

## License

MIT
