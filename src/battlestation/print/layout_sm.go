package main

var smallLayout = layout{
	"name": &cell{
		x: 27, y: 13, w: 65, h: 8,
		align:      "LM",
		fontFamily: "Courier",
		fontSize:   12,
	},
	"profession": &cell{
		x: 27, y: 21, w: 65, h: 8,
		align:      "LM",
		fontFamily: "Courier",
		fontSize:   12,
	},
	"species": &cell{
		x: 27, y: 29, w: 65, h: 8,
		align:      "LM",
		fontFamily: "Courier",
		fontSize:   12,
	},
	"alien_ability": &cell{
		x: 27, y: 37, w: 65, h: 5,
		align:        "LM",
		fontFamily:   "Courier",
		fontSize:     10,
		fontWeight:   "I",
		overflowRows: 2,
	},

	"athletics": &cell{
		x: 107, y: 13, w: 17, h: 8,
		align:      "CM",
		fontFamily: "Courier",
		fontSize:   16,
		fontWeight: "B",
	},
	"combat": &cell{
		x: 107, y: 21, w: 17, h: 8,
		align:      "CM",
		fontFamily: "Courier",
		fontSize:   16,
		fontWeight: "B",
	},
	"engineering": &cell{
		x: 107, y: 29, w: 17, h: 8,
		align:      "CM",
		fontFamily: "Courier",
		fontSize:   16,
		fontWeight: "B",
	},
	"pilot": &cell{
		x: 107, y: 37, w: 17, h: 8,
		align:      "CM",
		fontFamily: "Courier",
		fontSize:   16,
		fontWeight: "B",
	},
	"science": &cell{
		x: 107, y: 45, w: 17, h: 8,
		align:      "CM",
		fontFamily: "Courier",
		fontSize:   16,
		fontWeight: "B",
	},

	"base_hp": &cell{
		x: 138, y: 13, w: 19, h: 8,
		align:      "CM",
		fontFamily: "Courier",
		fontSize:   16,
		fontWeight: "B",
	},
	"move": &cell{
		x: 138, y: 21, w: 19, h: 8,
		align:      "CM",
		fontFamily: "Courier",
		fontSize:   16,
		fontWeight: "B",
	},
	"luck": &cell{
		x: 138, y: 29, w: 19, h: 8,
		align:      "CM",
		fontFamily: "Courier",
		fontSize:   16,
		fontWeight: "B",
	},
	"target": &cell{
		x: 138, y: 37, w: 19, h: 8,
		align:      "CM",
		fontFamily: "Courier",
		fontSize:   16,
		fontWeight: "B",
	},
	"hands": &cell{
		x: 138, y: 45, w: 19, h: 8,
		align:      "CM",
		fontFamily: "Courier",
		fontSize:   16,
		fontWeight: "B",
	},
}
