package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func renderPDF(charData bsChar, useLargeTemplate bool) {
	orientation := "L"
	sheetSize := "A5"
	if useLargeTemplate {
		orientation = "P"
		sheetSize = "Letter"
	}
	pdf := gofpdf.New(orientation, "mm", sheetSize, baseDir)
	//pdf.SetMargins(0, 0, 0)
	pdf.SetAutoPageBreak(false, 0)
	loadBkgrndImg(pdf, useLargeTemplate)
	pdf.AddPage()
	options := gofpdf.ImageOptions{
		ImageType:             "png",
		ReadDpi:               true,
		AllowNegativePosition: true,
	}
	theLayout := smallLayout
	if useLargeTemplate {
		pdf.ImageOptions(lgImgNm, 5, 5, 200, 280, false, options, 0, "'")
		largeLayoutInit()
		theLayout = largeLayout
	} else {
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
					lk := keyName(ttSA, i, k)

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
					lk := keyName(ttEq, i, k)
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
