package cmd

import (
	"strconv"
	"strings"
)

const (
	AddFlag      = "-add"
	AFlag        = "-a"
	CompleteFlag = "-complete"
	CFlag        = "-c"
	RedoFlag     = "-redo"
	RFlag        = "-r"
	DelFlag      = "-del"
	DFlag        = "-d"
)

type command struct {
	Flag string
	Args string
}

func ParseCmdArgs(cmdArgs []string) []command {
	currentFlag := ""
	flagArg := ""
	execCommands := make([]command, 0)

	for _, arg := range cmdArgs {
		if validFlags[arg] {
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

func ExecFlag(c command, todos *Todos) error {
	var err error
	switch c.Flag {
	case AddFlag, AFlag:
		err = todos.Add(c.Args)
	case CompleteFlag, CFlag:
		err = todos.Complete(getIdFromIndex(c.Args, *todos))
	case RedoFlag, RFlag:
		err = todos.Redo(getIdFromIndex(c.Args, *todos))
	case DelFlag, DFlag:
		err = todos.Delete(getIdFromIndex(c.Args, *todos))
	}
	return err
}

func getIdFromIndex(arg string, todos Todos) string {
	index, err := strconv.Atoi(strings.TrimSpace(arg))
	id := ""
	if err != nil || index < 1 || index > len(todos) {
		return id
	} else {
		for i, todo := range todos {
			if i == index-1 {
				id = todo.ID
				break
			}
		}
		return id
	}
}

var validFlags = map[string]bool{
	AddFlag:      true,
	AFlag:        true,
	CompleteFlag: true,
	CFlag:        true,
	RedoFlag:     true,
	RFlag:        true,
	DelFlag:      true,
	DFlag:        true,
}
