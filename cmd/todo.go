package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

type todo struct {
	ID        string
	Title     string
	Completed bool
	CreatedAt time.Time
}

type Todos []todo

func (t *Todos) Print() {
	if len(*t) == 0 {
		fmt.Println("Empty!!")
		return
	}

	columns := []table.Column{
		{Title: "No", Width: 4},
		{Title: "Title", Width: 40},
		{Title: "Status", Width: 8},
		{Title: "Created At", Width: 20},
	}

	var rows []table.Row
	for i, todo := range *t {
		status := "🔳"
		if todo.Completed {
			status = "✅"
		}

		rows = append(rows, table.Row{
			fmt.Sprintf("%d", i+1),
			todo.Title,
			status,
			todo.CreatedAt.Format("2006-01-02 15:04"),
		})
	}

	tb := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithHeight((len(rows)*3)+1),
		table.WithFocused(false),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Cell = s.Cell.
		BorderForeground(lipgloss.Color("240")).
		PaddingTop(1).
		PaddingBottom(1)
	s.Selected = s.Cell.
		BorderForeground(lipgloss.Color("240")).
		Padding(0)

	tb.SetStyles(s)

	fmt.Println(
		lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			Render(tb.View()))
}

func (t *Todos) Add(title string) error {
	if title == "" {
		return errors.New("add: title cannot be empty")
	}
	id, err := GenerateUniqueID()
	if err != nil {
		return errors.New("add: something went wrong")
	}
	newTodo := todo{
		ID:        id,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	*t = append(*t, newTodo)

	return nil
}

func (t *Todos) Complete(id string) error {
	found := false

	for index, todo := range *t {
		if todo.ID == id {
			(*t)[index].Completed = true
			found = true
			break
		}
	}

	if found {
		return nil
	} else {
		return errors.New("complete: invalid #id")
	}
}

func (t *Todos) Redo(id string) error {
	found := false

	for index, todo := range *t {
		if todo.ID == id {
			(*t)[index].Completed = false
			found = true
			break
		}
	}

	if found {
		return nil
	} else {
		return errors.New("redo: invalid #id")
	}
}

func (t *Todos) Delete(id string) error {
	found := false

	for _, todo := range *t {
		if todo.ID == id {
			found = true
			break
		}
	}

	if found {
		filteredTodo := make(Todos, 0, len(*t)-1)
		for _, todo := range *t {
			if todo.ID != id {
				filteredTodo = append(filteredTodo, todo)
			}
		}
		*t = filteredTodo

		return nil
	} else {
		return errors.New("del: invalid #id")
	}
}
