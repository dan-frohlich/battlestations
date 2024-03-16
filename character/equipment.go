package character

type GearType int

func (gt GearType) String() string {
	switch gt {
	case CyberwareGearType:
		return "cyberware"
	case DrugGearType:
		return "drug"
	case ExplosiveGearType:
		return "explosive"
	case GeneralGearType:
		return "general equipment"
	case MeleeWeaponGearType:
		return "melee"
	case RangedWeaponGearType:
		return "ranged"
	case SlugAmmunitionType:
		return "slug ammo"
	case ToxinGearType:
		return "toxin"
	default:
		return ""
	}
}

const (
	CyberwareGearType = iota
	DrugGearType
	ExplosiveGearType
	GeneralGearType
	MeleeWeaponGearType
	RangedWeaponGearType
	SlugAmmunitionType
	ToxinGearType
)

type Gear struct {
	Cost   Credits
	Energy YN
	Mass   Mass
	Name   string
	Notes  string
	Type   GearType
}

var (
	ArmorGear   = Gear{Name: "Armor", Cost: 200, Mass: 10, Energy: N, Notes: "- 1 damage.", Type: GeneralGearType}
	BlasterGear = Gear{Name: "Blaster", Cost: 250, Mass: 4, Energy: Y, Notes: "2d6-1 damage", Type: RangedWeaponGearType}
	MedKitGear  = Gear{Name: "MedKit", Cost: 250, Mass: 5, Energy: Y, Notes: "Sci. vs. 11 heals 1 d6 dam.", Type: GeneralGearType}
	ToolkitGear = Gear{Name: "ToolKit", Cost: 100, Mass: 5, Energy: Y, Notes: "- 1 difficulty on Engineering checks to upgrade or repair.", Type: GeneralGearType}
)
