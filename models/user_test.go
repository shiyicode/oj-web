package models

import (
	"fmt"
	"testing"
)

func TestUserCreate(t *testing.T) {
	InitAllInTest()

	user := &User{AccountId: 1, NickName: "1521901996586326000", UserName: "asdf"}
	if _, err := user.Create(user); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestUserUpdate(t *testing.T) {
	InitAllInTest()

	user := &User{Id: 1, UserName: "adaad"}
	if err := user.Update(user); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestUserRemove(t *testing.T) {
	InitAllInTest()

	var user User
	if err := user.Remove(1); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestUserGetById(t *testing.T) {
	InitAllInTest()

	user := &User{AccountId: 1, UserName: "abcdfg", NickName: "hahaha", Description: "1111",
		Sex: "男", Birthday: "2011-10-01", DailyAddress: "西安"}
	User{}.Create(user)

	getUser, err := User{}.GetById(user.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getUser != *user {
		t.Error("GetById() failed:", "%v != %v", user, getUser)
	}
}
func TestUserQueryByName(t *testing.T) {
	InitAllInTest()

	user := &User{NickName: "ssd", UserName: "rrrrrr"}
	user1 := &User{NickName: "ssd", UserName: "tttttt"}
	user.Create(user)
	user.Create(user1)

	userList, err := user.QueryByName("ssd")
	if err != nil {
		t.Error("QueryByName() failed:", err)
	}
	if len(userList) != 2 {
		t.Error("QueryByName() failed:", "count is wrong!")
	}
}
func TestUserGetByAccountId(t *testing.T) {
	InitAllInTest()

	user := &User{AccountId: 20}
	user.Create(user)

	getUser, err := user.GetByAccountId(1)
	if err != nil {
		t.Error("GetByAccountId() failed:", err)
	}
	fmt.Println(getUser)
	//if getUser.AccountId != 20 {
	//	t.Error("GetByAccountId() failed:", "%v != %v", user, getUser)
	//}
}
