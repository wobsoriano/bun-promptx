package selection

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/mritd/bubbles/common"

	"github.com/mritd/bubbles/selector"

	tea "github.com/charmbracelet/bubbletea"
)

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

type Result struct {
	SelectedIndex string `json:"selectedIndex"`
	Error         string `json:"error"`
}

func Selection(jsonData, headerText, footerText string, perPage int) string {
	var item []ListItem
	json.Unmarshal([]byte(jsonData), &item)
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
			HeaderFunc: selector.DefaultHeaderFuncWithAppend(headerText),
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
				return common.FontColor(footerText, selector.ColorFooter)
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
			Error:         fmt.Sprintf("%s", err),
		})
		return string(result)
	}
	if !m.sl.Canceled() {
		result, _ := json.Marshal(&Result{
			SelectedIndex: strconv.Itoa(m.sl.Index()),
			Error:         "",
		})
		return string(result)
	} else {
		result, _ := json.Marshal(&Result{
			SelectedIndex: "",
			Error:         "Cancelled",
		})
		return string(result)
	}
}
