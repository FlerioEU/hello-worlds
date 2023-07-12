package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/exp/slices"
)

var numbers = []int{1, 4, 7}

// used to lock the seed for random integers
// so rand.New(s) does not create the same number in each goroutine
// rand.NewSource() is NOT threadsafe
var mu sync.Mutex
var seed = rand.NewSource(time.Now().Unix())

func main() {
	wg := sync.WaitGroup{}
	queue := make(chan int)

	for _, n := range numbers {
		wg.Add(1)

		go func(v int) {
			defer wg.Done()

			// lock the seed so we are not using the same seed value
			// in each goroutine
			mu.Lock()
			s := seed
			mu.Unlock()

			v = v + rand.New(s).Intn(100)
			queue <- v
		}(n)
	}

	go func() {
		wg.Wait()
		close(queue)
	}()

	results := []int{}
	for v := range queue {
		results = append(results, v)
	}

	slices.Sort(results)
	fmt.Printf("results: %v\n", results)

	fmt.Println("waiting for process to finish")
	<-blockingFn()
	fmt.Println("finished! :)")
}

func blockingFn() <-chan struct{} {
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 5)

		close(ch)
	}()

	return ch
}
