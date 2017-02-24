package main

import (
	"fmt"
	"github.com/rohit01/mahabharat/internal"
)

func main() {
	b, _ := internal.New("Rohit", 5, "1:3,2:3")
	b.BombedAt(1, 3)
	b.BombedAt(2, 3)
	b.BombedAt(2, 1)
	b.Print()
	fmt.Println(b.AllShipsDestroyed())
}
