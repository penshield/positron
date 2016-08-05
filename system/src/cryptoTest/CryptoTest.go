package main

import (
	"flag"
	"crypto/rand"
	"fmt"
	"encoding/base64"
	"crypto/aes"
	 "positron/logging"
	"crypto/cipher"

)

var (

	IV = []byte("Mohamed Ibrahims")
	message = flag.String("message","HelloWorld","Message to use for encryption, the default is HelloWorld")
	keySize = flag.Int("keySize",32,"Key size of the encryption algorithm to use : 16 , 32 , 64")
	keyContent = flag.String("key","","The default key to use during encryption or decryption")
	operation = flag.String("operation","encrypt","The operation to perform on the message either , Encrypt(The default) or decrypt")
)


func main() {


	//create a logging config file
	logConfig := logging.LoggingConfig{ Location:"/media/snouto/rest/projects/positron/logs/system/logs/positron.log", Name:"Positron"}
	logging.SetLoggingConfig(&logConfig)
	flag.Parse()
	var key []byte
	var err error
	fmt.Println("Generating Cryptographic Key........")
	key = makeKey()


	logging.Log("Using Key : " , base64.StdEncoding.EncodeToString(key))

	if *keyContent != "" {

		//get the key from the command prompt
		key , err = base64.StdEncoding.DecodeString(*keyContent)

		if err != nil {

			logging.Log("There was a problem getting the key from Command prompt")
			return
		}

	}

	result := operate([]byte(*message),key)

	logging.Log(fmt.Sprintf("Operation: %s , Result:%s",*operation,result))


}

func operate(data , key []byte) (result string) {

	switch *operation {

	case "encrypt":

		result= encrypt(data,key)

		break
	case "decrypt":
		result= decrypt(data,key)
		break

	}

	return result


}
func decrypt(data []byte, key []byte) string {

	c , err := aes.NewCipher(key)

	if err != nil {

		logging.Log("There was a problem decrypting the data")
		return ""
	}


	data , err = base64.StdEncoding.DecodeString(string(data))

	if err != nil {

		logging.Log("There was a problem decrypting the message")
		return ""
	}

	stream := cipher.NewCTR(c,IV)



	stream.XORKeyStream(data,data)

	return string(data)


}
func encrypt(data []byte, key []byte) string {

	c , err := aes.NewCipher(key)

	if err != nil {

		logging.Log("There was a problem creating the AES Cipher")
		return ""
	}




	stream := cipher.NewCTR(c,IV)

	stream.XORKeyStream(data,data)


	return base64.StdEncoding.EncodeToString(data)
}

func makeKey() []byte{

	//create a slice that takes the keysize into context
	key_slice := make([]byte,*keySize)
	//now generate a random key size
	n , err := rand.Read(key_slice)

	if err != nil {
		logging.Log("There was a problem generating the Cryptographic Key of Size : ",*keySize)
		return nil
	}

	if n < *keySize {

		logging.Log("Failed to generate the full cryptographic Key of Size : ",*keySize)

		return nil
	}

	return key_slice
}