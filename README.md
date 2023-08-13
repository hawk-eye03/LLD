
# TIC TAC TOE GAME

TicTacToe is a X/O game that two people play on 3x3 grid and three consecutive Xs or Os in a straight line determines the winner of the game.




## Run Locally

Clone the project

```bash
  git clone https://github.com/hawk-eye03/TicTacToe.git
```

Go to the project directory

```bash
  cd TicTacToe
```

Type on terminal 

```bash
  make run
```

Game should be up and running




## FEATURES
- **Move**: Players should be able to input their next move, in case of human player input has to be taken from terminal, in case of bot it moves automatically according to its playing strategy.
- **Undo**: Player should have the ability to undo the last move.


### Implementations
- **Move**: After every move we have to check if someone has won, and update the state of game if it happens
- **Undo**: When UNDO happens, clear cell, remove current move from list of moves and move turn to the previous player. 

#### Ways To Win
        One can win in three ways, same symbol in complete row or complete column or complete diagonal.
#### Algorithm to check winner
- There are three algorithms to check for the winner:
        
        1) Brute Force: One can win in three ways, same symbol in complete row or complete row or complete diagonal. TC: O(n3). [no_symbols x no_rows x no_columns]
        2) Some Optimization: By observation we can that the player who made the last move can only be a winner. We can therefore reduce TC to O(n2) because there is no need to check all the symbols, just the last one. [no_rows x no_columns]
        3) Most Optimised: Have a dictionary for all rows, columns and diagonals. Each dictionary has a symbol as a key and count of that symbol in that line as value. After each move, we update the all related dicts. Updating and checking dict requires O(1) time. So we can check winner where the value of a particular key becomes N. 
        As a result, TC: O(1), SC: O(n2) [no_of_maps x space_for_each_map]

        Note: We have used Algorithm 3 to check for the winner in the code

#### Interfaces
- WinningStrategies

        1) RowWinningStrategy 
        2) ColumnWinningStrategy 
        3) DiagonalWinningStrategy 

- BotPlayingStrategies

        1) EASYBotPlayingStrategy 
        2) MEDIUMBotPlayingStrategy 
        3) HARDBotPlayingStrategy 
        
        Note: As part of this code we have only implemented easy bot playing strategy
