package user

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// GetUserInfoByUid
// @Description 通过用户uid获取用户信息
// @Author 		xiaolong
// @Date		2022/11/28 19:50(create);
// @Param		uid		int		用户uid
// @Return		*user.UserInfo, error
func GetUserInfoByUid(uid int) (*UserInfo, error) {
	o := orm.NewOrm()
	// -------------------------------------
	u := &UserInfo{Uid: uid}

	err := o.Read(u)
	if err != nil {
		//fmt.Printf("%v\n", err.Error())
		return nil, err
	}

	return u, err
}

// GetUserMain
// @Description 通过指定字段获取用户核心数据
// @Author 		xiaolong
// @Date		2022/11/28 19:50(create);
// @Param		um		UserInfo	  在指定字段已填好的UserInfo
// @Param		by		string		  字段，"email", "phone", "uid"
// @Return		error
func GetUserMain(um *UserMain, by string) error {
	o := orm.NewOrm()
	// -------------------------------------
	err := o.Read(um, by)
	if err != nil {
		//fmt.Printf("%v\n", err.Error())
		return err
	}
	return nil
}

// AddUserMain
// @Description 插入一条用户核心数据，成功后在um中会保存uid
// @Author 		xiaolong
// @Date		2022/11/29 14:59(create);
// @Param		um		*UserMain		用户核心数据
// @Return		error
func AddUserMain(um *UserMain) error {
	if um == nil {
		return fmt.Errorf("um is nil")
	}

	o := orm.NewOrm()
	uid, err := o.Insert(um)
	if err != nil {
		return err
	}
	um.Uid = int(uid)
	return nil
}

// AddUserInfo
// @Description 插入一条用户核心数据，其中uid必须设置
// @Author 		xiaolong
// @Date		2022/11/29 14:59(create);
// @Param		ui		*UserInfo		用户信息
// @Return		error
func AddUserInfo(ui *UserInfo) error {
	if ui == nil {
		return fmt.Errorf("ui is nil")
	}

	o := orm.NewOrm()
	_, err := o.Insert(ui)
	if err != nil {
		return err
	}
	return nil
}
