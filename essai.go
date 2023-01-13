package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
func main() {
	matEnStr := ReadFile("test.txt")
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
