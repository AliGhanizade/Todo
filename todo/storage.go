package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

var FileName string = "todo.json"

func RemoveTaskFromFIle(id int) ([]Task, error) {
	data, err := os.ReadFile(FileName)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	var existingTasks []Task
	if len(data) > 0 {
		// If the file is not empty, unmarshal the existing JSON data.
		err = json.Unmarshal(data, &existingTasks)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal existing JSON: %w", err)
		}
	}
	for i, t := range existingTasks {
		if t.ID == id {
			return append(existingTasks[:i], existingTasks[i+1:]...), nil
			
		}
	}
	return nil, fmt.Errorf("cant find id : %d", id)
}

func SaveTasksToFile(tasks []Task) error {
	os.Remove(FileName)
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal tasks to JSON: %w", err)
	}

	err = os.WriteFile(FileName, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("cant write in file: %w", err)
	}
	return nil
}

func SaveTaskToFile(tasks Task) error {
	data, err := os.ReadFile(FileName)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error reading file: %w", err)
	}

	var existingTasks []Task
	if len(data) > 0 {
		err = json.Unmarshal(data, &existingTasks)
		if err != nil {
			return fmt.Errorf("could not unmarshal existing JSON: %w", err)
		}
	}

	existingTasks = append(existingTasks, tasks)

	jsonData, err := json.MarshalIndent(existingTasks, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal tasks to JSON: %w", err)
	}

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
