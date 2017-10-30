package main

import (
	"os"
	"log"
	"encoding/csv"
	"github.com/alldroll/go-datastructures/cdb"
)

func main() {
	var sourceFile *os.File

	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s source", os.Args[0])
	}

	sourceFile, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0)
	if err != nil {
		log.Fatalf("[Fail to open source file] %s", err)
	}

	defer sourceFile.Close()

	cdbHandle := cdb.New()
	cdbReader, err := cdbHandle.GetReader(sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	iterator, err := cdbReader.Iterator()
	if err != nil {
		log.Fatal(err)
	}

	csvWriter := csv.NewWriter(os.Stdout)
	defer csvWriter.Flush()

	for {
		key, value := iterator.Key(), iterator.Value()
		if key == nil {
			break
		}

		err = csvWriter.Write([]string{string(key), string(value)})
		if err != nil {
			log.Fatal(err)
		}

		ok, err := iterator.Next()
		if err != nil {
			log.Fatal(err)
		}

		if !ok {
			break
		}
	}
}
