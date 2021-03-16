package main

import (
	"fmt"
	"sync"
	"time"
)

//todo map 如何实现key不存在get操作等待直到 key存在或者超时，保证并发安全


//todo 想法 ：协程阻塞想到channel  并发安全想到锁

type Map struct {

	c map[string]*entry
	rmx *sync.RWMutex
}

type entry struct {
	ch chan struct{}
	value interface{}
	isExist bool
}


//要求存入key/value,如果该key读取的goroutine挂起则唤醒，此方法不会阻塞，时刻都可以执行并返回
func (m *Map)Out(key string,val interface{}){

	m.rmx.Lock()
	defer m.rmx.Unlock()

	if e,ok := m.c[key];ok{
		e.value = val
		e.isExist = true
		close(e.ch)
	}else {
		// e = nil
		e = &entry{ch:make(chan struct{}),isExist:true}
		m.c[key] = e
		close(e.ch)
	}
}

// 读取一个key,如果不存在则阻塞，等待key存在或者超时

func (m *Map)Rd(key string,timeout time.Duration)interface{} {

	m.rmx.Lock()
	if e,ok := m.c[key];ok && e.isExist{
		m.rmx.Unlock()
		return e.value
	}else if !ok{
		//不存在  e = nil
		e = &entry{ch:make(chan struct{}),isExist:false}
		m.c[key] = e
		m.rmx.Unlock()

		fmt.Println("协程阻塞 ->",key)
		select {
		case <-e.ch:  //读取关闭的通道 读到零值
			return e.value
		case <-time.After(timeout):
			fmt.Println("协程超时 ->",key)
			return nil
		}
	}else {
		m.rmx.Unlock()
		fmt.Println("协程阻塞 ->",key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			fmt.Println("协程超时 ->",key)
			return nil
		}
	}
}