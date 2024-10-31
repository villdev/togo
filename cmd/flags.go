package cmd

const (
	AddFlag      = "-add"
	CompleteFlag = "-complete"
	RedoFlag     = "-redo"
	DelFlag      = "-del"
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

func ExecFlag(c command, todos *Todos, offset int) error {
	var err error
	switch c.Flag {
	case AddFlag:
		err = todos.Add(c.Args)
	case CompleteFlag:
		err = todos.Complete(c.Args, offset)
	case RedoFlag:
		err = todos.Redo(c.Args, offset)
	case DelFlag:
		err = todos.Delete(c.Args, offset)
	}
	return err
}
