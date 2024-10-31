package main

import (
	"fmt"
	"os"

	"github.com/villdev/togo/cmd"
	"github.com/villdev/togo/store"
)

const todoFilePath = "./store/todo.json"

func main() {
	todos, err := store.Load(todoFilePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	flagQueue := cmd.ParseCmdArgs(os.Args[1:])

	offset := 0
	for _, c := range flagQueue {
		err = cmd.ExecFlag(c, &todos, offset)
		if c.Flag == cmd.DelFlag {
			offset++
		}
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
	}

	err = store.Save(todos, todoFilePath)
	if err != nil {
		fmt.Println("Something went wrong: ", err)
		return
	}

	todos.Print()
}
