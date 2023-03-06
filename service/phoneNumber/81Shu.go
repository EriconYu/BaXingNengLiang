package service

import (
	"strconv"

	"github.com/EriconYu/BaXingNengLiang/util"
)

func ComputeMingYunShu(phoneNumber string) (string, string) {

	// 81数
	var last4Sum int
	phoneNumberLenth := len(phoneNumber)
	l4, _ := strconv.Atoi(string(phoneNumber[phoneNumberLenth-1-3]))
	l3, _ := strconv.Atoi(string(phoneNumber[phoneNumberLenth-1-2]))
	l2, _ := strconv.Atoi(string(phoneNumber[phoneNumberLenth-1-1]))
	l1, _ := strconv.Atoi(string(phoneNumber[phoneNumberLenth-1-0]))
	last4Sum = l4*1000 + l3*100 + l2*10 + l1*1
	last4Sum = last4Sum % 81
	if last4Sum == 0 {
		last4Sum = 81
	}
	MingYunShu_81 := util.ShuArray_81[last4Sum]

	// 九星
	var sumVal int
	for i := 3; i < len(phoneNumber); i++ {
		num, _ := strconv.Atoi(string(phoneNumber[i]))
		sumVal = sumVal + num
	}
	sumVal = sumVal/10 + sumVal%10
	sumVal = sumVal/10 + sumVal%10
	MingYunShu_9Xing := util.MingYunShuArray_9Xing[sumVal]

	return MingYunShu_81, MingYunShu_9Xing
}
