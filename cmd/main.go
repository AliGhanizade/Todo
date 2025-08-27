package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todoCLI/todo"
)

func main() {
	for {

		fmt.Print("\tWrite Work\nShow-Create-Delete-exit :")
		reader := bufio.NewReader(os.Stdin)

		
		work, _ := reader.ReadString('\n')
		work = strings.TrimSpace(work)
		switch work {
		case "Show":
			todo.GetAll()
		case "Create":
			todo.Create()
		case "Delete":
			todo.Delete()
		case "exit":
			os.Exit(1)
		}
	}

}
