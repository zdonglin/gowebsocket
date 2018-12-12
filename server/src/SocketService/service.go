package SocketService

import (
	"fmt"
)

func SendMessage(session string, message []byte) error {

	fmt.Println("session:", session)
	t := GetType(session)
	fmt.Println("t:", t)
	c, err := GetConn(session)
	if err != nil {
		return err
	}
	err = c.WriteMessage(t, message)
	if err != nil {
		return err
	}
	return nil
}
