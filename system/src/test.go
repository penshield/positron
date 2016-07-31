package main

import ("fmt"

	tables "github.com/apcera/termtables"
	"positron/configuration"
)

func main() {

	table := tables.CreateTable()
	table.AddHeaders("Name","Email")
	table.AddRow("Mohamed Ibrahim","csharpizer@gmail.com")
	fmt.Println(table.Render())

	configuration.SayHello()



}
