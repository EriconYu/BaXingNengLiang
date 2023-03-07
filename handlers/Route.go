package handlers

import (
	"github.com/devfeel/dotweb"
)

/**

 * @Description 初始化路由

 * @Params

 * @Returns

 * @Creator  Ericon Freeman

 * @Date  2022-05-24

 */
func InitRoute(server *dotweb.HttpServer) {

	// 先做一个api v1版本的group
	apiV1Group := server.Group("/api/v1")

	// 1、手机号码分析类
	gPhoneNumber := apiV1Group.Group("/baxing/phoneNumber")
	gPhoneNumber.GET("/:phoneNumber", phoneNumberIndex)

	// 2、黄历类
	gCalendar := apiV1Group.Group("/calendar")
	gCalendar.GET("/today", TodayIndexFunc)  //今日宜忌

	// 3、抽签类
	gQian := apiV1Group.Group("/qian")
	gQian.GET("/64gua/random" , Random64GuaFunc)
	gQian.GET("/64gua/name/:name" , QueryName64GuaFunc)
	gQian.GET("/64gua/order/:order" , QueryOrder64GuaFunc)
}
