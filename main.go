package main

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"encoding/json"
	"fmt"
	"strconv"
	"unsafe"

	"github.com/mritd/bubbles/common"

	"github.com/mritd/bubbles/selector"

	tea "github.com/charmbracelet/bubbletea"
)

func ch(str string) *C.char {
	return C.CString(str)
}

func str(ch *C.char) string {
	return C.GoString(ch)
}

func sf(err error) string {
	if err == nil {
		return ""
	}

	return fmt.Sprintf("%s", err)
}

//export FreeString
func FreeString(str *C.char) {
	C.free(unsafe.Pointer(str))
}

type model struct {
	sl selector.Model
}

func (m model) Init() tea.Cmd {
	return nil
}
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// By default, the prompt component will not return a "tea.Quit"
	// message unless Ctrl+C is pressed.
	//
	// If there is no error in the input, the prompt component returns
	// a "common.DONE" message when the Enter key is pressed.
	switch msg {
	case common.DONE:
		return m, tea.Quit
	}

	_, cmd := m.sl.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.sl.View()
}

type ListItem struct {
	Text        string `json:"text"`
	Description string `json:"description"`
}

func main() {}

type Result struct {
	SelectedIndex string `json:"selectedIndex"`
	Error         string `json:"error"`
}

//export CreateSelection
func CreateSelection(jsonData, headerText, footerText *C.char, perPage int) *C.char {
	var item []ListItem
	json.Unmarshal([]byte(str(jsonData)), &item)
	data := []interface{}{}
	for _, val := range item {
		data = append(data, ListItem{Text: val.Text, Description: val.Description})
	}
	m := &model{
		sl: selector.Model{
			Data:    data,
			PerPage: perPage,
			// Use the arrow keys to navigate: ↓ ↑ → ←
			// Select Commit Type:
			HeaderFunc: selector.DefaultHeaderFuncWithAppend(str(headerText)),
			// [1] feat (Introducing new features)
			SelectedFunc: func(m selector.Model, obj interface{}, gdIndex int) string {
				t := obj.(ListItem)
				if t.Description != "" {
					return common.FontColor(fmt.Sprintf("[%d] %s (%s)", gdIndex+1, t.Text, t.Description), selector.ColorSelected)
				}
				return common.FontColor(fmt.Sprintf("[%d] %s", gdIndex+1, t.Text), selector.ColorSelected)
			},
			// 2. fix (Bug fix)
			UnSelectedFunc: func(m selector.Model, obj interface{}, gdIndex int) string {
				t := obj.(ListItem)
				if t.Description != "" {
					return common.FontColor(fmt.Sprintf(" %d. %s (%s)", gdIndex+1, t.Text, t.Description), selector.ColorUnSelected)
				}
				return common.FontColor(fmt.Sprintf(" %d. %s", gdIndex+1, t.Text), selector.ColorUnSelected)
			},
			FooterFunc: func(m selector.Model, obj interface{}, gdIndex int) string {
				return common.FontColor(str(footerText), selector.ColorFooter)
			},
			FinishedFunc: func(s interface{}) string {
				return ""
			},
		},
	}

	p := tea.NewProgram(m)
	err := p.Start()
	if err != nil {
		result, _ := json.Marshal(&Result{
			SelectedIndex: "",
			Error:         sf(err),
		})
		return ch(string(result))
	}
	if !m.sl.Canceled() {
		result, _ := json.Marshal(&Result{
			SelectedIndex: strconv.Itoa(m.sl.Index()),
			Error:         "",
		})
		return ch(string(result))
	} else {
		result, _ := json.Marshal(&Result{
			SelectedIndex: "",
			Error:         "Cancelled",
		})
		return ch(string(result))
	}
}
