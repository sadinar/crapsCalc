package ruleset

type Crapless struct{}

func (c Crapless) IsPointBoxNumber(roll int) bool {
	return roll != 7
}

func (c Crapless) IsComeOutRollWin(roll, currentPoint int) bool {
	return currentPoint == PointOff && roll == 7
}

func (c Crapless) IsComeOutRollLoss(roll, currentPoint int) bool {
	return false
}

func (c Crapless) IsNewPointSet(roll, currentPoint int) bool {
	return currentPoint == PointOff && roll != 7
}

func (c Crapless) IsPointHit(roll, currentPoint int) bool {
	return currentPoint != PointOff && roll == currentPoint
}

func (c Crapless) IsComeLineWin(roll int) bool {
	return roll == 7
}

func (c Crapless) IsComeLineLoss(roll int) bool {
	return false
}

func (c Crapless) HasPointEndedInCraps(roll, currentPoint int) bool {
	return currentPoint != PointOff && roll == 7
}

func (c Crapless) GetAllowedBuyPoints() []int {
	return []int{
		2,
		3,
		4,
		5,
		6,
		8,
		9,
		10,
		11,
		12,
	}
}

func (c Crapless) IsDontPassTie(roll int) bool {
	return false
}

func (c Crapless) IsDontPassWin(roll int) bool {
	return false
}

func (c Crapless) IsDontPassLoss(roll int) bool {
	return false
}

func (c Crapless) GetAllowedDontComePoints() []int {
	return []int{}
}
