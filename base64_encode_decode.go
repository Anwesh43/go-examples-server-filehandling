package main 

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	stdEncoding := b64.StdEncoding
	str1 := "hello" 
	enc := stdEncoding.EncodeToString(([]byte)(str1))
	fmt.Println(string(enc))
	dec,_ := stdEncoding.DecodeString(enc)
	fmt.Println(string(dec)) 
}	