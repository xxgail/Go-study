package main

import "fmt"

// Go的八皇后问题
type EightQueue struct {
	Column []int
}

func (eq *EightQueue) CalEightQueues(row int) {
	if row == 8 {
		fmt.Println(eq.Column)
		eq.PrintQueue()
		return
	}
	for column := 0; column < 8; column++ {
		if eq.IsOk(row, column) {
			eq.Column[row] = column
			eq.CalEightQueues(row + 1)
		}
	}
}

func (eq *EightQueue) IsOk(row, column int) bool {
	left := column - 1
	right := column + 1

	for i := row - 1; i >= 0; i-- {
		if eq.Column[i] == column {
			return false
		}

		if left >= 0 {
			if eq.Column[i] == left {
				return false
			}
		}

		if right < 8 {
			if eq.Column[i] == right {
				return false
			}
		}
		left--
		right++
	}
	return true
}

func (eq *EightQueue) PrintQueue() {
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if eq.Column[row] == column {
				fmt.Printf(" Q ")
			} else {
				fmt.Printf(" * ")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	eq := EightQueue{
		Column: make([]int, 8),
	}
	eq.CalEightQueues(0)
}
