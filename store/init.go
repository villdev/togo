package store

import (
	"fmt"
	"os"
	"path/filepath"
)

// ensures `homeDir/.togo/todo.json` exists and returns exact path
func GetJSONFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine the home directory: %w", err)
	}
	todoDir := filepath.Join(homeDir, ".togo")

	if _, err := os.Stat(todoDir); os.IsNotExist(err) {
		if err = os.MkdirAll(todoDir, 0755); err != nil {
			return "", fmt.Errorf("error creating directory: %w", err)
		}
	}

	todoFilePath := filepath.Join(todoDir, "todo.json")

	if _, err := os.Stat(todoFilePath); os.IsNotExist(err) {
		file, err := os.Create(todoFilePath)
		if err != nil {
			return "", fmt.Errorf("failed to create todo.json file: %w", err)
		}

		_, err = file.WriteString("[]")
		if err != nil {
			return "", fmt.Errorf("failed to create todo.json file: %w", err)
		}

		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close newly created file: %v\n", closeErr)
		}
	}
	return todoFilePath, nil
}
