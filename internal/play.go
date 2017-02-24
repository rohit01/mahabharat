// Module with structs and the game logic

package internal

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type (
	Player struct {
		PlayerNo     int
		Hits         int
		AttackMoves  []CoOrdinates
		Winner       bool
		Turn         chan struct{}
		battleGround []([]string)
	}

	CoOrdinates struct {
		X int
		Y int
	}
)

func (self *Player) GroundStatus() string {
	groundStatus := ""
	for rowNumber := range self.battleGround {
		for columnNumber := range self.battleGround[rowNumber] {
			groundStatus = fmt.Sprintf("%v%v ", groundStatus, self.battleGround[rowNumber][columnNumber])
		}
		groundStatus = fmt.Sprintf("%v\n", groundStatus)
	}
	return groundStatus
}

func (self *Player) Bomb(position CoOrdinates) (bool, error) {
	if (position.X >= len(self.battleGround)) || (position.Y >= len(self.battleGround)) {
		return false, fmt.Errorf("Missile attacked (at %v,%v) beyond the grid length: %v",
			position.X, position.Y, len(self.battleGround))
	}
	if (self.battleGround[position.X][position.Y] == "B") || (self.battleGround[position.X][position.Y] == "X") {
		self.battleGround[position.X][position.Y] = "X"
		return true, nil
	} else {
		self.battleGround[position.X][position.Y] = "O"
	}
	return false, nil
}

func (self *Player) AllShipsDestroyed() bool {
	for _, row := range self.battleGround {
		for _, item := range row {
			if item == "B" {
				return false
			}
		}
	}
	return true
}

func (self *Player) Attack(ctx context.Context, opponent *Player, gameOver chan struct{}) {
	for _, position := range self.AttackMoves {
		select {
		case <-self.Turn:
			opponent.Bomb(position)
			if opponent.AllShipsDestroyed() {
				self.Winner = true
				gameOver <- struct{}{}
			}
			opponent.Turn <- struct{}{}
		case <-ctx.Done():
			fmt.Println("Context done")
			gameOver <- struct{}{}
		}
	}
	gameOver <- struct{}{}
}

func New(playerNo int, gridSize int, shipGridString, attackGridString string) (*Player, error) {
	if gridSize <= 0 {
		return nil, errors.New("Grid Size must be a positive integer")
	}
	// Allocate memory & initialize
	player := new(Player)
	player.Turn = make(chan struct{})
	player.PlayerNo = playerNo
	player.battleGround = make([]([]string), gridSize)
	for rowNumber := 0; rowNumber < gridSize; rowNumber++ {
		player.battleGround[rowNumber] = make([]string, gridSize)
		for columnNumber := 0; columnNumber < gridSize; columnNumber++ {
			player.battleGround[rowNumber][columnNumber] = "_"
		}
	}
	// Process input positions & set attack moves
	initialPositions, err := ParseGridString(shipGridString)
	if err != nil {
		return nil, err
	}
	for _, position := range initialPositions {
		player.battleGround[position.X][position.Y] = "B"
	}
	attackMoves, err := ParseGridString(attackGridString)
	if err != nil {
		return nil, err
	}
	player.AttackMoves = attackMoves
	return player, nil
}

func ParseGridString(gridString string) (moves []CoOrdinates, err error) {
	moves = make([]CoOrdinates, 0)
	for _, position := range strings.Split(gridString, ",") {
		xyPosition := strings.Split(position, ":")
		x, err := strconv.ParseInt(xyPosition[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Invalid grid position data: %v", gridString)
		}
		y, err := strconv.ParseInt(xyPosition[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Invalid grid position data: %v", gridString)
		}
		moves = append(moves, CoOrdinates{int(x), int(y)})
	}
	return moves, nil
}
