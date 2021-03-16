package main

import (
	"fmt"
	"time"
)

func main(){

	go func() {

		//todo 定时任务
		t := time.NewTicker(2*time.Second)

		for {

			 <-t.C
				// 每秒执行
				func() {

					//todo panic 恢复执行 recover
					defer func() {
						if err := recover();err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()

		}

	}()



	select {
	}
}

func proc(){
	panic("ok")
}

func main2(){

	t := time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case <- t.C:
				go func() {
					fmt.Println("hello world")
				}()
			}
		}
	}()

	select {
	case <-time.After(10*time.Second):
		break
	}
}