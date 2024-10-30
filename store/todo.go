package store

import (
	"encoding/json"
	"os"

	"github.com/villdev/togo/cmd"
)

func Load(filename string, dst *cmd.Todos) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, dst)
	if err != nil {
		return err
	}

	return nil
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
