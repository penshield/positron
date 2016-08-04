package main

import ("encoding/xml"

	"fmt"
)

type Configuration struct {

	Name string `xml:"id"`
	Location string `xml:,"location"`
	Email Email `xml:,"Email"`
}

type Email struct {

	Type string `xml:,attr`
}


func (config Configuration) toString() string{

	repr := fmt.Sprintf(" Name : %s , Location: %s",config.Name,config.Location)

	return repr
}

func main() {

	//var myxml string = `<Configuration><id>Configuration Name </id><Location>Virtual Reality</Location></Configuration>`
	var Config Configuration = Configuration{Name:"Configuration Name",Location:"Virtual Reality"}
	mystring := marshal(&Config)

	fmt.Println(mystring)

}

func unmarshal(xmlString string) Configuration {

	var config Configuration
	if err := xml.Unmarshal([]byte(xmlString),&config) ; err != nil {

		fmt.Println("There was an error in parsing the xml")

	}

	return config
}

func marshal(config *Configuration) string {

	repr , err := xml.Marshal(config)

	if err != nil {

		fmt.Println("There was a problem marshalling the config Struct into string")
		return ""
	}

	return string(repr)
}
