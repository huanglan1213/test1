package main

import (
	"fmt"
	"strings"
)

func main(){

	fmt.Println(isReGroup("assda","asdsa"))

}


func isReGroup(s1,s2 string)bool{

	l1 := len([]rune(s1))
	l2 := len([]rune(s2))

	if l1 > 5000 || l2 >5000 || l1 != l2  {
		return false
	}

	for _,v1 := range s1 {
		if strings.Count(s1,string(v1)) != strings.Count(s2,string(v1)){
			return false
		}
	}
	return  true
}
