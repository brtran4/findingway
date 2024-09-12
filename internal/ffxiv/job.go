package ffxiv

type Job int

const (
	GNB Job = iota
	PLD
	GLD
	DRK
	WAR
	MRD
	SCH
	ACN // Arcanist
	SGE
	AST
	WHM
	CNJ
	SAM
	DRG
	NIN
	MNK
	RPR
	VPR
	BRD
	MCH
	DNC
	BLM
	BLU
	SMN
	PCT
	RDM
	LNC
	PUG
	ROG
	THM
	ARC // Archer
	Unknown
)

func JobFromAbbreviation(abbreviation string) Job {
	switch abbreviation {
	case "GNB":
		return GNB
	case "PLD":
		return PLD
	case "GLD":
		return GLD
	case "DRK":
		return DRK
	case "WAR":
		return WAR
	case "MRD":
		return MRD
	case "SCH":
		return SCH
	case "ACN":
		return ACN
	case "SGE":
		return SGE
	case "AST":
		return AST
	case "WHM":
		return WHM
	case "CNJ":
		return CNJ
	case "SAM":
		return SAM
	case "DRG":
		return DRG
	case "NIN":
		return NIN
	case "MNK":
		return MNK
	case "RPR":
		return RPR
	case "VPR":
		return VPR
	case "BRD":
		return BRD
	case "MCH":
		return MCH
	case "DNC":
		return DNC
	case "BLM":
		return BLM
	case "BLU":
		return BLU
	case "SMN":
		return SMN
	case "PCT":
		return PCT
	case "RDM":
		return RDM
	case "LNC":
		return LNC
	case "PUG":
		return PUG
	case "ROG":
		return ROG
	case "THM":
		return THM
	case "ARC":
		return ARC
	}
	return Unknown
}

func (j Job) Emoji() string {
	switch j {
	case GNB:
		return "<:GBR:1283573154131738685>"
	case PLD:
		return "<:PLD:1283573164789727302>"
	case GLD:
		return "<:gld:1277675881334046761>"
	case DRK:
		return "<:DRK:1283573142341685373>"
	case WAR:
		return "<:WAR:1283573176588177460>"
	case MRD:
		return "<:mrd:1277675899629600768>"
	case SCH:
		return "<:SCH:1283573266648137728>"
	case ACN:
		return "<:acn:1277675927374925824>"
	case SGE:
		return "<:SGE:1283573243474612234>"
	case AST:
		return "<:AST:1283573232527736884>"
	case WHM:
		return "<:WHM:1283573282250948648>"
	case CNJ:
		return "<:cnj:1277675946358345788>"
	case SAM:
		return "<:SAM:1283573764336124058>"
	case DRG:
		return "<:DRG:1283573524824592416>"
	case NIN:
		return "<:NIN:1283573727648415774>"
	case MNK:
		return "<:MNK:1283573715254382612>"
	case RPR:
		return "<:RPR:1283573738134311014>"
	case VPR:
		return "<:VPR:1283577676120064050>"
	case BRD:
		return "<:BRD:1283573481048510575>"
	case MCH:
		return "<:MCH:1283573699290726460>"
	case DNC:
		return "<:DNC:1283573511067140170>"
	case BLM:
		return "<:BLM:1283573494973464607>"
	case BLU:
		return "<:BLU:1283576210257084498>"
	case SMN:
		return "<:SMN:1283573786163281992>"
	case RDM:
		return "<:RDM:1283573751358689301>"
	case PCT:
		return "<:PCT:1283577649972908032>"
	case LNC:
		return "<:lnc:1277675968701661217>"
	case PUG:
		return "<:pgl:1277675984862052392>"
	case ROG:
		return "<:rog:1277676011386962064>"
	case THM:
		return "<:thm:1277676028113719296>"
	case ARC:
		return "<:arc:1277676046992408636>"
	}
	return "<:FSH:1283578184025116672>"
}
