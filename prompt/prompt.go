package prompt

import (
	"encoding/json"
	"fmt"

	"github.com/mritd/bubbles/common"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mritd/bubbles/prompt"
)

type model struct {
	input *prompt.Model
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

	_, cmd := m.input.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.input.View()
}

func (m model) Value() string {
	return m.input.Value()
}

type Result struct {
	Value string `json:"value"`
	Error string `json:"error"`
}

func Prompt(promptText, echoMode string, required bool, charLimit int) string {
	m := model{input: &prompt.Model{
		ValidateFunc: prompt.VFNotBlank,
		Prompt:       promptText,
		CharLimit:    charLimit,
		EchoMode:     prompt.EchoNormal,
	}}

	if echoMode == "none" {
		m.input.EchoMode = prompt.EchoNone
	} else if echoMode == "password" {
		m.input.EchoMode = prompt.EchoPassword
	} else {
		m.input.EchoMode = prompt.EchoNormal
	}

	if required {
		m.input.ValidateFunc = prompt.VFNotBlank
	} else {
		m.input.ValidateFunc = prompt.VFDoNothing
	}

	p := tea.NewProgram(&m)
	err := p.Start()
	if err != nil {
		result, _ := json.Marshal(&Result{
			Value: "",
			Error: fmt.Sprintf("%s", err),
		})
		return string(result)
	}
	result, _ := json.Marshal(&Result{
		Value: m.Value(),
		Error: "",
	})
	return string(result)
}
