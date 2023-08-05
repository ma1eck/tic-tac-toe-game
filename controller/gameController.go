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

func Play(game *gameModel.Game, x int, y int) bool {
	if (x>2 || x<0 || y>2 || y<0){
		return false
	}
	if (game.Map[x][y] == gameModel.Player{}) {
		game.Play(x, y)
		return true
	}
	return false
}

func WhoWins(gameMap [3][3]gameModel.Player) *gameModel.Player {
	// check rows
	for i := 0; i < 3; i++ {
		playersSetInARow := make(map[gameModel.Player]int)
		for j := 0; j < 3; j++ {
			player := gameMap[i][j]
			playersSetInARow[player]++
		}
		if playersSetInARow[gameModel.Player1] == 3 {
			return &gameModel.Player1
		}
		if playersSetInARow[gameModel.Player2] == 3 {
			return &gameModel.Player2
		}
	}
	// check colunm
	for i := 0; i < 3; i++ {
		playersSetInAColumn := make(map[gameModel.Player]int)
		for j := 0; j < 3; j++ {
			player := gameMap[j][i]
			playersSetInAColumn[player]++
		}
		if playersSetInAColumn[gameModel.Player1] == 3 {
			return &gameModel.Player1
		}
		if playersSetInAColumn[gameModel.Player2] == 3 {
			return &gameModel.Player2
		}
	}
	// check diameters
	playersSetInMainDiameter := make(map[gameModel.Player]int)
	for j := 0; j < 3; j++ {
		player := gameMap[j][j]
		playersSetInMainDiameter[player]++
	}
	if playersSetInMainDiameter[gameModel.Player1] == 3 {
		return &gameModel.Player1
	}
	if playersSetInMainDiameter[gameModel.Player2] == 3 {
		return &gameModel.Player2
	}
	playersSetInSecondDiameter := make(map[gameModel.Player]int)
	for j := 0; j < 3; j++{
		player := gameMap[j][2-j]
		playersSetInSecondDiameter[player]++
	}
	if playersSetInSecondDiameter[gameModel.Player1] == 3 {
		return &gameModel.Player1
	}
	if playersSetInSecondDiameter[gameModel.Player2] == 3 {
		return &gameModel.Player2

	}
	for _, row := range gameMap {
		for _, player := range row {
			if (player == gameModel.Player{}) {
				return nil
			}
		}
	}

	return &gameModel.Player{}
}
