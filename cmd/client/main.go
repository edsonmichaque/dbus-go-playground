package main

import (
	"fmt"
	"github.com/godbus/dbus/v5"
)

func main() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	var s string
	obj := conn.Object("com.michaque.DBusGoPlayground", "/com/michaque/DBusGoPlayground")
	err = obj.Call("com.michaque.DBusGoPlayground.Ping", 0).Store(&s)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response", s)

}
