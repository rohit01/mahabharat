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


Gophercon Coding Challenge
==========================

### Battleship problem
- Battleship is a game played between 2 players. Each player will be initialised with an M * M grid with ’S’ number of ships placed at specified positions on the grid. 
- One battleship occupies a single position on the grid.   
- It takes a single shot to bring down the battleship of an opponent.
- The objective of the game is to destroy opponent’s battleships. Each player will be given ’T’ number of missiles to bring down the fleet of opposition. The player who does the most damage/brings down the ships of the opposition wins
- This will be a simulation of an actual battle wherein each player is given ‘T’ number of moves as an input to the program
- Based on hits/misses of a missile on the opponent, either of the player might be victorious or the game might end as a draw

#### INPUT:
  The input for the game will be read from a file which contains the following (numbers below represent line numbers in input file)
  1. Contains the size of the battleground ‘M’ ( 0 < M < 10 )
  2. Contains the number of ships ’S’ which can be placed on the M*M grid ( 0 < S <= M/2)
  3. Player1 ship positions in the grid, position represented by x1:y1,x2:y2 and so on
  4. Player2 ship positions, they are placed in grid with the same format as above
  5. Tells the total number of missiles players have ( 0 < T < 100)
  6. Player 1 moves
  7. Player 2 moves

  These will be target positions (a, b) of each player’s missile attack on enemy's grid.

 ##### Input File Format:

  - M i.e GridSize [Matrix of M*M]
  - S i.e TotalShips
  - P1 Ship Positions: 1:1,2:0,2:3,3:4,.. (x:y pairs separated by commas)
  - P2 Ship Positions: 0:1,2:0,2:3,3:4,...
  - T i.e TotalMissiles
  - P1 moves: 1,1:2:0… (x,y pairs of length ’T’)
  - P2 moves: 1,2:3:0… (x,y pairs of length ‘T’)


  Sample Input File:
  ```
  5
  5
  1:1,2:0,2:3,3:4,4:3
  0:1,2:3,3:0,3:4,4:1
  5
  0,1:4,3:2,3:3,1:4,1
  0,1:0,0:1,1:2,3:4,3
  ```

####  OUTPUT:

  Output should be written to a file that should contain the following information:
  - Player 1 and Player 2 grids after the battleship simulation.
  - Alive Battleships denoted with “B”.
  - Dead Battleships with “X” (if missile hit the battleship) HIT
  - Missile Missed Locations “O” (if the missile location didn’t have a ship) MISS
  - Final result: P1:total hits P2:total hits 
  - Game Result

  Sample output file:
  ```
  Player1
  O O _ _ _
  _ X _ B _
  B _ _ X _
  _ _ _ _ B
  _ _ _ X _



  Player2
  _ X _ _ _
  _ _ _ _ _
  _ _ _ X _
  B O _ _ B
  _ X _ O _

  P1:3
  P2:3
  It is a draw
  ```

 ### Game result:

  - "Player 1 wins”, if player 1 does most damage to player 2’s ships
  - "Player 2 wins”, if player 2 does most damage to player 1’s ships
  - "It is a draw”, if both players inflict the same damage

  ### Game Rules:
  The solution to the problem should be object oriented. It should be solved in Golang.
  Some features of the solution carries more weightage which include:

  - Unit testing the solution
  - Idiomatic golang code
  - Good coding practices - naming, modular etc
  - Object oriented approach - Domain modelling
  - README

  P.S: You can reach out to GO-JEK folks for clarification of the above problem.
  We are not looking for the complete solution to the problem but one’s approach towards solving a problem
