package main

import (
	"flag"
	"log"
	"os"

	"github.com/dan-frohlich/battlestations/character"
	"github.com/dan-frohlich/battlestations/character/model"
)

func main() {
	printPtr := flag.Bool("print", false, "print action")
	fileNamePtr := flag.String("file", "", "path to character file")
	flag.Parse()

	if printPtr == nil || !*printPtr {
		flag.Usage()
		log.Fatal("print action required")
	}

	switch *printPtr {
	case true:
		//print command
		if fileNamePtr == nil || *fileNamePtr == "" {
			flag.Usage()
			log.Fatal("file param required for print action")
		}
		printCharAction(*fileNamePtr)
	default:
		flag.Usage()
		log.Fatal("no action specified")
	}

}

func printCharAction(fileName string) {
	b, e := os.ReadFile(fileName)
	checkFatal(e, "read "+fileName)

	c, e := model.LoadCharacter(b)
	checkFatal(e, "load charcter from "+fileName)

	issues, errs := model.NewCharacterValidator().ValidateAll(c)
	for _, e := range errs {
		log.Default().Printf("ERROR validation error: %s", e)
	}
	for _, issue := range issues {
		log.Default().Printf("WARN validation issue: %s", issue)
	}

	m := &character.Manager{}
	m.SetCharacter(c)
	m.Print()
}

func checkFatal(err error, msg string) {
	if err != nil {
		log.Fatal(msg, " ", err)
	}
}
