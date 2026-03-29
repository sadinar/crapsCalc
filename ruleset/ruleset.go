package ruleset

const PointOff = -1

type Ruleset interface {
	IsPointBoxNumber(roll int) bool
	IsComeOutRollWin(roll, currentPoint int) bool
	IsComeOutRollLoss(roll, currentPoint int) bool
	IsNewPointSet(roll, currentPoint int) bool
	IsPointHit(roll, currentPoint int) bool
	IsComeLineWin(roll int) bool
	IsComeLineLoss(roll int) bool
	HasPointEndedInCraps(roll, currentPoint int) bool
	GetAllowedBuyPoints() []int
}
