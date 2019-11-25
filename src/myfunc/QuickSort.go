package myfunc

func QuickSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	location := partition(nums, left, right)
	quickSort(nums, left, location-1)
	quickSort(nums, location+1, right)
}

func partition(nums []int, left int, right int) int {
	temp := nums[left]
	for left < right {
		for nums[right] >= temp && left < right {
			right--
		}
		nums[left] = nums[right]
		for nums[left] <= temp && left < right {
			left++
		}
		nums[right] = nums[left] 
	}
	nums[left] = temp
	return left
}