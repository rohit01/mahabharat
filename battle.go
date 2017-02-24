package main

import (
	"context"
	"fmt"
	"github.com/rohit01/mahabharat/internal"
	"time"
)

func main() {
	config, err := internal.ReadInputFile(internal.INPUT_FILENAME)
	if err != nil {
		fmt.Println("Error reading config...")
		panic(err)
	}
	player1, _ := internal.New(1, config.GridSize, config.P1ShipPositions, config.P1Moves)
	player2, _ := internal.New(2, config.GridSize, config.P2ShipPositions, config.P2Moves)

	ctx := context.Background()
	gameOver := make(chan struct{})
	go player1.Attack(ctx, player2, gameOver)
	go player2.Attack(ctx, player1, gameOver)
	fmt.Println("Starting game...")
	player2.Turn <- struct{}{}
	<-gameOver
	time.Sleep(100 * time.Millisecond)

	result := ""
	if player1.Winner {
		result = "Player 1 wins"
	} else if player2.Winner {
		result = "Player 2 wins"
	} else if player1.Hits > player2.Hits {
		result = "Player 1 wins"
	} else if player2.Hits > player1.Hits {
		result = "Player 2 wins"
	} else {
		result = "It is a draw"
	}

	outputMsg := fmt.Sprintf("Player 1\n%v\nPlayer 2\n%v\nP1:%v\nP2:%v\n%v\n",
		player1.GroundStatus(),
		player2.GroundStatus(),
		player1.Hits,
		player2.Hits,
		result,
	)
	internal.WriteResult(internal.OUTPUT_FILENAME, outputMsg)
	fmt.Println("Game over, thanks for playing :)")
}
