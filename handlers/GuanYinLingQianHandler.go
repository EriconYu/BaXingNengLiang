package handlers

import (
	"strconv"
	"time"

	"github.com/EriconYu/BaXingNengLiang/dto"
	"github.com/EriconYu/BaXingNengLiang/util"
	"github.com/devfeel/dotweb"
)

/**
 * @Description random一个观音灵签给用户返回去
 * @Params  [ctx]
 * @Returns  [baseResp]
 * @Creator  Ericon Freeman
 * @Date  2023/3/7 5:36 下午
 */
func RandomGuanYinLingQianFunc(ctx dotweb.Context) error {

	qianOrder := time.Now().UnixNano() % 739 % 100
	if qianOrder == 0 {
		qianOrder = 100
	}
	guanYinLingQian := util.GuanYinLingQian[int(qianOrder)]
	var baseResp dto.BaseResponse
	baseResp.Msg = "抽签完成，恭喜您抽得" + guanYinLingQian.(map[string]interface{})["qian"].(string)
	baseResp.Res = map[string]interface{}{
		"order": qianOrder,
		"name":  guanYinLingQian.(map[string]interface{})["qian"].(string),
		"qian":  guanYinLingQian,
	}
	return ctx.WriteJson(baseResp)
}

/**
 * @Description 通过签序号指定一个卦给用户返回去
 * @Params  [ctx]
 * @Returns  [baseResp]
 * @Creator  Ericon Freeman
 * @Date  2023/3/7 5:36 下午
 */
func QueryOrderGuanYinLingQianFunc(ctx dotweb.Context) error {

	qianOrder, _ := strconv.Atoi(ctx.GetRouterName("order"))
	guanYinLingQian := util.GuanYinLingQian[int(qianOrder)]
	var baseResp dto.BaseResponse
	baseResp.Msg = "抽签完成，恭喜您抽得" + guanYinLingQian.(map[string]interface{})["qian"].(string)
	baseResp.Res = map[string]interface{}{
		"order": qianOrder,
		"name":  guanYinLingQian.(map[string]interface{})["qian"].(string),
		"qian":  guanYinLingQian,
	}
	return ctx.WriteJson(baseResp)

}
