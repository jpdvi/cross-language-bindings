package main
import "C"

import (
	"fmt"
	"unsafe"
	"encoding/binary"
	"encoding/json"
)
type SampleJSON struct {
	Data string `json:"data"`
}

func somethingPrivate(msg string) string {
	return msg+"- GO wuz here"
}

//export DoSomething
func DoSomething(incomingData *C.char) unsafe.Pointer {
	// Serialize some json
	readSomeData := somethingPrivate(C.GoString(incomingData))
	s := SampleJSON{ Data: readSomeData }
	var jsonData []byte
	jsonData, err := json.Marshal(s)
	if err != nil {
		fmt.Println("Oops")
	}
	// Create a byte array
	length := make([]byte, 64)
	binary.LittleEndian.PutUint64(length, uint64(len(jsonData)))
	return C.CBytes(append(length, jsonData...))
}

func main() {
	fmt.Println("I'm pointless")
}
