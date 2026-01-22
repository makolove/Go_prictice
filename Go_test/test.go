package main

import (
	"fmt"
	"slices" // Go 1.21+ 必须导入
	"strings"
)

// 定义一个服务器结构体
type Server struct {
	Name   string
	IP     string
	Active bool
}

func main() {
	// 1. 初始化一个服务器切片 (故意打乱顺序)
	servers := []Server{
		{Name: "Web-02", IP: "192.168.1.2", Active: true},
		{Name: "DB-01", IP: "192.168.1.10", Active: false},
		{Name: "Web-01", IP: "192.168.1.1", Active: true},
		{Name: "Cache-01", IP: "192.168.1.5", Active: false},
	}

	fmt.Println("--- 原始数据 ---")
	printServers(servers)

	// 2. 【自定义逻辑排序】：按名称 (Name) 升序排列
	// 我们传入一个匿名函数告诉 SortFunc：对比两个服务器时，看它们的 Name
	slices.SortFunc(servers, func(a, b Server) int {
		return strings.Compare(a.Name, b.Name)
	})
	fmt.Println("\n--- 按名称排序后 ---")
	printServers(servers)

	// 3. 【自定义逻辑删除】：删掉所有 Active 为 false 的服务器
	// 这里的匿名函数就是“自定义逻辑变量”，它决定了谁该被删掉
	servers = slices.DeleteFunc(servers, func(s Server) bool {
		return s.Active == false // 如果返回 true，这个元素就会被删掉
	})
	fmt.Println("\n--- 删掉关机服务器后 ---")
	printServers(servers)

	// 4. 【基础函数演示】：检查是否包含某个 IP (仅演示简单切片)
	ips := []string{"192.168.1.1", "192.168.1.2"}
	hasIP := slices.Contains(ips, "192.168.1.1")
	fmt.Printf("\n检测 IP 192.168.1.1 是否在列表中: %v\n", hasIP)
}

// 辅助函数：打印服务器列表
func printServers(ss []Server) {
	for _, s := range ss {
		status := "运行中"
		if !s.Active {
			status = "已关机"
		}
		fmt.Printf("[%s] IP: %-12s 状态: %s\n", s.Name, s.IP, status)
	}
}
