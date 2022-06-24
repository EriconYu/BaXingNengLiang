package handlers

import (
	"log"
	"time"

	"github.com/6tail/lunar-go/calendar"
	"github.com/EriconYu/BaXingNengLiang/dto"
	"github.com/EriconYu/BaXingNengLiang/util"
	"github.com/devfeel/dotweb"
)

type YiJi struct {
	FullString string
	DayYi      []interface{}
	DayJi      []interface{}
	TimeYi     []interface{}
	TimeJi     []interface{}
}

func TodayIndexFunc(ctx dotweb.Context) error {
	lunarToday := calendar.NewLunarFromDate(time.Now())
	fullString := util.ToFullStringPretty(lunarToday)

	var yiji YiJi
	yiji.FullString = fullString
	for e := lunarToday.GetDayYi().Front(); e != nil; e = e.Next() {
		yiji.DayYi = append(yiji.DayYi, e.Value)
	}
	for e := lunarToday.GetDayJi().Front(); e != nil; e = e.Next() {
		yiji.DayJi = append(yiji.DayJi, e.Value)
	}
	for e := lunarToday.GetTimeYi().Front(); e != nil; e = e.Next() {
		yiji.TimeYi = append(yiji.TimeYi, e.Value)
	}
	for e := lunarToday.GetTimeJi().Front(); e != nil; e = e.Next() {
		yiji.TimeJi = append(yiji.TimeJi, e.Value)
	}
	log.Println(yiji)
	return ctx.WriteJson(dto.BaseResponse{
		Res: yiji,
	})
}
