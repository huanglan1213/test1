package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"context"
)

func main(){


	s := getSlice(1000)
	dataLen := len(s)
	target := 15
	size := 5

	res_ch := make(chan struct{})
	//todo 找到目标值或者超时后立刻结束所有goroutine的执行---- 需要借助计时器、通道和context才行
	//todo context.WithCancel创建一个上下文对象传递给每个执行任务的goroutine，外部在满足条件后（找到目标值或者已超时）通过调用上下文的取消函数来通知所有goroutine停止工作。

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < dataLen; i += size {
		end := i + size
		if end >= dataLen {
			end = dataLen - 1
		}
		go search(ctx, s[i:end], target, res_ch)
	}

	select {
	case <-res_ch:
		fmt.Fprintf(os.Stdout, "Found it!\n")
		cancel()
	case <- time.After(time.Second*5):
		fmt.Fprintf(os.Stdout, "Timeout! Not Found")
		cancel()
	}
	time.Sleep(time.Second * 2)
}

func search(ctx context.Context, data []int, target int, resultChan chan struct{}){

	for _,v := range data {
		select {
		case <-ctx.Done():
			fmt.Fprintf(os.Stdout, "Task cancelded! \n")
		default:
		}
		// 模拟一个耗时查找，这里只是比对值，真实开发中可以是其他操作
		fmt.Fprintf(os.Stdout, "v: %d \n", v)
		time.Sleep(time.Millisecond * 1500)
		if v == target {
			resultChan <- struct{}{}
		}
	}

}

func getSlice(num int)[]int{

	s := make([]int,0)

	for i := 0 ; i < num ; i++ {
		rand.Seed(time.Now().UnixNano())
		if i == 55 {
			s = append(s,15)
		}
		s = append(s, rand.Intn(100))
	}
	return s
}


