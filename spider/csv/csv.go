package csv

import (
	//0.1
	"log"

	//0.2
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

//0.3
func Write(path string, data [][]string) {
	file, err := os.Create(path)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

//0.3
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

//0.4
// read template CSV file
// parameter file path
// return a list
func Read(path string) (items_id []string, items_name []string) {
	f, _ := os.Open(path)
	const NAME_COLUMN = 0
	const ID_COLUMN = 2

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		items_name = append(items_name, record[NAME_COLUMN])
		items_id = append(items_id, record[ID_COLUMN])
	}
	return
}
