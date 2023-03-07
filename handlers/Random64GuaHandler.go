package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/EriconYu/BaXingNengLiang/dto"
	"github.com/EriconYu/BaXingNengLiang/util"
	"github.com/devfeel/dotweb"
)

/**
 * @Description random一个卦给用户返回去
 * @Params  [ctx]
 * @Returns  [baseResp]
 * @Creator  Ericon Freeman
 * @Date  2023/3/7 5:36 下午
 */
func Random64GuaFunc(ctx dotweb.Context) error {

	guaOrder := time.Now().UnixNano() % 64
	if guaOrder == 0 {
		guaOrder = 64
	}
	guaOrderString := strconv.Itoa(int(guaOrder))
	yaoci := util.Gua64_YaoCi[guaOrderString]
	jiexi := util.Gua64_JieXi[guaOrderString]
	var baseResp dto.BaseResponse
	baseResp.Res = map[string]interface{}{
		"order": guaOrder,
		"name":  yaoci.(map[string]interface{})["Name"],
		"yaoci": yaoci,
		"jiexi": jiexi,
	}
	return ctx.WriteJson(baseResp)
}

/**
 * @Description 通过卦名指定一个卦给用户返回去
 * @Params  [ctx]
 * @Returns  [baseResp]
 * @Creator  Ericon Freeman
 * @Date  2023/3/7 5:36 下午
 */
func QueryName64GuaFunc(ctx dotweb.Context) error {

	name := ctx.GetRouterName("name")
	var baseResp dto.BaseResponse
	for order, content := range util.Gua64_YaoCi {
		ret, _ := content.(map[string]interface{})
		if name == ret["Name"] {
			log.Println("find the Gua ： ", name)
			yaoci := util.Gua64_YaoCi[order]
			jiexi := util.Gua64_JieXi[order]
			baseResp.Res = map[string]interface{}{
				"order": order,
				"name":  yaoci.(map[string]interface{})["Name"],
				"yaoci": yaoci,
				"jiexi": jiexi,
			}
			return ctx.WriteJson(baseResp)
		}
	}
	log.Println("Cannot find the Gua ： ", name)
	ctx.Response().SetStatusCode(http.StatusBadRequest)
	baseResp.Msg = "未能找到卦名，请仔细核对输入是否有误"
	return ctx.WriteJson(baseResp)

}

/**
 * @Description 通过卦序号指定一个卦给用户返回去
 * @Params  [ctx]
 * @Returns  [baseResp]
 * @Creator  Ericon Freeman
 * @Date  2023/3/7 5:36 下午
 */
func QueryOrder64GuaFunc(ctx dotweb.Context) error {

	order := ctx.GetRouterName("order")
	var baseResp dto.BaseResponse
	yaoci := util.Gua64_YaoCi[order]
	jiexi := util.Gua64_JieXi[order]
	baseResp.Res = map[string]interface{}{
		"order": order,
		"name":  yaoci.(map[string]interface{})["Name"],
		"yaoci": yaoci,
		"jiexi": jiexi,
	}
	return ctx.WriteJson(baseResp)
}
