package service

import (
	"log"
	"testing"
)

func Test_computeMingYunShu(t *testing.T) {
	MingYunShu_81, MingYunShu_9Xing := ComputeMingYunShu("18562824664")
	log.Println("MingYunShu_81 : ", MingYunShu_81)
	log.Println("MingYunShu_9Xing : ", MingYunShu_9Xing)
}
