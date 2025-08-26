package cli

import (
	"fmt"
)

func GetUserInput() int {
	fmt.Print("1-Show all 2-show once(By ID) 3-Create 4-delete(By ID) 0-exit")
	var input int
	fmt.Scanf("%d", input)
	fmt.Println("You entered:", input)
	return input
}


