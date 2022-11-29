package types

/**
定义一些公共类型和接口
*/

// 基础类型集合
type BaseType interface {
	string | int | float64 | float32 | byte | int8 | int16 | int32 | int64 | uint16 | uint32 | uint64 | uint
}

// 可比较接口
type Comparable interface {
	CompareTo(o any) int
}
