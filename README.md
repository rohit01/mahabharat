Mahabharat
==========

A ancient era game of battleships written in go.

Sample Run #1
```
$ go run battle.go ; cat output.txt
Starting game...
Game over, thanks for playing :)
Player 1
O _ _ _ _
_ X O _ _
B _ _ X _
_ _ _ _ B
_ _ _ X _

Player 2
_ X _ _ _
_ _ _ _ _
_ _ _ X _
B O _ _ B
_ X _ O _

P1:3
P2:3
It is a draw
```

Sample Run #2 (different moves by player #1)
```
$ go run battle.go ; cat output.txt
Starting game...
Game over, thanks for playing :)
Player 1
O O _ _ _
_ B O _ _
B _ _ X _
_ _ _ _ B
_ _ _ X _

Player 2
_ X _ _ _
_ _ _ _ _
_ _ _ X _
B _ _ _ X
_ X _ O _

P1:4
P2:2
Player 1 wins
```

Sample Run #3 (different moves by player #1)
```
$ go run battle.go ; cat output.txt
Starting game...
Game over, thanks for playing :)
Player 1
O O _ _ _
_ B O _ _
B _ _ X _
_ _ _ _ B
_ _ _ X _

Player 2
_ X _ _ _
_ _ _ _ _
_ _ _ X _
X _ _ _ X
_ X _ _ _

P1:5
P2:2
Player 1 wins
$
```

Note:

* input filename is `input.txt` in the present working direcory.
* output filename is `output.txt` in the present working direcory.
* This project can be improved :-)
