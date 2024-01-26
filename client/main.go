package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/golang/protobuf/proto"
	"github.com/spikebike/proto/sum"
)

func main() {
	conn, err := net.Dial("tcp", ":4040")
	if err != nil {
		fmt.Printf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	numbers := &sum.Numbers{
		A: 3,
		B: 5,
	}

	out, err := proto.Marshal(numbers)
	if err != nil {
		fmt.Printf("Failed to encode message: %v", err)
	}

	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[0:], uint32(len(out)))
	conn.Write(buf[0:])
	conn.Write(out)

	_, err = conn.Read(buf[0:])
	if err != nil {
		fmt.Printf("Failed to read response: %v", err)
	}

	length := binary.LittleEndian.Uint32(buf[0:])
	data := make([]byte, length)

	_, err = io.ReadFull(conn, data)
	if err != nil {
		fmt.Printf("Failed to read response: %v", err)
	}

	response := &sum.Sum{}
	if err := proto.Unmarshal(data, response); err != nil {
		fmt.Printf("Failed to parse response: %v", err)
	}

	fmt.Printf("The sum is: %d\n", response.Result)
}