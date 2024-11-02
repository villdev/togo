package store

import (
	"encoding/json"
	"os"

	"github.com/villdev/togo/cmd"
)

func Load() (cmd.Todos, error) {
	var todos cmd.Todos
	filePath, err := GetJSONFilePath()
	if err != nil {
		return todos, err
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return todos, err
	}

	err = json.Unmarshal(file, &todos)
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func Save(src cmd.Todos) error {
	file, err := json.Marshal(src)
	if err != nil {
		return err
	}
	filePath, err := GetJSONFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
