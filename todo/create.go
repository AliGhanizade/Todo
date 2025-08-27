package todo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Create() {
	var task Task

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter title: ")
	title, _ := reader.ReadString('\n')
	task.Title = strings.TrimSpace(title) 

	fmt.Print("Enter description: ")
	description, _ := reader.ReadString('\n')
	task.Description = strings.TrimSpace(description)
	
	task.IsCompleted = false
	task.CreateTime = time.Now().Format(time.DateTime)

	task.ID = LenTasks() + 1

	err := SaveTaskToFile(task)
	if err != nil {
		fmt.Print(err)
		return
	}
}
