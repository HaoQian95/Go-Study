package test

import "fmt"

type sortArray interface {
	len() int
	less(i int, j int) bool
	swap(i int, j int)
}

type intArray []int

func (array intArray)len() int {
	return len(array)
}

func (array intArray)less(i int, j int) bool {
	return array[i] < array[j]
}

func (array intArray)swap(i int, j int) {
	array[i], array[j] = array[j], array[i]
}

type stringArray []string

func (array stringArray) len() int {
	return len(array)
}

func (array stringArray) less(i int, j int) bool {
	return len(array[i]) < len(array[j])
}

func (array stringArray) swap(i int, j int) {
	array[i], array[j] = array[j], array[i]
}

func bsort(array sortArray){
	for i := 0; i < array.len(); i++ {
		for j := 1; j < array.len() - i; j++ {
			if array.less(j, j-1) {
				array.swap(j ,j-1)
			}
		}
	}
}

// TestInterface 测试函数
func TestInterface(){
	array := intArray{5,7,2,1,8,10,4}
	bsort(array)
	fmt.Println(array)

	strArray := stringArray{"china", "I", "love", "family"}
	bsort(strArray)
	fmt.Println(strArray)
}