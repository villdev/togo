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

func ParseCmdArgs(cmdArgs []string, todos *Todos) []command {
	currentFlag := ""
	flagArg := ""
	execCommands := make([]command, 0)

	for _, arg := range cmdArgs {
		if validFlags[arg] {
			if currentFlag != "" {
				execCommands = append(execCommands, generateCommand(currentFlag, flagArg, *todos))
			}
			currentFlag = arg
			flagArg = ""
		} else {
			flagArg += " " + arg
		}
	}
	if currentFlag != "" {
		execCommands = append(execCommands, generateCommand(currentFlag, flagArg, *todos))
	}

	return execCommands
}

func ExecFlag(c command, todos *Todos) error {
	var err error
	switch c.Flag {
	case AddFlag, AFlag:
		err = todos.Add(c.Args)
	case CompleteFlag, CFlag:
		err = todos.Complete(c.Args)
	case RedoFlag, RFlag:
		err = todos.Redo(c.Args)
	case DelFlag, DFlag:
		err = todos.Delete(c.Args)
	}
	return err
}

func generateCommand(flag string, args string, todos Todos) command {
	c := command{flag, args}
	if flag == AddFlag || flag == AFlag {
		return c
	} else {
		c.Args = getIdFromIndex(c.Args, todos)
		return c
	}
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
