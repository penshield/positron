package logging

import (
	"log"
	"os"
	"io"
	"fmt"
)


var logger *log.Logger

const (

	//This is the name of the application that logger will output as a message  in the configuration section for logging
	LOGGER_CONF_NAME ="positron.logging.conf.name"
	//this is the full path of the file that will hold the log messages
	LOGGER_CONF_OUTPUT = "positron.logging.output.file"
	//logging info

)

type LoggingConfig struct {

	Name string //the name of the application
	Location string

}

//This is the function that will be called once the logger has been called
func init(){

	log.SetPrefix("positron - ")
	log.SetFlags(log.LUTC|log.Ltime|log.Ldate|log.LstdFlags)
	log.SetOutput(os.Stderr)
	logger = log.New(os.Stderr,"Positron - ",log.Flags())


}

func SetLoggingConfig (config *LoggingConfig){

	log.SetPrefix(config.Name)


	//open file for reading
	file , err := os.OpenFile(config.Location,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0777)

	if err != nil {

		logger.Fatal(fmt.Sprintf("Problem opening the logging File location at : %s\n",config.Location))
	}
	logger.SetFlags(log.Flags())
	logger.SetPrefix(log.Prefix() + " - ")
	SetOutput(file)
}


func SetOutput(output io.Writer){

	logger.SetOutput(output)
}

func Log(msg ...interface{}) {

	logger.Println(msg...)
}
