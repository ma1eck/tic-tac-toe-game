package main

import (
	// "fmt"
	// controller "ticTacToe/controller"
	// gameModel "ticTacToe/model/game"
	viweMainMenu "ticTacToe/view/mainMenu"
)

func main() {
	viweMainMenu.Run()
	// newGame := gameModel.NewGame()
	// for {
	// 	var x,y int
	// 	fmt.Scanf("%d %d\n", &x, &y)
	// 	controller.Play(&newGame, x, y)
	// 	controller.PrintMap(newGame)
	// 	winner := controller.WhoWins(newGame.Map)
	// 	if (winner != nil){
	// 		fmt.Println("winner is", *winner)
	// 		break	
	// 	}
	// }
}
