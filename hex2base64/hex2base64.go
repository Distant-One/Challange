package hex2base64

import (
	"fmt"
)

var infohex2base64 string = "hex2base64.go v002"

var base64Indextable string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

//**************************************
// Base64 Encode a binary block of data
//**************************************
func Bin2base64(inBin []byte) [64]byte {
	//print version to keep track of my upates
	fmt.Println(infohex2base64)
	// Need to check input bounds - maybe look at how hex.DecodeString() handles variable lenghts
	// Need to add check for rounds that are do not contain 3 inBin bytes
	// Do I need to add PAD characters in to indicate last round did not include 3 inBin bytes so decocde will work properly
	var outB64 [64]byte
	var rounds, i, o int

	i = 0
	o = 0
	rounds = len(inBin) / 3
	for i < rounds {
		outB64[o+0] = base64Indextable[((inBin[i*3] >> 2) & 0x3f)]
		outB64[o+1] = base64Indextable[((inBin[i*3+0]<<4)&0x30)|((inBin[i*3+1]>>4)&0x0f)]
		outB64[o+2] = base64Indextable[((inBin[i*3+1]<<2)&0x3c)|((inBin[i*3+2]>>6)&0x03)]
		outB64[o+3] = base64Indextable[(inBin[i*3+2] & 0x3f)]
		i++
		o += 4
	}
	return outB64
}
