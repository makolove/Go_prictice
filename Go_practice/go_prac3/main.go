package main

import (
	"fmt"
	"math/rand"
	"slices"
)

// Go语言中map的定义语法如下
// map[KeyType]ValueType

// 从value查找key，采用泛型写法
func v2k[K comparable, V comparable](m map[K]V, target V) (K, bool) {
	for k, v := range m {
		if v == target {
			return k, true
		}
	}
	var zero K
	return zero, false // 如果函数头只写了 K，这里就会报错
}
func main() {
	/* userInfo := map[string]string{
		"username1": "LWH",
		"password1": "lwh20216.",
		"username2": "ZJY",
		"password2": "ZJYKAWAYI",
	}
	fmt.Println(userInfo)
	//map自带的查找语法，从key查value
	_, ok := userInfo["username1"]
	if ok {
		fmt.Println("value exist")
	} else {
		fmt.Println("value empty")
	}
	findkey, flag := v2k(userInfo, "ZJY2")
	if flag {
		fmt.Printf("找到ZJY对应的key:%v", findkey)
	} else {
		fmt.Println("error")
	} */

	scoreMap := make(map[string]int, 200)
	for i := 0; i < 10; i++ {
		//key:=strconv.Itoa(i)
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}

	//map在保存的时候是无序随机的，并不会根据key的大小排序的，println会自动排序
	fmt.Println(scoreMap)
	//真实情况
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	// 取出map中的所有key存入切片keys
	keys := make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	// 对切片进行排序
	slices.Sort(keys)
	// 按照排序后的key遍历map
	//这里的key实际上是keys切片的value
	//为什么一直用：，因为短变量声明只作用在局域上（所以每个for要在再定义一次）
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
