package handlers

import (
	"strconv"
	"time"

	"github.com/EriconYu/BaXingNengLiang/dto"
	"github.com/EriconYu/BaXingNengLiang/util"
	"github.com/devfeel/dotweb"
)

/**
 * @Description random一个诸葛神签给用户返回去
 * @Params  [ctx]
 * @Returns  [baseResp]
 * @Creator  Ericon Freeman
 * @Date  2023/3/7 5:36 下午
 */
func RandomZhuGeShenQianFunc(ctx dotweb.Context) error {

	qianOrder := time.Now().UnixNano() % 64
	if qianOrder == 0 {
		qianOrder = 64
	}
	qianOrderString := strconv.Itoa(int(qianOrder))
	zhuGeShenQian := util.ZhuGeShenQian[qianOrderString]
	var baseResp dto.BaseResponse
	baseResp.Msg = "抽签完成，恭喜您抽得" + zhuGeShenQian.(map[string]interface{})["qian"].(string)
	baseResp.Res = map[string]interface{}{
		"order": qianOrder,
		"name":  zhuGeShenQian.(map[string]interface{})["qian"].(string),
		"qian":  zhuGeShenQian,
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
func QueryOrderZhuGeShenQianFunc(ctx dotweb.Context) error {

	qianOrder := ctx.GetRouterName("order")
	zhuGeShenQian := util.ZhuGeShenQian[qianOrder]
	var baseResp dto.BaseResponse
	baseResp.Msg = "抽签完成，恭喜您抽得" + zhuGeShenQian.(map[string]interface{})["qian"].(string)
	baseResp.Res = map[string]interface{}{
		"order": qianOrder,
		"name":  zhuGeShenQian.(map[string]interface{})["qian"].(string),
		"qian":  zhuGeShenQian,
	}
	return ctx.WriteJson(baseResp)

}
