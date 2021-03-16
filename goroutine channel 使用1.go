package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
/*
	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {

		defer  wg.Done()
		for i:= 0;i<5;i++{
			ch <- rand.Intn(10)
		}
		close(ch)

	}()

	go func(){
		for i := range ch{
			fmt.Println(i)
		}
		defer  wg.Done()
	}()
	wg.Wait()
*/
	test1()

}

func test1(){


	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1 )
	go func() {
		//设置种子数
		rand.Seed(time.Now().UnixNano())
		for i:= 0;i<5;i++{
			ch <- rand.Intn(5)
		}
		defer close(ch)
		wg.Done()
	}()

	go func() {

		for {
			select {
			case i := <-ch:
				fmt.Println(i)
			}
		}
	}()


	wg.Wait()




}

