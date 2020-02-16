//============================================================
// 描述:
// 作者: Yang
// 日期: 2020/2/17 00:15 上午
// 版权: 山东深链智能科技有限公司 @Since 2019
//
//============================================================
package utils

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

//去掉字符串最后一个字符(中英文)
func DelStrLastChar(str string) string {
	r, size := utf8.DecodeLastRuneInString(str)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return str[:len(str)-size]
}

//去掉字符串中的空格、换行符、回车符
func TrimFormatStr(str string) string {
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	// 去掉回车符
	str = strings.Replace(str, "\r", "", -1)
	return str
}

// 返回字符串中的字符数量
func LenRune(str string) int  {
	return  len([]rune(str))
}

// MD5 加密
func MD5Encrypt(str string) string  {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

//NewPasswdFormStr 生成密码密文,采用 bcrypt 加解密
func NewPasswdFormStr(passwd string) string {
	en, _ := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	return base64.StdEncoding.EncodeToString(en)
}


//ComparePasswd 密码比较
func ComparePasswd(passwd, encodingPwd string) bool {
	bb, err := base64.StdEncoding.DecodeString(encodingPwd)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword(bb, []byte(passwd))
	if err != nil {
		return false
	}
	return true
}

//CreateSMSCode 创建6位随机验证码
func CreateSMSCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

