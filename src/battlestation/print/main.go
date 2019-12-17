package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
	"gopkg.in/yaml.v2"
)

const (
	baseDir       = "./assets"
	fontFNm       = "vt323"
	fontNm        = "Roddenberry"
	sampleCharFNm = "sample_char.yml"
	smImgNm       = "background_small"
	smImgType     = "png"
	smImgFNm      = "sheet_sm.png"
)

func main() {

	pdf := gofpdf.New("L", "mm", "A5", baseDir)
	loadBkgrndImg(pdf)
	pdf.AddPage()
	options := gofpdf.ImageOptions{
		ImageType:             "png",
		ReadDpi:               true,
		AllowNegativePosition: true,
	}
	pdf.ImageOptions(smImgNm, 5, 5, 200, 140, false, options, 0, "'")

	t := map[string]string{}
	r, e := fileReader(baseDir, sampleCharFNm)
	check(e, "load "+sampleCharFNm)
	data, e := ioutil.ReadAll(r)
	check(e, "read  "+sampleCharFNm)
	e = yaml.Unmarshal(data, &t)
	check(e, "unmarshal  "+sampleCharFNm)

	for k, v := range t {
		smallLayout.draw(pdf, k, v, false)
	}

	//smallLayout.draw(pdf, "name", "Lt. Dan", false)
	//smallLayout.draw(pdf, "profession", "Scientist", false)
	//smallLayout.draw(pdf, "species", "Human", false)
	//smallLayout.draw(pdf, "alien_ability", "Willpower: re-roll both dice on professional skill checks, yada, yada, yada, yada, yada, yada", false)
	//
	//smallLayout.draw(pdf, "athletics", "1", false)
	//smallLayout.draw(pdf, "combat", "2", false)
	//smallLayout.draw(pdf, "engineering", "3", false)
	//smallLayout.draw(pdf, "pilot", "4", false)
	//smallLayout.draw(pdf, "science", "5", false)
	//
	//smallLayout.draw(pdf, "base_hp", "5", true)
	//smallLayout.draw(pdf, "move", "6", true)
	//smallLayout.draw(pdf, "luck", "6", true)
	//smallLayout.draw(pdf, "target", "6", true)
	//smallLayout.draw(pdf, "hands", "6", true)

	err := pdf.OutputFileAndClose("hello.pdf")
	check(err, "write pdf")
}

//
//func loadFont(pdf *gofpdf.Fpdf) {
//	r, err := fileReader(baseDir, fontFNm)
//	check(err, "loading Font")
//	pdf.AddFontFromReader(fontNm, "", r)
//}

func check(err error, msg string) {
	if err != nil {
		log.Fatal(msg, " ", err)
	}
}

func fileReader(baseDir, name string) (r io.Reader, err error) {
	path := filePath(baseDir, name)
	return os.Open(path)
}

func filePath(dir string, name string) string {
	path := fmt.Sprintf("%s%c%s", baseDir, os.PathSeparator, name)
	return path
}
