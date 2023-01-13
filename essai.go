package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	mat, err := os.Stdout.Write(data) // lis matrice dans fichier et l'écris dans terminal
	fmt.Println(" pause2")

	tab := int(data)
	i :=0
	for i in range (len(tab)):

}
