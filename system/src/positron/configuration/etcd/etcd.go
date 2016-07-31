package etcd

import (

	client "github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
	"golang.org/x/net/context"
)

/*
   This is the main struct for Etcd Daemon process for setting key-value configuration options
 */

type Etcd struct{

	client *client.Client
	state bool
	ctx context.Context
	cancel context.CancelFunc
}

func (etcd Etcd) GetState() bool {

	return etcd.state
}

func (etcd *Etcd) Close() {

	if (etcd.state){

		etcd.cancel()
		etcd.state = false
		etcd.client.Close()
	}
}

func (etcd *Etcd) OpenConnection() (err error) {

	//Check for errors at the end
	defer func(){

		exception := recover()

		if exception != nil {

			//log the exception in here
			fmt.Println("positron-Etcd:",exception)
			etcd.state = false

			err = fmt.Errorf("positron-etcd:%v",exception)


		}
	}()

	etcd.client , err = client.New(client.Config{

		Endpoints:[]string{"localhost:2379"},
		DialTimeout:5*time.Second,
	})

	etcd.state = true


	etcd.ctx , etcd.cancel = context.WithTimeout(context.Background(),5*time.Second)

	//set the etcd state to true
	etcd.state = true

	return err

}

func (etcd *Etcd) PutConfig(key,value string) error {

	if !etcd.state {

	   return fmt.Errorf("positron-etcd:%v","Try to open connection to the etcd daemon first")
	}

	_ , err := etcd.client.Put(etcd.ctx,key,value)

	if err != nil {

		return err
	}

	return nil
}

func (etcd Etcd) GetConfig(key string) (string,error) {

	response , err := etcd.client.Get(etcd.ctx,key)

		if err != nil {

			return "" , err
		}

		return string(response.Kvs[0].Value),nil
}

