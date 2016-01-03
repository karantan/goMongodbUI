package controllers

import (
	"encoding/json"
	"fmt"
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

// @Title create
// @Description create connection
// @Param	body		body 	models.Connection	true		"The connection content"
// @Success 200 {string} models.Connection.ConnectionId
// @Failure 400 body is empty
// @router / [post]
func (o *ConnectionController) Post() {
	// TODO: add validation
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
// @Failure 400 :connectionId is invalid
// @router /:connectionId [get]
func (o *ConnectionController) Get() {
	var err_msg string
	connectionId := o.Ctx.Input.Params[":connectionId"]
	if connectionId != "" {
		ob, err := models.GetOne(connectionId)
		if err != nil {
			err_msg = fmt.Sprintf("Error: (%s)", err)
		} else {
			o.Data["json"] = ob
		}
	}

	if err_msg != "" {
		o.CustomAbort(400, err_msg)
		return
	}

	o.ServeJson()
}

// @Title GetAll
// @Description get all connections
// @Success 200 {connection} models.Connection
// @Failure 500 :internal server error
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
// @Failure 400 :connectionId is invalid
// @router /:connectionId [put]
func (o *ConnectionController) Put() {
	var err_msg string
	connectionId := o.Ctx.Input.Params[":connectionId"]
	var ob models.Connection
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(connectionId, ob.Name, ob.Address, ob.Port)
	if err != nil {
		err_msg = fmt.Sprintf("Error: (%s)", err)
	} else {
		o.Data["json"] = "update success!"
	}

	if err_msg != "" {
		o.CustomAbort(400, err_msg)
		return
	}

	o.ServeJson()
}

// @Title delete
// @Description delete the connection
// @Param	connectionId		path 	string	true		"The connectionId you want to delete"
// @Success 200 {string} delete success!
// @Failure 400 connectionId is invalid
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
// @Failure 400 connectionId is invalid
// @router /:connectionId/databases [get]
func (o *ConnectionController) GetDatabases() {
	var err_msg string

	connectionId := o.Ctx.Input.Params[":connectionId"]
	ob, err := models.GetOne(connectionId)
	if err != nil {
		err_msg = fmt.Sprintf("Error: (%s)", err)
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
			err_msg = fmt.Sprintf("Error: (%s)", err)
		} else {
			// Optional. Switch the session to a monotonic behavior.
			session.SetMode(mgo.Monotonic, true)
			dbs, _ := session.DatabaseNames()
			o.Data["json"] = dbs
		}
	}

	if err_msg != "" {
		o.CustomAbort(400, err_msg)
		return
	}

	o.ServeJson()
}

// @Title collections
// @Description get all collections
// @Param connectionId path string true "Set connectionId"
// @Param database path string true "Set database name"
// @Success 200 {list} success!
// @Failure 400 connectionId or database is invalid
// @router /:connectionId/:database/collections [get]
func (o *ConnectionController) GetCollections() {
	var err_msg string

	connectionId := o.Ctx.Input.Params[":connectionId"]
	database := o.Ctx.Input.Params[":database"]

	ob, err := models.GetOne(connectionId)
	if err != nil {
		err_msg = fmt.Sprintf("Error: (%s)", err)
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
			err_msg = fmt.Sprintf("Error: (%s)", err)
		} else {
			// Optional. Switch the session to a monotonic behavior.
			session.SetMode(mgo.Monotonic, true)
			db := session.DB(database)
			cn, _ := db.CollectionNames()
			o.Data["json"] = cn
		}
	}

	if err_msg != "" {
		o.CustomAbort(400, err_msg)
		return
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
// @Failure 400 connectionId or database or collection is invalid
// @router /:connectionId/:database/:collection/query [post]
func (o *ConnectionController) QueryCollection() {
	var err_msg string

	connectionId := o.Ctx.Input.Params[":connectionId"]
	database := o.Ctx.Input.Params[":database"]
	collection := o.Ctx.Input.Params[":collection"]

	var query bson.M
	json.Unmarshal(o.Ctx.Input.RequestBody, &query)
	log.Printf("[*] Executing query: %s", query)

	ob, err := models.GetOne(connectionId)
	if err != nil {
		err_msg = fmt.Sprintf("Error: (%s)", err)
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
			err_msg = fmt.Sprintf("Error: (%s)", err)
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

	if err_msg != "" {
		o.CustomAbort(400, err_msg)
		return
	}

	o.ServeJson()
}

// @Title create collection
// @Description create collection
// @Param connectionId path string true "Set connectionId"
// @Param database path string true "Set database name"
// @Param collection path string true "Set collection name"
// @Success 200 {string} success!
// @Failure 400 connectionId or database or collection is invalid
// @router /:connectionId/:database/:collection/create [post]
func (o *ConnectionController) CreateCollection() {
	var err_msg string

	connectionId := o.Ctx.Input.Params[":connectionId"]
	database := o.Ctx.Input.Params[":database"]
	collection := o.Ctx.Input.Params[":collection"]

	ob, err := models.GetOne(connectionId)
	if err != nil {
		err_msg = fmt.Sprintf("Error: (%s)", err)
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
			err_msg = fmt.Sprintf("Error: (%s)", err)
		} else {
			// Optional. Switch the session to a monotonic behavior.
			session.SetMode(mgo.Monotonic, true)

			collection_info := mgo.CollectionInfo{}
			session.DB(database).C(collection).Create(&collection_info)
			o.Data["json"] = "Collection created"
		}
	}

	if err_msg != "" {
		o.CustomAbort(400, err_msg)
		return
	}

	o.ServeJson()
}

// @Title drop collection
// @Description drop collection
// @Param connectionId path string true "Set connectionId"
// @Param database path string true "Set database name"
// @Param collection path string true "Set collection name"
// @Success 200 {string} success!
// @Failure 400 connectionId or database or collection is invalid
// @router /:connectionId/:database/:collection/drop [delete]
func (o *ConnectionController) DropCollection() {
	var err_msg string

	connectionId := o.Ctx.Input.Params[":connectionId"]
	database := o.Ctx.Input.Params[":database"]
	collection := o.Ctx.Input.Params[":collection"]

	ob, err := models.GetOne(connectionId)
	if err != nil {
		err_msg = fmt.Sprintf("Error: (%s)", err)
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
			err_msg = fmt.Sprintf("Error: (%s)", err)
		} else {
			// Optional. Switch the session to a monotonic behavior.
			session.SetMode(mgo.Monotonic, true)

			session.DB(database).C(collection).DropCollection()
			o.Data["json"] = "Collection dropped"
		}
	}

	if err_msg != "" {
		o.CustomAbort(400, err_msg)
		return
	}

	o.ServeJson()
}

// @Title insert documents
// @Description insert documents. It must be list of dicts.
// @Param connectionId path string true "Set connectionId"
// @Param database path string true "Set database name"
// @Param collection path string true "Set collection name"
// @Param document body string true "MongoDB documents"
// @Success 200 {string} success!
// @Failure 400 connectionId or database or collection is invalid
// @router /:connectionId/:database/:collection/insert [post]
func (o *ConnectionController) InsertDocuments() {
	var err_msg string

	connectionId := o.Ctx.Input.Params[":connectionId"]
	database := o.Ctx.Input.Params[":database"]
	collection := o.Ctx.Input.Params[":collection"]

	// NOTE: use http://www.generatedata.com/ for test data.
	var documents []bson.M
	json.Unmarshal(o.Ctx.Input.RequestBody, &documents)
	log.Printf("[*] Inserting documents: %s", documents)

	ob, err := models.GetOne(connectionId)
	if err != nil {
		err_msg = fmt.Sprintf("Error: (%s)", err)
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
			err_msg = fmt.Sprintf("Error: (%s)", err)
		} else {
			// Optional. Switch the session to a monotonic behavior.
			session.SetMode(mgo.Monotonic, true)

			var json_msg []string
			var msg string
			for i, document := range documents {
				if len(document) == 0 {
					msg = fmt.Sprintf("Document #%d is empty. I refuse to insert empty document.", i+1)
				} else {
					err := session.DB(database).C(collection).Insert(document)
					if err != nil {
						err_msg = fmt.Sprintf("Error: (%s)", err)
					} else {
						msg = fmt.Sprintf("Document #%d inserted", i+1)
					}
				}
				json_msg = append(json_msg, msg)
			}
			o.Data["json"] = json_msg
		}
	}

	if err_msg != "" {
		o.CustomAbort(400, err_msg)
		return
	}

	o.ServeJson()
}

