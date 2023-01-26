package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

const (
	HOST       = "localhost"
	PORT       = "8080"
	TYPE       = "tcp"
	BUFFERSIZE = 1024
)

func fillString(retunString string, toLength int) string {
	for {
		lengtString := len(retunString)
		if lengtString < toLength {
			retunString = retunString + ":"
			continue
		}
		break
	}
	return retunString
}

func sendFileToServer(connection net.Conn, nomFile string) {
	fmt.Println("On envoie le fichier au serveur")
	file, err := os.Open(nomFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Fichier ouvert sans erreur")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Acces aux stats sans erreur")
	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fileName := fillString(fileInfo.Name(), 64)
	fmt.Println("Sending filename and filesize!")
	connection.Write([]byte(fileSize))
	connection.Write([]byte(fileName))
	sendBuffer := make([]byte, BUFFERSIZE)
	fmt.Println("Start sending file!")
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		connection.Write(sendBuffer)
	}
	fmt.Println("File has been sent, closing connection!")
	return
}

func main() {
	connection, err := net.Dial(TYPE, HOST+":"+PORT)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to server")
	sendFileToServer(connection, "matriceA.txt")
	sendFileToServer(connection, "matriceB.txt")
	defer connection.Close()

}
