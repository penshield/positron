package main

import (
	"gopkg.in/mcuadros/go-syslog.v2"
	"fmt"
)

func main() {


	fmt.Println("Running SysLog server")
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)
	server := syslog.NewServer()
	server.SetFormat(syslog.Automatic)
	server.SetHandler(handler)
	server.ListenUDP("0.0.0.0:5167")
	server.ListenTCP("0.0.0.0:5167")
	err := server.Boot()

	if err != nil {
		fmt.Println(err)
		return
	}



	fmt.Println(fmt.Sprintf("Listening on : %v with Port : %v","0.0.0.0","5167"))

	go func(channel syslog.LogPartsChannel) {
	    for logParts := range channel {
		fmt.Println(logParts)
		fmt.Println(logParts["message"])
	    }
	}(channel)

	server.Wait()

}
