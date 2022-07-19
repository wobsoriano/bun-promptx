# promptx

promptx is a terminal prompt library based on [bubbles](https://github.com/mritd/bubbles) via `bun:ffi`.

## Install

```bash
bun add promptx
```

## Usage

```js
import { createSelection } from 'promptx'

const result = createSelection([
  { text: 'feat', description: 'Introducing new features' },
  { text: 'fix', description: 'Bug fix' },
  { text: 'docs', description: 'Writing docs' },
  { text: 'style', description: 'Improving structure/format of the code' },
  { text: 'refactor', description: 'Refactoring code' }
], {
  headerText: 'Select Commit Type: ',
  perPage: 5
})

console.log(result)
// { selectedIndex: 2, error: null }
```

## License

MIT
