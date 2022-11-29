// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"bee-api-demo/controllers/user"
	"github.com/beego/beego/v2/server/web/filter/cors"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 跨域解决方案
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// 允许访问所有源
		AllowAllOrigins: true,
		// 可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 指的是允许的Header的种类
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))

	// todo: 登录验证Filter
	// todo: 权限检查Filter

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&user.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
