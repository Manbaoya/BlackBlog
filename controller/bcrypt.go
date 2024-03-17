package controller

import (
	"BLACKBLOG/log"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

//加密

func Encrypt(password string) (string, bool) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), 1)
	if err != nil {
		log.SugaredLogger.Errorf("加密失败:%s", err)
		return "", false
	}
	fmt.Println(string(pwd))
	return string(pwd), true
}

//密码验证

func Decrypt(hashPwd string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(password))
	if err != nil {
		return false
	}
	return true
}
