package main
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"

)

func main(){
	fmt.Println("Encryption program v0.01")



	fmt.Println("Enter Your text which is to be encrypted: ")
	var first string
    	fmt.Scanln(&first)
    	fmt.Println("Enter Your passphrase (passphrase should be 32 bytes): ")    	
	var passphrase string
    	fmt.Scanln(&passphrase)
    	text := []byte(first)
	key := []byte(passphrase)
	c,err := aes.NewCipher(key)
	if err != nil{
		fmt.Println(err)
	}
	
	gcm,err := cipher.NewGCM(c)
	if err != nil{
	fmt.Println(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	
	if _, err = io.ReadFull(rand.Reader,nonce); err != nil{
	
	fmt.Println(err)
	}
	//fmt.Println(gcm.Seal(nonce, nonce, text, nil))
	err = ioutil.WriteFile("myfile.data", gcm.Seal(nonce, nonce, text, nil), 0777)

	if err != nil {
	fmt.Println(err)
}
}
