package main

import (
	"fmt"
	"reflect"
)

//todo 反射的用法


//todo 1. 获取变量内部信息

func main1(){

	str := "this is a test"

	t := reflect.TypeOf(str)

	fmt.Println(t.Name())

	v := reflect.ValueOf(str)

	fmt.Println(v)
}


type Student struct {
	Id   int
	Name string
}
func (s Student) EchoName(name string){
	fmt.Println("我的名字是：", name)
}

//todo 2. 修改变量  必须是指针类型
func main(){
	s := &Student{Id: 1, Name: "咖啡色的羊驼"}
	v := reflect.ValueOf(s)
	// 修改值必须是指针类型否则不可行
	if v.Kind() != reflect.Ptr {
		fmt.Println("不是指针类型，没法进行修改操作")
		return
	}
	// 获取指针所指向的元素
	v = v.Elem()
	// 获取目标key的Value的封装
	name := v.FieldByName("Name")
	if name.Kind() == reflect.String {
		name.SetString("小学生")
	}
	fmt.Printf("%#v \n", *s)

	//todo 整型的话
	i := 2
	v1 := reflect.ValueOf(&i)
	v1.Elem().SetInt(3)
	fmt.Println(i)
}

//todo 3. 通过反射调用方法
func main3(){
	s := Student{Id: 1, Name: "咖啡色的羊驼"}
	v := reflect.ValueOf(s)
	// 获取方法控制权
	// 官方解释：返回v的名为name的方法的已绑定（到v的持有值的）状态的函数形式的Value封装
	mv := v.MethodByName("EchoName")
	// 拼凑参数
	args := []reflect.Value{reflect.ValueOf("咖啡色的羊驼")}
	// 调用函数
	mv.Call(args)
}


