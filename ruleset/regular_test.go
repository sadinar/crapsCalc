package ruleset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNewPointSet(t *testing.T) {
	rr := Regular{}
	assert.False(t, rr.IsNewPointSet(2, PointOff))
	assert.False(t, rr.IsNewPointSet(3, PointOff))
	assert.False(t, rr.IsNewPointSet(7, PointOff))
	assert.False(t, rr.IsNewPointSet(11, PointOff))
	assert.False(t, rr.IsNewPointSet(12, PointOff))

	assert.True(t, rr.IsNewPointSet(4, PointOff))
	assert.True(t, rr.IsNewPointSet(5, PointOff))
	assert.True(t, rr.IsNewPointSet(6, PointOff))
	assert.True(t, rr.IsNewPointSet(8, PointOff))
	assert.True(t, rr.IsNewPointSet(9, PointOff))
	assert.True(t, rr.IsNewPointSet(10, PointOff))

	assert.False(t, rr.IsNewPointSet(4, 6))
	assert.False(t, rr.IsNewPointSet(5, 6))
	assert.False(t, rr.IsNewPointSet(6, 6))
	assert.False(t, rr.IsNewPointSet(8, 6))
	assert.False(t, rr.IsNewPointSet(9, 6))
	assert.False(t, rr.IsNewPointSet(10, 6))
}

func TestIsPointHit(t *testing.T) {
	rr := Regular{}
	assert.False(t, rr.IsPointHit(2, PointOff))
	assert.False(t, rr.IsPointHit(3, PointOff))
	assert.False(t, rr.IsPointHit(4, PointOff))
	assert.False(t, rr.IsPointHit(5, PointOff))
	assert.False(t, rr.IsPointHit(6, PointOff))
	assert.False(t, rr.IsPointHit(7, PointOff))
	assert.False(t, rr.IsPointHit(8, PointOff))
	assert.False(t, rr.IsPointHit(9, PointOff))
	assert.False(t, rr.IsPointHit(10, PointOff))
	assert.False(t, rr.IsPointHit(11, PointOff))
	assert.False(t, rr.IsPointHit(12, PointOff))

	assert.False(t, rr.IsPointHit(2, 6))
	assert.False(t, rr.IsPointHit(3, 6))
	assert.False(t, rr.IsPointHit(4, 6))
	assert.False(t, rr.IsPointHit(5, 6))
	assert.False(t, rr.IsPointHit(7, 6))
	assert.False(t, rr.IsPointHit(8, 6))
	assert.False(t, rr.IsPointHit(9, 6))
	assert.False(t, rr.IsPointHit(10, 6))
	assert.False(t, rr.IsPointHit(11, 6))
	assert.False(t, rr.IsPointHit(12, 6))

	assert.True(t, rr.IsPointHit(6, 6))
}

func TestHasPointEndedInCraps(t *testing.T) {
	rr := Regular{}

	assert.False(t, rr.HasPointEndedInCraps(7, PointOff))
	assert.True(t, rr.HasPointEndedInCraps(7, 6))
}

func TestIsComeOutRollWin(t *testing.T) {
	rr := Regular{}
	assert.False(t, rr.IsComeOutRollWin(2, PointOff))
	assert.False(t, rr.IsComeOutRollWin(3, PointOff))
	assert.False(t, rr.IsComeOutRollWin(12, PointOff))
	assert.False(t, rr.IsComeOutRollWin(8, PointOff))

	assert.True(t, rr.IsComeOutRollWin(7, PointOff))
	assert.True(t, rr.IsComeOutRollWin(11, PointOff))
}

func TestIsComeOutRollLoss(t *testing.T) {
	rr := Regular{}
	assert.True(t, rr.IsComeOutRollLoss(2, PointOff))
	assert.True(t, rr.IsComeOutRollLoss(3, PointOff))
	assert.True(t, rr.IsComeOutRollLoss(12, PointOff))

	assert.False(t, rr.IsComeOutRollLoss(8, PointOff))
	assert.False(t, rr.IsComeOutRollLoss(7, PointOff))
	assert.False(t, rr.IsComeOutRollLoss(11, PointOff))
}

func TestIsPossiblePoint(t *testing.T) {
	rr := Regular{}
	assert.False(t, rr.IsPointBoxNumber(2))
	assert.False(t, rr.IsPointBoxNumber(3))
	assert.True(t, rr.IsPointBoxNumber(4))
	assert.True(t, rr.IsPointBoxNumber(5))
	assert.True(t, rr.IsPointBoxNumber(6))
	assert.True(t, rr.IsPointBoxNumber(8))
	assert.True(t, rr.IsPointBoxNumber(9))
	assert.True(t, rr.IsPointBoxNumber(10))
	assert.False(t, rr.IsPointBoxNumber(11))
	assert.False(t, rr.IsPointBoxNumber(7))
}

func TestIsComeLineWin(t *testing.T) {
	rr := Regular{}
	assert.True(t, rr.IsComeLineWin(7))
	assert.True(t, rr.IsComeLineWin(11))
	assert.False(t, rr.IsComeLineWin(12))
	assert.False(t, rr.IsComeLineWin(2))
	assert.False(t, rr.IsComeLineWin(8))
	assert.False(t, rr.IsComeLineWin(3))
}

func TestIsComeLineLoss(t *testing.T) {
	rr := Regular{}
	assert.False(t, rr.IsComeLineLoss(7))
	assert.False(t, rr.IsComeLineLoss(11))
	assert.False(t, rr.IsComeLineLoss(9))
	assert.True(t, rr.IsComeLineLoss(12))
	assert.True(t, rr.IsComeLineLoss(2))
	assert.True(t, rr.IsComeLineLoss(3))
}

func TestGetAllowedBuyPoints(t *testing.T) {
	rr := Regular{}
	assert.Equal(
		t,
		[]int{4, 5, 6, 8, 9, 10},
		rr.GetAllowedBuyPoints(),
	)
}
