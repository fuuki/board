package board

type BoardProfile interface {
}

type BoardProfileDefinition[BP BoardProfile] interface {
	New() BP
	Clone(BP) BP
}
