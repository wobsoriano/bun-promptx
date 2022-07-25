import fs from 'fs'
import { suffix } from 'bun:ffi'

const path = new URL('./release', import.meta.url).pathname
const files = fs.readdirSync(path)

const { platform, arch } = process
let filename

if (arch === 'x64') {
  filename = `promptx-${platform}-amd64.${suffix}`
} else {
  filename = `promptx-${platform}-${arch}.${suffix}`
}

files.forEach((file) => {
  if (file !== filename) {
    fs.unlinkSync(`./release/${file}`)
  }
})
