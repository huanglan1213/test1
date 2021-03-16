package main

import "fmt"

func main(){

	n := 0
	f := func() int {

		n = n + 1
		return n
	}
	fmt.Println(f())

	fmt.Println(f())

	fmt.Println(f1(0)())
	fmt.Println(f1(0)())

}


func f1(i int) func() int{

	return func() int {
		i++
		return i
	}
}
