package test

import (
	"bee-api-demo/models"
	"bee-api-demo/models/user"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func init() {
	// 测试固定写法------------------------------------
	projectName := "bee-api-demo"
	_, file, _, _ := runtime.Caller(0)
	for {
		file = filepath.Dir(file)
		if strings.HasSuffix(file, projectName) {
			break
		}
	}
	file, _ = filepath.Abs(file)
	os.Setenv("ROOTDIR", file)
	// ------------------------------------------

	models.RegisterAll()
}

func TestGetUserMainByXXX(t *testing.T) {
	u, _ := user.GetUserMainByEmail("1006814807@qq.com")
	fmt.Printf("%v\n", u)

	ui, _ := user.GetUserInfoByUid(u.Uid)
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
