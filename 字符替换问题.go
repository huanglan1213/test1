package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main(){


	s,ok := replace("da在dsa")
	if !ok {
		fmt.Println(ok)
	}
	fmt.Println(s)

}
/*
请编写一个方法，将字符串中的空格全部替换为“%20”。
假定该字符串有足够的空间存放新增的字符，并且知道字符串的真实长度(小于等于1000)，同时保证字符串由【大小写的英文字母组成】。
给定一个string 为原始的串，返回替换后的string。
 */

func replace(s string)(string,bool){

	l := len([]rune(s))
	if l > 1000 {
		return s,false
	}

	for _,v := range s {

		// 汉字也是一个字母字符，因此需要特殊判断
		if unicode.Is(unicode.Han,v){
			fmt.Println(string(v))
			return s,false
		}

		if string(v) != " " && !unicode.IsLetter(v){
			fmt.Println(string(v))
			return s,false
		}
	}
	return strings.Replace(s," ","%20",-1),true

}
