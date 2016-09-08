package main
//#cgo CFLAGS: -I/usr/local/include
//#cgo LDFLAGS: -L/usr/local/lib -lTestGoLibs
//#include<stdio.h>
//#include<stdlib.h>
//#include<TestGoLib.h>
import "C"
import (
	"fmt"
)

func main() {

	structure := C.struct_TestData{}

	C.LoadData(&structure)
	fmt.Println(C.GoString(structure.Name))
	fmt.Println(C.GoString(structure.Email))

	i :=0
	for i < 50 {

		message := structure.Messages[i]
		if message == nil {
			break
		}
		fmt.Println(C.GoString(message.Payload))
		i++
	}


	fmt.Println("Finished")

}
