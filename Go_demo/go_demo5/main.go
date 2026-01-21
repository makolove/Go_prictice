package main

import "fmt"

func main() {
	str1 := "LWHloveZJY"
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c(%v)\n", str1[i], str1[i])
	}
	//别忘了fmt，for循环的部分要用;隔开，字符串定义的时候要""char字符，var变量的ASC
	str2 := "lwh"
	bytestr2 := []byte(str2)
	bytestr2[1] = 'y'
	bytestr2[2] = 'l'
	fmt.Println(string(bytestr2))
}
