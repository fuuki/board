package board

const (
	samplePhaseNamePlay PhaseName = "play"
)

func newSampleBP(totalPlayer uint) *sBP {
	p := &sBP{
		points: make([]int, totalPlayer),
	}
	return p
}

type sBP struct {
	points []int
	turn   int
}
type sADP struct {
}
type sGame = Game[*sBP, *sADP]
type sPhase = Phase[*sBP, *sADP]
type sAP = ActionProfile[*sADP]
type sAR = ActionRequest[*sADP]

// samplePhasePlay returns a phase of rock-paper-scissors.
func samplePhasePlay() *sPhase {
	prepare := func(g *sGame) *sAR {
		apr := newSampleAR(g.TotalPlayer())
		g.BoardProfile().turn++
		return apr
	}
	execute := func(g *sGame, bp *sBP, ap *sAP) (PhaseName, *sBP) {
		if g.BoardProfile().turn > 3 {
			return "", bp
		}
		return samplePhaseNamePlay, bp
	}
	p := NewPhase(samplePhaseNamePlay, prepare, execute)
	return p
}

func newSampleAR(totalPlayer uint) *sAR {
	r := NewActionRequest[*sADP](totalPlayer)
	for i := 0; i < int(totalPlayer); i++ {
		r.RegisterValidator(Player(i), func(a *sADP) error { return nil })
	}
	return r
}
