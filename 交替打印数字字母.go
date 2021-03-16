package main

// TODO chanel 交替输出 循环打印

import (
	"fmt"
	"strings"
	"sync"
)

func main(){

	var wg sync.WaitGroup
	DogChan := make(chan struct{})
	CatChan := make(chan struct{})
	PigChan := make(chan struct{})

	num := 10

	wg.Add(1)
	go func(){
		i := 0
		for {
			select {
			case <-DogChan:
				if i > num -1 {
					wg.Done()
					return
				}
				fmt.Println("Dog")
				i++
				CatChan <- struct{}{}
				break
			}
		}

	}()
	go func(){
		for {
			select {
			case <-CatChan:
				fmt.Println("Cat")
				PigChan <- struct{}{}
				break
			}
		}
	}()
	go func(){
		for {
			select {
			case <-PigChan:
				fmt.Println("Pig")
				DogChan <- struct{}{}
				break
			}
		}
	}()
	DogChan <- struct{}{}
	wg.Wait()
}

// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
func CircleNumberAndLetter(){

	var wg sync.WaitGroup
	letter := make(chan bool)
	number := make(chan bool)

	go func(){
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()
	wg.Add(1)
	go func(){
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0

		for {
			select{
			case <- letter:

				//todo 判断越界
				if i >= strings.Count(str,"") -1 {
					wg.Done()
					return
				}
				fmt.Print(str[i:i+1])
				i++
				fmt.Print(str[i:i+1])
				i++
				number <- true
				break
			default:
				break
			}
		}
	}()
	number <- true
	wg.Wait()
}
