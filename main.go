package main

import (
	"fmt"
	"os"

	"github.com/villdev/togo/cmd"
	"github.com/villdev/togo/store"
)

const (
	AddFlag      = "-add"
	CompleteFlag = "-complete"
	RedoFlag     = "-redo"
	DelFlag      = "-del"
)

type command struct {
	flag string
	args string
}

func main() {
	todos, err := store.Load("./db.json")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	flagQueue := parseCmdArgs(os.Args[1:])

	for _, c := range flagQueue {
		if err == nil {
			err = execFlag(c.flag, c.args, &todos)
		}
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

func parseCmdArgs(cmdArgs []string) []command {
	currentFlag := ""
	flagArg := ""
	execCommands := make([]command, 0)

	for _, arg := range cmdArgs {
		if arg == AddFlag || arg == CompleteFlag || arg == RedoFlag || arg == DelFlag {
			if currentFlag != "" {
				execCommands = append(execCommands, command{currentFlag, flagArg})
			}
			currentFlag = arg
			flagArg = ""
		} else {
			flagArg += " " + arg
		}
	}
	if currentFlag != "" {
		execCommands = append(execCommands, command{currentFlag, flagArg})
	}

	return execCommands
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
