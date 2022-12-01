package test

import (
	"bee-api-demo/models/user"
	"bee-api-demo/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func init() {
	//models.RegisterAll()
}

func TestGetUserMainByXXX(t *testing.T) {
	um := &user.UserMain{Email: "1006814807@qq.com"}
	err := user.GetUserMain(um, "email")
	if err != nil {
		panic("")
	}
	fmt.Printf("%v\n", um)

	ui, _ := user.GetUserInfoByUid(um.Uid)
	fmt.Printf("%v\n", ui)
}

func TestGetUserMain(t *testing.T) {
	um := &user.UserMain{Phone: "1001010010"}
	err := user.GetUserMain(um, "phone")
	if err != nil {
		panic("")
	}

	fmt.Printf("%v\n", um)
}

func TestAddUser(t *testing.T) {
	um := &user.UserMain{
		Email:    "zhangsan@qq.com",
		Phone:    "1001010010",
		Password: "zhangsan123",
	}

	err := user.AddUserMain(um)
	if err != nil {
		panic("")
	}
	fmt.Printf("%v\n", um)

	ui := &user.UserInfo{
		Uid:      um.Uid,
		Username: "zhangsan",
		Region:   "usa",
		Balance:  "500",
		Intro:    "hello world",
	}
	err = user.AddUserInfo(ui)
	if err != nil {
		panic("")
	}

	fmt.Printf("%v\n", ui)
}

func TestUtils(t *testing.T) {
	um := &user.UserMain{
		Uid:      0,
		Email:    "123456@qq.com",
		Phone:    "123456",
		Password: "654321",
	}

	ui := &user.UserInfo{
		Uid:      1,
		Avatar:   "sdgsddvdfdfgsd",
		Username: "xiaolong",
		Region:   "ningbo",
		Balance:  "12031",
		Intro:    "nothing",
	}

	m := utils.MergeModelsToMap(ui, um)
	fmt.Printf("%v\n", m)

	m["gender"] = "man"
	m["old"] = 23

	fmt.Printf("%v\n", m)

	//jsonBytes, _ := json.Marshal(m)
	um2 := &user.UserMain{}
	ui2 := &user.UserInfo{}
	err := utils.ParseMapToModels(&m, um2, ui2)
	if err != nil {
		panic(fmt.Sprintf("%s", err.Error()))
	}

	fmt.Printf("%v\n", um2)
	fmt.Printf("%v\n", ui2)

	w := utils.Wrapper("ok", 0, m)
	fmt.Printf("%v\n", w)

}
