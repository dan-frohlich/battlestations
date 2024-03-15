package main

import (
	"flag"
	"log"
	"os"

	"github.com/dan-frohlich/battlestations/character/print"
)

func main() {
	fileNamePtr := flag.String("file", "", "path to character file")
	flag.Parse()

	if fileNamePtr == nil || *fileNamePtr == "" {
		flag.Usage()
		log.Fatal("web mode not available")
	}

	fileName := *fileNamePtr
	r, e := os.Open(fileName)
	checkFatal(e, "load "+fileName)

	charData, e := print.LoadCharFromReader(r)
	checkFatal(e, "load character")

	e = print.WritePDFFile(charData)
	checkFatal(e, "write pdf")
}

func checkFatal(err error, msg string) {
	if err != nil {
		log.Fatal(msg, " ", err)
	}
}
