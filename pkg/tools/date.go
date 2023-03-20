//@Author: wulinlin
//@Description:
//@File:  date
//@Version: 1.0.0
//@Date: 2023/03/10 03:34

package tools

import (
	"fmt"
	"time"
)

const (
	timeLayout = "15:04:05"
	dateLayout = "2006-01-02"
)

// 将日期字符串转换为时间对象
func ParseTime(timeStr string) (time.Time, error) {
	return time.Parse(timeLayout, timeStr)
}

// 将时间字符串转换为日期对象
func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse(dateLayout, dateStr)
}

// 将日期对象格式化为日期字符串
func FormatDate(date time.Time) string {
	return date.Format(dateLayout)
}

// 将时间对象格式化为时间字符串
func FormatTime(time time.Time) string {
	return time.Format(timeLayout)
}

// 将日期时间对象格式化为日期时间字符串
func FormatDateTime(dateTime time.Time) string {
	return dateTime.Format(fmt.Sprintf("%s %s", dateLayout, timeLayout))
}
