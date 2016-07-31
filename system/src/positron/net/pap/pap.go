package pap

import "fmt"

type Packet struct {

	Command string
	Name string
	Version string
	Headers map[string]interface{}
	Payload []byte
}

func (packet *Packet) SetCommand(command string) error {

	packet.Command = command

	return nil
}

func (packet *Packet) SetName(name string) error {

	packet.Name = name
	return nil
}

func (packet *Packet) SetVersion(version string) error {

	packet.Version = version
	return nil
}

func (packet *Packet) SetHeader(headerName string , headerValue interface{}) error {

	if packet.Headers == nil {

		packet.Headers = map[string]interface{}{}
	}

	packet.Headers[headerName] = headerValue

	return nil
}

func (packet *Packet) SetPayload(payload []byte) error {

	packet.Payload = payload
	return nil
}

func (packet Packet) DisplayInfo() string {

	info := fmt.Sprintf("Command : %v , Proto-Name : %v , Version : %v\n",packet.Command,packet.Name,packet.Version)

	for key , value := range packet.Headers {

		info +=fmt.Sprintf("Key:%v,Value:%v\n",key,value)

	}

	return info
}