// @Title update document
// @Description update one document that matches doc_selector.
// @Param connectionId path string true "Set connectionId"
// @Param database path string true "Set database name"
// @Param collection path string true "Set collection name"
// @Param document body string true "{"selector": doc_selector, "document": MongoDB_document}"
// @Success 200 {string} success!
// @Failure 400 connectionId or database or collection or document_selector is invalid
// @router /:connectionId/:database/:collection/update [put]
func (o *ConnectionController) UpdateDocuments() {
	var err_msg string

	connectionId := o.Ctx.Input.Params[":connectionId"]
	database := o.Ctx.Input.Params[":database"]
	collection := o.Ctx.Input.Params[":collection"]

	var document bson.M
	json.Unmarshal(o.Ctx.Input.RequestBody, &document)

	ob, err := models.GetOne(connectionId)
	if err != nil {
		err_msg = fmt.Sprintf("Error: (%s)", err)
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
			err_msg = fmt.Sprintf("Error: (%s)", err)
		} else {
			// Optional. Switch the session to a monotonic behavior.
			session.SetMode(mgo.Monotonic, true)
			change := mgo.Change{
				Update:    document["document"],
				ReturnNew: true,
			}
			var result bson.M
			info, err := session.DB(database).C(collection).Find(document["selector"]).Apply(change, &result)
			if err != nil {
				err_msg = fmt.Sprintf("Error: (%s)", err)
			} else {
				o.Data["json"] = info
			}
		}
	}

	if err_msg != "" {
		o.CustomAbort(400, err_msg)
		return
	}

	o.ServeJson()
}

