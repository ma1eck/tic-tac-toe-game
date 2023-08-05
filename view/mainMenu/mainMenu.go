package mainMenu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	gameMenu "ticTacToe/view/gameMenu"
)

func Run() {
	for {
		input := gettingInput() 
		if input == "exit"{
			fmt.Println("You have exited")
			break
		}else if input == "start game"{
			fmt.Println("The game has been started")
			gameMenu.StartGame()
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
