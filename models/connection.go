package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	Connections map[string]*Connection
)

type Connection struct {
	ConnectionId string
	Name         string
	Address      string
	Port         int
}

func init() {
	Connections = make(map[string]*Connection)
	Connections["default"] = &Connection{"default", "local", "localhost", 27017}
}

func AddOne(Connection Connection) (ConnectionId string) {
	Connection.ConnectionId = strconv.FormatInt(time.Now().UnixNano(), 10)
	Connections[Connection.ConnectionId] = &Connection
	return Connection.ConnectionId
}

func GetOne(ConnectionId string) (Connection *Connection, err error) {
	if v, ok := Connections[ConnectionId]; ok {
		return v, nil
	}
	return nil, errors.New("ConnectionId Not Exist")
}

func GetAll() map[string]*Connection {
	return Connections
}

func Update(ConnectionId string, Name string, Address string, Port int) (err error) {
	if v, ok := Connections[ConnectionId]; ok {
		v.Name = Name
		v.Address = Address
		v.Port = Port
		return nil
	}
	return errors.New("ConnectionId Not Exist")
}

func Delete(ConnectionId string) {
	delete(Connections, ConnectionId)
}
