package main

import(
	"os"
	"fmt"
	"log"
	"encoding/csv"
	"encoding/json"
	"io"
)

type Person struct {
	Name string
	Age string
	Address *Address
}

type Address struct {
	City string
	Country string
}

func main() {
	csvFile, error := os.Open("data.csv")

	if error != nil {
		log.Fatal(error)
	}

	csvReader := csv.NewReader(csvFile)
	var pasrsedData []Person

	for {
		line, error := csvReader.Read()

		if error == io.EOF {
			break
		}

		if error != nil {
			log.Fatal(error)
		}

		pasrsedData = append(pasrsedData, Person{
			Name: line[0],
			Age: line[1],
			Address: &Address{
				City: line[2],
				Country: line[3],
			},
		})
	}

	personJson, error:= json.Marshal(pasrsedData)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(string(personJson))
}