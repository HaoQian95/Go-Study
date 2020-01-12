package temp

import "fmt"

func changePointer(x int) {
	x = x + 1
}

func Sss() {
	value := 10
	p := &value
	changePointer(*p)
	fmt.Println(value)
}

func RangeTest() {
	array := []int{3,4,5,6,7}
	for num := range(array) {
		fmt.Println(num)
	}
}

func SliceAppendNil() {
	s := make([]*int, 0)
	s = append(s, nil)
	fmt.Println(len(s))
}