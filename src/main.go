package main

import (
	"fmt"
	"sync"
)

var WG = sync.WaitGroup{}

func main() {

	WG.Add(2)
	go func() {
		i := 2

		fmt.Println(i)
		WG.Done()
	}()
	go func() {
		WG.Done()
	}()
	WG.Wait()
}

// func main() {
// 	t := time.Now()
// 	timeToDieInMilliseconds := 2000

// 	time.Sleep(2 * time.Second)
// 	x := time.Since(t) + time.Duration(timeToDieInMilliseconds)*time.Millisecond
// 	fmt.Printf("%v, %T", x, x)
// }
