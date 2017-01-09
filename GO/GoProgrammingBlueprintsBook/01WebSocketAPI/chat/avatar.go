package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

var ErrNoAvatarURL = errors.New("chat: URL获取失败")

type Avatar interface {
	GetAvatarURL(u ChatUser) (string, error)
}

type AuthAvatar struct {}

var UseAuthAvatar AuthAvatar

func (_ AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

type GravatarAvatar struct {
}

var UserGravatar GravatarAvatar

func(_ GravatarAvatar) GetAvatarURL(u ChatUser) (string, error){
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

type FileSystemAvatar struct {
}

var UserFileSystemAvatar FileSystemAvatar

func(_ FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {

	//文件查找
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := filepath.Match(u.UniqueID() + "*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}

	return "", ErrNoAvatarURL
}

type TryAvatars []Avatar

func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}

	return "", ErrNoAvatarURL
}
