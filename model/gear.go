package model

import (
	"fmt"
	"strings"
	"sync"
)

type YN bool

func (v YN) String() string {
	if v {
		return "y"
	}
	return "n"
}

const (
	Y YN = true
	N YN = false
)

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
	Cost   int
	Energy YN
	Mass   int
	Name   string
	Notes  string
	Type   GearType
}

var (
	indexGearOnce     = &sync.Once{}
	gearIndex         = make(map[string]Gear)
	gearIndexedByType = make(map[GearType][]Gear)
)

func indexGear() {
	fn := func() {
		gg := make([]Gear, 0, len(drugs)+len(cyberware)+len(explosives)+len(generalGear)+len(meleeWeapons)+len(rangedWeapons)+len(toxins))
		gg = append(gg, cyberware...)
		gg = append(gg, drugs...)
		gg = append(gg, explosives...)
		gg = append(gg, generalGear...)
		gg = append(gg, meleeWeapons...)
		gg = append(gg, rangedWeapons...)
		gg = append(gg, slugAmmo...)
		gg = append(gg, toxins...)

		for _, g := range gg {
			gearIndex[strings.TrimSpace(g.Name)] = g
			gearIndexedByType[g.Type] = append(gearIndexedByType[g.Type], g)
		}
	}
	indexGearOnce.Do(fn)
}

func GetKits() []Gear {
	indexGear()
	return []Gear{
		gearIndex["MedKit"],
		gearIndex["ToolKit"],
	}
}

func GetGetGearByType(gt GearType) []Gear {
	indexGear()
	return gearIndexedByType[gt]
}

func GetGear(name string) Gear {
	indexGear()
	if g, ok := gearIndex[strings.TrimSpace(name)]; ok {
		return g
	}
	return Gear{}
}

func GearDetails(name string) string {
	indexGear()
	if g, ok := gearIndex[strings.TrimSpace(name)]; ok {
		return g.Details()
	}
	return ""
}

func IsType(name string, gt GearType) bool {
	indexGear()
	if g, ok := gearIndex[strings.TrimSpace(name)]; ok {
		return g.Type == gt
	}
	return false
}

func (g Gear) Details() string {
	return fmt.Sprintf("%s: [%s e? %s, %d cr, %d kg - %s]", g.Name, g.Type, g.Energy, g.Cost, g.Mass, g.Notes)
}

var cyberware = []Gear{
	{Name: "Autonurse", Cost: 500, Mass: 6, Energy: Y, Notes: "+1 point of healing per die.", Type: CyberwareGearType},
	{Name: "Cyberfoot", Cost: 2000, Mass: 8, Energy: Y, Notes: "+1 move.", Type: CyberwareGearType},
	{Name: "Cybergyros", Cost: 1000, Mass: 10, Energy: Y, Notes: "Treat OOC for you as one lower.", Type: CyberwareGearType},
	{Name: "Cyberhand", Cost: 2500, Mass: 6, Energy: Y, Notes: "+1 hand.", Type: CyberwareGearType},
	{Name: "Cyberhook", Cost: 500, Mass: 8, Energy: Y, Notes: "Built in vibrakife (1d6 damage). Other actions with it at +1 difficulty", Type: CyberwareGearType},
	{Name: "Medjack", Cost: 150, Mass: 2, Energy: Y, Notes: "Carries drugs** at half weight. -1 difficulty to heal you.", Type: CyberwareGearType},
	{Name: "Pro Chip", Cost: 1000, Mass: 2, Energy: Y, Notes: "-1 difficulty for Pro skill checks", Type: CyberwareGearType},
	{Name: "Set of Skill Chips", Cost: 6000, Mass: 12, Energy: Y, Notes: "-1 difficulty for ALL skill checks", Type: CyberwareGearType},
	{Name: "Skeletal Enhancement", Cost: 1500, Mass: 0, Energy: Y, Notes: "Increase carry by 10. Upgrade increases carry by 5 more.", Type: CyberwareGearType},
}

