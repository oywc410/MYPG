package main

import (
	"testing"
	"path/filepath"
	"io/ioutil"
	"os"
	gomniauthtest "github.com/stretchr/gomniauth/test"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar

	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return("", ErrNoAvatarURL)
	testChatUser := &chatUser{User: testUser}
	_, err := authAvatar.GetAvatarURL(testChatUser)
	if err != ErrNoAvatarURL {
		t.Error("URL获取失败时应该返回 ErrNoAvatarURL错误")
	}

	//测试值
	testUrl := "aaaa"
	testUser = &gomniauthtest.TestUser{}
	testChatUser.User = testUser
	testUser.On("AvatarURL").Return(testUrl, nil)
	url, err := authAvatar.GetAvatarURL(testChatUser)

	if err != nil {
		t.Error("获取URL成功时不能返回错误")
	} else {
		if url != testUrl {
			t.Error("返回值错误")
		}
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	user := &chatUser{uniqueID: "abc"}
	url, err := gravatarAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("获取URL成功时不能返回错误")
	}

	if url != "//www.gravatar.com/avatar/abc" {
		t.Error("url返回错误值")
	}

}

func TestFileSystemAvatar(t *testing.T) {

	//测试用文件
	filename := filepath.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename, []byte{}, 0777)
	defer func() {
		os.Remove(filename)
	}()

	var fileSystemAvatar FileSystemAvatar
	user := &chatUser{uniqueID: "abc"}
	url, err := fileSystemAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("获取URL成功时不能返回错误")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("url返回错误值 %s", url)
	}
}