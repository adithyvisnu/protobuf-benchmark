package main

import (
	"encoding/json"
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Hello Protocol Buffers!")

	m := &UserData{
		User: "Adithya Visnu",
		Age:  19,
	}

	jsonMarshal, _ := json.Marshal(m)
	fmt.Printf("JSON Marshal: %v, length: %d", jsonMarshal, len(jsonMarshal))
	fmt.Println()
	protoMarshal, _ := proto.Marshal(m)
	fmt.Printf("Protobuf Marshal: %v, length: %d", protoMarshal, len(protoMarshal))
	fmt.Println()
	fmt.Println()
	err := json.Unmarshal(jsonMarshal, m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("JSON Unmarshal: %+v", m)
	fmt.Println()
	err = proto.Unmarshal(protoMarshal, m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Protobuf Unmarshal: %+v", m)
	fmt.Println()
}
