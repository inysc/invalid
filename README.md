# invalid

结构体字段校验

## 支持规则

| 规则 | 语法 | 说明 | 示例 |
| :-- | :--: | :-- | :-- |
| 禁止 | ! | 字段不可等于所给值 | !nil<非空> |
| 枚举 | {} | 字段需被包含在枚举值中 | {'1', "2"} |
| 区间 | `[` 或 `(` 、 `]` 或 `)` | 数值或时间的区间 | [0, )<非负> |
| 正则 | r'' | 正则匹配 | r'[\u4E00-\u9FA5]'<包含中文> |
| 默认值 | d... | 字段默认值 | d"1"<字符串1> |
| 引用 | <字段> | 指代结构体中的某个字段值 | i.Role.Type |
| 标注 | v | 表明改字段的基础类型实现了接口 `interface { Invalid() error }` | v |

## 规则详细说明

```go
// comparable is an interface that is implemented by all comparable types
// (booleans, numbers, strings, pointers, channels, arrays of comparable types,
// structs whose fields are all comparable types).
// The comparable interface may only be used as a type parameter constraint,
// not as the type of a variable.
// copy $GOROOT/src/builtin/builtin.go
type comparable interface{ comparable }

// copy ./demo/proto/proto.go
type NewReq struct {
	ID    int64
	Score *float64 `iv:"!nil; !20; [1,)"`
	Name  string   `iv:"s(,10]"`
	Phone string
	Role  struct{ Type string }
}
```

### 禁止

字段类型应当是 `comparable` 及其切片。对于指针类型，可以约束其值。 比如 `Score` 字段既约束了指针也约束了其值

### 枚举

字段类型应当是 `comparable` 及其切片。

### 范围

整体是按照数学的写法，使用小括号和中括号控制区间的开闭

 - 数字类型（包含指针）：
 - 切片/字符串类型：前面加上 `l` 表示此范围特指长度
 - 时间类型（字符串）：这里应该有三个数据，依次是最晚时间，时间格式，最早时间

### 正则


### 默认值


### 引用


## 特性

 - 默认值

 * [x] : 普通类型的默认值
 * [x] : 指针类型的默认值

 - 枚举值

 * [x] : 普通类型的枚举值
 * [x] : 普通类型指针的枚举值
 * [x] : 普通类型切片的枚举值
 * [x] : 普通类型切片的指针的枚举值

 - 禁止值

 * [x] : 普通类型的禁止值
 * [x] : 普通类型指针的禁止值
 * [x] : 普通类型切片的禁止值
 * [x] : 普通类型切片的指针的禁止值

 - 标记

 * [x] : 标记普通类型
 * [x] : 标记普通类型切片
 * [x] : 标记普通类型切片的指针

 - 正则匹配

 * [x] : string
 * [x] : *string
 * [x] : []string
 * [x] : *[]string

 - 区间限制

 * [x] : 数值类型的
 * [x] : 数值类型指针
 * [x] : 数值类型切片
 * [x] : 数值类型切片指针
 * [x] : 时间类型<string>
 * [x] : 事件类型指针<*string>

## TODO

* [ ] : 内置部分正则表达式
