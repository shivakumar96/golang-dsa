package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan *int)
	wg := sync.WaitGroup{}
	arr := []int{1, 2, 3, 4, 5}

	wg.Add(3)
	//producer
	go func() {
		for i := range arr {
			ch <- &arr[i]
		}
		close(ch)
		wg.Done()
	}()

	//consumer
	go func() {
		for val, open := <-ch; open; {
			fmt.Println(*val)
			val, open = <-ch
		}
		wg.Done()
	}()
	go func() {
		for val := range ch {
			fmt.Println(*val)
		}
		wg.Done()
	}()

	wg.Wait()
}
