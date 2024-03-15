package character

type Rank int

const (
	_ Rank = iota
	EnsignRank
	LtJrGradeRank
	LieutenantRank
	CommanderRank
	CaptainRank
	MajorRank
	ColonelRank
	CommodoreRank
	AdmiralRank
	FleetAdmiralRank
	CommanderInChiefRank
	SenatorOfTheRepublicRank
)

func (r Rank) Int() int {
	return int(r)
}

func (r Rank) String() string {
	if s, ok := rankNames[r]; ok {
		return s.name
	}
	if r > SenatorOfTheRepublicRank {
		return rankNames[SenatorOfTheRepublicRank].name
	}
	return ""
}

func (r Rank) AbrvString() string {
	if s, ok := rankNames[r]; ok {
		return s.abbreviation
	}
	if r > SenatorOfTheRepublicRank {
		return rankNames[SenatorOfTheRepublicRank].abbreviation
	}
	return ""
}

type rankName struct {
	name         string
	abbreviation string
}

var rankNames = map[Rank]rankName{
	EnsignRank:               {name: "ensign", abbreviation: "ENS"},
	LtJrGradeRank:            {name: "lt. jr. grade", abbreviation: "LTJG"},
	LieutenantRank:           {name: "lieutenant", abbreviation: "LT"},
	CommanderRank:            {name: "commander", abbreviation: "CDR"},
	CaptainRank:              {name: "captain", abbreviation: "CAP"},
	MajorRank:                {name: "major", abbreviation: "MAG"},
	ColonelRank:              {name: "colonel", abbreviation: "COL"},
	CommodoreRank:            {name: "commodore", abbreviation: ""},
	AdmiralRank:              {name: "admiral", abbreviation: "ADM"},
	FleetAdmiralRank:         {name: "fleet admiral", abbreviation: "FADM"},
	CommanderInChiefRank:     {name: "commander in chief", abbreviation: "CNC"},
	SenatorOfTheRepublicRank: {name: "senator", abbreviation: "SEN"},
}
