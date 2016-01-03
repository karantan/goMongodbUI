package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"] = append(beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"] = append(beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"],
		beego.ControllerComments{
			"Get",
			`/:connectionId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"] = append(beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"] = append(beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"],
		beego.ControllerComments{
			"Put",
			`/:connectionId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"] = append(beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"],
		beego.ControllerComments{
			"Delete",
			`/:connectionId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"] = append(beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"],
		beego.ControllerComments{
			"GetDatabases",
			`/:connectionId/databases`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"] = append(beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"],
		beego.ControllerComments{
			"GetCollections",
			`/:connectionId/:database/collections`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"] = append(beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"],
		beego.ControllerComments{
			"QueryCollection",
			`/:connectionId/:database/:collection/query`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"] = append(beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"],
		beego.ControllerComments{
			"CreateCollection",
			`/:connectionId/:database/:collection/create`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"] = append(beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"],
		beego.ControllerComments{
			"DropCollection",
			`/:connectionId/:database/:collection/drop`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"] = append(beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"],
		beego.ControllerComments{
			"InsertDocuments",
			`/:connectionId/:database/:collection/insert`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"] = append(beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"],
		beego.ControllerComments{
			"UpdateDocuments",
			`/:connectionId/:database/:collection/update`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"] = append(beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"],
		beego.ControllerComments{
			"UpdateIdDocument",
			`/:connectionId/:database/:collection/:document_id/update`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"] = append(beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"],
		beego.ControllerComments{
			"RemoveDocuments",
			`/:connectionId/:database/:collection/remove`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"] = append(beego.GlobalControllerRouter["goMongodbAPI/controllers:ConnectionController"],
		beego.ControllerComments{
			"RemoveIdDocument",
			`/:connectionId/:database/:collection/:document_id/remove`,
			[]string{"delete"},
			nil})

}
