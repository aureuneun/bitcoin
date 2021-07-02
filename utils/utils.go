package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func ToBytes(i interface{}) []byte {
	var buffer bytes.Buffer
	HandleErr(gob.NewEncoder(&buffer).Encode(i))
	return buffer.Bytes()
}

func FromBytes(i interface{}, data []byte) {
	HandleErr(gob.NewDecoder(bytes.NewReader(data)).Decode(i))
}
