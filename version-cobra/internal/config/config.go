package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type InputTarget struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Owner string `json:"owner"`
}

func LoadTargetFromFile(filePath string) ([]InputTarget, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Impossible de lire le fichier")
	}
	var targets []InputTarget
	if err = json.Unmarshal(data, &targets); err != nil {
		return nil, fmt.Errorf("Impossible de lire le fichier %s: %w", filePath)
	}
	return targets, nil
}

func SaveTargetToFile(filePath string, target []InputTarget) error {
	data, err := json.MarshalIndent(target, "", " ")
	if err != nil {
		return fmt.Errorf("Impossible de lire le fichier")
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("Impossible d'écrire dans le fichier")
	}
	return nil
}
