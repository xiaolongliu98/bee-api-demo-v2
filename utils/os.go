package utils

import (
	"path/filepath"
	"runtime"
	"strings"
)

// GetProjectRoot
// @Description   获取当前项目根目录
// @Author        xiaolong
// @Date          2022/11/29 16:11(create);
// @Return        string
func GetProjectRoot() string {
	_, file, _, _ := runtime.Caller(0)
	file, _ = filepath.Abs(filepath.Dir(filepath.Dir(file)))
	file = strings.ReplaceAll(file, "\\", "/")
	file = strings.ReplaceAll(file, "\\\\", "/")
	return file
}