var drugs = []Gear{
	{Name: "Aggro", Cost: 25, Mass: 1, Energy: N, Notes: "Combat skill check difficulties reduced by 1. [+1 penaly on other skill checks]", Type: DrugGearType},
	{Name: "Charme", Cost: 25, Mass: 1, Energy: N, Notes: "Diplomacy skill check difficulties reduced by 1. [+1 penaly on other skill checks]", Type: DrugGearType},
	{Name: "Detox", Cost: 25, Mass: 1, Energy: N, Notes: "Eliminate all drugs and toxins in the recipient's system.", Type: DrugGearType},
	{Name: "Dull", Cost: 25, Mass: 1, Energy: N, Notes: "-1 point of damage per die. [+1 penaly on other skill checks]", Type: DrugGearType},
	{Name: "Equilout", Cost: 25, Mass: 1, Energy: N, Notes: "Ignore OOC.", Type: DrugGearType},
	{Name: "FlyBoy", Cost: 25, Mass: 1, Energy: N, Notes: "Piloting skill check difficulties reduced by 1. [+1 penaly on other skill checks]", Type: DrugGearType},
	{Name: "Numb", Cost: 25, Mass: 1, Energy: N, Notes: "Reroll highest die in your skill check and May Reroll highest damage die you suffer.", Type: DrugGearType},
	// {Name: "Patches", Cost: 25, Mass: 1, Energy: N, Notes: "Drug patch technology mass +1 but need not be in hand.", Type: DrugGearType},
	{Name: "Roid", Cost: 25, Mass: 1, Energy: N, Notes: "Athletics skill check difficulties reduced by 1. +100 to your Carry and +1 melee damage. [+1 penaly on other skill checks]", Type: DrugGearType},
	{Name: "Stim", Cost: 25, Mass: 1, Energy: N, Notes: "+3 to movement even if unconscious or dying (up to -20 hp) [+1 penaly on other skill checks]", Type: DrugGearType},
	{Name: "StunGone", Cost: 25, Mass: 1, Energy: N, Notes: "Reroll vs stun effects. Athletics check of 11 or be stunned when administered.", Type: DrugGearType},
	{Name: "SupSci", Cost: 25, Mass: 1, Energy: N, Notes: "Science skill check difficulties reduced by 1. [+1 penaly on other skill checks]", Type: DrugGearType},
	{Name: "TecKnow", Cost: 25, Mass: 1, Energy: N, Notes: "Engineering skill check difficulties reduced by 1. [+1 penaly on other skill checks]", Type: DrugGearType},
}

var explosives = []Gear{
	{Name: "EMP Grenade", Notes: "EMP all energied equipment and cyberware in L.O.S.", Energy: Y, Cost: 25, Mass: 1, Type: ExplosiveGearType},
	{Name: "Energy Grenade", Notes: "2d6 damage.", Energy: Y, Cost: 25, Mass: 1, Type: ExplosiveGearType},
	{Name: "Flare Grenade", Notes: "Makes person or object holding 3 easier to be targeted.", Energy: N, Cost: 25, Mass: 1, Type: ExplosiveGearType},
	{Name: "Frag Grenade", Notes: "2d6-1 damage.", Energy: N, Cost: 25, Mass: 1, Type: ExplosiveGearType},
	{Name: "Fritzer Grenade", Notes: "1d6 personal damage Roll 4d6 to break module.", Energy: Y, Cost: 25, Mass: 1, Type: ExplosiveGearType},
	{Name: "FrostBomb Grenade", Notes: "Frost effect on all in L.O.S.", Energy: Y, Cost: 25, Mass: 1, Type: ExplosiveGearType},
	{Name: "Gas Grenade", Notes: "Contains needler toxin*. All in cloud affected as if shot by needler (see page 96).", Energy: N, Cost: 25, Mass: 1, Type: ExplosiveGearType},
	{Name: "Ion Grenade", Notes: "No damage but all in L.O.S. have ionization level raised by one.", Energy: Y, Cost: 25, Mass: 1, Type: ExplosiveGearType},
	{Name: "Neutron Grenade", Notes: "No damage to modules, equipment. Every being in L.O.S. “Oucho”ed.", Energy: Y, Cost: 25, Mass: 1, Type: ExplosiveGearType},
	{Name: "Plasma Grenade", Notes: "1d6 fire damage to all in L.O.S. And within 6 squares.", Energy: Y, Cost: 25, Mass: 1, Type: ExplosiveGearType},
	{Name: "Smoke Grenade", Notes: "Cuts off LOS and Life Support in the module.", Energy: N, Cost: 25, Mass: 1, Type: ExplosiveGearType},
	{Name: "Stun Grenade", Notes: "All in L.O.S. make passive Athletics vs.11. Difference is stun markers.", Energy: Y, Cost: 25, Mass: 1, Type: ExplosiveGearType},
	{Name: "Satchel Charge", Notes: "1d6 hull. 4,5,6 breaks module. 2D6 personal damage.", Energy: Y, Cost: 50, Mass: 3, Type: ExplosiveGearType},
}

