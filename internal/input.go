// Module to take input from user

package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type (
	Config struct {
		GridSize, TotalShips, TotalMissiles                int
		P1ShipPositions, P2ShipPositions, P1Moves, P2Moves string
	}
)

func ReadInputFile(filename string) (*Config, error) {
	config := Config{}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		if !strings.HasPrefix(filename, "/") {
			dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
			if err != nil {
				return nil, err
			}
			filename = fmt.Sprintf("%v/%v", dir, filename)
			return ReadInputFile(filename)
		} else {
			return nil, err
		}
	}
	data := string(b)
	lines := strings.Split(data, "\n")
	var tempInt int64
	if tempInt, err = strconv.ParseInt(lines[0], 10, 64); err != nil {
		return nil, err
	}
	config.GridSize = int(tempInt)
	if tempInt, err = strconv.ParseInt(lines[1], 10, 64); err != nil {
		return nil, err
	}
	config.TotalShips = int(tempInt)
	if tempInt, err = strconv.ParseInt(lines[4], 10, 64); err != nil {
		return nil, err
	}
	config.TotalMissiles = int(tempInt)
	config.P1ShipPositions = strings.TrimSpace(lines[2])
	config.P2ShipPositions = strings.TrimSpace(lines[3])
	config.P1Moves = strings.TrimSpace(lines[5])
	config.P2Moves = strings.TrimSpace(lines[6])
	return &config, nil
}
