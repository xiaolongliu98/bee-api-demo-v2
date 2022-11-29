package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// ReadConfigJSON
// @Description   将json配置文件注入到Struct中
// @Author        xiaolong
// @Date          [2022/11/25 16:49](create);
// @Param         fileName       string      json文件路径
// @Param         targetPtr      any         目标结构体指针
// @Return        失败返回错误信息
func ReadConfigJSON(fileName string, targetPtr any) error {
	if targetPtr == nil {
		return fmt.Errorf("targetPtr is nil")
	}

	fd, err := os.Open(fileName)
	defer fd.Close()

	if err != nil {
		return fmt.Errorf("failed to open file, %s", err.Error())
	}
	// 创建json解码器
	decoder := json.NewDecoder(fd)
	err = decoder.Decode(targetPtr)
	if err != nil {
		return fmt.Errorf("decode error")
	}
	return nil
}
