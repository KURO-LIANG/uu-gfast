package utils

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"regexp"
	"time"
)

// CheckMobile 校验手机号
func CheckMobile(mobile string) bool {
	regRuler := "^1[3456789]{1}\\d{9}$"
	// 正则调用规则
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(mobile)
}

// 获取间隔时间
func GetDuration(startTime *gtime.Time, endTime *gtime.Time) (duration string, per string, err error) {
	// 相差秒数
	diffSecond := endTime.Sub(startTime)
	// 相差分钟数
	diffMinutes := int(diffSecond.Minutes())
	if diffMinutes == 0 {
		seconds := diffSecond.Seconds()
		if seconds < 10 {
			duration = "刚刚"
			return
		}
		duration = gconv.String(int(seconds))
		per = "秒"
		return
	}
	// 相差小时数
	diffHours := int(diffSecond.Hours())
	if diffHours == 0 {
		duration = gconv.String(diffMinutes)
		per = "分钟"
		return
	}
	// 计算相差的天数
	diffDays := diffHours / 24
	if diffDays == 0 {
		duration = gconv.String(diffHours)
		per = "小时"
		return
	}
	// 计算相差的月数
	diffMonths := endTime.Month() - startTime.Month() + (endTime.Year()-startTime.Year())*12
	if diffMonths == 0 {
		duration = gconv.String(diffDays)
		per = "天"
		return
	}
	// 计算相差的年数
	diffYears := endTime.Year() - startTime.Year()
	if diffYears == 0 {
		duration = gconv.String(diffMonths)
		per = "月"
		return
	}
	duration = gconv.String(diffYears)
	per = "年"
	return
}

func GetWeekend() (monday string, sunday string) {
	// 获取本周时间
	now := time.Now()
	weekday := now.Weekday()
	mondayTime := now.AddDate(0, 0, -int(weekday)+1)
	sundayTime := now.AddDate(0, 0, 7-int(weekday))
	monday = fmt.Sprint(mondayTime.Format("2006-01-02"), " 00:00:00")
	sunday = fmt.Sprint(sundayTime.Format("2006-01-02"), " 23:59:59")
	return
}
