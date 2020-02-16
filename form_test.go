//============================================================
// 描述:
// 作者: Yang
// 日期: 2020/2/17 00:26 上午
// 版权: 山东深链智能科技有限公司 @Since 2019
//
//============================================================
package utils

import (
	"fmt"
	"testing"
)

func TestChangePhoneToStar(t *testing.T) {
	str:="13145310988"
	s := ChangePhoneToStar(str)
	fmt.Println(s)
}
