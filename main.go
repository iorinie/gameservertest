//需要在 protoc-gen-go目录执行go build和go install，生成 protoc-gen-go.exe
//protoc -I proto/ proto/gameservertest.proto --go_out=plugins=grpc:proto

package main

import "gameservertest/network"

func main() {
	go network.StartServer()

	exit := make(chan int)
	<- exit
}