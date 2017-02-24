// Module with structs and the game logic

package internal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type (
	Player struct {
		Name         string
		Hits         int
		battleGround []([]string)
	}

	Position struct {
		X int
		Y int
	}
)

func (self *Player) Print() {
	for rowNumber := range self.battleGround {
		for columnNumber := range self.battleGround[rowNumber] {
			fmt.Printf("%v ", self.battleGround[rowNumber][columnNumber])
		}
		fmt.Print("\n")
	}
}

func (self *Player) BombedAt(x, y int) (bool, error) {
	if (x >= len(self.battleGround)) || (y >= len(self.battleGround)) {
		return false, fmt.Errorf("Missile attacked (at %v,%v) beyond the grid length: %v", x, y, len(self.battleGround))
	}
	if (self.battleGround[x][y] == "B") || (self.battleGround[x][y] == "X") {
		self.battleGround[x][y] = "X"
		return true, nil
	} else {
		self.battleGround[x][y] = "O"
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

func New(name string, gridSize int, shipPositions string) (*Player, error) {
	if gridSize <= 0 {
		return nil, errors.New("Grid Size must be a positive integer")
	}
	// Allocate memory & initialize
	player := new(Player)
	player.battleGround = make([]([]string), gridSize)
	for rowNumber := 0; rowNumber < gridSize; rowNumber++ {
		player.battleGround[rowNumber] = make([]string, gridSize)
		for columnNumber := 0; columnNumber < gridSize; columnNumber++ {
			player.battleGround[rowNumber][columnNumber] = "_"
		}
	}
	// Process input positions
	initialPositions, err := ParseGridString(shipPositions)
	if err != nil {
		return nil, err
	}
	for _, position := range initialPositions {
		player.battleGround[position.X][position.Y] = "B"
	}
	return player, nil
}

func ParseGridString(gridString string) (moves []Position, err error) {
	moves = make([]Position, 0)
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
		moves = append(moves, Position{int(x), int(y)})
	}
	return moves, nil
}
