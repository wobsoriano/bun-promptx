import { ptr } from 'bun:ffi'
import { symbols } from './ffi'
import { encode, toString } from './utils'

export type PromptOptions = {
  charLimit?: number
  required?: boolean
  echoMode?: 'normal' | 'password' | 'none'
}

export type PromptResult = {
  value: string | null
  error: string | null
}

export function createPrompt(prompt: string, options: PromptOptions = {}): PromptResult {
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
