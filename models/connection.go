package models

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// the proper way to use mgo is to have one sesson
// and clone it every time you need to do something
var (
	Connections map[string]*Connection
	mgoSession  *mgo.Session
)

type Connection struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string
	Address  string
	Port     int
	Username string
	Password string
}

func init() {
	// init connection with mongodb
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:   []string{"localhost"},
		Timeout: 5 * time.Second,
		// Username: ob.Username,
		// Password: ob.Password,
	}
	sesssion, err := mgo.DialWithInfo(mongoDBDialInfo)

	if err != nil {
		log.Fatalf("Error: (%s)\n", err)
	}
	mgoSession = sesssion
}

func AddOne(Connection Connection) Connection {
	session := mgoSession.Clone()
	defer session.Close()
	session.DB("goMongodbUI").C("connections").Insert(Connection)

	return Connection
}

func GetOne(ConnectionId string) (Connection, error) {
	session := mgoSession.Clone()
	defer session.Close()
	var connection Connection

	if !bson.IsObjectIdHex(ConnectionId) {
		return connection, fmt.Errorf("connectionId: %s is not ObjectId", ConnectionId)
	}

	oid := bson.ObjectIdHex(ConnectionId)
	session.DB("goMongodbUI").C("connections").FindId(oid).One(&connection)

	return connection, nil
}

func GetAll() []Connection {
	session := mgoSession.Clone()
	defer session.Close()
	var connections []Connection
	session.DB("goMongodbUI").C("connections").Find(bson.M{}).All(&connections)
	return connections
}

func Update(connectionId string, name string, address string, port int, username string, password string) error {
	session := mgoSession.Clone()
	defer session.Close()

	if !bson.IsObjectIdHex(connectionId) {
		return fmt.Errorf("connectionId: %s is not ObjectId", connectionId)
	}
	document := Connection{
		Name:     name,
		Address:  address,
		Port:     port,
		Username: username,
		Password: password,
	}
	oid := bson.ObjectIdHex(connectionId)
	err := session.DB("goMongodbUI").C("connections").UpdateId(oid, document)
	if err != nil {
		return err
	}
	return nil
}

func Delete(connectionId string) error {
	session := mgoSession.Clone()
	defer session.Close()

	if !bson.IsObjectIdHex(connectionId) {
		return fmt.Errorf("connectionId: %s is not ObjectId", connectionId)
	}
	oid := bson.ObjectIdHex(connectionId)
	err := session.DB("goMongodbUI").C("connections").RemoveId(oid)
	if err != nil {
		return err
	}
	return nil
}
