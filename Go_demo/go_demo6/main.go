package main

import (
	"fmt"
)

// float，int类型要同样位数才可以相加，转化：folat64(变量名)
// go中默认int，float和电脑位数一样
//alt+shift+a 一键块注释
//Sprintf强制转化成字符串,且不保存，所以要和赋值或者返回连在一起用
//go中++，--属于语句不是运算符只能单独使用，且不能放前面，a:=i++是不对的
//注意字符串一定要加""
// && || ! 与或非，^异或注意&&和&的区别，&是2进制位与
//<<左移>>右移
//比较运算符返回的是布尔变量，注意赋值运算符用法：+=，/=，%=。。。。。

/*
func main() {
	a := 10
	b := 0.1
	fmt.Println(float64(a) + b)
}
*/

var a int = 10

func test() bool {
	fmt.Println("test on")
	return true
}

/*
	func main() {
		if a >= 11 && test() {
			fmt.Println("pass")
		}
		//当&&不成立的时候，与门短路，与的函数不执行
	}
*/
func main() {
	if a >= 9 || test() {
		fmt.Println("pass")
	}
	//当||成立的时候，或门短路，或的函数不执行也可以通，那就不执行
}
