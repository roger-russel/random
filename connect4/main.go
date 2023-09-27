package main

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	None Player = 0
	playerOne Player = 1
	playerTwo Player = 2

	Rows int = 6
	Columns int = 7
)

type Player int

var errFullColumn = errors.New("the column is full choice another column")

func main() {
	game := NewGame()
	game.Start()
	
}

func NewGame() ConnectFour{
	return &GameConnectFour{}
}

type ConnectFour interface {
	Start()
	CheckWinner(player Player, column, row int) bool
	DropDisk(player Player, column int) (row int, err error)
	RandomColumn() int
	GetNextPlayer(player Player) Player
	CheckDirection(player Player, column, row, x, y int) int
}

type GameConnectFour struct { 
	board [Rows][Columns]Player
}

func (g *GameConnectFour) Start(){
	currentPlayer := playerOne

	for {
		var column int 
		var row int
		var err error

		for {
			column = g.RandomColumn()

			row, err = g.DropDisk(currentPlayer, column)
			if err == nil {
				break
			}
			fmt.Println("invalid moviment from player", currentPlayer, err.Error())
		}

		if g.CheckWinner(currentPlayer, column, row) {
			fmt.Println("winner: ", currentPlayer)
			break
		}	

		currentPlayer = g.GetNextPlayer(currentPlayer)
	}
}

func (g *GameConnectFour) GetNextPlayer(currPlayer Player) Player{
	if currPlayer == playerOne {
		return playerTwo
	}
	return playerOne
}

func (g *GameConnectFour) RandomColumn() int {
	return rand.Intn(Columns)
}

func (g *GameConnectFour) DropDisk(currPlayer Player, column int) (row int, err error) {
	for row = 0; row < Rows ; row++ {
		if g.board[row][column] == None {
			g.board[row][column] = currPlayer
			return row, nil
		}
	}
	return 0, errFullColumn
}

func (g *GameConnectFour) CheckWinner(player Player, column, row int) bool {
	directionsToCheckPositions := [][]int{
		{0, 1}, // Horizontal
		{1, 0}, // Vertical
		{1, 1}, // Diagnal \
		{1, -1}, // Diagonal /
	}

	for _, direction := range directionsToCheckPositions {
		total := g.CheckDirection(player, column, row, direction[0], direction[1]) + 
				 g.CheckDirection(player, column, row, -direction[0], -direction[1])

		if total >= 3 {
			return true
		}
	}
	return false
}

func (g *GameConnectFour) CheckDirection(player Player, column,  row, dx, dy int) int {
	count := 0
	for i := 0; i < 4 ; i++ {
		column += dx
		row += dy

		if 	row > 0 && 
			column > 0 && 
			row < Rows && 
			column < Columns && 
			g.board[row][column] == player {

			count++
		}
	}
	return count
}
