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
	//pdf.SetMargins(0, 0, 0)
	pdf.SetAutoPageBreak(false, 0)
	loadBkgrndImg(pdf)
	pdf.AddPage()
	options := gofpdf.ImageOptions{
		ImageType:             "png",
		ReadDpi:               true,
		AllowNegativePosition: true,
	}
	pdf.ImageOptions(smImgNm, 5, 5, 200, 140, false, options, 0, "'")

	t := loadChar()

	smallLayoutInit()

	DrawBorder := false

	for k, v := range t {
		if s, ok := v.(string); ok {
			smallLayout.draw(pdf, k, s, DrawBorder)
		}
		if i, ok := v.(int); ok {
			smallLayout.draw(pdf, k, fmt.Sprintf("%d", i), DrawBorder)
		}
	}

	saAttr := []string{"name", "notes", "pool"}
	sa := t["special_abilities"]
	if sal, ok := sa.([]interface{}); ok {
		for i, sai := range sal {
			if saim, ok := sai.(map[interface{}]interface{}); ok {
				for _, k := range saAttr {
					v := saim[k]
					lk := fmt.Sprintf("sa.%d.%s", i, k)
					if vs, ok := v.(string); ok {
						smallLayout.draw(pdf, lk, vs, DrawBorder)
					}
					if vi, ok := v.(int); ok {
						smallLayout.draw(pdf, lk, fmt.Sprintf("%d", vi), DrawBorder)
					}
				}
			}
		}
	}

	eqAttr := []string{"name", "notes", "mass", "status"}
	eq := t["equipment"]
	if sal, ok := eq.([]interface{}); ok {
		for i, eqi := range sal {
			if eqim, ok := eqi.(map[interface{}]interface{}); ok {
				for _, k := range eqAttr {
					v := eqim[k]
					lk := fmt.Sprintf("eq.%d.%s", i, k)
					if vs, ok := v.(string); ok {
						smallLayout.draw(pdf, lk, vs, DrawBorder)
					}
					if vi, ok := v.(int); ok {
						smallLayout.draw(pdf, lk, fmt.Sprintf("%d", vi), DrawBorder)
					}
				}
			}
		}
	}

	name := fmt.Sprintf("%s_%v_%v.pdf", t["name"], t["rank"], t["prestige"])
	err := pdf.OutputFileAndClose(name)
	check(err, "write pdf")
}

func loadChar() map[string]interface{} {
	var r io.Reader
	var e error
	var fileName string
	if len(os.Args) > 1 {
		fileName = os.Args[1]
		r, e = os.Open(fileName)
		check(e, "load "+fileName)
	} else {
		fileName = sampleCharFNm
		r, e = fileReader(baseDir, fileName)
		check(e, "load "+fileName)
	}
	t := map[string]interface{}{}
	data, e := ioutil.ReadAll(r)
	check(e, "read  "+fileName)
	e = yaml.Unmarshal(data, &t)
	check(e, "unmarshal  "+fileName)
	return t
}

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
