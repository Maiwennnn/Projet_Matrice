package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
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

func fillString(returnString string, toLength int) string { // ajoute des ":" pour completer la taille du fichier transmis car elle ne fait pas forcement 9 octets
	for {
		lengtString := len(returnString)
		if lengtString < toLength {
			returnString = returnString + ":"
			continue
		}
		break
	}
	return returnString
}

func sendFileToClient(connection net.Conn, nomFile string) { //envoie le fichier au client
	fmt.Println("On envoie le fichier au client")
	file, err := os.Open(nomFile) // on ouvre le fichier
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Fichier ouvert sans erreur")

	fileInfo, err := file.Stat() // on extrait la taille du fichier
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Acces aux stats sans erreur")
	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fileName := fillString(fileInfo.Name(), 64)
	fmt.Println("Sending filesize!")
	connection.Write([]byte(fileSize)) // on envoie la taille du fichier au client
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

func ReadFile(fileName string, c chan string) { // on lit le fichier

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	s := string(data)
	c <- s
}

func ecritDansFichier(mat [][]int, nameFile string) { // on ecrit la matrice dans un fichier
	file, err := os.Create(nameFile)
	defer file.Close() // on ferme automatiquement a la fin de notre programme
	for i := 0; i < len(mat); i++ {
		s := ""
		for j := 0; j < len(mat); j++ {
			s = s + strconv.Itoa(mat[i][j]) + " "
		}

		_, err = file.WriteString(s) // ecrire dans le fichier
		s = "\n"
		_, err = file.WriteString(s)
	}
	fmt.Print(err)
}
func printMat(mat [][]int) { // on affiche la matrice
	fmt.Println("Matrice")
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			fmt.Print((mat)[i][j], " ")
		}
		fmt.Println("")
	}
}

func createEmptyMat(n int) [][]int { // on cree un matrice vide carree de taille n
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
	}
	return mat
}

func matProduct1(mat1 [][]int, mat2 [][]int, c chan []int, i int) {
	//this function calculates and returns the product of mat1 and mat2
	res := make([]int, len(mat1))

	for j := 0; j < len(mat1); j++ {
		for k := 0; k < len(mat1); k++ {
			res[j] += mat1[i][k] * mat2[k][j]
		}
	}
	c <- res

}

func matriceEnInt(data string) [][]int { // on convertit un string en matrice
	// On separe la chaine de caracteres quand on rencontre des symboles "a la ligne"
	data2 := strings.Split(strings.ReplaceAll(data, "\r\n", "\n"), "\n")
	n := len(data2)

	//On fait une matrice de strings
	data3 := make([][]string, n)
	for i := range data3 {
		data3[i] = make([]string, n)
	}

	// On separe les chaines de caracteres quand on rencontre des espaces
	for i := range data2 {
		splitage := strings.Split(data2[i], " ")
		for j := range splitage {
			data3[i][j] = splitage[j]
		}
	}

	//On convertit les strings du tableau en entiers
	matrice_finale := createEmptyMat(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			conversion, err := strconv.Atoi(data3[i][j])
			matrice_finale[i][j] = conversion
			if err != nil {
				fmt.Println("Error during conversion")
			}
		}
	}

	return matrice_finale

}

func traitement_matrices() { //on lit les matrices recues, puis on les multiplie et on ecrit le resultat dans un nouveau fichier
	cA := make(chan string)
	cB := make(chan string)
	go ReadFile("matriceA.txt", cA)
	go ReadFile("matriceB.txt", cB)
	matA := <-cA
	matB := <-cB

	matAint := matriceEnInt(matA)
	matBint := matriceEnInt(matB)

	//on cree autant de goroutines que de colonnes dans la matrice resultat
	//on fait un tableau de channel pour retourner chaque colonne

	channel_second := make([]chan []int, len(matAint))
	for i := range channel_second {
		channel_second[i] = make(chan []int)
	}

	matC := make([][]int, len(matAint))
	for i := 0; i < len(matAint); i++ {
		go matProduct1(matAint, matBint, channel_second[i], i)
		matC[i] = <-channel_second[i]
	}
	ecritDansFichier(matC, "matriceC.txt")
}

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
	for i := 0; i < 2; i++ {
		connection, err := server.Accept()
		print(" on a accepte une connexion")
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		fmt.Println("Client connected")
		connection.Read(bufferFileSize)
		fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)
		connection.Read(bufferFileName)
		fileName := strings.Trim(string(bufferFileName), ":")

		fmt.Println("On va creer le file")

		newFile, err := os.Create(fileName)

		if err != nil {
			panic(err)
		}
		defer newFile.Close()
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
		connection.Close()
		if i == 1 {
			break
			server.Close()
		}

	}
	//server, err = net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		fmt.Println("Error listening: ", err)
		os.Exit(1)
	}
	fmt.Println("Server started! Waiting for connections...")

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		fmt.Println("Client connected")

		traitement_matrices()
		go sendFileToClient(connection, "matriceC.txt")

	}

}
