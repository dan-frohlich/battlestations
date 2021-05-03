package main

import "fmt"

const (
	smallLayoutFont = "Times"
)

var smallLayout = layout{
	"name": &cell{
		x: 27, y: 13, w: 65, h: 8,
		align:      "LM",
		fontFamily: smallLayoutFont,
		fontSize:   16,
		fontWeight: "B",
	},
	"profession": &cell{
		x: 27, y: 21, w: 65, h: 8,
		align:      "LM",
		fontFamily: smallLayoutFont,
		fontSize:   12,
	},
	"species": &cell{
		x: 27, y: 29, w: 65, h: 8,
		align:      "LM",
		fontFamily: smallLayoutFont,
		fontSize:   12,
	},
	"alien_ability": &cell{
		x: 27, y: 37, w: 65, h: 5,
		align:        "LM",
		fontFamily:   smallLayoutFont,
		fontSize:     10,
		fontWeight:   "I",
		overflowRows: 2,
	},

	"athletics": &cell{
		x: 107, y: 13, w: 17, h: 8,
		align:      "CM",
		fontFamily: smallLayoutFont,
		fontSize:   16,
		fontWeight: "B",
	},
	"combat": &cell{
		x: 107, y: 21, w: 17, h: 8,
		align:      "CM",
		fontFamily: smallLayoutFont,
		fontSize:   16,
		fontWeight: "B",
	},
	"engineering": &cell{
		x: 107, y: 29, w: 17, h: 8,
		align:      "CM",
		fontFamily: smallLayoutFont,
		fontSize:   16,
		fontWeight: "B",
	},
	"pilot": &cell{
		x: 107, y: 37, w: 17, h: 8,
		align:      "CM",
		fontFamily: smallLayoutFont,
		fontSize:   16,
		fontWeight: "B",
	},
	"science": &cell{
		x: 107, y: 45, w: 17, h: 8,
		align:      "CM",
		fontFamily: smallLayoutFont,
		fontSize:   16,
		fontWeight: "B",
	},

	"base_hp": &cell{
		x: 138, y: 13, w: 19, h: 8,
		align:      "CM",
		fontFamily: smallLayoutFont,
		fontSize:   16,
		fontWeight: "B",
	},
	"hp": &cell{
		x: 171, y: 13, w: 32, h: 8,
		align:      "CM",
		fontFamily: smallLayoutFont,
		fontSize:   16,
		fontWeight: "B",
	},
	"move": &cell{
		x: 138, y: 21, w: 19, h: 8,
		align:      "CM",
		fontFamily: smallLayoutFont,
		fontSize:   16,
		fontWeight: "B",
	},
	"luck": &cell{
		x: 138, y: 29, w: 19, h: 8,
		align:      "CM",
		fontFamily: smallLayoutFont,
		fontSize:   16,
		fontWeight: "B",
	},
	"target": &cell{
		x: 138, y: 37, w: 19, h: 8,
		align:      "CM",
		fontFamily: smallLayoutFont,
		fontSize:   16,
		fontWeight: "B",
	},
	"hands": &cell{
		x: 138, y: 45, w: 19, h: 8,
		align:      "CM",
		fontFamily: smallLayoutFont,
		fontSize:   16,
		fontWeight: "B",
	},
	"rank": &cell{
		x: 171, y: 37, w: 32, h: 8,
		align:      "CM",
		fontFamily: smallLayoutFont,
		fontSize:   16,
		fontWeight: "B",
	},
	"carry": &cell{
		x: 171, y: 45, w: 32, h: 8,
		align:      "CM",
		fontFamily: smallLayoutFont,
		fontSize:   16,
		fontWeight: "B",
	},

	"prestige": &cell{
		x: 40, y: 55, w: 32, h: 6,
		align:      "CR",
		fontFamily: smallLayoutFont,
		fontSize:   14,
		fontWeight: "I",
	},
	"experience": &cell{
		x: 105, y: 55, w: 32, h: 6,
		align:      "CR",
		fontFamily: smallLayoutFont,
		fontSize:   14,
		fontWeight: "I",
	},
	"credits": &cell{
		x: 170, y: 55, w: 32, h: 6,
		align:      "CR",
		fontFamily: smallLayoutFont,
		fontSize:   14,
		fontWeight: "I",
	},
}

func smallLayoutInit() {
	//add special abilities
	dy := 6.2
	offsetX := 7.0
	height := 6.0
	tableType := ttSA
	for row := 0; row < 10; row++ {
		key := keyName(tableType, row, "name")
		offset := float64(row) * dy
		smallLayout[key] = &cell{
			x: offsetX, y: 76 + offset, w: 30, h: height,
			align: "LM", fontFamily: smallLayoutFont, fontSize: 10, fontWeight: "B",
		}
		key = keyName(tableType, row, "notes")
		smallLayout[key] = &cell{
			x: 31 + offsetX, y: 76 + offset, w: 38, h: height,
			align: "LM", fontFamily: smallLayoutFont, fontSize: 8, fontWeight: "I",
		}
		key = keyName(tableType, row, "pool")
		smallLayout[key] = &cell{
			x: 69 + offsetX, y: 76 + offset, w: 12, h: height,
			align: "CM", fontFamily: smallLayoutFont, fontSize: 12, fontWeight: "B",
		}
	}

	tableType = ttEq
	offsetX = 101.5
	for row := 0; row < 10; row++ {
		key := keyName(tableType, row, "name")
		offset := float64(row) * dy
		smallLayout[key] = &cell{
			x: offsetX, y: 76 + offset, w: 30, h: height,
			align: "LM", fontFamily: smallLayoutFont, fontSize: 10, fontWeight: "B",
		}
		key = keyName(tableType, row, "notes")
		smallLayout[key] = &cell{
			x: 30.5 + offsetX, y: 76 + offset, w: 38.5, h: height,
			align: "LM", fontFamily: smallLayoutFont, fontSize: 8, fontWeight: "I",
		}
		key = keyName(tableType, row, "mass")
		smallLayout[key] = &cell{
			x: 69 + offsetX, y: 76 + offset, w: 14.5, h: height,
			align: "CM", fontFamily: smallLayoutFont, fontSize: 12, fontWeight: "B",
		}
		key = keyName(tableType, row, "status")
		smallLayout[key] = &cell{
			x: 84 + offsetX, y: 76 + offset, w: 19, h: height,
			align: "LM", fontFamily: smallLayoutFont, fontSize: 8, fontWeight: "I",
		}
	}
}

type tableType string
const (
	ttSA = tableType("sa")
	ttEq = tableType("eq")
)
func keyName(tableType tableType, row int, field string) string {
	return fmt.Sprintf("%s.%d.%s", tableType, row, field)
}
