package network

import (
	"encoding/binary"
	"fmt"
	gameServerTest "gameservertest/proto"
	"github.com/golang/protobuf/proto"
	"io"
	"net"
)

type Head struct {
	Len uint16
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
		head := new(Head)
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
			msg := new(gameServerTest.Msg)
			err = proto.Unmarshal(body, msg)
			if err != nil {
				fmt.Printf("SERVER|unmarshal body fail, err = %s\n", err)
				return
			}
			fmt.Printf("SERVER|body = %#v\n", msg)

			go handleRequest(conn, msg)
		}
	}
}

func handleRequest(conn net.Conn, msg *gameServerTest.Msg) {
	switch msg.Name {
	case "Login" :
		loginReq := new(gameServerTest.LoginReq)
		err := proto.Unmarshal([]byte(msg.Content), loginReq)
		if err != nil {
			fmt.Printf("SERVER|unmarshal login content fail, err = %s\n", err)
			break
		}
		fmt.Printf("SERVER|login info = %#v\n", loginReq)
	default:
		fmt.Printf("SERVER|handleRequest wrong msg type")
		break
	}
}