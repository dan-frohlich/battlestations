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
	lgImgNm       = "background_large"
	lgImgType     = "png"
	smImgFNm      = "sheet_sm.png"
)

func main() {

	charData := loadChar()
	isLarge := useLargeTemplate(charData)

	orientation := "L"
	sheetSize := "A5"
	if isLarge {
		orientation = "P"
		sheetSize = "Letter"
	}
	pdf := gofpdf.New(orientation, "mm", sheetSize, baseDir)
	//pdf.SetMargins(0, 0, 0)
	pdf.SetAutoPageBreak(false, 0)
	loadBkgrndImg(pdf, isLarge)
	pdf.AddPage()
	options := gofpdf.ImageOptions{
		ImageType:             "png",
		ReadDpi:               true,
		AllowNegativePosition: true,
	}
	theLayout := smallLayout
	if isLarge {
		pdf.ImageOptions(lgImgNm, 5, 5, 200, 280, false, options, 0, "'")
		largeLayoutInit()
		theLayout = largeLayout
	}else{
		pdf.ImageOptions(smImgNm, 5, 5, 200, 140, false, options, 0, "'")
		smallLayoutInit()
	}


	DrawBorder := false

	for k, v := range charData {
		if s, ok := v.(string); ok {
			theLayout.draw(pdf, k, s, DrawBorder)
		}
		if i, ok := v.(int); ok {
			theLayout.draw(pdf, k, fmt.Sprintf("%d", i), DrawBorder)
		}
	}

	saAttr := []string{"name", "notes", "pool"}
	sa := charData["special_abilities"]
	if sal, ok := sa.([]interface{}); ok {
		for i, sai := range sal {
			if saim, ok := sai.(map[interface{}]interface{}); ok {
				for _, k := range saAttr {
					v := saim[k]
					lk := fmt.Sprintf("sa.%d.%s", i, k)
					if vs, ok := v.(string); ok {
						theLayout.draw(pdf, lk, vs, DrawBorder)
					}
					if vi, ok := v.(int); ok {
						theLayout.draw(pdf, lk, fmt.Sprintf("%d", vi), DrawBorder)
					}
				}
			}
		}
	}

	eqAttr := []string{"name", "notes", "mass", "status"}
	eq := charData["equipment"]
	if sal, ok := eq.([]interface{}); ok {
		for i, eqi := range sal {
			if eqim, ok := eqi.(map[interface{}]interface{}); ok {
				for _, k := range eqAttr {
					v := eqim[k]
					lk := fmt.Sprintf("eq.%d.%s", i, k)
					if vs, ok := v.(string); ok {
						theLayout.draw(pdf, lk, vs, DrawBorder)
					}
					if vi, ok := v.(int); ok {
						theLayout.draw(pdf, lk, fmt.Sprintf("%d", vi), DrawBorder)
					}
				}
			}
		}
	}

	name := fmt.Sprintf("%s_%v_%v.pdf", charData["name"], charData["rank"], charData["prestige"])
	err := pdf.OutputFileAndClose(name)
	check(err, "write pdf")
}

func useLargeTemplate(charData map[string]interface{}) bool {
	if val, ok := charData["special_abilities"]; ok {
		if sas, ok := val.([]interface{}); ok {
			if len(sas) > 10 {
				return true
			}
		}
	}

	if val, ok := charData["equipment"]; ok {
		if equip, ok := val.([]interface{}); ok {
			if len(equip) > 10 {
				return true
			}
		}
	}
	return false
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

func filePath(baseDir string, name string) string {
	path := fmt.Sprintf("%s%c%s", baseDir, os.PathSeparator, name)
	return path
}
