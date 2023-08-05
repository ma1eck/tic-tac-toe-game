package mainMenu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {
	for {
		input := gettingInput() 
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

func gettingInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadBytes('\n')
	trimedInput := strings.TrimSpace(string(input))
	return	trimedInput
}
