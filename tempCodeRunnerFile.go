or {
		var x,y int
		fmt.Scanf("%d %d\n", &x, &y)
		controller.Play(&newGame, x, y)
		controller.PrintMap(newGame)
		winner := controller.WhoWins(newGeme.Map)
		if (winner != gameModel.Player{}){
			fmt.Println("winner is", winner)	
		}
	}