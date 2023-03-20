package main

import (
    "fmt"
    "flag"
    "time"
    "math/rand"
)

var board [][]bool

func main() {
    fmt.Println("Starting game of life")

    width := flag.Int("width", 10, "width of the game board")
    height := flag.Int("height", 10, "height of the game board")

    flag.Parse()
    fmt.Println("width: ", *width)
    fmt.Println("height: ", *height)

    board = initBoard(*width, *height)

    gameLoop()
}

func initBoard(width, height int) [][]bool {
    board := make([][]bool, height)
    for i := range board {
        board[i] = make([]bool, width)

        for j := 0; j < len(board[i]); j++ {
            r := rand.Intn(100)

            if r > 25 {
                board[i][j] = true
            }
        }
    }
    return board
}

func gameLoop() {
   for {
    printBoard()

    updateBoard()

    time.Sleep(500 * time.Millisecond)
   }
}

func updateBoard() {
    for y := 0; y < len(board); y++ {
        for x := 0; x < len(board[y]); x++ {
            var position bool = board[y][x]
            var aliveCount int = numberAliveNeighbors(x, y)

            if position { //alive
                if aliveCount < 2 || aliveCount > 3 {
                    board[y][x] = false
                } 
            } else {
                if aliveCount == 3 {
                    board[y][x] = true
                }
            }
        }
    }
}

func numberAliveNeighbors(x, y int) int {
    aliveCount := 0
    //top
    if y - 1 >= 0 {
        //top
        if board[y - 1][x] {
            aliveCount++
        }

        //top left
        if x - 1 >= 0 {
            if board[y - 1][x - 1] {
                aliveCount++
            }
        }

        //top right
        if x + 1 < len(board[y - 1]) {
            if board[y - 1][x + 1] {
                aliveCount++
            }
        }
    }

    

    //right
    if x + 1 < len(board[y]) {
        if board[y][x + 1] {
            aliveCount++
        }
    }

    //left
    if x - 1 >= 0 {
        if board[y][x - 1] {
            aliveCount++
        }
    }

    //bottom
    if y + 1 < len(board) {
        //bottom
        if board[y + 1][x] {
            aliveCount++
        }

        //bottom left
        if x - 1 >= 0 {
            if board[y + 1][x - 1] {
                aliveCount++
            }
        }

        //bottom right
        if x + 1 < len(board[y + 1]) {
            if board[y + 1][x + 1] {
                aliveCount++
            }
        }

    }

    return aliveCount
}

func printBoard() {
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            if board[i][j] {
                fmt.Print("\033[31m" + "O" + "\033[0m")
            } else {
                fmt.Print(" ")
            }
        }

        fmt.Println()
    }
}
