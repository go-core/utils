//============================================================
// 描述:
// 作者: Yang
// 日期: 2020/2/17 00:16 上午
// 版权: 山东深链智能科技有限公司 @Since 2019
//
//============================================================
package utils

import (
	"fmt"
	"testing"
)

func TestDelStrLastChar(t *testing.T) {
	str:="abcdef"
	s := DelStrLastChar(str)
	fmt.Println(s)
}

func TestTrimFormatStr(t *testing.T) {
	str:="你好啊 中国 "
	s:=TrimFormatStr(str)
	fmt.Println(s)
}

func TestLenRune(t *testing.T) {
	str:="你好中国"
	l := LenRune(str)
	fmt.Println(l)
}

func TestMD5Encrypt(t *testing.T) {
	str:="admin"
	s := MD5Encrypt(str)
	fmt.Println(s)
}
