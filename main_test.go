package main

import (
	"encoding/json"
	"testing"

	"github.com/golang/protobuf/proto"
)

func BenchmarkJSON(b *testing.B) {
	user := &UserData{
		User: "Adithya Visnu",
		Age:  10,
	}

	for i := 0; i < b.N; i++ {
		json.Marshal(&user)
	}
}
func BenchmarkProtobuf(b *testing.B) {
	user := &UserData{
		User: "Adithya Visnu",
		Age:  10,
	}

	for i := 0; i < b.N; i++ {
		proto.Marshal(user)
	}
}