var generalGear = []Gear{
	{Name: "JetPack", Cost: 500, Mass: 4, Energy: Y, Notes: "Move up to 10 squares as one point of your movement with pilot skill check.", Type: GeneralGearType},
	{Name: "MagBoots", Cost: 150, Mass: 4, Energy: Y, Notes: "-1 movement when active but allows movement along hull without other penalties.", Type: GeneralGearType},
	{Name: "MedKit", Cost: 250, Mass: 5, Energy: Y, Notes: "Science Skill check vs. 11 to heal 1 die of damage", Type: GeneralGearType},
	{Name: "Scope", Cost: 25, Mass: 1, Energy: Y, Notes: "Consider distance to target as half.", Type: GeneralGearType},
	{Name: "ToolKit", Cost: 100, Mass: 5, Energy: Y, Notes: "- 1 difficulty on Engineering checks to upgrade or repair.", Type: GeneralGearType},
}

var meleeWeapons = []Gear{
	{Name: "Butt", Cost: 100, Mass: 4, Energy: N, Notes: "Add to ranged weapon to add 5 to range band and allow use as a melee weapon 1d6 melee.", Type: MeleeWeaponGearType},
	{Name: "Energy Blade", Cost: 1300, Mass: 8, Energy: Y, Notes: "3d6-3 damage. 1/6 break module. [☠ dangerous]", Type: MeleeWeaponGearType},
	{Name: "Knife", Cost: 5, Mass: 2, Energy: N, Notes: "1d6 damage.", Type: MeleeWeaponGearType},
	{Name: "Lightning Rod", Cost: 525, Mass: 5, Energy: Y, Notes: "1d6 damage and Stun Effect. [☠ dangerous]", Type: MeleeWeaponGearType},
	{Name: "Phase Pick", Cost: 300, Mass: 10, Energy: Y, Notes: "Melee Disintegrator: 1d6 damage and Athletics check of 8 or disintegrate", Type: MeleeWeaponGearType},
	{Name: "Plasma Dagger", Cost: 275, Mass: 5, Energy: Y, Notes: "Roll 2d6 fire damage and count either one. [☠ dangerous]", Type: MeleeWeaponGearType},
	{Name: "Sword", Cost: 10, Mass: 4, Energy: Y, Notes: "2d6-2 damage.", Type: MeleeWeaponGearType},
	{Name: "VibraKnife", Cost: 250, Mass: 4, Energy: Y, Notes: "1d6 damage. Ignore Damage Reduction.", Type: MeleeWeaponGearType},
}

