package main

import (
	"flag"
	"log"
	"os"

	"github.com/dan-frohlich/battlestations/character"
	"github.com/dan-frohlich/battlestations/character/model"
)

func main() {
	fileNamePtr := flag.String("file", "", "path to character file")
	flag.Parse()

	if fileNamePtr == nil || *fileNamePtr == "" {
		flag.Usage()
		log.Fatal("web mode not available")
	}

	fileName := *fileNamePtr
	b, e := os.ReadFile(fileName)
	checkFatal(e, "read "+fileName)

	c, e := model.LoadCharacter(b)
	checkFatal(e, "load charcter from "+fileName)
	m := &character.Manager{}
	m.SetCharacter(c)
	m.Print()
}

func checkFatal(err error, msg string) {
	if err != nil {
		log.Fatal(msg, " ", err)
	}
}
