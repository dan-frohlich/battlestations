package model

import (
	"strings"
	"sync"
)

type AbilityType int

func (gt AbilityType) String() string {
	switch gt {
	case AthleticAbilityType:
		return "athletic"
	case CombatAbilityType:
		return "combat"
	case EngineerAbilityType:
		return "engineering"
	case PilotAbilityType:
		return "pilot"
	case PsychicAbilityType:
		return "psychic"
	case ScienceAbilityType:
		return "science"
	default:
		return ""
	}
}

const (
	AthleticAbilityType = iota
	CombatAbilityType
	DiplomacyAbilityType
	EngineerAbilityType
	PilotAbilityType
	PsychicAbilityType
	ScienceAbilityType
)

type SpecialAbility struct {
	Name            string
	Summary         string
	OutputSummary   string
	FullDescription string
	Types           []AbilityType
	// poolFunc        func(c *Character)
}

var (
	abilityIndex map[string]SpecialAbility
	indexSync    = &sync.Once{}
)

func indexAbility() {
	indexSync.Do(func() {
		abilityIndex := make(map[string]SpecialAbility)
		for _, a := range abilities {
			abilityIndex[a.Name] = a
		}
	})
}

func GetAbility(name string) SpecialAbility {
	indexAbility()
	if g, ok := abilityIndex[strings.TrimSpace(name)]; ok {
		return g
	}
	return SpecialAbility{}
}

