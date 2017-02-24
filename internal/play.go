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

func New(name string, gridSize int, positions string) (*Player, error) {
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
	for _, shipPosition := range strings.Split(positions, ",") {
		position := strings.Split(shipPosition, ":")
		x, err := strconv.ParseInt(position[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Invalid position data: %v", positions)
		}
		y, err := strconv.ParseInt(position[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Invalid position data: %v", positions)
		}
		player.battleGround[x][y] = "B"
	}
	return player, nil
}
