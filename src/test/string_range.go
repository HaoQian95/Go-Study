package test

import (
	"fmt"
	"strings"
)

func StringRange() {
	s := "abcdefg"
	for ss := range(s) {
		fmt.Println(ss)
	}
}

func String1() {
	s := "ABCDEFG"
	string2(s)
	fmt.Println(s)
}

func string2(s string) {
	s = strings.ToLower(s)
	fmt.Println(s)
}

func String2Rune() {
	s := "abcdefg"
	ss := []rune(s)
	fmt.Println(s)
	fmt.Println(ss)
}