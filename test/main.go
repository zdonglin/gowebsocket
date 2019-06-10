package main

import (
	"SocketService"
	"encoding/json"
	"fmt"
	"time"
)

type Test struct {
	Id   string
	Data string
}

func main() {
	url := "localhost:8080"
	SocketService.Init(url)

	time.Sleep(time.Second * 10)

	err := SocketService.SendMessage("abc", []byte("zeng"))
	if err != nil {
		fmt.Println("1:", err)
	}

	t := Test{"Id", "Data"}

	d, _ := json.Marshal(t)

	SocketService.SendMessage("abcd", d)

	if err != nil {
		fmt.Println("2:", err)
	}
	time.Sleep(time.Second * 1000)

}
