package handlers

import (
	"log"

	"github.com/EriconYu/BaXingNengLiang/dto"
	service "github.com/EriconYu/BaXingNengLiang/service/phoneNumber"
	"github.com/devfeel/dotweb"
)

/**
 * @Description 手机号码吉凶分析
 * @Params  [ctx]
 * @Returns
 * @Creator  Ericon Freeman
 * @Date  2022/5/24 3:09 下午
 */
func phoneNumberIndex(ctx dotweb.Context) error {
	phoneNumber := ctx.GetRouterName("phoneNumber")
	var baseResp dto.BaseResponse
	var err error
	if len(phoneNumber) != 11 {
		baseResp = dto.BaseResponse{
			Msg:    "提交的phoneNumber错误，请提交11位手机号",
			Reason: err.Error(),
		}
		return ctx.WriteJson(baseResp)
	}
	log.Println("phoneNumber is ", phoneNumber)

	// 接下来开始分析
	baXingArray, baXingContent := service.ParsePhoneNumberByBaXing(phoneNumber)
	MingYunShu_81, MingYunShu_9Xing := service.ComputeMingYunShu(phoneNumber)
	baseResp.Reason = baXingArray
	result := make(map[string]string)
	result["baXingContent"] = baXingContent
	result["MingYunShu_81"] = MingYunShu_81
	result["MingYunShu_9Xing"] = MingYunShu_9Xing
	baseResp.Res = result

	return ctx.WriteJson(baseResp)
}
