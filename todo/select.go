package todo

import "fmt"

func GetAll() {

	tasks, err := LoadTasksFromFile()
	if err != nil {
		return
	}
	// fmt.Print(tasks)
	for _, t := range tasks {
		fmt.Printf("%d\t-%s-\t|%t|\n%s\n---------------------\n", t.ID, t.Title, t.IsCompleted, t.Description)
	}
}
