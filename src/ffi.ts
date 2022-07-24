import { dlopen, FFIType, ptr, suffix } from 'bun:ffi'
import { encode, toString } from './utils'

const { platform, arch } = process

let filename: string

if (platform === 'linux' && arch === 'x64') {
  filename = `../release/promptx-${platform}-amd64.${suffix}`
} else {
  filename = `../release/promptx-${platform}-${arch}.${suffix}`
}

const location = new URL(filename, import.meta.url).pathname
export const { symbols } = dlopen(location, {
  CreateSelection: {
    args: [FFIType.ptr, FFIType.ptr, FFIType.ptr, FFIType.int],
    returns: FFIType.ptr
  },
  CreatePrompt: {
    args: [FFIType.ptr, FFIType.ptr, FFIType.bool, FFIType.int],
    returns: FFIType.ptr
  },
  FreeString: {
    args: [FFIType.ptr],
    returns: FFIType.void
  }
})

export type SelectionItem = {
  text: string
  description?: string
}

export type SelectionOptions = {
  perPage?: number
  headerText?: string
  footerText?: string
}

export type SelectionReturn = {
  selectedIndex: number | null
  error: string | null
}

export function createSelection(items: SelectionItem[], options: SelectionOptions = {}): SelectionReturn {
  const stringifiedItems = JSON.stringify(items.map((item) => {
    return {
      text: item.text,
      description: item.description || ''
    }
  }))
  const returnedPtr = symbols.CreateSelection(
    ptr(encode(stringifiedItems)),
    ptr(encode(options.headerText || 'Select an item: ')),
    ptr(encode(options.footerText || '')),
    options.perPage || 5
  )
  const { selectedIndex, error } = JSON.parse(toString(returnedPtr)) as {
    selectedIndex: string
    error: string
  }
  if (error !== "") {
    return {
      selectedIndex: null,
      error
    }
  }
  return {
    selectedIndex: Number(selectedIndex),
    error: null
  }
}

export type PromptOptions = {
  charLimit?: number
  required?: boolean
  echoMode?: 'normal' | 'password' | 'none'
}

export function createPrompt(prompt: string, options: PromptOptions = {}) {
  const returnedPtr = symbols.CreatePrompt(
    ptr(encode(prompt)),
    ptr(encode(options.echoMode || 'normal')),
    options.required ?? true,
    options.charLimit || 0
  )
  const { value, error } = JSON.parse(toString(returnedPtr)) as {
    value: string
    error: string
  }
  if (error !== "") {
    return {
      value: null,
      error
    }
  }
  return {
    value,
    error: null
  }
}
