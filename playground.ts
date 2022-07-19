import { createSelection } from "./src";

const result = createSelection([
  { text: "feat", description: "Introducing new features" },
  { text: "fix", description: "Bug fix" },
  { text: "docs", description: "Writing docs" },
  { text: "style", description: "Improving structure/format of the code" },
  { text: "refactor", description: "Refactoring code" }
], {
  headerText: 'Select Commit Type: '
})

console.log(result)
