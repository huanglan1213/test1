package main

import (
	"fmt"
	"runtime"
	"time"
)

func main1(){


	ch := make(chan struct{})

	go func() {

		select {
		case i := <-ch:
			fmt.Println(i)
		case <- time.After(time.Second):
			fmt.Println("超时了")
		}
	}()

	close(ch)

	time.Sleep(2*time.Second)
}

//0AB12CD34EF56G  或者 A01BC23DE45FG6
func main2(){
	ch := make(chan int)
	go func() {
		s := "ABCDEFG"

		for i:=0;i<len(s);i++{
			ch <-i
			fmt.Print(string(s[i]))
		}
	}()
	go func() {
		for i:=0;i<26;i++{
			<- ch
			fmt.Print(i)
		}
	}()
	time.Sleep(time.Second)
}

func main(){


	runtime.GC()

}