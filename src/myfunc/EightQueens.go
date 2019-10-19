package myfunc

import (
	"fmt"
)

var array [8][8]int
var count int

func TestEightQueens(){
	count = 0
	fmt.Println("Eight queens question:")
	eightQueens(0)
	fmt.Println("Eight queens question has", count, "solutions.")
}

func eightQueens(row int) {
	if row > 7 {
		count++
		print()
		return
	}
	for i := 0; i < 8; i++ {
		if check(row, i) {
			array[row][i] = 1
			eightQueens(row + 1)
			array[row][i] = 0
		}
	}
}

func check(row int, col int) bool {
	for i := 0; i < row; i++ {		//检查列冲突
		if array[i][col] == 1 {
			return false
		}
	}
	for i,j := row-1, col-1; i >= 0 && j >= 0; i-- {	//检查左上对角线冲突
		if array[i][j] == 1 {
			return false
		}
		j--
	}
	for i,j := row-1, col+1; i >= 0 && j <= 7; i-- {	//检查右上对角线冲突
		if array[i][j] == 1 {
			return false
		}
		j++
	}
	return true
}

func print() {
	fmt.Println("No", count)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if array[i][j] == 0 {
				fmt.Print("o ")
			} else {
				fmt.Print("x ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}