//============================================================
// 描述:
// 作者: Yang
// 日期: 2020/2/16 23:44 上午
// 版权: 山东深链智能科技有限公司 @Since 2019
//
//============================================================
package utils

import (
	"fmt"
	"testing"
)

func TestGetFormatTime(t *testing.T) {
	fmt.Println(GetFormatTime())
}

func TestGetDate(t *testing.T) {
	fmt.Println(GetDate())
}

func TestFormatTime2Timestamp(t *testing.T) {
	datetime:="2018-02-01 02:10:25"
	timestamp := FormatTime2Timestamp(datetime)
	fmt.Println(timestamp)
}

func TestFormatDate2Timestamp(t *testing.T) {
	datetime:="2018-02-01"
	timestamp := FormatDate2Timestamp(datetime)
	fmt.Println(timestamp)
}
