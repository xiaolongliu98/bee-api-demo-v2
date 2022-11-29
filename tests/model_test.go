package test

import (
	"bee-api-demo/models"
	"bee-api-demo/models/user"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func init() {
	models.RegisterAll()
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
