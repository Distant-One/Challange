package main

import (
	"encoding/hex"
	"fmt"
)

var hexString string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

var base64KnownString string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

var base64Indextable [64]byte

var base64Alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func main() {
	fmt.Println("Input vector:"+hexString+" Len:", len(hexString))
	fmt.Println("Known answer base64 result:"+base64KnownString+" Len:", len(base64KnownString))

	bs, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", bs)

	//	var sl string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	copy(base64Indextable[:], base64Alphabet)

	bin2base64(bs)

}

func bin2base64(inBin []byte) {
	var outB64 [64]byte
	var rounds, i, o int
	i = 0
	o = 0
	rounds = len(inBin) / 3
	// Need to check input bounds - maybe look at how hex.DecodeString() handles variable lenghts
	// Need to add check for rounds that are do not contain 3 inBin bytes
	// Do I need to add PAD characters in to indicate last round did not include 3 inBin bytes so decocde will work properly
	for i < rounds {
		outB64[o+0] = base64Indextable[((inBin[i*3] >> 2) & 0x3f)]
		outB64[o+1] = base64Indextable[((inBin[i*3+0]<<4)&0x30)|((inBin[i*3+1]>>4)&0x0f)]
		outB64[o+2] = base64Indextable[((inBin[i*3+1]<<2)&0x3c)|((inBin[i*3+2]>>6)&0x03)]
		outB64[o+3] = base64Indextable[(inBin[i*3+2] & 0x3f)]
		i++
		o += 4
	}
	fmt.Printf("Input Vector         : %x\n", inBin)
	fmt.Printf("Output Base64 Encoded: %s\n", outB64)
	fmt.Printf("Known Answer         : %s\n", base64KnownString)
}
