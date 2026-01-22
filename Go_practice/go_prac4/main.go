//slices包
// --- 查找与检查 ---
// slices.Contains(s, v): 判断切片 s 是否包含值 v (返回 bool)
// slices.Index(s, v): 返回值 v 在 s 中第一次出现的索引，未找到返回 -1
// slices.BinarySearch(s, v): 在【已排序】切片中二分查找 v (返回 索引, 是否找到)
// --- 元素操作 (会改变原切片内容) ---
// s = slices.Insert(s, i, v...): 在索引 i 处插入一个或多个元素，返回新切片
// s = slices.Delete(s, i, j): 删除索引 [i, j) 范围内的元素，返回新切片
// s = slices.Replace(s, i, j, v...): 将 [i, j) 范围替换为新元素 v...，返回新切片
// s = slices.Compact(s): 将【连续重复】的元素收缩为一个（常配合 Sort 使用去重）
// --- 排序与顺序 ---
// slices.Sort(s): 对切片进行升序排序 (直接修改原切片)
// slices.Reverse(s): 将切片元素顺序反转 (直接修改原切片)
// slices.IsSorted(s): 检查切片是否已经是升序排列 (返回 bool)
// --- 比较与极值 ---
// slices.Equal(s1, s2): 检查两个切片长度和对应位置元素是否完全相等
// slices.Compare(s1, s2): 字典序比较两个切片大小 (返回 -1, 0, 或 1)
// slices.Max(s): 返回切片中的最大值 (切片不能为空)
// slices.Min(s): 返回切片中的最小值 (切片不能为空)
// --- 内存管理 ---
// slices.Clone(s): 创建并返回切片的一个副本（浅拷贝）
// s = slices.Grow(s, n): 扩充容量，确保至少能额外容纳 n 个元素
// s = slices.Clip(s): 移除未使用的容量，使 cap 等于 len
//可以在参数里加一个自定义的方法函数f，比如slices.ContainsFunc(s, f): 使用自定义逻辑 f 判断是否包含满足条件的元素

package main

import (
	"fmt"
	"slices"
)

// 定义一个支持加法的类型约束
type Number interface {
	int | int64 | float64
}

// 实现n个int和float之间的相加，不用强制转换类型
// 函数返回值如果命名必须写在()里，输入不能光有类型，还要赋值
// 泛型[类型，接口]
func Sum[T Number](X ...T) (result T) {
	for _, v := range X {
		result += v
	}
	return // 自动返回 total
}

// F2一个变量名字实现全局改写
func Mult[T Number](X ...T) T {
	var result T = 1
	for _, v := range X {
		result *= v
	}
	return result
}

type Cal[T Number] func(X ...T) T

// 函数作为函数输入的意义
// 封装了一个函数输入Sum/mult实现累加或者累乘，n个数值型变量
// cumulative add cumulative mult
func CA_CM[T Number](f Cal[T], X ...T) T {
	return f(X...)
}

func SortMapKey(m1 map[string]string) (sortedkey []string) {
	for k := range m1 {
		sortedkey = append(sortedkey, k)
	}
	slices.Sort(sortedkey)
	return
}

func main() {
	var inputNums []float64 // 统一用 float64 接收，因为它兼容性最强
tabel:
	fmt.Println("请输入数字（输入非数字结束）：")
	for {
		var val float64
		// Scan 会返回成功读取的个数,数字之间用空格，tab，回车隔开。当读到字母或者文件结尾返回非空
		_, err := fmt.Scan(&val)
		if err != nil {
			break
		}
		inputNums = append(inputNums, val)
	}
	// 检查用户是否输入了数据
	if len(inputNums) == 0 {
		fmt.Println("未检测到有效输入，请重新输入")
		goto tabel
	}

	fmt.Println("\n--- 计算结果 ---")
	// 注意：因为 CA_CM 需要解构切片，必须使用 inputNums...
	// 显式指定类型参数 [float64] 确保类型安全
	fmt.Printf("累加和 (Sum): %.3f\n", CA_CM(Sum[float64], inputNums...))
	fmt.Printf("累乘积 (Mult): %.3f\n", CA_CM(Mult[float64], inputNums...))
}
