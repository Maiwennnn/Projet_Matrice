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

func lecture(connection net.Conn, bufferFileSize []byte, bufferFileName []byte) {
	connection.Read(bufferFileSize)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

	connection.Read(bufferFileName)
	fileName := strings.Trim(string(bufferFileName), ":")

	fmt.Println("On va creer le file")

	newFile, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	var receivedBytes int64
	fmt.Println("File cree sans erreur")

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
	receivedBytes = 0
}

func main() {
	server, err := net.Listen(TYPE, HOST+":"+PORT)
	fmt.Println("Server started! Waiting for connections...")

	if err != nil {
		fmt.Println("Error listening: ", err)
		os.Exit(1)
	}
	defer server.Close()
	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		fmt.Println("Connected to client, start receiving the file name and file size")
		lecture(connection, bufferFileSize, bufferFileName)
		bufferFileName = nil
		bufferFileSize = nil
		fmt.Println("Deuxieme matrice")
		lecture(connection, bufferFileSize, bufferFileName)

	}

}
