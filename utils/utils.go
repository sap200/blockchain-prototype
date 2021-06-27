package utils

// This is utils package for the blockchain
// This package contains the necessary helper functions that are used by almost all packages

import (
	"crypto/sha256"
	"fmt"
	"github.com/tidwall/pretty"
	"encoding/json"
	"blockchain/types"
)

// This function is used to find SHA256 hash of any string data passed as an argument
// It returns the hashed data as a slice of uint8 which is an alias of byte
// The return type from sha256.sum256() is [32]byte which is an array of length 32
// To make it a slice just use data[:] which is slice of the array containing full array
func Hash(data string) []byte {
	hashedBytes := sha256.Sum256([]byte(data))
	return hashedBytes[:]
}

// PrettyPrint function helps visualising data in terminal
// It prints json format beautifully
// This is only for development and debugging purpose
// The API and the endpoint is responsible for pretty printing
func PrettyPrint(data string) {
	res1 := pretty.Pretty([]byte(data))
	res2 := pretty.Color(res1, nil)
	fmt.Println(string(res2))
}

func EncodeMessage(objectToEncode types.Message) string {
	objInByteSlice, _ := json.Marshal(objectToEncode)
	return string(objInByteSlice)
}

func DecodeMessage(objToDecode string) types.Message {
	var s types.Message
	json.Unmarshal([]byte(objToDecode), &s)
	return s
}

func EncodePeers(object []string) string {
	objInByteSlice, _ := json.Marshal(object)
	return string(objInByteSlice)	
}

func DecodePeers(object string) []string {
	var s []string
	json.Unmarshal([]byte(object), &s)
	return s
}
