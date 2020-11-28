package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func main() {

	data := `fish,blue,water
dog,brown,house
tiger,orange,jungle
crocodrile,green,river`

	lector := csv.NewReader(strings.NewReader(data))

	contenido, err := lector.Read()

	if err != nil {
		log.Fatal(err)
	}

	archivo, err := os.OpenFile("databases.csv", os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer archivo.Close()
	escritor := csv.NewWriter(archivo)

	err = escritor.Write(contenido)

	if err != nil {
		log.Fatal(err)
	}

}
