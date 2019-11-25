package myfunc

import (
	"testing"
	"fmt"
	"math/rand"
	"time"
)

func TestQuickSort(t *testing.T) {
	var nums []int
	rand.Seed(time.Now().UnixNano())
	for  i := 0; i < 10000; i++{
		 nums = append(nums, rand.Intn(10000))
	}
	//fmt.Println(nums)
	QuickSort(nums)
	fmt.Println(nums)
}