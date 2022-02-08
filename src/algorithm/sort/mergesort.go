package main

import "math"

func main() {

}

func mergesort(nums []int) {
	//int len = array.length;
	//int[] result = new int[len];
	//int block, start;
	//for (block = 1; block < len * 2; block *= 2) {
	//	for (start = 0; start < len; start += block * 2) {
	//		int low = start;
	//		int mid = (start + block) < len ? (start + block) : len;
	//		int end = (start + block * 2) < len ? (start + block * 2) : len;
	//		int start1 = low, end1 = mid;
	//		int start2 = mid, end2 = end;
	//		while (start1 < end1 && start2 < end2) {
	//			result[low++] = array[start1] < array[start2] ? array[start1++] : array[start2++];
	//}
	//while (start1 < end1) {
	//result[low++] = array[start1++];
	//}
	//while (start2 < end2) {
	//result[low++] = array[start2++];
	//}
	//}
	//int[] temp = array;
	//array = result;
	//result = temp;
	//}
	//result = array;
	//return result;
	var array = []int{}
	length := len(nums)
	block := 0
	for ; block < length*2; block *= 2 {
		for start := 0; start < length; start += block * 2 {
			low := start
			mid := int(math.Min(float64((start + block)), float64(length)))
			end := int(math.Min(float64((start + block*2)), float64(length)))
			start1, end1 := low, mid
			start2, end2 := mid, end
			for start1 < end1 && start2 < end2 {
				if nums[start1] < nums[start2] {
					array[low] = nums[start]
				}
			}
		}
	}
}
