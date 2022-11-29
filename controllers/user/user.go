package user

import (
	"bee-api-demo/services/user"
	"bee-api-demo/utils"
)

/**
完成用户注册登录功能
*/

// For Example
// 启动后，测试访问：http://localhost:8080/v1/user/uid/6
// 若提示nomatch，请关闭后在命令行使用bee fix修复后再进行启动
// 注意：默认数据库配置需要开启玉泉VPN，表结构参考SNOS

// GetUserMainAndInfoById
// @Title GetUser
// @Author xiaolong
// @Date 2022/11/24 21:26(create);
// @Description 通过用户id获取用户信息
// @Param uid query int true 用户id
// @Success 200 {object} UserMain+UserInfo
// @Failure 500 {string}
// @router /uid/:uid [get]
func (u *UserController) GetUserMainAndInfoById() {
	uid, err := u.GetInt(":uid")
	if err != nil {
		u.Data["json"] = err.Error()
		u.ServeJSON()
		return
	}

	um, ui, code := user.GetUserMainAndInfoByUid(uid)
	if code == 0 {
		u.Data["json"] = "get user failed"
		u.ServeJSON()
		return
	}

	u.Data["json"] = utils.MergeModelsToMap(um, ui)
	u.ServeJSON()
}
