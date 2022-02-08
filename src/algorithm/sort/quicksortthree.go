package main

import "fmt"

func main() {

	array := []int{3, 1, 2, 5, 7, 23, 123, 45, 2, 15, 12}
	quicksortthree(array, 0, len(array)-1)
	fmt.Println(array)
}

func quicksortthree(nums []int, left int, right int) {
	if left >= right {
		return
	}
	site := partition1(nums, left, right)
	quicksortthree(nums, left, site[0])
	quicksortthree(nums, site[1], right)
}

func partition1(nums []int, left int, right int) [2]int {
	less, more := left-1, right+1
	cur, site := left, nums[left]

	for cur < more {
		if nums[cur] < site {
			less++
			nums[less], nums[cur] = nums[cur], nums[less]
			cur++
		} else if nums[cur] > site {
			more--
			nums[more], nums[cur] = nums[cur], nums[more]
		} else {
			cur++
		}
	}
	return [2]int{less, more}
}
