package main

import (
	"fmt"
	"io"

	"github.com/jung-kurt/gofpdf"
)

func writePDFFile(charData bsChar) error {
	pdf, _ := renderPDF(charData)
	name := fmt.Sprintf("%s_%v_%v.pdf", charData.Name, charData.Rank, charData.Prestige)
	return pdf.OutputFileAndClose(name)
}

func writePDF(charData bsChar, output io.WriteCloser) error {
	pdf, _ := renderPDF(charData)
	return pdf.OutputAndClose(output)
}

func renderPDF(charData bsChar) (*gofpdf.Fpdf, error) {
	orientation := "L"
	sheetSize := "A5"
	useLargeTemplate := charData.isLarge()
	if useLargeTemplate {
		orientation = "P"
		sheetSize = "Letter"
	}
	pdf := gofpdf.New(orientation, "mm", sheetSize, baseDir)
	//pdf.SetMargins(0, 0, 0)
	pdf.SetAutoPageBreak(false, 0)
	err := loadBkgrndImg(pdf, useLargeTemplate)
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
	chm := charData.toMap()
	for k, v := range chm {
		theLayout.draw(pdf, layoutKey(k), v, DrawBorder)
	}
	saAttr := []string{"name", "notes", "pool"}
	sal := charData.SpecialAbilities
	for i, sa := range sal {
		sam := sa.toMap()
		for _, k := range saAttr {
			v := sam[k]
			lk := layoutTableKeyName(ttSA, i, k)
			theLayout.draw(pdf, lk, v, DrawBorder)
		}
	}
	eqAttr := []string{"name", "notes", "mass", "status"}
	eql := charData.Equipment
	for i, eq := range eql {
		eqm := eq.toMap()
		for _, k := range eqAttr {
			v := eqm[k]
			lk := layoutTableKeyName(ttEq, i, k)
			theLayout.draw(pdf, lk, v, DrawBorder)
		}
	}
	return pdf, err
}
