package main

import "fmt"

var name string = "卢文浩" //可以只声明不赋值
/*
func getvartype(tin any)(string){
	//func f (输入)(输出类型)，没有输出就只有一个()，没输入则输入()空着
	tin=fmt.SPrintf("%T",tin)
	//这里是把tin的类型赋值给tin，要用=。注意fmt.Printf
	return tin
}
	//错误原因是你不能把一个any变量返回成一个string
*/

func getvartype(tin any) string {
	return fmt.Sprintf("%T", tin)
	//直接返回不赋值就好了
	//Sprintf是用来不输出字符串，赋值变量的
}

func main() {
	age := 22
	gentel := "male"

	school := "bjtu"         //短变量声明法只能用在函数内部，全局变量不行
	a, b, _ := 1, 2, "three" //匿名变量不占命名空间，不分配内存，可以重复使用
	fmt.Printf("my name is %s\nage:%d\ngentel:%s\nschool:%s\n", name, age, gentel, school)
	fmt.Printf("%v,%v\n", a, b)
	nametype := getvartype(name)
	agetype := getvartype(age)
	fmt.Printf("Name type is %s\nAge type is %s", nametype, agetype)
}
