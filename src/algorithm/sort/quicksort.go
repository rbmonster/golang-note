package main

import "fmt"

func main() {
	array := []int{112, 23, 121, 23, 1123}
	quicksort(array, 0, len(array)-1)
	fmt.Println(array)
}

func quicksort(array []int, start int, end int) {
	if start > end {
		return
	}
	site := partition(array, start, end)
	quicksort(array, start, site-1)
	quicksort(array, site+1, end)
}

func partition(array []int, start int, end int) int {
	i, j := start, end
	site := array[start]
	for i < j {
		for i < j && array[j] >= site {
			j--
		}
		for i < j && array[i] <= site {
			i++
		}
		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i], array[start] = array[start], array[i]
	return i
}
