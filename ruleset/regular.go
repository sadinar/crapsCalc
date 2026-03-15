package ruleset

type Regular struct {
	name string
}

func (r Regular) IsPossiblePoint(roll int) bool {
	if roll < 4 {
		return false
	}
	if roll == 7 {
		return false
	}
	return roll < 11
}

func (r Regular) IsComeOutRollWin(roll, currentPoint int) bool {
	return currentPoint == PointOff && (roll == 11 || roll == 7)
}

func (r Regular) IsComeOutRollLoss(roll, currentPoint int) bool {
	return currentPoint == PointOff && (roll == 2 || roll == 3 || roll == 12)
}

func (r Regular) IsNewPointSet(roll, currentPoint int) bool {
	return currentPoint == PointOff && (roll > 3 && roll != 7 && roll < 11)
}

func (r Regular) IsPointHit(roll, currentPoint int) bool {
	return currentPoint != PointOff && roll == currentPoint
}

func (r Regular) HasPointEndedInCraps(roll, currentPoint int) bool {
	return currentPoint != PointOff && roll == 7
}

func (r Regular) IsComeLineWin(roll int) bool {
	return roll == 11 || roll == 7
}

func (r Regular) IsComeLineLoss(roll int) bool {
	return roll == 2 || roll == 3 || roll == 12
}

func (r Regular) GetAllowedBuyPoints() []int {
	return []int{
		4,
		5,
		6,
		8,
		9,
		10,
	}
}
