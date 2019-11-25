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
	for  i := 0; i < 100; i++{
		 nums = append(nums, rand.Intn(100))
	}
	fmt.Println(nums)
	QuickSort(nums)
	fmt.Println(nums)
}