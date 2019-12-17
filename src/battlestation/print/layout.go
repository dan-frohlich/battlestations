package main

import (
	"github.com/jung-kurt/gofpdf"
	"log"
)

type layout map[string]*cell

type cell struct {
	x, y, w, h   float64
	align        string
	border       string
	fontFamily   string
	fontSize     float64
	fontWeight   string
	overflowRows int
}

func (l layout) draw(pdf *gofpdf.Fpdf, name string, text string, border bool) bool {
	c, ok := l[name]
	if ok {
		c.drawCell(pdf, text, border)
	}
	return ok
}

func (c *cell) drawCell(pdf *gofpdf.Fpdf, text string, border bool) {
	pdf.SetXY(c.x, c.y)
	brdr := ""
	if border {
		brdr = "1"
	}
	pdf.SetFont(c.fontFamily, c.fontWeight, c.fontSize)
	lines := pdf.SplitText(text, c.w)
	rows := c.overflowRows + 1
	if len(lines) > rows {
		log.Printf("WARN cell needs [%d] lines to handle text [%s]", len(lines), text)
	}
	for i, line := range lines {
		if i == rows {
			break
		}
		if i != 0 {
			line = " " + line
		}
		pdf.CellFormat(c.w, c.h, line, brdr, 2, c.align, false, -1, "")
		//pdf.Cell(c.w, c.h, line)
	}
}

func loadBkgrndImg(pdf *gofpdf.Fpdf) {
	r, err := fileReader(baseDir, smImgFNm)
	check(err, "loading background image")
	pdf.RegisterImageReader(smImgNm, smImgType, r)
}
