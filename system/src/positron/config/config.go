package config

import (
	"positron/config/central"
	"positron/logging"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"io/ioutil"
	"bytes"
)

var (
	server etcd.Etcd
	logger logging.LoggingConfig
	config *viper.Viper
)

const (

	MAIN_CONF_FILE_NAME = "POSITRONCONF"
)

func openConnection() error {

	config = viper.New()
	config , err := ReadLocalConfig()
	endpoint := fmt.Sprintf("http://%v:%v",config.GetString("positron.config.remote.server"),
	config.GetInt("positron.config.remote.port"))
	timeout := config.GetInt("positron.config.remote.timeout")
	err = server.OpenConnectionWith(endpoint,timeout)
	if err != nil {

		return fmt.Errorf("There Was a problem opening the connection with the central configuration server")
	}

	return nil
}

func InitialSetup() error {

	openConnection()
	defer server.Close()

	if server.GetState() {

		//that means the server is already open and ready to set configuration variables
		config.AutomaticEnv()
		//get the configuration named "positron.main"
		positron_main := config.GetString(MAIN_CONF_FILE_NAME)

		if positron_main != "" {

			//put the configuration inside the etcd
			//server.PutConfig(MAIN_CONF_FILE_NAME,positron_main)
			//that means the variable was set and ready to be read , so read it
			file , err := os.Open(positron_main)
			defer file.Close()
			if err != nil {

				return fmt.Errorf("There was a problem opening file pointed in '%s' environment variable",MAIN_CONF_FILE_NAME)
			}
			contents ,err := ioutil.ReadAll(file)

			if err != nil {

				return fmt.Errorf("There was a problem reading the contents of the file pointed by '%s' environment variable.",MAIN_CONF_FILE_NAME)
			}

			//now read all the contents of that file
			config.SetConfigType("json")
			config.ReadConfig(bytes.NewBuffer([]byte(contents)))

			//get all the keys within the configuration file and add them into the etcd daemon
			for key , value := range config.AllSettings() {

				readRecursively(&server,&key,value)
			}

			return nil

		}
	}

	return fmt.Errorf("There was a problem in writing the central Etcd daemon")
}

func readRecursively(server *etcd.Etcd , key *string ,value interface{}){

	switch value.(type){
		case map[string]interface{}:
		 for another,val := range value.(map[string]interface{}) {
			 another = fmt.Sprintf("%s.%s",*key,another)
			 readRecursively(server,&another,val)
		 }
	case string:
		server.PutConfig(*key,value.(string))
		*key =""
	case int:
		server.PutConfig(*key,value.(string))
		*key =""
	case float64:
		value_converted := fmt.Sprintf("%v",value)
		server.PutConfig(*key,value_converted)
		*key =""
	}
}

func WriteRemoteConfig(key string , value interface{}) error {

	openConnection()



	switch value.(type) {
		case string:
		server.PutConfig(key,value.(string))
		case int:
		server.PutConfig(key,value.(string))
		case float64:
		value_converted := fmt.Sprintf("%v",value)
		server.PutConfig(key,value_converted)
	}

	return nil


}

func ReadRemoteConfig(key string) (interface{},error) {

	server.Close()
	localConfig , err := ReadLocalConfig()
	if err != nil {

		return nil , fmt.Errorf("There was a problem opening the local file")
	}

	daemon := new(etcd.Etcd)
	defer daemon.Close()
	endpoint := fmt.Sprintf("%v:%v",localConfig.Get("positron.config.remote.server"),localConfig.Get("positron.config.remote.port"))
	timeout := localConfig.GetInt("positron.config.remote.timeout")

	daemon.OpenConnectionWith(endpoint,timeout)


	value , err := daemon.GetConfig(key)

	if err != nil {

		return nil , fmt.Errorf("Error getting the value from Etcd Daemon for key %s",key)
	}

	return value,nil
}

func ReadLocalConfig() (*viper.Viper,error) {

	//now read the contents of the configuration file
	//from etcd by using positron.main
	localFile := os.Getenv(MAIN_CONF_FILE_NAME)
	file , err := os.Open(localFile)
	if err != nil {

		return nil , fmt.Errorf("There was a problem getting the main positron configuration from Local path in the environment variable")
	}

	contents , err := ioutil.ReadAll(file)

	if err != nil {

		return nil , fmt.Errorf("There was a problem reading the main positron configuration from Local path in the environment variable")

	}
	config.SetConfigType("json")
	config.ReadConfig(bytes.NewBuffer([]byte(contents)))

	//get the host:port of the remote etcd daemon
	return config , nil
}