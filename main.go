package main

import (
	"encoding/hex"
	"fmt"
	"github.com/Distant-One/Challenge/hex2base64"
)

//var hexString string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
var hexString string = "49276d"

var base64KnownString string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

//var base64Indextable string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcportdefghijklmnopqrstuvwxyz0123456789+/"

func main() {

	// For base64 encoding convert ascii hex representation to binary
	bs, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}

	//encode binary block to bas64
	enc64 := hex2base64.Bin2base64(bs)

	fmt.Printf("Input Vector         : %x Len: %d\n", bs, len(bs))
	fmt.Printf("Output Base64 Encoded: %s len: %d\n", enc64, len(enc64))
	fmt.Printf("Known Answer         : %s len: %d\n", base64KnownString, len(base64KnownString))

}
