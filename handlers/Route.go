package handlers

import "github.com/devfeel/dotweb"

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

}
