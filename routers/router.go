// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"goMongodbAPI/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/connection",
			beego.NSInclude( // NSInclude is designed for Annotation Router
				&controllers.ConnectionController{},
			),
		),
	)
	beego.AddNamespace(ns)

	// beego.Router("/database/:connectionId", &controllers.DatabaseCallsController{}, "get:GetAll")
	// beego.Router("/collection/:connectionId/:database", &controllers.CollectionCallsController{}, "get:GetAll")
}
