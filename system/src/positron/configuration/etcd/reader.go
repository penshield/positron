package etcd


import "fmt"




var etcd Etcd


func init(){
	err := etcd.OpenConnection()

	if err != nil {

		fmt.Println(fmt.Sprintf("Got Exception during opening connection to Etcd : %v",err))

		return
	}


}

func ReadConfig(key string) (string,error) {


	defer etcd.Close()

	defer func(){

		exception := recover()

		if exception != nil {

			fmt.Println(fmt.Sprint("Got Exception : %v", exception))

			//return "", fmt.Errorf("Got Exception :%v", exception)
		}
	}()

	return etcd.GetConfig(key)


}

