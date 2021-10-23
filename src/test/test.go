package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

type Point struct {
	X, Y float64
}

type Circle struct {
	Point  Point
	Redius int
}

type Wheel struct {
	Circle Circle
	Spokes int
	owner  string
}

type DoInterface interface {
	doRead() (msg string, err error)
	doWrite(msg string) (ok bool, err error)
}

func (w Wheel) doWrite(msg string) (ok bool, err error) {
	fmt.Println("write", msg)
	return true, nil
}

func (w Wheel) doRead() (msg string, err error) {
	msg = "sdaf"
	fmt.Println("read", msg)
	return "sdf", nil
}

func (w Wheel) Write(p []byte) (n int, err error) {
	return 12, nil
}

type HandlerFunc func(x int, ptr *[3]int, f func(inter interface{}))

func (w Wheel) list(x int, ptr *[3]int, f func(inter interface{})) {
	fmt.Println(x, *ptr)
	f(1)
	fmt.Println("========>")
	fmt.Println(w.owner)
}

// teetsetes
func main() {

	//nums := []int{3,1,6,3,87,42,2,7,12,31,43}
	//quicksort(nums, 0, len(nums)-1)
	//fmt.Println(nums)
	nums := []int{1, 3, 1, 5, 4}
	findPairs(nums, 0)
}

func maxAbsValExpr(arr1 []int, arr2 []int) int {
	var aMax = float64(math.MinInt)
	var aMin = math.MaxFloat64
	var bMax float64
	var bMin = math.MaxFloat64
	var cMax float64
	var cMin = math.MaxFloat64
	var dMax float64
	var dMin = math.MaxFloat64
	/**
	A = arr1[i] + arr2[i] + i
	B = arr1[i] + arr2[i] - i
	C = arr1[i] - arr2[i] + i
	D = arr1[i] - arr2[i] - i
	*/
	for i := range arr1 {
		aMax = math.Max(float64(arr1[i]+arr2[i]+i), aMax)
		bMax = math.Max(float64(arr1[i]+arr2[i]-i), bMax)
		cMax = math.Max(float64(arr1[i]-arr2[i]+i), cMax)
		dMax = math.Max(float64(arr1[i]-arr2[i]-i), dMax)

		aMin = math.Min(float64(arr1[i]+arr2[i]+i), aMin)
		bMin = math.Max(float64(arr1[i]+arr2[i]-i), bMin)
		cMin = math.Max(float64(arr1[i]-arr2[i]+i), cMin)
		dMin = math.Max(float64(arr1[i]-arr2[i]-i), dMin)
	}

	return int(math.Max(math.Max(aMax-aMin, bMax-bMin), math.Max(cMax-cMin, dMax-dMin)))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	newHead := ListNode{}
	p := head
	num := 0
	for p != nil {
		tmp := p.Next
		p.Next = newHead.Next
		newHead.Next = p
		p = tmp
		num++
	}
	res := make([]int, num)
	p = newHead.Next
	for p != nil {
		res = append(res, p.Val)
		p = p.Next
	}
	return res
}

func findPairs(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	note := make(map[int]int)
	var res int
	for i, v := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			note[v]++
			continue
		}
		val, ok := note[i-k]
		if ok && val > 0 {
			res++
		}
		val2, ok := note[i+k]
		if ok && val2 > 0 {
			res++
		}
		note[v]++
	}
	if k == 0 {
		res = 0
		for _, v := range note {
			if v > 1 {
				res++
			}
		}
	}
	return res
}

func quickSort(nums []int, start int, end int) {
	if start > end {
		return
	}
	site := partition(nums, start, end)
	quickSort(nums, start, site-1)
	quickSort(nums, site+1, end)
}

func partition(nums []int, start int, end int) int {
	i, j := start, end
	tmp := nums[start]
	for i < j {
		for i < j && nums[j] >= tmp {
			j--
		}
		for i < j && nums[i] <= tmp {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[start] = nums[start], nums[i]
	return i
}

func tick() func() {
	now := time.Now()
	log.Println("tick time, start:", now)
	return func() {
		log.Println("tick stop cost:", time.Since(now))
	}
}
