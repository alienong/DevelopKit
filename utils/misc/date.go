/**
 * @Author: alienongwlx@gmail.com
 * @Description: Some Datetime Methods
 * @Version: 1.0.0
 * @Date: 2020/4/15 19:50
 */

package misc

import (
	"strings"
	"time"
)

const (
	MM_DD                      = "MM-dd"
	YYYYMM                     = "yyyyMM"
	YYYY_MM                    = "yyyy-MM"
	YYYY_MM_DD                 = "yyyy-MM-dd"
	YYYYMMDD                   = "yyyyMMdd"
	YYYYMMDDHHMMSS             = "yyyyMMddHHmmss"
	YYYYMMDDHHMM               = "yyyyMMddHHmm"
	YYYYMMDDHH                 = "yyyyMMddHH"
	YYMMDDHHMM                 = "yyMMddHHmm"
	MM_DD_HH_MM                = "MM-dd HH:mm"
	MM_DD_HH_MM_SS             = "MM-dd HH:mm:ss"
	YYYY_MM_DD_HH_MM           = "yyyy-MM-dd HH:mm"
	YYYY_MM_DD_HH_MM_SS        = "yyyy-MM-dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_SS_SSS    = "yyyy-MM-dd HH:mm:ss.SSS"
	MM_DD_EN                   = "MM/dd"
	YYYY_MM_EN                 = "yyyy/MM"
	YYYY_MM_DD_EN              = "yyyy/MM/dd"
	MM_DD_HH_MM_EN             = "MM/dd HH:mm"
	MM_DD_HH_MM_SS_EN          = "MM/dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_EN        = "yyyy/MM/dd HH:mm"
	YYYY_MM_DD_HH_MM_SS_EN     = "yyyy/MM/dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_SS_SSS_EN = "yyyy/MM/dd HH:mm:ss.SSS"
	MM_DD_CN                   = "MM月dd日"
	YYYY_MM_CN                 = "yyyy年MM月"
	YYYY_MM_DD_CN              = "yyyy年MM月dd日"
	MM_DD_HH_MM_CN             = "MM月dd日 HH:mm"
	MM_DD_HH_MM_SS_CN          = "MM月dd日 HH:mm:ss"
	YYYY_MM_DD_HH_MM_CN        = "yyyy年MM月dd日 HH:mm"
	YYYY_MM_DD_HH_MM_SS_CN     = "yyyy年MM月dd日 HH:mm:ss"
	HH_MM                      = "HH:mm"
	HH_MM_SS                   = "HH:mm:ss"
	HH_MM_SS_MS                = "HH:mm:ss.SSS"
)

/**
@description: Return Time Day's Timespan
@param tm: Time
@return: timespan
*/
func Date(tm time.Time) int64 {
	year, month, day := tm.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local).Unix()
}

/**
@description: Return Time Month's Timespan
@param tm: Time
@return: timespan
*/
func Month(tm time.Time) int64 {
	year, month, _ := tm.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, time.Local).Unix()
}

/**
@description: Return Time Year's Timespan
@param tm: Time
@return: timespan
*/
func Year(tm time.Time) int64 {
	year, _, _ := tm.Date()
	return time.Date(year, 1, 1, 0, 0, 0, 0, time.Local).Unix()
}

/**
@description: Return Time Year's Timespan
@param date: time string
@param format: time format
@return: timespan
*/
func DateToUnix(date string, format string) int64 {
	theTime, _ := time.ParseInLocation(format, date, time.Local)
	return theTime.Unix()
}

/**
@description:  Return Timespan's Format String
@param ts: timespasn
@param format: time format e.g. 2006-01-02 15:04:05
@return: Timespan
*/
func UnixToDate(ts int64, format string) string {
	return time.Unix(ts, 0).Format(format)
}

/**
@description: Return First Day of Month
@param month: time string
@param layout: time format e.g. 2006-01-02 15:04:05
@return: Timespan
*/
func FirstDayOfMonth(month string, layout string) int {
	theTime, _ := time.ParseInLocation(layout, month, time.Local)
	currentYear, currentMonth, _ := theTime.Date()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.Local)
	return firstOfMonth.Day()
}

/**
@description: Return Last Day of Month
@param month: time string
@param layout: time format e.g. 2006-01-02 15:04:05
@return: Timespan
*/
func LastDayOfMonth(month string, layout string) int {
	theTime, _ := time.ParseInLocation(layout, month, time.Local)
	currentYear, currentMonth, _ := theTime.Date()
	currentLocation := theTime.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	return lastOfMonth.Day()
}

/**
@description: Return First Day And Last Day of Month
@param month: time string
@param layout: time format e.g. 2006-01-02 15:04:05
@return: Timespan
*/
func FirstAndLastUnixOfMonth(month string, layout string) (s int64, e int64) {
	theTime, _ := time.ParseInLocation(layout, month, time.Local)
	currentYear, currentMonth, _ := theTime.Date()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.Local)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	return firstOfMonth.Unix(), lastOfMonth.Unix()
}

/**
@description: Return First Day And Last Day of Year
@param month: time string
@param layout: time format e.g. 2006-01-02 15:04:05
@return: Timespan
*/
func FirstAndLastUnixOfYear(month string, layout string) (s int64, e int64) {
	theTime, _ := time.ParseInLocation(layout, month, time.Local)
	currentYear, _, _ := theTime.Date()
	firstOfYear := time.Date(currentYear, 1, 1, 0, 0, 0, 0, time.Local)
	lastOfYear := firstOfYear.AddDate(1, 0, -1)
	return firstOfYear.Unix(), lastOfYear.Unix()
}

/**
@description: Return  Add Month To Time String For Month
@param month: time string
@param layout: time format e.g. 2006-01-02 15:04:05
@param num: add nums
@return: Timespan
*/
func AddMonthForUnix(month string, layout string, num int) int64 {
	theTime, _ := time.ParseInLocation(layout, month, time.Local)
	year := theTime.AddDate(0, num, 0)
	return year.Unix()
}

/**
@description: Return First Day And Last Day of Week
@param date: timespan
@return: Week's First TimeSpan and Last TimeSpan
*/
func FirstAndLastUnixOfWeek(date int64) (s int64, e int64) {
	theTime := time.Unix(date, 0)
	w := int(theTime.Weekday())
	if w == 0 {
		s = date - int64(6*86400)
		e = date
	} else {
		s = date - int64((w-1)*86400)
		e = date + int64((7-w)*86400)
	}
	return s, e
}

/**
@description: FormatDate By dateStyle
@param date: time
@param dateStyle: dateStyle
@return: Format String
*/
func FormatDate(date time.Time, dateStyle string) string {
	layout := string(dateStyle)
	layout = strings.Replace(layout, "yyyy", "2006", 1)
	layout = strings.Replace(layout, "yy", "06", 1)
	layout = strings.Replace(layout, "MM", "01", 1)
	layout = strings.Replace(layout, "dd", "02", 1)
	layout = strings.Replace(layout, "HH", "15", 1)
	layout = strings.Replace(layout, "mm", "04", 1)
	layout = strings.Replace(layout, "ss", "05", 1)
	layout = strings.Replace(layout, "SSS", "000", -1)
	return date.Format(layout)
}
