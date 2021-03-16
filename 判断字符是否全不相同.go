package main

import (
	"fmt"
	"strings"
)

func main(){


	str := "asdfghjklqwertyuiopzxcvbnm,./;'q"


	fmt.Println(judge(str))
}

func judge(str string) bool {

	if strings.Count(str,"") > 3000 {
		return false
	}

	for _,v := range str {
		if v > 127 {
			return false
		}
		if strings.Count(str,string(v)) >1 {
			return false
		}
	}
	return true





}