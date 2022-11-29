package user

// 用户核心信息
type UserMain struct {
	Uid int `orm:"pk;"`

	Email    string // 邮箱 unique
	Phone    string // 电话 NotNull, unique
	Password string // 密码 NotNuLL
}

// 用户信息
type UserInfo struct {
	Uid int `orm:"pk;"`

	Avatar   string // 头像
	Username string // 用户名 Not Null, unique
	Region   string // 地区
	Balance  string // 余额 Not Null, default(0)
	Intro    string // 简介 varchar(1024)
}

// 指定User结构体默认绑定的表名

func (um *UserMain) TableName() string {
	return "user_main"
}

func (ui *UserInfo) TableName() string {
	return "user_info"
}

// 以下提供一些基础功能

func (um *UserMain) Check() bool {
	if um == nil {
		return false
	}
	if um.Phone == "" {
		return false
	}
	if um.Password == "" {
		return false
	}
	return true
}

func (ui *UserInfo) Check() bool {
	if ui == nil {
		return false
	}
	if ui.Username == "" {
		return false
	}
	if ui.Balance == "" {
		return false
	}
	return true
}
