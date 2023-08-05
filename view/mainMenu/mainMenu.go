package mainMenu

import (
	"fmt"
	"strings"
)

func Run() {
	for {
		var input string
		fmt.Scanf("%[^\n]s",&input)
		input = strings.TrimSpace(input) 
		if input == "exit"{
			fmt.Println("You have exited")
			break
		}else if input == "start game"{
			fmt.Println("The game has been started")
		}else {
			fmt.Println("Invalid input")
		}
	}
}
