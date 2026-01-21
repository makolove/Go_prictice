package main

import "fmt"

//C=(F-32)/1.8,输入F，求C
var F float64
var C float64

/* func main() {
	fmt.Println("请输入华氏度")
	fmt.Scanf("%f", &F) //Scanf使用要告知输入类型，指针
	C = (F - 32) / 1.8
	fmt.Printf("摄氏度：%f", C)
}
//输入不是float也会输出，当scanf输入不合法，F会保存为0，tip：err变量nil(空，零值)，非nil两种
*/

func main() {
	fmt.Println("请输入华氏度（请输入数字）：")
	for {
		// Scan会返回成功读取的数量和是否有错误，不要求格式化读
		_, err := fmt.Scan(&F)
		if err != nil {
			// 如果出错了（比如输入了 abc）
			fmt.Println("请输入温度！")
			// 重要：清空输入缓冲区
			// 如果不清理，缓冲区里的 "abc" 会一直留在那里，导致无限循环报错
			var dump string
			fmt.Scanln(&dump)
			continue // 跳过本次循环，重新开始
		}
		// 如果没有错误，跳出循环
		break
	}
	C = (F - 32) / 1.8
	fmt.Printf("华氏度 %.2f 对应的摄氏度为：%.2f\n", F, C)
}

// fmt.Scan 和 fmt.Scanln 的主要区别：
//
// 1. 分隔符：
//    - Scan:   会将“空格”和“回车”都当作分隔符。如果你读两个数，中间敲回车或空格没区别。
//    - Scanln: 只将“空格”当作分隔符，一旦遇到“回车”就会立刻停止读取。
//
// 2. 缓冲区的处理（最关键）：
//    - Scan:   在读取完目标数据后，可能会把最后那个“回车符”留在缓冲区里。
//    - Scanln: 读取到行尾时，会顺手把最后的“回车符”也从缓冲区里带走并消耗掉。
//
// 结论：当你需要“清空这一行”或者“确保下次读取不受干扰”时，首选 Scanln。
