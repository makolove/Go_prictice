package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "lwh-zjy"
	arr1 := strings.Split(str1, "-")
	//注意调用格式，join,split,contains,HasPrefix(前缀),HasSuffix，Index(从0开始)，LastIndex
	fmt.Println(arr1)
	str2 := strings.Join(arr1, "*")
	fmt.Println(str2)
}
