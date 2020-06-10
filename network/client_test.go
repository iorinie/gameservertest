package network

import (
	"encoding/json"
	"fmt"
	gameServerTest "gameservertest/proto"
	"github.com/golang/protobuf/proto"
	"net"
	"testing"
)

func TestConnectServer(t *testing.T) {
	conn, err := net.Dial("tcp", ":10010")
	if err != nil {
		fmt.Printf("CLIENT|connect to server fail, err = %s\n", err)
		return
	}

	defer conn.Close()

	req := new(gameServerTest.LoginReq)
	reqBytes, err := proto.Marshal(req)
	if err != nil {
		fmt.Printf("CLIENT|marshal login data fail, err = %s\n", err)
		return
	}
	head := Head{
		Len: uint16(len(reqBytes)),
	}
	headBytes, err := json.Marshal(head)
	if err != nil {
		fmt.Printf("CLIENT|marshal head fail, err = %s\n", err)
		return
	}
	_, err = conn.Write(headBytes)
	if err != nil {
		fmt.Printf("CLIENT|send head data to server fail, err = %s\n", err)
		return
	}
	_, err = conn.Write(reqBytes)
	if err != nil {
		fmt.Printf("CLIENT|send login data to server fail, err = %s\n", err)
		return
	}
}