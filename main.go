package main

import (
	"fmt"
	controller "ticTacToe/controller"
	gameModel "ticTacToe/model/game"
)

func main() {
	newGame := gameModel.NewGame()
	for {
		var x,y int
		fmt.Scanf("%d %d\n", &x, &y)
		controller.Play(&newGame, x, y)
		controller.PrintMap(newGame)
	}
}
