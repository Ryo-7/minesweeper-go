package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type state struct {
	bomb       bool
	opend      bool
	aroundBomb int
}

func main() {
	var matrixSize int
	fmt.Println("マインスイーパーの盤面サイズを指定してください(n四方になります)")
	fmt.Scan(&matrixSize)
	board := make([][]state, matrixSize)
	for i := range board {
		board[i] = make([]state, matrixSize)
	}
	closedCellNum := matrixSize * matrixSize

	var bombNum int
	fmt.Println("ボムの数を指定してください")
	fmt.Scan(&bombNum)
	i := 0
	for i < bombNum {
		x := rand.Intn(matrixSize)
		y := rand.Intn(matrixSize)
		if !board[x][y].bomb {
			board[x][y].bomb = true
			for j := -1; j < 2; j++ {
				if x == 0 && j == -1 {
					continue
				}
				if x == matrixSize-1 && j == 1 {
					continue
				}
				for k := -1; k < 2; k++ {
					if j == 0 && k == 0 {
						continue
					}
					if y == 0 && k == -1 {
						continue
					}
					if y == matrixSize-1 && k == 1 {
						continue
					}
					board[x+j][y+k].aroundBomb++
				}
			}
			i++
		}
	}
	for {
		var x, y int
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				if board[i][j].opend {
					fmt.Print(strconv.Itoa(board[i][j].aroundBomb) + " ")
				} else {
					fmt.Print("x ")
				}
			}
			fmt.Println("")
		}
		if closedCellNum == bombNum {
			fmt.Println("game clear")
			break
		}
		fmt.Println("空白区切りで開けたい座標をx yの順で指定してください")
		fmt.Scan(&x, &y)
		if board[x][y].bomb {
			fmt.Println("game over")
			break
		}
		board[x][y].opend = true
		closedCellNum--
	}
}
