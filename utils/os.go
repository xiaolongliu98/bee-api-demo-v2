package utils

import (
	"path/filepath"
	"runtime"
)

// GetRunningAbsDir
// @Description   获取当前文件运行的目录
// @Author        xiaolong
// @Date          2022/11/29 16:11(create);
// @Return        string
func GetRunningAbsDir() string {
	_, file, _, _ := runtime.Caller(0)
	file, _ = filepath.Abs(filepath.Dir(file))
	return file
}
