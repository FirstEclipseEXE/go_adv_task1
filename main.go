package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	code := make(chan int)
	mode := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func ()  {
		arr(code)
		wg.Done()
	}()
	wg.Add(1)
	go func ()  {
		cube(code, mode)
		wg.Done()
	}()
	go func ()  {
		wg.Wait()
		close(code)
		close(mode)
	}()
	for range 10 {
		num := <- mode
		fmt.Println(num) 
	}

}


func arr(code chan int) {
	arr := [10]int{}
	for i := range arr {
        arr[i] = rand.Intn(100)
		code <- arr[i]
    }
}

func cube(code, mode chan int) {
	for range 10 {
		x := <-code
		mode <- x*x 
	}
}