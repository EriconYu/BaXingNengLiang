package util

import (
	"strings"

	"github.com/6tail/lunar-go/calendar"
)

func ToFullStringPretty(lunar *calendar.Lunar) string {
	s := ""
	s += lunar.String()
	s += " "
	s += lunar.GetYearInGanZhi()
	s += "("
	s += lunar.GetYearShengXiao()
	s += ")年 "
	s += lunar.GetMonthInGanZhi()
	s += "("
	s += lunar.GetMonthShengXiao()
	s += ")月 "
	s += lunar.GetDayInGanZhi()
	s += "("
	s += lunar.GetDayShengXiao()
	s += ")日 "
	s += lunar.GetTimeZhi()
	s += "("
	s += lunar.GetTimeShengXiao()
	s += ")时 \n纳音["
	s += lunar.GetYearNaYin()
	s += " "
	s += lunar.GetMonthNaYin()
	s += " "
	s += lunar.GetDayNaYin()
	s += " "
	s += lunar.GetTimeNaYin()
	s += "] 星期"
	s += lunar.GetWeekInChinese() + "\n"
	for i := lunar.GetFestivals().Front(); i != nil; i = i.Next() {
		s += " ("
		s += i.Value.(string)
		s += ")"
	}
	for i := lunar.GetOtherFestivals().Front(); i != nil; i = i.Next() {
		s += " ("
		s += i.Value.(string)
		s += ")"
	}

	jq := lunar.GetJieQi()
	if len(jq) > 0 {
		s += " ["
		s += jq
		s += "]"
	}
	s += lunar.GetGong()
	s += "方"
	s += lunar.GetShou()
	s += " 星宿["
	s += lunar.GetXiu()
	s += lunar.GetZheng()
	s += lunar.GetAnimal()
	s += "]("
	s += lunar.GetXiuLuck()
	s += ") 彭祖百忌["
	s += lunar.GetPengZuGan()
	s += " "
	s += lunar.GetPengZuZhi()
	s += "] 喜神方位["
	s += lunar.GetDayPositionXi()
	s += "]("
	s += lunar.GetDayPositionXiDesc()
	s += ") 阳贵神方位["
	s += lunar.GetDayPositionYangGui()
	s += "]("
	s += lunar.GetDayPositionYangGuiDesc()
	s += "阴贵神方位["
	s += lunar.GetDayPositionYinGui()
	s += "]("
	s += lunar.GetDayPositionYinGuiDesc()
	s += ") 福神方位["
	s += lunar.GetDayPositionFu()
	s += "]("
	s += lunar.GetDayPositionFuDesc()
	s += ") 财神方位["
	s += lunar.GetDayPositionCai()
	s += "]("
	s += lunar.GetDayPositionCaiDesc()
	s += ") 冲["
	s += lunar.GetDayChongDesc()
	s += "] 煞["
	s += lunar.GetDaySha()
	s += "]"
	s = strings.Replace(s, " [", "\r\n [", -1)
	s = strings.Replace(s, "] ", "]\r\n", -1)
	s = strings.Replace(s, " (", "\r\n (", -1)
	s = strings.Replace(s, ") ", ")\r\n", -1)
	s = strings.Replace(s, "[", " [", -1)
	s = strings.Replace(s, "]", "] ", -1)
	return s
}
