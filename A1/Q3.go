package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func RandomGenerator(wg *sync.WaitGroup, stop <-chan bool, m int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer func() { wg.Done() }()
		defer close(intStream)
		for {
			select {
			case <-stop:
				return
			case intStream <- rand.Intn(1000000) * m:
			}
		}
	}()
	return intStream
}

func Mutiple(x int, m int) bool {
	if x%m == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	stop := make(chan bool, 3)
	m5 := RandomGenerator(&wg, stop, 5)
	count5 := 0
	m13 := RandomGenerator(&wg, stop, 13)
	count13 := 0
	m97 := RandomGenerator(&wg, stop, 97)
	count97 := 0
	// three count variable & generator for each number
	for i := 0; i < 100; i++ {
		var i int
		select {
		case i = <-m5:
		case i = <-m13:
		case i = <-m97:
		}
		//run the generator

		switch {

		case Mutiple(i, 5):
			{
				count5++
			}

		case Mutiple(i, 13):
			{
				count13++
			}
		case Mutiple(i, 97):
			{
				count97++
			}
		}
		//check the Multiple for each situation
	}

	fmt.Printf("total number of generated multiples of 5: %d\n", count5)
	fmt.Printf("total number of generated multiples of 13: %d\n", count13)
	fmt.Printf("total number of generated multiples of 97: %d\n", count97)

	var j int = 0
	for j < 3 {
		stop <- true
		j++
	}

	wg.Wait()

}
