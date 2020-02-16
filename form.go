//============================================================
// 描述:
// 作者: Yang
// 日期: 2020/2/17 00:25 上午
// 版权: 山东深链智能科技有限公司 @Since 2019
//
//============================================================
package utils

import "strings"

//对电话号码中间四位显示星号
func ChangePhoneToStar(s string) string {
	old := ""
	for k, v := range s {
		if k >= 3 && k <= 6 {
			old = old + string(v)
		}
	}
	phone := strings.Replace(s, old, "****", -1)
	return phone
}

