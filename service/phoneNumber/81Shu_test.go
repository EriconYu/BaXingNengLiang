package service

import (
	"fmt"
	"testing"
)

func Test_computeMingYunShu(t *testing.T) {
	MingYunShu_81, MingYunShu_9Xing := ComputeMingYunShu("18562824664")
	fmt.Println("MingYunShu_81 : ", MingYunShu_81)
	fmt.Println("MingYunShu_9Xing : ", MingYunShu_9Xing)
}
