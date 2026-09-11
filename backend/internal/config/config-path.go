package config

import (
	"os"
	"path/filepath"
)

func ConfigPath() (string, error) {
	binPath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Join(filepath.Dir(binPath), "..", "config.toml"), nil
}
