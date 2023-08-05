package gameMenu

import (
	"fmt"
	controller "ticTacToe/controller"
	gameModel "ticTacToe/model/game"
)


func StartGame(){
	newGame := gameModel.NewGame()
	for {
		var x,y int
		fmt.Scanf("%d %d\n", &x, &y)
		controller.Play(&newGame, x, y)
		controller.PrintMap(newGame)
		winner := controller.WhoWins(newGame.Map)
		if (winner == nil){ // todo clean return value 
			continue
		}
		if ( *winner == gameModel.Player{}) {
			fmt.Println("Game ended with a draw")
			break	
		}else {
			fmt.Println("winner is", *winner)
			break	
		}
	}
}