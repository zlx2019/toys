/**
  @author: Zero
  @date: 2023/3/28 15:19:16
  @desc: 日期时间相关函数库

**/

package toys

import "time"

// FormatMap 日期时间格式化风格映射
var FormatMap map[string]string

// 常用格式
const (
	DTFormat   = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"
	TimeFormat = "15:04:05"
)

// 初始化
func init() {
	FormatMap = map[string]string{
		"yyyy-mm-dd hh:mm:ss": "2006-01-02 15:04:05",
		"yyyy-mm-dd hh:mm":    "2006-01-02 15:04",
		"yyyy-mm-dd hh":       "2006-01-02 15:04",
		"yyyy-mm-dd":          "2006-01-02",
		"yyyy-mm":             "2006-01",
		"mm-dd":               "01-02",
		"dd-mm-yy hh:mm:ss":   "02-01-06 15:04:05",
		"yyyy/mm/dd hh:mm:ss": "2006/01/02 15:04:05",
		"yyyy/mm/dd hh:mm":    "2006/01/02 15:04",
		"yyyy/mm/dd hh":       "2006/01/02 15",
		"yyyy/mm/dd":          "2006/01/02",
		"yyyy/mm":             "2006/01",
		"mm/dd":               "01/02",
		"dd/mm/yy hh:mm:ss":   "02/01/06 15:04:05",
		"yyyy":                "2006",
		"mm":                  "01",
		"hh:mm:ss":            "15:04:05",
		"mm:ss":               "04:05",
	}
}

// Now 获取当前时间,并且格式化为 yyyy-mm-dd hh:mm:ss
func Now() string {
	return time.Now().Format(DTFormat)
}

// NowDate 获取当前日期,并且格式化为 yyyy-mm-dd
func NowDate() string {
	return time.Now().Format(DateFormat)
}

// NowTime 获取当前时间,并且格式化为 hh:mm:ss
func NowTime() string {
	return time.Now().Format(TimeFormat)
}

// Format 格式化日期时间 格式为:yyyy-mm-dd hh:mm:ss
func Format(datetime time.Time) string {
	return FormatDateTime(datetime, DTFormat)
}

// Parse 解析日期时间 格式为: yyyy-mm-dd hh:mm:ss
func Parse(datetime string) (time.Time, error) {
	return ParseDateTime(datetime, DTFormat)
}

// ParseDateTime 解析日期时间,根据指定的格式
func ParseDateTime(datetime, format string) (time.Time, error) {
	return time.Parse(format, datetime)
}

// FormatDateTime 格式化时间,根据指定的格式
func FormatDateTime(datetime time.Time, format string) string {
	return datetime.Format(format)
}

// AddMinute 对时间以`分钟`为单位,进行增加。
func AddMinute(datetime time.Time, minute int64) time.Time {
	return datetime.Add(time.Minute * time.Duration(minute))
}

// AddHour 对时间以`小时`为单位,进行增加。
func AddHour(datetime time.Time, hour int64) time.Time {
	return datetime.Add(time.Hour * time.Duration(hour))
}

// AddDay 对时间以`天`为单位,进行增加。
func AddDay(datetime time.Time, day int64) time.Time {
	return datetime.Add((24 * time.Hour) * time.Duration(day))
}

// BeginOfMinute 获取一个时间以`分钟`为单位的起始时间
// 2023-03-28 16:17:00
func BeginOfMinute(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), 0, 0, t.Location())
}

// EndOfMinute 获取一个时间以`分钟`为单位的结束时间
// 2023-03-28 16:17:59
func EndOfMinute(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfHour 获取一个时间以`小时`为单位的开始时间
// 2023-03-28 16:00:00
func BeginOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

// EndOfHour 获取一个时间以`小时`为单位的结束时间
// 2023-03-28 16:59:59
func EndOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfDay 获取一个时间的当天开始时间
// 2023-03-28 00:00:00
func BeginOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// EndOfDay 获取一个时间的当天结束时间
// 2023-03-28 23:59:59
func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfMonth 获取一个时间的当月开始时间
// 2023-03-01 00:00:00
func BeginOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth 获取一个时间当月结束时间
// 2023-03-31 23:59:59
func EndOfMonth(t time.Time) time.Time {
	return BeginOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// BeginOfYear 获取一个时间的年份开始时间
// 2023-01-01 00:00:00
func BeginOfYear(t time.Time) time.Time {
	y, _, _ := t.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear 获取一个时间的年份结束时间
// 2023-12-31 23:59:59
func EndOfYear(t time.Time) time.Time {
	return BeginOfYear(t).AddDate(1, 0, 0).Add(-time.Nanosecond)
}
