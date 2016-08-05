package main

import (


	"positron/config"
	"fmt"


)

func main() {

	err := config.InitialSetup()

	if err != nil {

		fmt.Println(err)
		return
	}

	configuration , err := config.ReadLocalConfig()

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println("positron.logging.output : ",configuration.Get("positron.logging.output"))
	value , err := config.ReadRemoteConfig("positron.logging.output")

	if err != nil {

		fmt.Println(err)
	}

	fmt.Println("Remote positron.logging.output : ",value)
	//var etcd etcd.Etcd
	//
	////open the connection
	//err := etcd.OpenConnection()
	//
	//if err != nil {
	//	logger.Log("Got exception:",err)
	//}
	//
	////write some data
	//var mymap map[string]string = make(map[string]string)
	//
	//mymap["saidyia"] = "Saidyia Yousef"
	//mymap["ibrahim"] = "ibrahim Fawzy"
	//mymap["morooj"] = "morooj Aldeeb"
	//mymap["serene"] = "Serene Fawzy"
	//mymap["yasmin"] = "Yasmin Fawzy"
	//mymap["elene"] = "Elene Fawzy"
	//
	//for key , value := range mymap {
	//
	//	//now insert all these keys and values
	//	result , err := writer.WriteConfig(key,value)
	//
	//	if err != nil && result == false {
	//
	//		logger.Log("Got Error : ",err)
	//	}
	//}
	//
	//
	////now closes everything
	//defer etcd.Close()
	//
	//logger.Log("Finished inserting data into the Etcd Daemon")
	//
	//logger.Log("Getting values for all the keys")
	//
	//for key , _ := range mymap{
	//
	//
	//	value , err := etcd.GetConfig(key)
	//
	//	if err != nil {
	//		logger.Log("Got error :",err)
	//		continue
	//	}
	//
	//	logger.Log(fmt.Sprintf("Key:%v , Value : %v\n",key,value))
	//
	//
	//}




}
