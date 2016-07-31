package main

import (


	etcd "positron/configuration"
	logger "positron/logging"

	"fmt"
)

func main() {



	var etcd etcd.Etcd

	//open the connection
	err := etcd.OpenConnection()

	if err != nil {
		logger.Log("Got exception:",err)
	}

	//write some data
	var mymap map[string]string = make(map[string]string)

	mymap["saidyia"] = "Saidyia Yousef"
	mymap["ibrahim"] = "ibrahim Fawzy"
	mymap["morooj"] = "morooj Aldeeb"
	mymap["serene"] = "Serene Fawzy"
	mymap["yasmin"] = "Yasmin Fawzy"
	mymap["elene"] = "Elene Fawzy"

	for key , value := range mymap {

		//now insert all these keys and values
		err = etcd.PutConfig(key,value)

		if err != nil {

			logger.Log("Got Error : ",err)
		}
	}


	//now closes everything
	defer etcd.Close()

	logger.Log("Finished inserting data into the Etcd Daemon")

	logger.Log("Getting values for all the keys")

	for key , _ := range mymap{


		value , err := etcd.GetConfig(key)

		if err != nil {
			logger.Log("Got error :",err)
			continue
		}

		logger.Log(fmt.Sprintf("Key:%v , Value : %v\n",key,value))


	}




}
