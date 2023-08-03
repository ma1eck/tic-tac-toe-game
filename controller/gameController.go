package controller

import (
	gameModel "ticTacToe/model/game"
)

func PrintMap(game gameModel.Game) {
	for _, row := range game.Map {
		for _, player := range row {
			if (player == gameModel.Player{}) {
				print("# ")
			} else {
				print(player.Symbol + " ")
			}
		}
		println()
	}
}
func Play(game *gameModel.Game, x int, y int) bool{
	if (game.Map[x][y] == gameModel.Player{}) {
		game.Play(x, y)
		return true
	}
	return false
}
