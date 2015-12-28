package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/connection","description":"Operations about connection\n"}],"info":{"title":"beego Test API","description":"beego has a very cool tools to autogenerate documents for your API","contact":"astaxie@gmail.com","termsOfServiceUrl":"http://beego.me/","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}}`
    Subapi string = `{"/connection":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/connection","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"create","type":"","summary":"create connection","parameters":[{"paramType":"body","name":"body","description":"\"The connection content\"","dataType":"Connection","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Connection.ConnectionId","responseModel":""},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/:connectionId","description":"","operations":[{"httpMethod":"GET","nickname":"Get","type":"","summary":"find connection by connectionid","parameters":[{"paramType":"path","name":"connectionId","description":"\"the connectionid you want to get\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Connection","responseModel":""},{"code":403,"message":":connectionId is empty","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"GetAll","type":"","summary":"get all connections","responseMessages":[{"code":200,"message":"models.Connection","responseModel":""},{"code":403,"message":":connectionId is empty","responseModel":""}]}]},{"path":"/:connectionId","description":"","operations":[{"httpMethod":"PUT","nickname":"update","type":"","summary":"update the connection","parameters":[{"paramType":"path","name":"connectionId","description":"\"The connectionid you want to update\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"body","name":"body","description":"\"The body\"","dataType":"Connection","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Connection","responseModel":""},{"code":403,"message":":connectionId is empty","responseModel":""}]}]},{"path":"/:connectionId","description":"","operations":[{"httpMethod":"DELETE","nickname":"delete","type":"","summary":"delete the connection","parameters":[{"paramType":"path","name":"connectionId","description":"\"The connectionId you want to delete\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"success!","responseModel":""},{"code":403,"message":"connectionId is empty","responseModel":""}]}]},{"path":"/:connectionId/check","description":"","operations":[{"httpMethod":"GET","nickname":"Check","type":"","summary":"tries to connect to mongodb","parameters":[{"paramType":"path","name":"connectionId","description":"\"the connectionId you want to connect\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Connection","responseModel":""},{"code":403,"message":":connectionId is empty","responseModel":""}]}]}]}}`
    BasePath string= "/v1"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.ApiDeclaration

func init() {
	if beego.EnableDocs {
		err := json.Unmarshal([]byte(Rootinfo), &rootapi)
		if err != nil {
			beego.Error(err)
		}
		err = json.Unmarshal([]byte(Subapi), &apilist)
		if err != nil {
			beego.Error(err)
		}
		beego.GlobalDocApi["Root"] = rootapi
		for k, v := range apilist {
			for i, a := range v.Apis {
				a.Path = urlReplace(k + a.Path)
				v.Apis[i] = a
			}
			v.BasePath = BasePath
			beego.GlobalDocApi[strings.Trim(k, "/")] = v
		}
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
