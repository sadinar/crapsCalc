package ruleset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPossibleCraplessPoint(t *testing.T) {
	ct := Crapless{}
	assert.True(t, ct.IsPointBoxNumber(2))
	assert.True(t, ct.IsPointBoxNumber(3))
	assert.True(t, ct.IsPointBoxNumber(8))
	assert.True(t, ct.IsPointBoxNumber(11))
	assert.True(t, ct.IsPointBoxNumber(12))
	assert.False(t, ct.IsPointBoxNumber(7))
}

func TestIsCraplessComeOutWin(t *testing.T) {
	ct := Crapless{}
	assert.True(t, ct.IsComeOutRollWin(7, PointOff))
	assert.False(t, ct.IsComeOutRollWin(7, 6))
	assert.False(t, ct.IsComeOutRollWin(2, PointOff))
}

func TestIsCraplessComeOutLoss(t *testing.T) {
	ct := Crapless{}
	assert.False(t, ct.IsComeOutRollLoss(7, PointOff))
	assert.False(t, ct.IsComeOutRollLoss(7, 6))
	assert.False(t, ct.IsComeOutRollLoss(2, PointOff))
	assert.False(t, ct.IsComeOutRollLoss(-1, -1))
}

func TestIsCraplessNewPointSet(t *testing.T) {
	ct := Crapless{}
	assert.False(t, ct.IsNewPointSet(7, PointOff))
	assert.False(t, ct.IsNewPointSet(7, 6))
	assert.True(t, ct.IsNewPointSet(2, PointOff))
	assert.True(t, ct.IsNewPointSet(3, PointOff))
	assert.True(t, ct.IsNewPointSet(11, PointOff))
	assert.True(t, ct.IsNewPointSet(12, PointOff))
	assert.True(t, ct.IsNewPointSet(8, PointOff))
}

func TestIsCraplessPointHit(t *testing.T) {
	ct := Crapless{}
	assert.False(t, ct.IsPointHit(7, PointOff))
	assert.False(t, ct.IsPointHit(7, 3))
	assert.True(t, ct.IsPointHit(2, 2))
	assert.True(t, ct.IsPointHit(3, 3))
	assert.True(t, ct.IsPointHit(11, 11))
	assert.True(t, ct.IsPointHit(12, 12))
	assert.True(t, ct.IsPointHit(8, 8))
}

func TestIsCraplessComeLineWin(t *testing.T) {
	ct := Crapless{}
	assert.True(t, ct.IsComeLineWin(7))
	assert.False(t, ct.IsComeLineWin(11))
	assert.False(t, ct.IsComeLineWin(6))
	assert.False(t, ct.IsComeLineWin(2))
}

func TestIsCraplessComeLineLoss(t *testing.T) {
	ct := Crapless{}
	assert.False(t, ct.IsComeLineLoss(7))
	assert.False(t, ct.IsComeLineLoss(6))
	assert.False(t, ct.IsComeLineLoss(2))
	assert.False(t, ct.IsComeLineLoss(3))
	assert.False(t, ct.IsComeLineLoss(11))
	assert.False(t, ct.IsComeLineLoss(12))
}

func TestCraplessHasPointEnded(t *testing.T) {
	ct := Crapless{}
	assert.False(t, ct.HasPointEndedInCraps(7, PointOff))
	assert.True(t, ct.HasPointEndedInCraps(7, 2))
	assert.True(t, ct.HasPointEndedInCraps(7, 3))
	assert.True(t, ct.HasPointEndedInCraps(7, 11))
	assert.True(t, ct.HasPointEndedInCraps(7, 12))
	assert.True(t, ct.HasPointEndedInCraps(7, 5))
}

func TestGetCraplessAllowedBuyPoints(t *testing.T) {
	ct := Crapless{}
	assert.Equal(
		t,
		[]int{2, 3, 4, 5, 6, 8, 9, 10, 11, 12},
		ct.GetAllowedBuyPoints(),
	)
}
