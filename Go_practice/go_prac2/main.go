package main

import (
	"fmt"
	"slices"
	"strconv"
)

//数组，切片，map（映射）
//复合数据类型
//函数
//指针，结构体
//package

//数组定义：arr1 := [...]int{1, 3, 5, 7, 8}，注意数组的长度属于数据类型，[3]int,[4]int是两种数据类型
//基本数据类型和数组都是值类型。切片是引用类型，slice1:=slice2共享指针,副本变了本身也变
//var arr1 [10][10]int  2维数组,注意函数体外不可以:=

//切片：var name=[]type{}
//切片的caps是从当前元素到底层数组最后一个元素个数，caps>=len
//切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。
//make([]T, size, cap),已知切片的内容不用make，make创建后只能对单个位置赋值，name[0]=
//不能说一个长度和容量都是0的切片一定是nil, 要判断一个切片是否是空的，要是用len(s)==0来判断

var arr1 = [2][2]int{{1, 2}, {2, 4}}
var arr2 = [2]string{"suzhou", "chengdu"}

/*
func main() {
	arr1 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(arr1); i++ {
		for j := i; j < len(arr1); j++ {
			if arr1[i]+arr1[j] == 8 {
				fmt.Printf("%d和%d之和=8，索引：%v,%v\n", arr1[i], arr1[j], i, j)
			}
		}
	}
}
*/

// Go 语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。代码如下：
func delete() {
	// 从切片中删除元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[30 31 33 34 35 36 37]
}

func main() {
	// copy()复制切片
	a := []int{1, 2, 3, 4, 5}

	//错误声明
	//var b []int，这样声明，b是nil的切片，必须先append
	//b = append(b, a...)
	//append会自动扩容，也就是说允许超过最开始定义的容量

	//冗余警报，len和caps长度一样的时候只写一个就好
	b := make([]int, 5)

	//使用copy()函数将切片a中的元素复制到切片c
	//copy不会扩容，所以不能copyz到一个nil切片
	copy(b, a)
	b[0] = 1000
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(b) //[1000 2 3 4 5]
	//注意copy和直接赋值的区别，直接赋值会共用指针
	c := a
	c[0] = 1000
	fmt.Println(a)
	fmt.Println(c)

	var A = make([]string, 5, 10)
	//append 函数的作用是在切片的 尾部（即当前 len 之后）追加元素,所以最后输入会有五个空格
	println(len(A))
	for i := 0; i < 12; i++ {
		//注意切片A的类型是string，所以这里不能直接append一个i
		A = append(A, strconv.Itoa(i))
	}
	fmt.Println(A)

	//slice包
	var num = []int{3, 7, 5, 9, 1}
	slices.Sort(num)
	fmt.Println(num)

}
