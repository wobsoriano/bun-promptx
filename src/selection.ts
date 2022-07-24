import { ptr } from 'bun:ffi'
import { symbols } from './ffi'
import { encode, toString } from './utils'

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
