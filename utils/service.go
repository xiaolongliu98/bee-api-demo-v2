package utils

import "encoding/json"

// MergeModelsToMap
// @Description   将多个结构体属性与值转换成map，并合并
// @Author        xiaolong
// @Date          2022/11/29 16:11(create);
// @Param         modelsPtr       ...any      结构体数组
// @Return        map[string]any
func MergeModelsToMap(modelsPtr ...any) map[string]any {
	var m map[string]any
	for _, model := range modelsPtr {
		jsonBytes, _ := json.Marshal(model)
		_ = json.Unmarshal(jsonBytes, &m)
	}

	return m
}
