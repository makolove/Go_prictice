package main

import "fmt"

// 定义一个自定义函数
func sayHi(name string) {
    fmt.Println("Hi,", name)
}

func main() {
    // 调用上面定义的函数
    sayHi("Golang666") 
}