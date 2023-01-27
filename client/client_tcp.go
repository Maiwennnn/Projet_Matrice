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

func fillString(retunString string, toLength int) string { // ajoute des ":" pour completer la taille du fichier transmis car elle ne fait pas forcement 9 octets
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

func sendFileToServer(connection net.Conn, nomFile string) { // envoie une matrice au serveur
	file, err := os.Open(nomFile) // on ouvre le fichier
	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfo, err := file.Stat() // on recupere la taille et le nom
	if err != nil {
		fmt.Println(err)
		return
	}
	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fileName := fillString(fileInfo.Name(), 64)
	connection.Write([]byte(fileSize))
	connection.Write([]byte(fileName))
	sendBuffer := make([]byte, BUFFERSIZE)
	for {
		_, err = file.Read(sendBuffer) // on envoie le fichier
		if err == io.EOF {
			break
		}
		connection.Write(sendBuffer)
	}
	return
}

func main() {
	print("on entre dans le main")
	connection, err := net.Dial(TYPE, HOST+":"+PORT) // on se connecte au serveur
	print("on s'est connecte au serveur")
	if err != nil {
		panic(err)
	}
	sendFileToServer(connection, "matriceA.txt") // on lui envoie la matrice A

	defer connection.Close()

	connection, err = net.Dial(TYPE, HOST+":"+PORT) // on se reconnecte au serveur
	if err != nil {
		panic(err)
	}
	sendFileToServer(connection, "matriceB.txt") // on lui envoie la matrice B

	defer connection.Close()

	connection, err = net.Dial(TYPE, HOST+":"+PORT) // on se connecte au serveur et on attend la reponse
	if err != nil {
		panic(err)
	}

	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)

	connection.Read(bufferFileSize)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

	connection.Read(bufferFileName)
	fileName := strings.Trim(string(bufferFileName), ":")

	newFile, err := os.Create(fileName) // on cree le fichier qui contiendra la reponse

	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	var receivedBytes int64

	for {
		if (fileSize - receivedBytes) < BUFFERSIZE {
			io.CopyN(newFile, connection, (fileSize - receivedBytes))
			connection.Read(make([]byte, (receivedBytes+BUFFERSIZE)-fileSize)) // on recoit la reponse
			break
		}
		io.CopyN(newFile, connection, BUFFERSIZE)
		receivedBytes += BUFFERSIZE
	}
	fmt.Println("Received file completely!")
	connection.Close()

}
