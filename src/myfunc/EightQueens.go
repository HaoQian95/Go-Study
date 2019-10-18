package myfunc

import "fmt"

var board [8][8]int
//count := 0
var count int

func TestEightQueens(){
	fmt.Println("Eight Queens question:")
	findQueen(0)
	fmt.Println("Eight Queens question has %d solutions:", count)
}

func findQueen(row int){
	if row > 7 {
		count++
		print()
		return
	}
	for i := 0; i < 8; i++ {
		if check(row, i){
			board[row][i] = 1
			findQueen(row + 1)
			board[row][i] = 0
		}
	}
}

func check(i int, j int) bool {

}

func print(){

}