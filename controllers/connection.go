package controllers

import (
	"encoding/json"
	"goMongodbAPI/models"
	"log"
	"time"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

// Operations about connection
type ConnectionController struct {
	beego.Controller
}

// @Title create
// @Description create connection
// @Param	body		body 	models.Connection	true		"The connection content"
// @Success 200 {string} models.Connection.ConnectionId
// @Failure 403 body is empty
// @router / [post]
func (o *ConnectionController) Post() {
	var ob models.Connection
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	connectionid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ConnectionId": connectionid}
	o.ServeJson()
}

// @Title Get
// @Description find connection by connectionid
// @Param	connectionId		path 	string	true		"the connectionid you want to get"
// @Success 200 {connection} models.Connection
// @Failure 403 :connectionId is empty
// @router /:connectionId [get]
func (o *ConnectionController) Get() {
	connectionId := o.Ctx.Input.Params[":connectionId"]
	if connectionId != "" {
		ob, err := models.GetOne(connectionId)
		if err != nil {
			o.Data["json"] = err
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJson()
}

// @Title GetAll
// @Description get all connections
// @Success 200 {connection} models.Connection
// @Failure 403 :connectionId is empty
// @router / [get]
func (o *ConnectionController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJson()
}

// @Title update
// @Description update the connection
// @Param	connectionId		path 	string	true		"The connectionid you want to update"
// @Param	body		body 	models.Connection	true		"The body"
// @Success 200 {connection} models.Connection
// @Failure 403 :connectionId is empty
// @router /:connectionId [put]
func (o *ConnectionController) Put() {
	connectionId := o.Ctx.Input.Params[":connectionId"]
	var ob models.Connection
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(connectionId, ob.Name, ob.Address, ob.Port)
	if err != nil {
		o.Data["json"] = err
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJson()
}

// @Title delete
// @Description delete the connection
// @Param	connectionId		path 	string	true		"The connectionId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 connectionId is empty
// @router /:connectionId [delete]
func (o *ConnectionController) Delete() {
	connectionId := o.Ctx.Input.Params[":connectionId"]
	models.Delete(connectionId)
	o.Data["json"] = "delete success!"
	o.ServeJson()
}

type Person struct {
	Name  string
	Phone string
}

// @Title Check
// @Description tries to connect to mongodb
// @Param	connectionId		path 	string	true		"the connectionId you want to connect"
// @Success 200 {connection} models.Connection
// @Failure 403 :connectionId is empty
// @router /:connectionId/check [get]
func (o *ConnectionController) Check() {
	// http://blog.mongodb.org/post/80579086742/running-mongodb-queries-concurrently-with-go
	connectionId := o.Ctx.Input.Params[":connectionId"]
	if connectionId != "" {
		ob, _ := models.GetOne(connectionId)
		// We need this object to establish a session to our MongoDB.
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:   []string{ob.Address},
			Timeout: 60 * time.Second,
			// Username: AuthUserName,
			// Password: AuthPassword,
		}

		// Create a session which maintains a pool of socket connections
		// to our MongoDB.
		session, err := mgo.DialWithInfo(mongoDBDialInfo)
		if err != nil {
			log.Fatalf("CreateSession: %s\n", err)
		}
		defer session.Close()

		// Optional. Switch the session to a monotonic behavior.
		session.SetMode(mgo.Monotonic, true)

		// c := session.DB("test")
		// cn, _ := c.CollectionNames()
		dbs, _ := session.DatabaseNames()
		o.Data["json"] = dbs
	}
	o.ServeJson()
}
