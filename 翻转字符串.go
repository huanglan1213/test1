package main

import "fmt"

func main(){

	str := "abcda"
	s,_ := reverse(str)
	fmt.Println(s)


	i := 7

	s1 := "hello 你好"

	fmt.Println(s1[i])

	fmt.Println(string(s1[i]))

	fmt.Println([]rune(s1)[i])

	fmt.Println(string([]rune(s1)[i]))


}


func reverse(str string) (string,bool) {

	s := []rune(str)
	l := len(s)

	if l > 5000 {
		return "",false
	}
	i := 0
	for {
		if i >= l/2 {
			break
		}
		s[i],s[l-1-i] = s[l-1-i],s[i]
		i++
	}
	return string(s),true
}

