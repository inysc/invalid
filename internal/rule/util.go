package rule

import "strings"

// 获取类型分类
//	[0]: 基础类型 int string float32...
//	[1]: 本包类型 当前包内定义的结构体，重命名的类型
//	[2]: 它包类型 别的包内定义的结构体，重命名的类型
//	[3]: 基础类型指针
//	[4]: 本包类型指针
//	[5]: 它包类型指针
//	[6]: 基础类型切片
//	[7]: 本包类型切片
//	[8]: 它包类型切片
// TODO: 指针
func KindOfType(typ string) int {

	return 0
}

// 移除默认值标记
func mvDefault(str string) string {
	return strings.TrimPrefix(str, "d")
}

func If[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func getr[T any](t T) *T {
	return &t
}
