package main

import (
    "crypto/aes"
    "crypto/cipher"
    "fmt"
    "io/ioutil"
)

func main() {
    fmt.Println("Decryption Program v0.01")
    fmt.Println("Enter Your passphrase (passphrase should be 32 bytes): ")    	
    var passphrase string
    fmt.Scanln(&passphrase)
    fmt.Println("Enter Your data filename (format = xyz.data): ")    	
    var filename string
    fmt.Scanln(&filenam)
    
    key := []byte(passphrase)
    
    ciphertext, err := ioutil.ReadFile(filename.data)

    if err != nil {
        fmt.Println(err)
    }

    c, err := aes.NewCipher(key)
    if err != nil {
        fmt.Println(err)
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        fmt.Println(err)
    }

    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        fmt.Println(err)
    }

    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(plaintext))
}
