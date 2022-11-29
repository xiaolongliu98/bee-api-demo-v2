package utils

import "bee-api-demo/types"

func ArrayContains(arr []types.Comparable, target types.Comparable) bool {
	return ArrayFind(arr, target) != -1
}

// ArrayFind
// @Description   在切片中查找目标位置index，未找到返回-1
// @Author        xiaolong
// @Date          2022/11/29 16:11(create);
// @Param         arr       	[]Comparable     实现Comparable接口的对象切片
// @Param         target        Comparable		 实现Comparable接口的目标对象
// @Return        index	    int
func ArrayFind(arr []types.Comparable, target types.Comparable) int {
	for i, e := range arr {
		if e.CompareTo(target) == 0 {
			return i
		}
	}
	return -1
}
