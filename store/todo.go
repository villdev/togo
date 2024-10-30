package store

import (
	"encoding/json"
	"os"

	"github.com/villdev/togo/cmd"
)

func Load(filename string) (cmd.Todos, error) {
	var todos cmd.Todos
	file, err := os.ReadFile(filename)
	if err != nil {
		return todos, err
	}

	err = json.Unmarshal(file, &todos)
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func Save(src cmd.Todos, filename string) error {
	file, err := json.Marshal(src)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
