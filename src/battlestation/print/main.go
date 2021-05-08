package main

import (
	"flag"
	"log"
	"os"
)

const (
	baseDir       = "./assets"
	fontFNm       = "vt323"
	fontNm        = "Roddenberry"
	sampleCharFNm = "sample_char.yml"
	smImgNm       = "background_small"
	smImgType     = "png"
	lgImgNm       = "background_large"
	lgImgType     = "png"
	smImgFNm      = "sheet_sm.png"
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

	charData, e := loadCharFromReader(r)
	checkFatal(e, "load character")

	e = writePDFFile(charData)
	checkFatal(e, "write pdf")
}

func checkFatal(err error, msg string) {
	if err != nil {
		log.Fatal(msg, " ", err)
	}
}
