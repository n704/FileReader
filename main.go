package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	fileReader "github.com/n704/file_reader/lib"
)

func main() {
	// sample command
	// file_reader -file github.com/n704/file_reader/main.go
	fileName := flag.String("file", "", "Enter path/to/file user want to read")
	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	lineReader, ok := fileReader.GetLineReader(file)
	if !ok {
		log.Fatalln("FileReader not created")
	}
	defer lineReader.Close()
	output, ok := lineReader.ReadLine()
	for ok {
		fmt.Println(output)
		output, ok = lineReader.ReadLine()
		lineReader.Close()
	}
}
