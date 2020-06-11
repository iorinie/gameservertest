package network

import (
	"bytes"
	"encoding/binary"
	"gameservertest/proto"
	"github.com/golang/protobuf/proto"
	"net"
	"testing"
)

func TestConnectServer(t *testing.T) {
	conn, err := net.Dial("tcp", ":10010")
	if err != nil {
		t.Errorf("CLIENT|connect to server fail, err = %s\n", err)
		return
	}

	defer conn.Close()

	req := new(gameServerTest.LoginReq)
	req.Name = "iorinie"
	reqBytes, err := proto.Marshal(req)
	if err != nil {
		t.Errorf("CLIENT|marshal login data fail, err = %s\n", err)
		return
	}

	msg := new(gameServerTest.Msg)
	msg.Name = "Login"
	msg.Content = string(reqBytes)
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		t.Errorf("CLIENT|marshal msg data fail, err = %s\n", err)
		return
	}

	head := &Head{
		Len: uint16(len(msgBytes)),
	}
	bf := new(bytes.Buffer)
	err = binary.Write(bf, binary.BigEndian, head)
	if err != nil {
		t.Errorf("CLIENT|build head data fail, err = %s\n", err)
		return
	}
	bf.Write(msgBytes)
	_, err = conn.Write(bf.Bytes())
	if err != nil {
		t.Errorf("CLIENT|send msg data to server fail, err = %s\n", err)
		return
	}
}