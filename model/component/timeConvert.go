package component

import "time"

//时间戳转日期 秒级
func TimestampToDate_Sec(timestamp int64) string {
	date := time.Unix(timestamp, 0)
	return date.Format("2006-01-02 15:04:05")
}

//时间戳转日期 毫秒级
func TimestampToDate_MillSec(timestamp int64) string {
	date := time.Unix(timestamp/1000, 0)
	return date.Format("2006-01-02 15:04:05")
}

//时间戳转日期 秒级 自定义
func TimestampToDate_Sec_Custom(timestamp int64, cusStr string) string {
	date := time.Unix(timestamp, 0)
	return date.Format(cusStr)
}

//时间戳转日期 毫秒级 自定义
func TimestampToDate_MillSec_Custom(timestamp int64, cusStr string) string {
	date := time.Unix(timestamp/1000, 0)
	return date.Format(cusStr)
}
