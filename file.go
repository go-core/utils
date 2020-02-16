//============================================================
// 描述:
// 作者: Yang
// 日期: 2020/2/17 00:09 上午
// 版权: 山东深链智能科技有限公司 @Since 2019
//
//============================================================
package utils

import "os"

//检查文件或目录是否存在
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

