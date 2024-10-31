package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

type todo struct {
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

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "No.\tTitle\tCompleted\tCreated At")

	for i, ti := range *t {
		fmt.Fprintf(w, "%d.\t%s\t%v\t%s\n", i+1, ti.Title, ti.Completed, ti.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	w.Flush()
}

func (t *Todos) Add(title string) error {
	if title == "" {
		return errors.New("add: title cannot be empty")
	}

	newTodo := todo{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	*t = append(*t, newTodo)

	return nil
}

func (t *Todos) Complete(id string, offset int) error {
	index, err := strconv.Atoi(strings.TrimSpace(id))
	index -= offset
	if err != nil || index < 1 || index > len(*t) {
		return errors.New("complete: invalid #id")
	}

	(*t)[index-1].Completed = true
	return nil
}

func (t *Todos) Redo(id string, offset int) error {
	index, err := strconv.Atoi(strings.TrimSpace(id))
	index -= offset
	if err != nil || index < 1 || index > len(*t) {
		return errors.New("redo: invalid #id")
	}

	(*t)[index-1].Completed = false
	return nil
}

func (t *Todos) Delete(id string, offset int) error {
	index, err := strconv.Atoi(strings.TrimSpace(id))
	index -= offset
	if err != nil || index < 1 || index > len(*t) {
		return errors.New("del: invalid #id")
	}

	newT := make(Todos, 0, len(*t)-1)
	for i, ti := range *t {
		if i != index-1 {
			newT = append(newT, ti)
		}
	}
	*t = newT

	return nil
}