var abilities = []SpecialAbility{
	{Name: "Acrobatic", Summary: "Pop 2 squares. Add OOC for movement. +3 target number when in motion. Enemies don't get free attack", Types: []AbilityType{AthleticAbilityType}},
	{Name: "Adaptable", Summary: "Spend a luck and an action to have another ability instead", Types: []AbilityType{}},
	{Name: "Adv. Combat Expert", Summary: "+1 target number, +1 Hit point, opponents must reroll highest die in attacks against you", Types: []AbilityType{}},
	{Name: "Assembly Line Worker", Summary: "Each Upgrade you make can be applied to three like objects", Types: []AbilityType{}},
	{Name: "Assistant", Summary: "May add +3 to the difficulty skill check for each additional bonus assisting", Types: []AbilityType{}},
	{Name: "Battle Frenzied", Summary: "Bonus action every Phase but may only melee, grapple or move. Must spend action to calm", Types: []AbilityType{}},
	{Name: "Blaze of Glory", Summary: "1/ Phase may force survival checks for all except your fighter in hex with your fighter", Types: []AbilityType{}},
	{Name: "Bloodlusted", Summary: "Bonus action when you put somebody down. +1 diffiuclty penalty on action per use", Types: []AbilityType{}},
	{Name: "Bosun", Summary: "Remote penalty 1 lower. Reroll pool of 5 for battlestation actions", Types: []AbilityType{}},
	{Name: "Bot Genius", Summary: "2 bonus bot upgrades. Reroll die when attacking, grappling, upgrading, repairing or damaging bots", Types: []AbilityType{}},
	{Name: "Braced", Summary: "May reroll one die any time you sustain damage", Types: []AbilityType{}},
	{Name: "Brutal", Summary: "You may reroll one of the dice each time you deal personal damage with a direct attack", Types: []AbilityType{}},
	{Name: "Calm", Summary: "May take an “8” on a skill check instead of (and before) rolling the dice. Pool of Rank x2", Types: []AbilityType{}},
	{Name: "Cannon Expert", Summary: "Free reroll to attack with, repair, or reconfigure cannon. Reconfigure a cannon as a free, automatic action", Types: []AbilityType{}},
	{Name: "Cannon Specialist", Summary: "Pool to reroll attacks, reconfigure, or repair a cannon module. Pool of Combat x2", Types: []AbilityType{}},
	{Name: "Charger", Summary: "Free unarmed (or natural weapon) melee attack at the end of your move action", Types: []AbilityType{}},
	{Name: "Coach", Summary: "You may use Diplomacy instead of the listed skill for your assist actions", Types: []AbilityType{}},
	{Name: "Connected", Summary: "You get an additional requisition and double your pay for each mission", Types: []AbilityType{}},
	{Name: "Contortionist", Summary: "Move & act in slagged squares as though not slagged. +1 target number except for grappling", Types: []AbilityType{}},
	{Name: "Cortex Overloader", Summary: "Pool to blow up people's heads Psionically. Pool of Combat + Psionics", Types: []AbilityType{}},
	{Name: "Courtier", Summary: "Move & act in slagged squares as though not slagged. Others may spend your luck", Types: []AbilityType{}},
	{Name: "Crafty", Summary: "Checkup, Focus Sensors, Tune Shields. Pool of Science", Types: []AbilityType{}},
	{Name: "Cutthroat", Summary: "Spend from this pool to reroll a damage die you deal directly in personal combat. Pool of Science x2", Types: []AbilityType{}},
	{Name: "Death Marcher", Summary: "+3 Hit Points. May take simple move actions when unconscious. \"Recover\" immediately", Types: []AbilityType{}},
	{Name: "Death Striker", Summary: "Unarmed / natural weapon attacks force target to Athletics check 8 or be brought instantly to -6 Hit Points", Types: []AbilityType{}},
	{Name: "Dervish", Summary: "As an action, melee attack on all adjacent targets", Types: []AbilityType{}},
	{Name: "Destroyer", Summary: "Free action to release a Psionic energy blast like satchel charge in your square. Pool of Athletics + Psionics", Types: []AbilityType{}},
	{Name: "Dirty Fighter", Summary: "Targets you hit are +3 on all active skill check difficulty. Coup de grace without preparing", Types: []AbilityType{}},
	{Name: "Displaced", Summary: "Psionically shimmer to add +3 to your target number until end of phase. Pool of Athletics + Psionics", Types: []AbilityType{}},
	{Name: "Doctor", Summary: "Pool to reroll skill checks to heal, detoxify, etc. or on the healing dice. Pool of Science x2", Types: []AbilityType{}},
	{Name: "Dogfighter", Summary: "Add Piloting to Combat when shooting Fighter's gun.+1 difficulty to target your fighter. -1 difficulty survival", Types: []AbilityType{}},
	{Name: "Empowered", Summary: "Spend Ship's power as though it were luck. Reallocate power. Pool of Engineering + Psionics", Types: []AbilityType{}},
	{Name: "EMT", Summary: "Pre-emptively heal incoming damage. Pool of Combat x2", Types: []AbilityType{}},
	{Name: "Energy Deflector", Summary: "Combat 11 to negate direct attack. Difficulty goes up by 1 for each use", Types: []AbilityType{}},
	{Name: "Engine Overloader", Summary: "Pool to make Engine pump +2 per extra power but satchel charge detonates. Pool of Engineering x2", Types: []AbilityType{}},
	{Name: "Engine Specialist", Summary: "Pool of rerolls to use when pumping, transferring power, or repairing an engine. Pool of Engineering x2", Types: []AbilityType{}},
	{Name: "Enraged", Summary: "+3 Combat skill until end of Round when injured by an enemy", Types: []AbilityType{}},
	{Name: "Fast Healer", Summary: "Any skill check to heal or treat you is at a bonus of +3 to the Skill check. Heal +1 point per die", Types: []AbilityType{}},
	{Name: "Fast Learner", Summary: "Increase your Experience awards by 10%. Perform all 4 skills to get a bonus", Types: []AbilityType{}},
	{Name: "Fated", Summary: "Choose the result of the first and last points of Luck instead of rerolling", Types: []AbilityType{}},
	{Name: "Field Surgeon", Summary: "Reduce healing difficulty by 3 and get a reroll but target must make an Athletics check of 8 or dies", Types: []AbilityType{}},
	{Name: "Fighter Jock", Summary: "Board, Launch, and disembark from fighter as a free action. Reroll survival checks in microships", Types: []AbilityType{}},
	{Name: "Fighter Mechanic", Summary: "2 free upgrade actions on fighters and a reroll when repairing, upgrading, or survival in a fighter", Types: []AbilityType{}},
	{Name: "Fire Starter", Summary: "Start fires with your mind. Pool of Combat + Psionics", Types: []AbilityType{}},
	{Name: "First Mate", Summary: "May let others spend your luck as if it were their own. Reroll on all attempts to assist. +1 Luck", Types: []AbilityType{}},
	{Name: "Florentine Fighter", Summary: "Reduce your penalty to attack with two personal weapons by 3", Types: []AbilityType{}},
	{Name: "Foresighted", Summary: "Start retroactively on overwatch. You can go generically on overwatch", Types: []AbilityType{}},
	{Name: "Forethinker", Summary: "Roll your skill check before declaring an action", Types: []AbilityType{}},
	{Name: "Fortunate", Summary: "You may spend 1 Luck to nudge a luckable die upwards by one instead of rerolling", Types: []AbilityType{}},
	{Name: "Fume-Runner", Summary: "Take actions that require power without it (causing 1 point of hull damage)", Types: []AbilityType{}},
	{Name: "Ghost in The Machine", Summary: "Operate a Battlestation from anywhere aboard ship at no remote penalty. Pool of Engineering + Psionics", Types: []AbilityType{}},
	{Name: "Grease Monkey", Summary: "Reroll on any skill check to repair, reconfigure or upgrade. Move between Battlestations", Types: []AbilityType{}},
	{Name: "Grenadier", Summary: "Rerolls with grenades, Free actions to arm, draw, or detonate. Explosives weigh half for you", Types: []AbilityType{}},
	{Name: "Gunner's Mate", Summary: "Fighter's guns can be fired 2x/phase. Damages your fighter and the occupants", Types: []AbilityType{}},
	{Name: "Hacker", Summary: "Hack as a free action 1/phase with difficulty -3 and a reroll. Your cyberware behaves as upgraded", Types: []AbilityType{}},
	{Name: "Hardened", Summary: "Pool to reroll a die that deals damage to you. Pool of 5", Types: []AbilityType{}},
	{Name: "Healer", Summary: "Whenever you heal anybody for at least one die, heal them for an additional die", Types: []AbilityType{}},
	{Name: "Hot Dog", Summary: "Doubles succeeds on your piloting checks but causes OOC as if it had failed Hull Stress", Types: []AbilityType{}},
	{Name: "Empath", Summary: "You may suffer to reroll hull damage dice your ship is taking. Pool of Engineering + Psionics", Types: []AbilityType{}},
	{Name: "Hunch Follower", Summary: "+2 luck and Reduce the difficulty to use the Science bay by 1 for each time you've asked questions", Types: []AbilityType{}},
	{Name: "Hyper-Do UV Belt", Summary: "Reroll Unarmed Attacks, Damage, Grappling Checks. Move into Enemy squares with no penalty", Types: []AbilityType{}},
	{Name: "Hyper-Physicist", Summary: "Pool to reroll when using the Hyperdrive and facing on warp in. Science x2", Types: []AbilityType{}},
	{Name: "Intuitive", Summary: "Add Psionics skill to scan or ask questions. Pool to ask yes/no or enemy's action. Pool of Science + Psionics", Types: []AbilityType{}},
	{Name: "Jack of All Trades", Summary: "You get 2 rerolls per campaign turn in each skill", Types: []AbilityType{}},
	{Name: "Jet-Pack Jockey", Summary: "You get a reroll on your JetPack Piloting skill checks and reduce the difficulty by 3", Types: []AbilityType{}},
	{Name: "Jury Rigger", Summary: "Once per phase, you may take a free action to attempt to repair the module you are in", Types: []AbilityType{}},
	{Name: "Killer Instinct", Summary: "+1 difficulty on passive checks you cause. No prepare before Coup de Grace. May spend Luck on damage", Types: []AbilityType{}},
	{Name: "Lucky", Summary: "Add +3 to your Luck", Types: []AbilityType{}},
	{Name: "Mechanical Empath", Summary: "Repair remotely as a free action. Pool of Engineering + Psionics", Types: []AbilityType{}},
	{Name: "Mentally Shielded", Summary: "Retard personal energy damage. Pool of Athletics + Psionics", Types: []AbilityType{}},
	{Name: "Mind Mender", Summary: "Absrob others damage onto yourself. Pool of Athletics + Psionics", Types: []AbilityType{}},
	{Name: "Miracle Worker", Summary: "Choose a die roll instead of rolling it once per campaign turn. Also +1 Luck", Types: []AbilityType{}},
	{Name: "Mobile", Summary: "+2 Move. Ignore OOC for movement. Reroll on any attempt to move extra squares", Types: []AbilityType{}},
	{Name: "Mr. Fixit", Summary: "Pool to reroll repairs and upgrades or repair as a free action. Pool of Engineering + Science", Types: []AbilityType{}},
	{Name: "Multi-Shot Expert", Summary: "Fire multibarrel twice as a single action. Reconfigure cannons as an automatic action", Types: []AbilityType{}},
	{Name: "Nimble", Summary: "Take bonus actions at +3 difficulty max once per phase. Pool of Athletics", Types: []AbilityType{}},
	{Name: "Noble", Summary: "You get an allowance. Use others luck. Difficulty to assist you is 1 easier and 1 more effective", Types: []AbilityType{}},
	{Name: "Numb Runner", Summary: "Ignore toxins and drug side effects. Dose as free action. Your toxins hard to resist. Pool of Athletics + Rank", Types: []AbilityType{}},
	{Name: "Obsessive", Summary: "Spending a second or subsequent Luck on a reroll gives you 3 rerolls each instead of 1", Types: []AbilityType{}},
	{Name: "Omniscient", Summary: "Pool to ask ANY questions. Add Psionics when hacking or gathering data. Pool of Science + Psionics", Types: []AbilityType{}},
	{Name: "Overloader", Summary: "Pool to fire with bonus guns power but satchel charge in your square. Pool of Engineering x2", Types: []AbilityType{}},
	{Name: "Pack Mule", Summary: "Double carry capacity. -3 difficulty and reroll to quickdraw. Reduce Penalty to act after quickdrawing by 1", Types: []AbilityType{}},
	{Name: "Patient", Summary: "Preparing reduces difficulty by 3. May convert prepare to Overwatch. Moving doesn't disrupt prepare", Types: []AbilityType{}},
	{Name: "Persevering", Summary: "When you fail Skill Checks, reduce difficulty by 3 and get a reroll on your next action if it uses that skill", Types: []AbilityType{}},
	{Name: "Plasma Wizard", Summary: "Fire weapons aren't dangerous to you. You may reroll skill checks and damage with fire weapons", Types: []AbilityType{}},
	{Name: "Polarizer", Summary: "Pool to EMP or De-EMP objects in L.O.S. Pool of Science + Psionics", Types: []AbilityType{}},
	{Name: "Power Slider", Summary: "You may combine different kinds of Helm maneuvers", Types: []AbilityType{}},
	{Name: "Powered Armor Expert", Summary: "Reduce PA penalties by 1. Add +1 Piloting for jetting. May move in phase you equip. Power up automatic", Types: []AbilityType{}},
	{Name: "Preconceived", Summary: "Spend 2 luck to select the result of a die that you would luck instead of rolling the die. (special pool)", Types: []AbilityType{}},
	{Name: "Prestidigitator", Summary: "Extradimensional pockets. Reduced quickdraw penalties. Add Psionics for Quickdraw checks", Types: []AbilityType{}},
	{Name: "Psychic Blaster", Summary: "Pool to deal 1d6+successes as damage ignoring damage reduction. Pool of Combat + Psionics", Types: []AbilityType{}},
	{Name: "Psychic Shifter", Summary: "Pool to move yourself to another module with your mind. Pool of Piloting + Psionics", Types: []AbilityType{}},
	{Name: "Psychic Stunner", Summary: "Pool to add a stun effect to a direct attack. Pool of Athletics + Psionics", Types: []AbilityType{}},
	{Name: "Puppeteer", Summary: "Pool to select target's next action. Pool of Diplomacy + Psionics", Types: []AbilityType{}},
	{Name: "Quartermaster", Summary: "Assist in upgrades and requisitions. Don’t pay for upgrades. Repair personal equipment as automatic action", Types: []AbilityType{}},
	{Name: "Quick on the Stick", Summary: "Shooting the Fighter's cannon is a free action for you. (It still uses the Fighter's cannon for the phase)", Types: []AbilityType{}},
	{Name: "Quick-Minded", Summary: "Spend 2 from a Psionic Ability pool to perform that Psionic action as a free action", Types: []AbilityType{}},
	{Name: "Reckless", Summary: "Add 1 die of damage to a personal attack 1/phase but suffer the lowest die without reduction if it is odd", Types: []AbilityType{}},
	{Name: "Reflexive", Summary: "You may dodge or ram once per phase during ship movement or missile resolution as a free action", Types: []AbilityType{}},
	{Name: "Research Spec.", Summary: "Count 1/2 used markers on Science Bay. Pool to reroll any use of the Science Bay. Science x2", Types: []AbilityType{}},
	{Name: "Resourceful", Summary: "Sub Science or Engineering for any other skill. Pool of Engineering or Science (whichever is lower)", Types: []AbilityType{}},
	{Name: "Rolls With It", Summary: "Pool to count OOC as a bonus insteaAbilityd of a penalty to your action or move. Pool of Piloting", Types: []AbilityType{}},
	{Name: "Saboteur", Summary: "Free action to attempt to break something. Reroll all dice you would deal to equipment", Types: []AbilityType{}},
	{Name: "Seer", Summary: "Line of sight to anywhere aboard the ship or get scans as a free action. Pool of Science + Psionics", Types: []AbilityType{}},
	{Name: "Sharpshooter", Summary: "Reroll allocation with ship's weapons or attacks with direct personal ranged attacks. Pool of Combat x2", Types: []AbilityType{}},
	{Name: "Shield Harmonizer", Summary: "Target's shields become 1 lower, yours are 1 higher", Types: []AbilityType{}},
	{Name: "Shock Trooper", Summary: "You may go on overwatch as a free action at the end of your move action. Pool of Combat x2", Types: []AbilityType{}},
	{Name: "Slipster", Summary: "Move through walls during movement. Pool of Athletics + Psionics", Types: []AbilityType{}},
	{Name: "Smooth", Summary: "Take actions that don't require skill check or include movement. Pool of Athletics + Rank", Types: []AbilityType{}},
	{Name: "Smuggler", Summary: "Move and act in slagged square at no penalty. Free reroll to use Cargo Bay Equipment", Types: []AbilityType{}},
	{Name: "Sniper", Summary: "Count range as closer. Reroll damage die if attack exceeded target by 3", Types: []AbilityType{}},
	{Name: "Spacelegs", Summary: "Ignore OOC", Types: []AbilityType{}},
	{Name: "Spacer", Summary: "5 rerolls for use in any skill check. Pool of 5", Types: []AbilityType{}},
	{Name: "Speed Demon", Summary: "Reroll acceleration and Fighter movement checks. Pool of Piloting x2", Types: []AbilityType{}},
	{Name: "Steady Handed", Summary: "Maneuvers always generate half OOC (round up). Steady the ship as an automatic action", Types: []AbilityType{}},
	{Name: "Stunner", Summary: "Unarmed and Natural weapon attacks generate a stun effect", Types: []AbilityType{}},
	{Name: "Sure Handed", Summary: "Ignore the “dangerous” effects. Reroll the damage die vs friendlies, quickdraw, repair. +1 Luck", Types: []AbilityType{}},
	{Name: "Swashbuckler", Summary: "You may take your actions during your move. Pop without penalty", Types: []AbilityType{}},
	{Name: "Sympathetic", Summary: "Your attacker also suffers the highest die of damage. Pool of Athletics + Psionics", Types: []AbilityType{}},
	{Name: "Tailgunner", Summary: "Reroll Combat checks when firing a fighter's gun. Go on overwatch with Fighter. Pool of Combat + Piloting", Types: []AbilityType{}},
	{Name: "Telekinetic", Summary: "Pool to move stuff with your mind. Pool of Athletics + Psionics", Types: []AbilityType{}},
	{Name: "Teleporter Specialist", Summary: "Count 1/2 used markers on Teleporter. Reroll pool for skill checks, allocation. Pool of Science x2", Types: []AbilityType{}},
	{Name: "Tinkerer", Summary: "Additional upgrade attempt. Reroll on all upgrade or repair attempts", Types: []AbilityType{}},
	{Name: "Tough", Summary: "Athletics counts as 1 higher and free reroll on any Athletics check", Types: []AbilityType{}},
	{Name: "Tough Silicoid", Summary: "+1 hit point. When you roll your Silicoid damage reduction, roll an extra die and count the higher one", Types: []AbilityType{}},
	{Name: "Trampler", Summary: "Free melee attack as you move through enemies", Types: []AbilityType{}},
	{Name: "Triage Medic", Summary: "Heal additional targets at +1 difficulty each", Types: []AbilityType{}},
	{Name: "Trick Shot", Summary: "Ignore Peeking, Popping, OOC, Shields, Cover for ranged attacks. Bank shots around corners", Types: []AbilityType{}},
	{Name: "Tricky", Summary: "Pool to distract microships with trash, focus sensors, or steady ship. Pool of Engineering", Types: []AbilityType{}},
	{Name: "Turn Specialist", Summary: "Pool to reroll maneuvers to turn, dodge, ram, or sideslip. Pool of Piloting x2", Types: []AbilityType{}},
	{Name: "Unconventional", Summary: "Pool to roll an additional die alongside your skill check. Odd it adds. Even it subtracts. Pool of Science", Types: []AbilityType{}},
	{Name: "Unflappable", Summary: "Pool to ignore a total of up to +3 in penalties each time you use this ability. Pool of Rank", Types: []AbilityType{}},
	{Name: "Uniminded", Summary: "You may use an ally's skill instead of your own as long as that ally is within Line Of Sight. Pool of 5", Types: []AbilityType{}},
	{Name: "Unlimited", Summary: "Once per campaign turn, you may reset up to two of your special abilities to their starting values", Types: []AbilityType{}},
	{Name: "Unpredictable", Summary: "Random skill each Phase gives you -1 difficulty with a free reroll", Types: []AbilityType{}},
	{Name: "Unsinkable", Summary: "Pool to reroll passive checks such as ship’s hull check, disintegration, stun, death, etc.. Pool of 10", Types: []AbilityType{}},
	{Name: "Wake Rider", Summary: "Explosions move your microship instead of damaging it. Also may move along with a ship or fighter", Types: []AbilityType{}},
	{Name: "Weapons Officer", Summary: "Reroll hit allocation and damage die in each shot you take with a fighter or cannon", Types: []AbilityType{}},
	{Name: "Wild Flyer", Summary: "Spend a Helm power, reroll a die, and apply an additional OOC (that takes effect after the maneuver)", Types: []AbilityType{}},
	{Name: "Wingman", Summary: "Reduce difficulty for survival checks or shots by friendlies", Types: []AbilityType{}},
	{Name: "Wrestler", Summary: "Choose to apply up to three different effects. Also a pool of rerolls to use in grapples. Pool of Athletics x2", Types: []AbilityType{}},
	{Name: "Xenobiologist", Summary: "Your direct attacks ignore alien damage reduction abilities", Types: []AbilityType{}},
	{Name: "Zone Controller", Summary: "All adjacent squares are considered occupied by you (slagged) for enemies", Types: []AbilityType{}},
}
