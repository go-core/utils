//============================================================
// 描述:
// 作者: Yang
// 日期: 2020/2/16 23:42 上午
// 版权: 山东深链智能科技有限公司 @Since 2019
//
//============================================================
package utils

import (
	"fmt"
	"time"
)

//格式化输出时间(2020-02-16 23:46:07)
func GetFormatTime() string {
	res := time.Now().Format("2006-01-02 15:04:05")
	return res
}

//获取年月日(2020年2月16日)
func GetDate() string {
	t := time.Now()
	year, month, day := t.Date()
	return fmt.Sprintf("%d年%d月%d日", year, month, day)
}

//格式化时间(年月日时分秒转时间戳)
func FormatTime2Timestamp(datetime string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	timestamp, _ := time.ParseInLocation("2006-01-02 15:04:05", datetime, loc)
	return timestamp.Unix()
}

//格式化时间(年月日转时间戳)
func FormatDate2Timestamp(datetime string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	timestamp, _ := time.ParseInLocation("2006-01-02", datetime, loc)
	return timestamp.Unix()
}

