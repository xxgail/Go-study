package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	test := &Student{
		Name:   "gaigai",
		Male:   true,
		Scores: []int32{99, 88, 77},
	}
	data, err := proto.Marshal(test)
	fmt.Println(test)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}
	newTest := &Student{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error:", err)
	}
	if test.GetName() != newTest.GetName() {
		log.Fatalf("data mismatch %q != %q", test.GetName(), newTest.GetName())
	}
}
