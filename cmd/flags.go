package cmd

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

func ExecFlag(c command, todos *Todos, offset int) error {
	var err error
	switch c.Flag {
	case AddFlag, AFlag:
		err = todos.Add(c.Args)
	case CompleteFlag, CFlag:
		err = todos.Complete(c.Args, offset)
	case RedoFlag, RFlag:
		err = todos.Redo(c.Args, offset)
	case DelFlag, DFlag:
		err = todos.Delete(c.Args, offset)
	}
	return err
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
