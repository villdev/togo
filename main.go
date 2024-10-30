package main

import (
	"fmt"
	"os"

	"github.com/villdev/togo/cmd"
	"github.com/villdev/togo/store"
)

const (
	AddFlag      = "add"
	CompleteFlag = "complete"
	RedoFlag     = "redo"
	DelFlag      = "del"
)

func main() {
	var todos cmd.Todos
	err := store.Load("./db.json", &todos)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	cmdArgs := os.Args[1:]

	currentFlag := ""
	flagArg := ""
	for _, arg := range cmdArgs {
		if err == nil {
			switch arg {
			case "-" + AddFlag:
				if currentFlag != "" {
					err = execFlag(currentFlag, flagArg, &todos)
				}
				currentFlag = AddFlag
				flagArg = ""
			case "-" + CompleteFlag:
				if currentFlag != "" {
					err = execFlag(currentFlag, flagArg, &todos)
				}
				currentFlag = CompleteFlag
				flagArg = ""
			case "-" + RedoFlag:
				if currentFlag != "" {
					err = execFlag(currentFlag, flagArg, &todos)
				}
				currentFlag = RedoFlag
				flagArg = ""
			case "-" + DelFlag:
				if currentFlag != "" {
					err = execFlag(currentFlag, flagArg, &todos)
				}
				currentFlag = DelFlag
				flagArg = ""
			default:
				flagArg += " " + arg
			}
		}
	}
	if currentFlag != "" && err == nil {
		execFlag(currentFlag, flagArg, &todos)
	}

	if err != nil {
		fmt.Println("Error: ", err)
	}

	err = store.Save(todos, "./db.json")
	if err != nil {
		fmt.Println("Something went wrong: ", err)
		return
	}
	todos.Print()
}

func execFlag(flag string, arg string, t *cmd.Todos) error {
	var err error
	switch flag {
	case AddFlag:
		err = t.Add(arg)
	case CompleteFlag:
		err = t.Complete(arg)
	case RedoFlag:
		err = t.Redo(arg)
	case DelFlag:
		err = t.Delete(arg)
	}
	return err
}
