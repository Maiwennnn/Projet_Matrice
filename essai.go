package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func ReadFile(fileName string) string {

	fmt.Printf("\n\nReading a file in Go lang\n")
	//fileName := "test.txt"

	// The ioutil package contains inbuilt
	// methods like ReadFile that reads the
	// filename and returns the contents.
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data) // conversion de byte en string
	s := string(data)
	return s
}

func ecritDansFichier(mat [][]int, nameFile string) {
	file, err := os.OpenFile(nameFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer file.Close() // on ferme automatiquement à la fin de notre programme
	for i := 0; i < len(mat); i++ {
		s := ""
		for j := 0; j < len(mat); j++ {
			s = s + strconv.Itoa(mat[i][j]) + " "
		}

		_, err = file.WriteString(s) // écrire dans le fichier
		s = "\n"
		_, err = file.WriteString(s)
	}
	fmt.Print(err)
}
func main() {
	a := [][]int{
		{1, 3},
		{6, 9},
	}

	go secritDansFichier(a, "test.txt")
	matEnStr := go ReadFile("test.txt")
	fmt.Printf("2eme print")
	fmt.Printf(matEnStr)

	/*data, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	mat, err := os.Stdout.Write(data) // lis matrice dans fichier et l'écris dans terminal
	fmt.Println(" pause2")
	tab := make([][]int, len(data))

	for i := 0; i < len(data); i++ {
		tab[i] = int(data[i])
	}*/

}
