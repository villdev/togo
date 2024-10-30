package cmd

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

func ParseCmdArgs(cmdArgs []string) []command {
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

func ExecFlag(c command, todos *Todos) error {
	var err error
	switch c.flag {
	case AddFlag:
		err = todos.Add(c.args)
	case CompleteFlag:
		err = todos.Complete(c.args)
	case RedoFlag:
		err = todos.Redo(c.args)
	case DelFlag:
		err = todos.Delete(c.args)
	}
	return err
}
