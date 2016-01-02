package controllers

import (
	"encoding/json"
	"goMongodbAPI/models"
	"log"
	"time"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Operations about connection
type ConnectionController struct {
	beego.Controller
}

// Error JSON resp
type ErrResponse struct {
	Message string
	Code    int
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

// @Title databases
// @Description get all databases
// @Param connectionId path string true "Fetch databases from the connectionId"
// @Success 200 {list} success!
// @Failure 403 connectionId is empty
// @router /:connectionId/databases [get]
func (o *ConnectionController) GetDatabases() {
	connectionId := o.Ctx.Input.Params[":connectionId"]
	ob, err := models.GetOne(connectionId)
	if err != nil {
		o.Data["json"] = "Connection not found"
		log.Printf("connectionId: %s\n", err)
	} else {
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{ob.Address},
			Timeout:  5 * time.Second,
			Username: ob.Username,
			Password: ob.Password,
		}

		session, err := mgo.DialWithInfo(mongoDBDialInfo)
		defer session.Close()
		if err != nil {
			o.Data["json"] = err
			log.Printf("CreateSession: %s\n", err)
		} else {
			// Optional. Switch the session to a monotonic behavior.
			session.SetMode(mgo.Monotonic, true)
			dbs, _ := session.DatabaseNames()
			o.Data["json"] = dbs
		}
	}

	o.ServeJson()
}

// @Title collections
// @Description get all collections
// @Param connectionId path string true "Set connectionId"
// @Param database path string true "Set database name"
// @Success 200 {list} success!
// @Failure 403 connectionId or database is empty
// @router /:connectionId/:database/collections [get]
func (o *ConnectionController) GetCollections() {
	connectionId := o.Ctx.Input.Params[":connectionId"]
	database := o.Ctx.Input.Params[":database"]

	ob, err := models.GetOne(connectionId)
	if err != nil {
		o.Data["json"] = "Connection not found"
	} else {
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{ob.Address},
			Timeout:  5 * time.Second,
			Username: ob.Username,
			Password: ob.Password,
		}

		session, err := mgo.DialWithInfo(mongoDBDialInfo)
		defer session.Close()
		if err != nil {
			o.Data["json"] = err
		} else {
			// Optional. Switch the session to a monotonic behavior.
			session.SetMode(mgo.Monotonic, true)
			db := session.DB(database)
			cn, _ := db.CollectionNames()
			o.Data["json"] = cn
		}
	}

	o.ServeJson()
}

// @Title query collection
// @Description query collection
// @Param connectionId path string true "Set connectionId"
// @Param database path string true "Set database name"
// @Param collection path string true "Set collection name"
// @Param query body string false "MongoDB query"
// @Success 200 {list} success!
// @Failure 403 connectionId or database or collection is empty
// @router /:connectionId/:database/:collection/query [post]
func (o *ConnectionController) QueryCollection() {
	connectionId := o.Ctx.Input.Params[":connectionId"]
	database := o.Ctx.Input.Params[":database"]
	collection := o.Ctx.Input.Params[":collection"]

	var query bson.M
	json.Unmarshal(o.Ctx.Input.RequestBody, &query)
	log.Printf("[*] Executing query: %s", query)

	ob, err := models.GetOne(connectionId)
	if err != nil {
		o.Data["json"] = "Connection not found"
	} else {
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{ob.Address},
			Timeout:  5 * time.Second,
			Username: ob.Username,
			Password: ob.Password,
		}

		session, err := mgo.DialWithInfo(mongoDBDialInfo)
		defer session.Close()
		if err != nil {
			o.Data["json"] = err
		} else {
			// Optional. Switch the session to a monotonic behavior.
			session.SetMode(mgo.Monotonic, true)

			//  convenient alias for a map[string]interface{} map, useful for dealing with BSON in a native way
			var m []bson.M
			db := session.DB(database).C(collection)
			db.Find(query).All(&m)
			o.Data["json"] = m
		}
	}

	o.ServeJson()
}

// @Title create collection
// @Description create collection
// @Param connectionId path string true "Set connectionId"
// @Param database path string true "Set database name"
// @Param collection path string true "Set collection name"
// @Success 200 {string} success!
// @Failure 403 connectionId or database or collection is empty
// @router /:connectionId/:database/:collection/create [post]
func (o *ConnectionController) CreateCollection() {
	connectionId := o.Ctx.Input.Params[":connectionId"]
	database := o.Ctx.Input.Params[":database"]
	collection := o.Ctx.Input.Params[":collection"]

	ob, err := models.GetOne(connectionId)
	if err != nil {
		o.Data["json"] = "Connection not found"
	} else {
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{ob.Address},
			Timeout:  5 * time.Second,
			Username: ob.Username,
			Password: ob.Password,
		}

		session, err := mgo.DialWithInfo(mongoDBDialInfo)
		defer session.Close()
		if err != nil {
			o.Data["json"] = err
		} else {
			// Optional. Switch the session to a monotonic behavior.
			session.SetMode(mgo.Monotonic, true)

			collection_info := mgo.CollectionInfo{}
			session.DB(database).C(collection).Create(&collection_info)
			o.Data["json"] = "Collection created"
		}
	}

	o.ServeJson()
}

// @Title insert document
// @Description insert document
// @Param connectionId path string true "Set connectionId"
// @Param database path string true "Set database name"
// @Param collection path string true "Set collection name"
// @Param document body string true "MongoDB document"
// @Success 200 {string} success!
// @Failure 403 connectionId or database or collection is empty
// @router /:connectionId/:database/:collection/insert [post]
func (o *ConnectionController) InsertDocument() {
	connectionId := o.Ctx.Input.Params[":connectionId"]
	database := o.Ctx.Input.Params[":database"]
	collection := o.Ctx.Input.Params[":collection"]

	var document bson.M
	json.Unmarshal(o.Ctx.Input.RequestBody, &document)
	log.Printf("[*] Inserting document: %s", document)

	ob, err := models.GetOne(connectionId)
	if err != nil {
		o.Data["json"] = "Connection not found"
	} else {
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{ob.Address},
			Timeout:  5 * time.Second,
			Username: ob.Username,
			Password: ob.Password,
		}

		session, err := mgo.DialWithInfo(mongoDBDialInfo)
		defer session.Close()
		if err != nil {
			o.Data["json"] = err
		} else {
			// Optional. Switch the session to a monotonic behavior.
			session.SetMode(mgo.Monotonic, true)
			if len(document) == 0 {
				o.Data["json"] = "I refuse to insert empty document."
			} else {
				err := session.DB(database).C(collection).Insert(document)
				if err != nil {
					o.Data["json"] = err
				} else {
					o.Data["json"] = "Document inserted"
				}
			}

		}
	}

	o.ServeJson()
}
