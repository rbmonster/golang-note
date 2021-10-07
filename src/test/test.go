package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
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

func (w Wheel) Write(p []byte) (n int, err error) {
	return 12, nil
}

// teetsetes
func main() {
	//defer trace("main trace")
	//var done = make(chan int)
	//for i := 0; i < 100; i++ {
	//	finish++
	//	go crawl(done, strconv.Itoa(i))
	//}
	//
	//for ;finish > 0; finish-- {
	//	var res = <-done
	//	fmt.Println(strconv.Itoa(res) + "done")
	//}
	//selectSend()
	//
	//var mu sync.Mutex
	//mu.Lock()
	//defer mu.Unlock()
	//// unlock
	//mu.Unlock()
	//
	//
	//var mw sync.RWMutex
	//
	//mw.RLock() // 读锁
	//mw.RUnlock()
	//mw.Lock() // 写锁
	//mw.Unlock()

	//var once sync.Once
	//
	//var f = func() {
	//	fmt.Println("init")
	//}
	//once.Do(f)
	//
	//once.Do(f)
	fmt.Print("hello world")

	var x int
	fmt.Println(reflect.TypeOf(x))
}

type entry struct {
	ready chan int
	value int
}

var syn sync.WaitGroup

func selectSend() {
	//sync.WaitGroup{}
	ch := &entry{ready: make(chan int)}
	for i := 0; i < 4; i++ {
		go sub(ch)
	}
	ch.value = 3
	ch.ready <- 3
	close(ch.ready)
	syn.Wait()
}

func sub(ch *entry) {
	syn.Add(1)
	<-ch.ready
	fmt.Println(ch.value)
	fmt.Println("finish")
	syn.Done()
}

var tokens = make(chan struct{}, 20)
var finish = 0

func crawl(done chan<- int, msg string) {
	//defer
	defer trace(msg + " crawl message")()
	tokens <- struct{}{}
	time.Sleep(1 * time.Second)
	<-tokens
	atoi, err := strconv.Atoi(msg)
	if err != nil {
		return
	}
	done <- atoi
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}
