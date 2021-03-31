package main

import (
	"context"
	"fmt"
	"time"
)

//todo context 程序上下文


//todo 1. 通过 cancel 主动关闭
func ctxCancel(){

	ctx,cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		select {


		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case <-time.After(10*time.Millisecond):
			fmt.Println("	time out ")
		}
	}(ctx)

	cancel()
}

// todo 2. 通过超时，自动触发
func ctxTimeout(){

	ctx,cancel := context.WithTimeout(context.Background(),time.Second*10)  //ctx 10s到期

	//主动执行cancel,也会让协程收到消息
	defer cancel()
	go func(ctx context.Context) {
		select {
		//使用select调用<-ctx.Done()判断是否要结束，如果接受到值的话，就可以返回结束goroutine了；如果接收不到，就会继续进行监控。
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case <-time.After(5*time.Second):
			fmt.Println("	time out ")
		}
	}(ctx)

}
//todo 使用场景
// 打开外卖的订单页，地图上显示外卖小哥的位置，而且是每秒更新 1 次。app 端向后台发起 websocket 连接（现实中可能是轮询）请求后，后台启动一个协程，
// 每隔 1 秒计算 1 次小哥的位置，并发送给端。如果用户退出此页面，则后台需要“取消”此过程，退出 goroutine，系统回收资源。
func main(){
	ctx,cancel := context.WithTimeout(context.Background(),time.Hour)
	go perform(ctx)
	//app 端返回页面，调用cancel 函数
	cancel()
}

func perform(ctx context.Context) {
	for {
		calculatePos() //计算位置
		sendResult()  //发送到app

		select {
		case <-ctx.Done():
			// 被取消，直接返回
			return
		case <-time.After(time.Second):
			// block 1 秒钟
		}
	}
}

func calculatePos(){}

func sendResult(){}

// todo 3. 传递共享参数
func ctxWithValue(ctx context.Context) {
	traceId, ok := ctx.Value("traceId").(string)
	if ok {
		fmt.Printf("process over. trace_id=%s\n", traceId)
	} else {
		fmt.Printf("process over. no trace_id\n")
	}
}

func main1() {
	ctx := context.Background()
	ctxWithValue(ctx)

	ctx = context.WithValue(ctx, "traceId", "qcrao-2019")
	ctxWithValue(ctx)
}


