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
func RandomZhuGeShenSuanFunc(ctx dotweb.Context) error {

	qianOrder := time.Now().UnixNano() % 384
	if qianOrder == 0 {
		qianOrder = 384
	}
	qianOrderString := strconv.Itoa(int(qianOrder))
	zhuGeShenSuan := util.ZhuGeShenSuan[qianOrderString]
	var baseResp dto.BaseResponse
	baseResp.Msg = "抽签完成，恭喜您抽得" + zhuGeShenSuan.(map[string]interface{})["qian"].(string)
	baseResp.Res = map[string]interface{}{
		"order": qianOrder,
		"name":  zhuGeShenSuan.(map[string]interface{})["qian"].(string),
		"qian":  zhuGeShenSuan,
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
func QueryOrderZhuGeShenSuanFunc(ctx dotweb.Context) error {

	qianOrder := ctx.GetRouterName("order")
	zhuGeShenSuan := util.ZhuGeShenSuan[qianOrder]
	var baseResp dto.BaseResponse
	baseResp.Msg = "抽签完成，恭喜您抽得" + zhuGeShenSuan.(map[string]interface{})["qian"].(string)
	baseResp.Res = map[string]interface{}{
		"order": qianOrder,
		"name":  zhuGeShenSuan.(map[string]interface{})["qian"].(string),
		"qian":  zhuGeShenSuan,
	}
	return ctx.WriteJson(baseResp)

}
