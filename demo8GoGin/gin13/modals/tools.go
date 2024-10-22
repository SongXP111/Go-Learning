package modals

import "time"

// 自定义Modal

// 时间戳转化成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func GetUnix() int64 {
	return time.Now().Unix()
}

func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}
