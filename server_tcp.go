package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	HOST       = "localhost"
	PORT       = "8080"
	TYPE       = "tcp"
	BUFFERSIZE = 1024
)

func main() {
	server, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		fmt.Println("Error listening: ", err)
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Connected to client, start receiving the file name and file size")
	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)

	fmt.Println("Server started! Waiting for connections...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		fmt.Println("Client connected")
		connection.Read(bufferFileSize)
		fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

		connection.Read(bufferFileName)
		fileName := strings.Trim(string(bufferFileName), ":")

		newFile, err := os.Create(fileName)

		if err != nil {
			panic(err)
		}
		defer newFile.Close()
		var receivedBytes int64

		for {
			if (fileSize - receivedBytes) < BUFFERSIZE {
				io.CopyN(newFile, connection, (fileSize - receivedBytes))
				connection.Read(make([]byte, (receivedBytes+BUFFERSIZE)-fileSize))
				break
			}
			io.CopyN(newFile, connection, BUFFERSIZE)
			receivedBytes += BUFFERSIZE
		}
		fmt.Println("Received file completely!")

	}

}
