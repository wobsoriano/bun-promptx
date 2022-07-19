import { createSelection } from "./src";

const result = createSelection([
  { text: "fix", description: "Bug fix" },
  { text: "docs", description: "Writing docs" },
  { text: "style", description: "Improving structure/format of the code" }
])

console.log(result)
