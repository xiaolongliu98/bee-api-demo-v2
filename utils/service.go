package utils

import "encoding/json"

// MergeModelsToMap
// @Description   将多个结构体属性与值转换成map，并合并，相同属性会被后面的覆盖
// @Author        xiaolong
// @Date          2022/11/29 16:11(create);
// @Param         modelsPtr       ...any      结构体数组
// @Return        map[string]any
func MergeModelsToMap(modelsPtr ...any) map[string]any {
	var m map[string]any
	for _, model := range modelsPtr {
		if model == nil {
			continue
		}
		jsonBytes, _ := json.Marshal(model)
		_ = json.Unmarshal(jsonBytes, &m)
	}

	return m
}

// ParseJsonToModels
// @Description   将Json拆分，并注入到多个相关的Model中
// @Author        xiaolong
// @Date          2022/11/29 16:11(create);
// @Param         modelsPtr       ...any      结构体数组
// @Return        map[string]any
func ParseJsonToModels(jsonBytes []byte, modelsPtr ...any) error {
	//var m map[string]any
	for _, model := range modelsPtr {
		if model == nil {
			continue
		}
		//jsonBytes, _ := json.Marshal(model)
		//_ = json.Unmarshal(jsonBytes, &m)
		err := json.Unmarshal(jsonBytes, model)
		if err != nil {
			return err
		}
	}

	return nil
}

// ParseMapToModels
// @Description   将Map拆分，并注入到多个相关的Model中
// @Author        xiaolong
// @Date          2022/11/29 16:11(create);
// @Param         modelsPtr       ...any      结构体数组
// @Return        map[string]any
func ParseMapToModels(m *map[string]any, modelsPtr ...any) error {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return ParseJsonToModels(jsonBytes, modelsPtr...)
}

// Wrapper
// @Description   封装接口返回数据
// @Author        xiaolong
// @Date          2022/11/29 16:11(create);
// @Param         msg        string   中文信息
// @Param         code       int      返回代码
// @Param         result     any      主要数据
// @Return        map[string]any
func Wrapper(msg string, code int, result any) map[string]any {
	m := make(map[string]any)
	m["msg"] = msg
	m["code"] = code
	m["result"] = result
	return m
}

// WrapperError
// @Description   封装接口返回的错误内容
// @Author        xiaolong
// @Date          2022/11/29 16:11(create);
// @Param         msg        string   中文信息
// @Param         code       int      返回代码
// @Return        map[string]any
func WrapperError(msg string, code int) map[string]any {
	m := make(map[string]any)
	m["msg"] = msg
	m["code"] = code
	m["result"] = struct{}{}
	return m
}
