package main

import (
	"fmt"
	"log"
	"protobuf-lesson/pb"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	employee := &pb.Employee{
		Id:           1,
		Name:         "Gopher",
		Email:        "test@example.com",
		Occupation:   pb.Occupation_ENGINEER,
		PhoneNumbers: []string{"080-1234-5678", "090-1234-5678"},
		Projects: map[string]*pb.Company_Project{
			"ProjectX": &pb.Company_Project{},
		},
		Profile: &pb.Employee_Text{
			Text: "Hello, World!",
		},
		Birthday: &pb.Date{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}

	// シリアライズ&&デシリアライズ
	// binData, err := proto.Marshal(employee)
	// if err != nil {
	// 	log.Fatalln("Failed to encode employee:", err)
	// }

	// if err := ioutil.WriteFile("test.bin", binData, 0644); err != nil {
	// 	log.Fatalln("Failed to write employee:", err)
	// }

	// in, err := ioutil.ReadFile("test.bin")
	// if err != nil {
	// 	log.Fatalln("Error reading file:", err)
	// }

	// readEmployee := &pb.Employee{}
	// err = proto.Unmarshal(in, readEmployee)
	// if err != nil {
	// 	log.Fatalln("Failed to parse employee:", err)
	// }

	// fmt.Println("Read employee:", readEmployee)

	// JSON
	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee)
	if err != nil {
		log.Fatalln("Failed to encode employee:", err)
	}
	// fmt.Println(out)

	readEmployee := &pb.Employee{}
	if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
		log.Fatalln("Failed to parse employee:", err)
	}
	fmt.Println("Read employee:", readEmployee)
}
