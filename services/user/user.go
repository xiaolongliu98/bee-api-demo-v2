package user

import "bee-api-demo/models/user"

/**
业务逻辑层 - user/user业务
*/

// GetUserMainAndInfoByUid
// @Description 通过用户id获取用户信息
// @Author 		xiaolong
// @Date		2022/11/29 15:46(create);
// @Param		uid		int		用户uid
// @Return		*user.UserMain, *user.UserInfo, int(0:失败，1:仅UserMain有效，2:成功)
func GetUserMainAndInfoByUid(uid int) (*user.UserMain, *user.UserInfo, int) {
	um := &user.UserMain{
		Uid: uid,
	}
	err := user.GetUserMain(um, "uid")
	if err != nil {
		return nil, nil, 0
	}

	ui, err := user.GetUserInfoByUid(um.Uid)
	if err != nil {
		return um, nil, 1
	}
	return um, ui, 2
}
