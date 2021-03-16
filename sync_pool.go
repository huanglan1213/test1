package main

import (
	"fmt"
	"sync"
)

//TODO  保存和复用临时对象

func main(){

	p := &sync.Pool{
		New: func() interface{} {
			return "hello"
		},
	}


	a := p.Get()
	p.Put("kko")
	p.Put("kko1")
	b := p.Get()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println( p.Get())
	fmt.Println( p.Get())
	fmt.Println( p.Get())
	fmt.Println( p.Get())

}



