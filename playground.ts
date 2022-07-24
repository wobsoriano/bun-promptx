import { createSelection, createPrompt } from './src'

const username = createPrompt("Enter username: ")
console.log(username.value)

const password = createPrompt("Enter password: ", {
  echoMode: 'password'
})
console.log(password.value)

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
