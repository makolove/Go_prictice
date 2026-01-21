package main

import "fmt"

//go中没有while

var n int
var result int

//一.求阶乘
/* fmt.Println("输入n")
_, err := fmt.Scanf("%v", &n)
if err != nil {
	fmt.Println("请输入数字！")
	var dumb string
	fmt.Scanln(&dumb)
} else {
	fmt.Scanln(&n)
	result = n
	for ; n >= 2; n-- {
		result *= n - 1
	}
}
fmt.Printf("阶乘的结果是%v", result) */

//画正方形（循环嵌套）
/* row := 4
column := 4
for i := 1; i <= row; i++ {
	for j := 1; j <= column; j++ {
		fmt.Print("*")
	}
	fmt.Print("\n") */

//for range 是Go语言中专门用来遍历各种容器对象的“万能钥匙”,可用于字符串，切片，映射等等所有
//range返回索引和值

func main() {
	for i := 1; i <= 2; i++ {
		for j := 0; j <= 9; j++ {
			if j == 2 {
				continue
			}
			fmt.Println(i, j)
		}

	}
	str1 := "I love My grandmother"
	for k, v := range str1 {
		fmt.Printf("key is %d,value is %c\n", k, v)
	}
}

//golang中的switch，case语句不写break也行，case可以加条件语句
//fallthrough 可以穿透到当前case的下一个case
//break是跳出当前循环的，可以break label，即跳出到label不执行
//continue是跳过当前循环，进入下次循环,continue label同理，跳到label继续执行
//goto label 跳到label
//label:  定义label，名字随意
