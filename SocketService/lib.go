package SocketService

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

//types
var TypeLock *sync.Mutex = new(sync.Mutex)
var Types map[string]int = make(map[string]int, 1000)

func SetType(session string, sessionType int) {
	TypeLock.Lock()
	defer TypeLock.Unlock()
	Types[session] = sessionType
}

func GetType(session string) int {
	TypeLock.Lock()
	defer TypeLock.Unlock()
	return Types[session]
}

//conns
var Conns map[string]*websocket.Conn = make(map[string]*websocket.Conn, 1000)
var ConnLock *sync.Mutex = new(sync.Mutex)

func SetConn(session string, conn *websocket.Conn) {
	ConnLock.Lock()
	defer ConnLock.Unlock()
	Conns[session] = conn
}

func GetConn(session string) (*websocket.Conn, error) {
	ConnLock.Lock()
	defer ConnLock.Unlock()
	if _, ok := Conns[session]; ok {
		return Conns[session], nil
	}
	return nil, errors.New("Conns not exist")
}
