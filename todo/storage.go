package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

var FileName string = "todo.json"

func SaveTasksToFile(tasks Task) error {
	// First, read the existing tasks from the file.
	data, err := os.ReadFile(FileName)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error reading file: %w", err)
	}

	var existingTasks []Task
	if len(data) > 0 {
		// If the file is not empty, unmarshal the existing JSON data.
		err = json.Unmarshal(data, &existingTasks)
		if err != nil {
			return fmt.Errorf("could not unmarshal existing JSON: %w", err)
		}
	}

	// Append the new task to the slice of existing tasks.
	existingTasks = append(existingTasks, tasks)

	// Marshal the entire slice (including the new task) back to JSON.
	jsonData, err := json.MarshalIndent(existingTasks, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal tasks to JSON: %w", err)
	}

	// Write the entire, updated JSON array back to the file, overwriting the old content.
	err = os.WriteFile(FileName, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("cant write in file: %w", err)
	}

	return nil

}

func LoadTasksFromFile() ([]Task, error) {

	file, err := os.OpenFile(FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("error in open file : %w", err)
	}
	defer file.Close()
	//
	fileContent, err := os.ReadFile(FileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)

	}

	var tasks []Task
	//
	err = json.Unmarshal(fileContent, &tasks)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}
	//
	return tasks, nil
}

func LenTasks() int {
	tasks, err := LoadTasksFromFile()
	if err != nil {
		return 0
	}
	return len(tasks)
}
