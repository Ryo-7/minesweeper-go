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

// ゲームの初期化処理
func initGame(boardSize, bombNum, x, y int) [][]state {
	board := createGameBoard(boardSize)
	setBomb(board, boardSize, bombNum, x, y)
	return board
}

// 盤面作成
func createGameBoard(boardSize int) [][]state {
	board := make([][]state, boardSize)
	for i := range board {
		board[i] = make([]state, boardSize)
	}
	return board
}

// 爆弾を盤面にセット
func setBomb(gameBoard [][]state, boardSize int, bombNum, playerSelectX, playerSelectY int) {
	i := 0
	for i < bombNum {
		x := rand.Intn(boardSize)
		y := rand.Intn(boardSize)
		if playerSelectX == x && playerSelectY == y {
			continue
		}
		if !gameBoard[x][y].bomb {
			gameBoard[x][y].bomb = true
			for j := -1; j < 2; j++ {
				if x == 0 && j == -1 {
					continue
				}
				if x == boardSize-1 && j == 1 {
					continue
				}
				for k := -1; k < 2; k++ {
					if j == 0 && k == 0 {
						continue
					}
					if y == 0 && k == -1 {
						continue
					}
					if y == boardSize-1 && k == 1 {
						continue
					}
					gameBoard[x+j][y+k].aroundBomb++
				}
			}
			i++
		}
	}
}

// 盤面を表示する関数
//
// すでに開かれているところはその周辺にある爆弾の数を表示する
func printGameBoard(board [][]state, finished bool) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if finished {
				if board[i][j].bomb {
					fmt.Print("B ")
				} else {
					fmt.Print(strconv.Itoa(board[i][j].aroundBomb) + " ")
				}

			} else {
				if board[i][j].opend {
					fmt.Print(strconv.Itoa(board[i][j].aroundBomb) + " ")
				} else {
					fmt.Print("x ")
				}
			}
		}
		fmt.Println("")
	}
}

// 開けたところが0だった場合、連鎖的に開ける
func chainOpen(board *[][]state, x int, y int) int {
	openCells := 0
	if x > 0 && y > 0 {
		if !(*board)[x-1][y-1].opend {
			(*board)[x-1][y-1].opend = true
			openCells++
			if (*board)[x-1][y-1].aroundBomb == 0 {
				openCells += chainOpen(board, x-1, y-1)
			}
		}
	}
	if x < len(*board)-1 && y < len(*board)-1 {
		if !(*board)[x+1][y+1].opend {
			(*board)[x+1][y+1].opend = true
			openCells++
			if (*board)[x+1][y+1].aroundBomb == 0 {
				openCells += chainOpen(board, x+1, y+1)
			}
		}
	}
	if x > 0 && y < len(*board)-1 {
		if !(*board)[x-1][y+1].opend {
			(*board)[x-1][y+1].opend = true
			openCells++
			if (*board)[x-1][y+1].aroundBomb == 0 {
				openCells += chainOpen(board, x-1, y+1)
			}
		}
	}
	if x < len(*board)-1 && y > 0 {
		if !(*board)[x+1][y-1].opend {
			(*board)[x+1][y-1].opend = true
			openCells++
			if (*board)[x+1][y-1].aroundBomb == 0 {
				openCells += chainOpen(board, x+1, y-1)
			}
		}
	}
	if x > 0 {
		if !(*board)[x-1][y].opend {
			(*board)[x-1][y].opend = true
			openCells++
			if (*board)[x-1][y].aroundBomb == 0 {
				openCells += chainOpen(board, x-1, y)
			}
		}
	}
	if x < len(*board)-1 {
		if !(*board)[x+1][y].opend {
			(*board)[x+1][y].opend = true
			openCells++
			if (*board)[x+1][y].aroundBomb == 0 {
				openCells += chainOpen(board, x+1, y)
			}
		}
	}
	if y > 0 {
		if !(*board)[x][y-1].opend {
			(*board)[x][y-1].opend = true
			openCells++
			if (*board)[x][y-1].aroundBomb == 0 {
				openCells += chainOpen(board, x, y-1)
			}
		}
	}
	if y < len(*board)-1 {
		if !(*board)[x][y+1].opend {
			(*board)[x][y+1].opend = true
			openCells++
			if (*board)[x][y+1].aroundBomb == 0 {
				openCells += chainOpen(board, x, y+1)
			}
		}
	}
	return openCells
}

func main() {
	var matrixSize int
	var bombNum int
	var x, y int

	for {
		fmt.Println("マインスイーパーの盤面サイズを指定してください(n四方になります)")
		fmt.Scan(&matrixSize)
		if matrixSize > 0 {
			break
		} else {
			fmt.Println("0以上の数値を指定してください")
		}
	}

	closedCellNum := matrixSize * matrixSize

	for {
		fmt.Println("ボムの数を指定してください")
		fmt.Scan(&bombNum)
		if bombNum > 0 && bombNum < closedCellNum {
			break
		} else {
			fmt.Printf("0以上の%d未満の数値を指定してください\n", closedCellNum)

		}
	}

	for {
		fmt.Println("空白区切りで開けたい座標をx yの順で指定してください")
		fmt.Scan(&x, &y)
		if 0 > x || x >= matrixSize {
			fmt.Printf("x座標は0以上%d未満で指定してください\n", matrixSize)
		} else if 0 > y || y >= matrixSize {
			fmt.Printf("y座標は0以上%d未満で指定してください\n", matrixSize)
		} else {
			break
		}
	}

	board := initGame(matrixSize, bombNum, x, y)

	for {
		if board[x][y].bomb {
			printGameBoard(board, true)
			fmt.Println("game over")
			break
		}
		if board[x][y].opend {
			fmt.Println("すでに開いています。")
			for {
				fmt.Println("空白区切りで開けたい座標をx yの順で指定してください")
				fmt.Scan(&x, &y)
				if 0 > x || x >= matrixSize {
					fmt.Printf("x座標は0以上%d未満で指定してください\n", matrixSize)
				} else if 0 > y || y >= matrixSize {
					fmt.Printf("y座標は0以上%d未満で指定してください\n", matrixSize)
				} else {
					break
				}
			}
			continue
		}
		board[x][y].opend = true
		closedCellNum--
		openCells := 0
		if board[x][y].aroundBomb == 0 {
			openCells = chainOpen(&board, x, y)
		}
		closedCellNum -= openCells
		printGameBoard(board, false)
		if closedCellNum == bombNum {
			fmt.Println("game clear")
			break
		}
		for {
			fmt.Println("空白区切りで開けたい座標をx yの順で指定してください")
			fmt.Scan(&x, &y)
			if 0 > x || x >= matrixSize {
				fmt.Printf("x座標は0以上%d未満で指定してください\n", matrixSize)
			} else if 0 > y || y >= matrixSize {
				fmt.Printf("y座標は0以上%d未満で指定してください\n", matrixSize)
			} else {
				break
			}
		}
	}
}
