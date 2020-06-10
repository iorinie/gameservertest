package network

import (
	"encoding/binary"
	"fmt"
	gameServerTest "gameservertest/proto"
	"github.com/golang/protobuf/proto"
	"io"
	"net"
)

type Msg struct {
	Head Head
	Body Body
}

type Head struct {
	Len uint16
}

type Body struct {
	Content []byte
}

func StartServer() {
	fmt.Println("SERVER|listening...")

	listener, err := net.Listen("tcp", ":10010")
	if err != nil {
		fmt.Printf("SERVER|tcp listen to port fail, err = %s\n", err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("SERVER|listener accept fail, err = %s\n", err)
			continue
		}

		fmt.Printf("SERVER|a client is connected, remote = %v\n", conn.RemoteAddr())

		go handleCConn(conn)
	}
}

func handleCConn(conn net.Conn) {
	for {
		head := &Head{}
		err := binary.Read(conn, binary.BigEndian, head)
		if err != nil {
			fmt.Printf("SERVER|read head fail, err = %s\n", err)
			return
		}
		count := head.Len
		if count > 0 {
			body := make([]byte, count)
			_, err := io.ReadFull(conn, body)
			if err != nil {
				fmt.Printf("SERVER|read body fail, err = %s\n", err)
				return
			}
			loginReq := new(gameServerTest.LoginReq)
			err = proto.Unmarshal(body, loginReq)
			if err != nil {
				fmt.Printf("SERVER|unmarshal body fail, err = %s\n", err)
				return
			}
			fmt.Printf("SERVER|body = %v", loginReq)
		}
	}
}