package todo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Delete() {
	var id string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter ID: ")
	id, _ = reader.ReadString('\n')
	id = strings.TrimSpace(id)

	intId, _ := strconv.Atoi(id)
	tasks, err := RemoveTaskFromFIle(intId)
	if err != nil {
		return
	}
	SaveTasksToFile(tasks)

}
