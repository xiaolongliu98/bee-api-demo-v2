package models

import "sync"

var (
	// 全局model实例申明
	once sync.Once
)

func RegisterAll() {
	// 全局model实例初始化
	once.Do(func() {
		registerORM()
	})
}
