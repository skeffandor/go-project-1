package manager

import (
	"encoding/json"
	"os"
)

func Save(tm *TaskManager, fileName string) error {
	data, err := json.MarshalIndent(tm, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}

func Load(tm *TaskManager, fileName string) error {
	data, err := os.ReadFile(fileName)

	if os.IsNotExist(err) {
		return Save(tm, fileName)
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(data, tm)
}