// @Title update document
// @Description update document by ID
// @Param connectionId path string true "Set connectionId"
// @Param database path string true "Set database name"
// @Param collection path string true "Set collection name"
// @Param document_id path string true "MongoDB document _id"
// @Param document body string true "MongoDB document"
// @Success 200 {string} success!
// @Failure 400 connectionId or database or collection or document_id is invalid
// @router /:connectionId/:database/:collection/:document_id/update [put]
func (o *ConnectionController) UpdateIdDocument() {
	var err_msg string

	connectionId := o.Ctx.Input.Params[":connectionId"]
	database := o.Ctx.Input.Params[":database"]
	collection := o.Ctx.Input.Params[":collection"]
	document_id := o.Ctx.Input.Params[":document_id"]
	if !bson.IsObjectIdHex(document_id) {
		o.CustomAbort(400, "document_id is not ObjectId")
		return
	}
	oid := bson.ObjectIdHex(document_id)

	var document bson.M
	json.Unmarshal(o.Ctx.Input.RequestBody, &document)

	ob, err := models.GetOne(connectionId)
	if err != nil {
		err_msg = fmt.Sprintf("Error: (%s)", err)
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
			err_msg = fmt.Sprintf("Error: (%s)", err)
		} else {
			// Optional. Switch the session to a monotonic behavior.
			session.SetMode(mgo.Monotonic, true)

			err := session.DB(database).C(collection).UpdateId(oid, document)
			if err != nil {
				err_msg = fmt.Sprintf("Error: (%s)", err)
			} else {
				o.Data["json"] = "Document updated"
			}

		}
	}

	if err_msg != "" {
		o.CustomAbort(400, err_msg)
		return
	}

	o.ServeJson()
}

// @Title remove documents
// @Description remove documents
// @Param connectionId path string true "Set connectionId"
// @Param database path string true "Set database name"
// @Param collection path string true "Set collection name"
// @Param document_selector body string true "MongoDB document selector"
// @Success 200 {string} success!
// @Failure 400 connectionId or database or collection or document selector is invalid
// @router /:connectionId/:database/:collection/remove [delete]
func (o *ConnectionController) RemoveDocuments() {
	var err_msg string

	connectionId := o.Ctx.Input.Params[":connectionId"]
	database := o.Ctx.Input.Params[":database"]
	collection := o.Ctx.Input.Params[":collection"]

	var document_selector bson.M
	json.Unmarshal(o.Ctx.Input.RequestBody, &document_selector)
	log.Printf("[*] Removing documents: %s", document_selector)

	ob, err := models.GetOne(connectionId)
	if err != nil {
		err_msg = fmt.Sprintf("Error: (%s)", err)
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
			err_msg = fmt.Sprintf("Error: (%s)", err)
		} else {
			// Optional. Switch the session to a monotonic behavior.
			session.SetMode(mgo.Monotonic, true)
			if len(document_selector) == 0 {
				o.Data["json"] = "I refuse to remove all documents. Use drop collection if you really want to remove all documents."
			} else {
				changeInfo, err := session.DB(database).C(collection).RemoveAll(document_selector)
				if err != nil {
					err_msg = fmt.Sprintf("Error: (%s)", err)
				} else {
					o.Data["json"] = changeInfo
				}
			}

		}
	}

	if err_msg != "" {
		o.CustomAbort(400, err_msg)
		return
	}

	o.ServeJson()
}

// @Title remove document
// @Description remove document by ID
// @Param connectionId path string true "Set connectionId"
// @Param database path string true "Set database name"
// @Param collection path string true "Set collection name"
// @Param document_id path string true "MongoDB document _id"
// @Success 200 {string} success!
// @Failure 400 connectionId or database or collection or document_id is invalid
// @router /:connectionId/:database/:collection/:document_id/remove [delete]
func (o *ConnectionController) RemoveIdDocument() {
	var err_msg string

	connectionId := o.Ctx.Input.Params[":connectionId"]
	database := o.Ctx.Input.Params[":database"]
	collection := o.Ctx.Input.Params[":collection"]
	document_id := o.Ctx.Input.Params[":document_id"]
	if !bson.IsObjectIdHex(document_id) {
		o.CustomAbort(400, "document_id is not ObjectId")
		return
	}
	oid := bson.ObjectIdHex(document_id)

	ob, err := models.GetOne(connectionId)
	if err != nil {
		err_msg = fmt.Sprintf("Error: (%s)", err)
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
			err_msg = fmt.Sprintf("Error: (%s)", err)
		} else {
			// Optional. Switch the session to a monotonic behavior.
			session.SetMode(mgo.Monotonic, true)

			err := session.DB(database).C(collection).RemoveId(oid)
			if err != nil {
				err_msg = fmt.Sprintf("Error: (%s)", err)
			} else {
				o.Data["json"] = "Document removed"
			}

		}
	}

	if err_msg != "" {
		o.CustomAbort(400, err_msg)
		return
	}

	o.ServeJson()
}
