package utils

import "time"

//时间戳转化为时间
//unix int 64, format 当为""时使用默认值"2006-01-02 15:04:05"
//@param: unix int64    时间戳
//@param: format string 使用默认值为"":2006-01-02 15:04:05
//@return: string
func UnixToTime(unix int64, format string) string  {
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	return time.Unix(unix, 0).Format(format)
}
