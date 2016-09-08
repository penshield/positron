package main
//#cgo LDFLAGS: -I/usr/local/include/glib-2.0 -Lusr/local/lib -lopenvas_nasl -I/usr/local/include/glib-2.0
//#include<stdio.h>
//#include<stdlib.h>
//#include<openvas/nasl/nasl.h>
import "C"
import (
	"fmt"
)
const (
	MAX_COUNT = 999
)
func main() {

	defer  func(){

		recover()
	}()

	 arguments := &C.struct_ExternalData{file:C.CString("/media/snouto/rest/projects/openvas/nvts/gb_default_smb_credentials.nasl"),target:C.CString("192.168.1.8"),authenticated:1}
	 info := &C.struct_DaltonScriptInfo{}

	fmt.Println("Target : ",C.GoString(arguments.target));
	C.executeNasl(arguments,info)
	fmt.Printf("Script Version : %s\n",C.GoString(info.ScriptVersion))
	//printing the cve ID of the current file
	var i int =0

	for(i <MAX_COUNT){

		if(info.ScriptCveIds[i] != nil){
			fmt.Printf("Script CVE ID : %s\n",C.GoString(info.ScriptCveIds[i].Contents));
		}else{
			break
		}

		i++
	}

	fmt.Println("Done")

}



