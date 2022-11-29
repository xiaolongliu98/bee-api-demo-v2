package utils

import "bee-api-demo/types"

// WithDefault
// @Description   将指定的返回值替换成默认值y
// @Author        xiaolong
// @Date          2022/11/29 16:11(create);
// @Param         x       	T     value
// @Param         undesired T     想要替换的值
// @Param         y 		T     默认值
// @Return        T
func WithDefault[T types.BaseType](x, undesired, y T) T {
	if x == undesired {
		return y
	}
	return x
}
