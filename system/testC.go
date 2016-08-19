package main


//#cgo LDFLAGS: -I/usr/local/include/glib-2.0 -Lusr/local/lib -lopenvas_nasl -I/usr/local/include/glib-2.0
//#include<stdio.h>
//#include<stdlib.h>
//#include<openvas/nasl/nasl.h>
import "C"
import (
	"fmt"

)

func main() {

	defer  func(){

		recover()
	}()

	 arguments := &C.struct_ExternalData{file:C.CString("/media/snouto/rest/projects/openvas/nvts/http_version.nasl"),target:C.CString("192.168.1.14"),authenticated:1}

	fmt.Println("Target : ",C.GoString(arguments.target));
	C.executeNasl(arguments)

	fmt.Println("Done")

}



