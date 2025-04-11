package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func Init() {
	_, err := createAppConfigDirectory()
	if err != nil {
		fmt.Println("Error creating config: %w", err)
	}
}

func createAppConfigDirectory() (string, error) {
	cdir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user config dir: %w", err)
	}

	appConfigDir := filepath.Join(cdir, "schemasnap")

	if _, err := os.Stat(appConfigDir); os.IsNotExist(err) {
		if err := os.MkdirAll(appConfigDir, 0755); err != nil {
			return "", fmt.Errorf("failed to create config dir: %w", err)
		}
	}

	return appConfigDir, nil
}