var rangedWeapons = []Gear{
	{Name: "Arc Laser", Cost: 600, Mass: 7, Energy: Y, Notes: "1d6 damage (ignore Damage Reduction). Area Effect (indirect) [☠ dangerous]", Type: RangedWeaponGearType},
	{Name: "Blaster", Cost: 250, Mass: 4, Energy: Y, Notes: "2d6-1 damage", Type: RangedWeaponGearType},
	{Name: "Disintegrator", Cost: 550, Mass: 9, Energy: Y, Notes: "1d6 damage and Athletics check of 8 or disintegrate", Type: RangedWeaponGearType},
	{Name: "EMP Pistol", Cost: 300, Mass: 5, Energy: Y, Notes: "EMPs a target [🚫 can't damage module]", Type: RangedWeaponGearType},
	{Name: "Flyntlock", Cost: 25, Mass: 2, Energy: N, Notes: "2d6 damage (breaks itself after each use)", Type: RangedWeaponGearType},
	{Name: "Froster", Cost: 100, Mass: 3, Energy: Y, Notes: "Combat vs 8 to frost target square and each adjacent square [🚫 can't damage module]", Type: RangedWeaponGearType},
	{Name: "Ion Bore", Cost: 450, Mass: 9, Energy: Y, Notes: "Raises ionization level. 1d6 per ionization level (max 4d6)", Type: RangedWeaponGearType},
	{Name: "Laser", Cost: 300, Mass: 5, Energy: Y, Notes: "1d6 +1 Ignores damage reduction. -1 difficulty skill check to use", Type: RangedWeaponGearType},
	{Name: "Needler", Cost: 300, Mass: 4, Energy: N, Notes: "Deals 1 point of damage and delivers toxin*. Damager reduction applies to target # [🚫 can't damage module]", Type: RangedWeaponGearType},
	{Name: "Nerve Disruptor", Cost: 300, Mass: 5, Energy: Y, Notes: "2d6-4. If any damage gets through target drops everything in hands", Type: RangedWeaponGearType},
	{Name: "Particle Gun", Cost: 800, Mass: 5, Energy: Y, Notes: "1d6 + successes damage (max 5). Roll 1d6 when shot. \"6\" breaks module shooting from", Type: RangedWeaponGearType},
	{Name: "Plasma Pistol", Cost: 300, Mass: 6, Energy: Y, Notes: "1d6 Fire damage [☠ dangerous]", Type: RangedWeaponGearType},
	{Name: "Plasma Projector", Cost: 1100, Mass: 12, Energy: Y, Notes: "1d6 fire Area of Effect [☠ dangerous]", Type: RangedWeaponGearType},
	{Name: "Slug Pistol", Cost: 150, Mass: 5, Energy: N, Notes: "Select one type of ◇ ammo when you target", Type: RangedWeaponGearType},
	{Name: "Slug Machine Gun", Cost: 750, Mass: 15, Energy: N, Notes: "Slug gun area of effect. Select one type of ◇ ammo when you target (pg84) [☠ dangerous]", Type: RangedWeaponGearType},
	{Name: "Sonic Beam", Cost: 400, Mass: 4, Energy: Y, Notes: "Indirect, Area of Effect, 1d6 damage, Ignore DR [☠ dangerous] [🚫 can't damage module]", Type: RangedWeaponGearType},
	{Name: "Stun Gun", Cost: 200, Mass: 3, Energy: Y, Notes: "Athletics check of 11 or stunned for difference. [🚫 can't damage module]", Type: RangedWeaponGearType},
	{Name: "Voltrex", Cost: 675, Mass: 10, Energy: Y, Notes: "1d6, up to 4 targets, +1 difficulty per # of targets [☠ dangerous]", Type: RangedWeaponGearType},
}

var toxins = []Gear{
	{Name: "Death", Cost: 25, Mass: 1, Energy: N, Notes: "Target makes an Athletics check of 8 or drop to -6 hit points.", Type: ToxinGearType},
	{Name: "Goof", Cost: 25, Mass: 1, Energy: N, Notes: "Target must reroll the highest die of each active skill check.", Type: ToxinGearType},
	{Name: "Ionizer", Cost: 25, Mass: 1, Energy: N, Notes: "Raise the target's ionization level by 2.", Type: ToxinGearType},
	{Name: "Kayo", Cost: 25, Mass: 1, Energy: N, Notes: "Athletics check of 8 or target knocked out until damaged.", Type: ToxinGearType},
	{Name: "Nervo", Cost: 25, Mass: 1, Energy: N, Notes: "Target makes Athletics check of 11 or drops everything in hand.", Type: ToxinGearType},
	{Name: "Oucho", Cost: 25, Mass: 1, Energy: N, Notes: "Target suffers 1d6-1 damage at the end of each phase.", Type: ToxinGearType},
	{Name: "Paino", Cost: 25, Mass: 1, Energy: N, Notes: "Target takes +2 points per die from all dice-based damage sources.", Type: ToxinGearType},
	{Name: "Slowgo", Cost: 25, Mass: 1, Energy: N, Notes: "Target's move attribute is reduced to 1.", Type: ToxinGearType},
	{Name: "Stun", Cost: 25, Mass: 1, Energy: N, Notes: "Target must make an Athletics check of 11 or be stunned see page 85.", Type: ToxinGearType},
	{Name: "Suscepto", Cost: 25, Mass: 1, Energy: N, Notes: "Target must reroll the highest die of any otherwise successful Athletics check", Type: ToxinGearType},
}

var slugAmmo = []Gear{
	{Name: "Standard", Notes: "1d6+1", Type: SlugAmmunitionType},
	{Name: "AP", Notes: "1d6-1 ignore all damage reduction", Type: SlugAmmunitionType},
	{Name: "Scattershot", Notes: "1d6 (reduce Combat skill check difficulty by 3 but apply damage reduction and range penalties twice)", Type: SlugAmmunitionType},
}
