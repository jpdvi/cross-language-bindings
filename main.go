package main
import "C"

import (
	"fmt"
	"unsafe"
	"encoding/binary"
	"encoding/json"
	"sync"
	"time"
)
type SampleJSON struct {
	Data string `json:"data"`
}

func somethingPrivate(msg string) string {
	return msg+"- GO wuz here"
}

func somethingConcurrent(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("Waiting", i)
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

	// Do some concurrent stuff and await the result
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go somethingConcurrent(i, &wg)
	}
	wg.Wait()

	return C.CBytes(append(length, jsonData...))
}

func main() {
	fmt.Println("I'm pointless")
}
