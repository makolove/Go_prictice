package main

//const类型定义的时候就必须赋值
//iota是行计数器，只能用在const中,const出现的时候iota置为0
import "fmt"

func main() {
	const (
		a = iota
		_
		c
		d = iota
	)
	const e = iota
	fmt.Println(a, c, d, e)
	const (
		n1, n2 = iota + 1, iota + 2
		n3, n4 = iota + 1, iota + 2
		n5, n6 = iota + 1, iota + 2
	)
	fmt.Println(n1, n2, n3, n4, n5, n6) //Println输出自动换行，直接输出变量，Printf要先格式化
}
