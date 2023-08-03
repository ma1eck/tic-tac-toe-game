package game

type Game struct {
	Map [3][3]Player
	currentPlayer Player
}
func NewGame() Game{
	// set player1 and player2

	var game Game = Game{ currentPlayer: Player1}
	return game
}


func (game *Game) Play(x int, y int) {
	game.Map[x][y] = game.currentPlayer
	game.nextPlayer()
}

func (game *Game) nextPlayer(){
	if game.currentPlayer == Player1 {
		game.currentPlayer = Player2
	}else {
		game.currentPlayer = Player1
	}
}