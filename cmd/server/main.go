package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
)

var spec = `<node>
	<interface name="com.michaque.DBusGoPlayground">
		<method name="Ping">
			<arg direction="out" type="s" />
		</method>
	</interface>` + introspect.IntrospectDataString + `</node>`

type dbusGoPlayground struct{}

func (d dbusGoPlayground) Ping() (string, *dbus.Error) {
	return "Pong 1", nil
}

func main() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	srv := &dbusGoPlayground{}

	if err := conn.Export(srv, "/com/michaque/DBusGoPlayground", "com.michaque.DBusGoPlayground"); err != nil {
		panic(err)
	}

	if err := conn.Export(introspect.Introspectable(spec), "/com/michaque/DBusGoPlayground", "org.freedesktop.DBus.Introspectable"); err != nil {
		panic(err)
	}

	reply, err := conn.RequestName("com.michaque.DBusGoPlayground", dbus.NameFlagDoNotQueue)
	if err != nil {
		panic(err)
	}

	if reply != dbus.RequestNameReplyPrimaryOwner {
		fmt.Fprintln(os.Stderr, "name already taken")
		os.Exit(1)
	}
	fmt.Println("Listening on com.michaque.DBusGoPlayground / /com/michaque/DBusGoPlayground ...")
	select {}

}
