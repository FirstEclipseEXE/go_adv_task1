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
		defer wg.Done()
		close(code)
	}()
	wg.Add(1)
	go func ()  {
		square(code, mode)
		defer wg.Done()
		close(mode)
	}()
	for num := range mode {
        fmt.Println(num)
    }
	wg.Wait()
}


func arr(code chan int) {
	arr := [10]int{}
	for i := range arr {
        arr[i] = rand.Intn(100)
		code <- arr[i]
    }
}

func square(code, mode chan int) {
	for range 10 {
		x := <-code
		mode <- x*x 
	}
}