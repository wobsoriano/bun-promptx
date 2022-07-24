package main

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"unsafe"

	"github.com/wobsoriano/promptx/prompt"
	"github.com/wobsoriano/promptx/selection"
)

func ch(str string) *C.char {
	return C.CString(str)
}

func str(ch *C.char) string {
	return C.GoString(ch)
}

func main() {}

//export FreeString
func FreeString(str *C.char) {
	C.free(unsafe.Pointer(str))
}

//export CreateSelection
func CreateSelection(jsonData, headerText, footerText *C.char, perPage int) *C.char {
	result := selection.Selection(str(jsonData), str(headerText), str(footerText), perPage)
	return ch(result)
}

//export CreatePrompt
func CreatePrompt(prompText, echoMode, validateOkPrefix, validateErrPrefix *C.char, required bool, charLimit int) *C.char {
	result := prompt.Prompt(str(prompText), str(echoMode), str(validateOkPrefix), str(validateErrPrefix), required, charLimit)
	return ch(result)
}
