package main

import (
	"fmt"
	"log"
	"sync"
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
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(taskNo int) {
			time.Sleep(1 * time.Second)
			log.Println("done for subtask:", taskNo)
			wg.Done()
		}(i)
	}
	defer tick()()
	wg.Wait()
	log.Println("main task start")
	time.Sleep(1 * time.Second)
	log.Println("main task done")
}

func tick() func() {
	now := time.Now()
	log.Println("tick time, start:", now)
	return func() {
		log.Println("tick stop cost:", time.Since(now))
	}
}
