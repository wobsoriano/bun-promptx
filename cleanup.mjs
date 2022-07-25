import fs from 'fs'
import { suffix } from 'bun:ffi'

const files = fs.readdirSync('./release')

const { platform, arch } = process
let filename

if (platform === 'linux' && arch === 'x64') {
  filename = `promptx-${platform}-amd64.${suffix}`
} else {
  filename = `promptx-${platform}-${arch}.${suffix}`
}

files.forEach((file) => {
  if (file !== filename) {
    fs.unlinkSync(`./release/${file}`)
  }
})
