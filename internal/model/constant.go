package model

// TODO:
// 指针切片类型
// 切片指针类型
// 。。。

// 类型枚举值
const (
	Unknow = iota

	// ----- 普通类型 -----
	// 基础类型
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	String
	Bool

	// 当前包内类型
	Pkg_Type

	// 非当前包的类型
	Pke_Type_3

	// ----- 指针类型 -----
	// 基础类型的指针类型
	Ptr_Int
	Ptr_Int8
	Ptr_Int16
	Ptr_Int32
	Ptr_Int64
	Ptr_Uint
	Ptr_Uint8
	Ptr_Uint16
	Ptr_Uint32
	Ptr_Uint64
	Ptr_Float_32
	Ptr_Float_64
	Ptr_String
	Ptr_Bool

	// 当前包内的
	Ptr_Pkg_Type

	// 其他包的
	Ptr_Pke_Type_3

	// ----- 切片/数组类型 -----
	// 基础数据类型
	Arr_Int
	Arr_Int8
	Arr_Int16
	Arr_Int32
	Arr_Int64
	Arr_Uint
	Arr_Uint8
	Arr_Uint16
	Arr_Uint32
	Arr_Uint64
	Arr_Float32
	Arr_Float64
	Arr_String
	Arr_Bool

	// 当前包内的
	Arr_Pkg_Type

	// 其他包的
	Arr_Pke_Type_3
)
