package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

// protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative proto/processor_message.proto
