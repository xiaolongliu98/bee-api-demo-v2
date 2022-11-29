package database

import (
	"bee-api-demo/utils"
	"fmt"
)

// 数据库配置信息
type DBConfig struct {
	Username string `json:"username"` // 数据库账号
	Password string `json:"password"` // 数据库密码
	Driver   string `json:"driver"`   // 数据库驱动名
	DBName   string `json:"db_name"`  // 数据库名称
	IPPort   string `json:"ip_port"`  // 地址端口
	Alias    string `json:"alias"`    // ORM数据库别名
	Timeout  string `json:"timeout"`  // 超时时间
}

// 以下提供一些基础功能

// NewDBConfig
// @Description   实例化一个DBConfig
// @Author        xiaolong
// @Date          [2022/11/24 16:00](create);
// @Return        *DBConfig
func NewDBConfig() *DBConfig {
	return &DBConfig{}
}

// NewDBConfigAndBind
// @Description   实例化一个DBConfig并且绑定具体配置文件
// @Author        xiaolong
// @Date          [2022/11/24 16:00](create);
// @Param         configFile       string      DB-JSON配置文件地址
// @Return        *DBConfig, error
func NewDBConfigAndBind(configFile string) (*DBConfig, error) {
	conf := &DBConfig{}
	err := conf.BindConfig(configFile)
	return conf, err
}

// GetDataSource
// @Description   生成DataSource链接，用于ORM连接数据库
// @Author        xiaolong
// @Date          [2022/11/24 16:00](create);
// @Return        string
func (conf *DBConfig) GetDataSource() string {
	s := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&timeout=%v&loc=Local&parseTime=true",
		conf.Username, conf.Password, conf.IPPort, conf.DBName, conf.Timeout)
	//s := root:123456@tcp(123.180.11.30:3306)/tizi?charset=utf8&timeout=5s&loc=Local&parseTime=true
	return s
}

// BindConfig
// @Description   （切换）绑定具体配置文件
// @Author        xiaolong
// @Date          [2022/11/24 16:00](create);
// @Param         configFile       string      DB-JSON配置文件地址
// @Return        error
func (conf *DBConfig) BindConfig(configFile string) error {
	err := utils.ReadConfigJSON(configFile, conf)
	return err
}
