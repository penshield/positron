package etcd

import "fmt"





func WriteConfig(key string , value string) (bool,error) {

	etcd.OpenConnection()
	defer etcd.Close()

	defer func(){

		exception := recover()

		if exception != nil {

			fmt.Println(fmt.Sprint("Got Exception : %v", exception))

			//return "", fmt.Errorf("Got Exception :%v", exception)
		}
	}()


	err := etcd.PutConfig(key,value)

	if err != nil {

		fmt.Println(fmt.Sprintf("There was a problem inserting the config key :%s with value :%s into the central configuration daemon",key,value))

		return false , fmt.Errorf("There was a problem inserting the config key :%s with value :%s into the central configuration daemon",key,value)
	}

	return true , nil
}
