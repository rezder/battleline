package flag

import (
	"github.com/rezder/go-battleline/battbot/combi"
	"github.com/rezder/go-battleline/battleline/cards"
	"testing"
)

func TestWedgeOneMorale(t *testing.T) {
	flagCards := []int{68}
	handCards := []int{}
	deck := []int{9, 10, 1, 11, 22}
	formationSize := 3
	expRank := 1
	rank := calcMaxRankNewFlag(flagCards, handCards, createSetFrom(deck), formationSize)
	if rank != expRank {
		t.Errorf("Expect rank %v got %v\n", expRank, rank)
	}
}
func TestPhalanx(t *testing.T) {
	flagCards := []int{53, 33}
	handCards := []int{}
	deck := []int{56, 18, 43, 58, 52, 21, 57, 51, 59, 5, 22, 55, 54, 30}
	formationSize := 3
	expRank := 16
	rank := CalcMaxRank(flagCards, handCards, createSetFrom(deck), 10, formationSize == 4)
	if rank != expRank {
		t.Logf("Combination %v", combi.Combinations(formationSize)[expRank-1])
		t.Errorf("Expect rank %v got %v\n", expRank, rank)
	}
}
func TestPhalanxNewFlag(t *testing.T) {
	flagCards := []int{}
	handCards := []int{}
	deck := []int{19, 10, 31, 41, 52, 59, 49}
	formationSize := 3
	expRank := 10
	rank := calcMaxRankNewFlag(flagCards, handCards, createSetFrom(deck), formationSize)
	if rank != expRank {
		t.Logf("Combination %v", combi.Combinations(formationSize)[rank-1])
		t.Errorf("Expect rank %v got %v\n", expRank, rank)
	}
}
func TestBattalionOneMorale(t *testing.T) {
	flagCards := []int{68}
	handCards := []int{}
	deck := []int{4, 10, 31, 41, 52}
	formationSize := 3
	expRank := 24
	rank := calcMaxRankNewFlag(flagCards, handCards, createSetFrom(deck), formationSize)
	if rank != expRank {
		t.Logf("Combination %v", combi.Combinations(formationSize)[rank-1])
		t.Errorf("Expect rank %v got %v\n", expRank, rank)
	}
}
func TestBattalion3Morale(t *testing.T) {
	flagCards := []int{68, 69, 67}
	handCards := []int{}
	deck := []int{4, 10, 31, 41, 52}
	formationSize := 3
	expRank := 25
	rank := calcMaxRankNewFlag(flagCards, handCards, createSetFrom(deck), formationSize)
	if rank != expRank {
		t.Logf("Combination %v", combi.Combinations(formationSize)[rank-1])
		t.Errorf("Expect rank %v got %v\n", expRank, rank)
	}
}
func TestLineOneMorale(t *testing.T) {
	flagCards := []int{68}
	handCards := []int{}
	deck := []int{19, 10, 31, 41, 52}
	formationSize := 3
	expRank := 40
	rank := calcMaxRankNewFlag(flagCards, handCards, createSetFrom(deck), formationSize)
	if rank != expRank {
		t.Logf("Combination %v", combi.Combinations(formationSize)[rank-1])
		t.Errorf("Expect rank %v got %v\n", expRank, rank)
	}
}
func createSetFrom(troopixs []int) map[int]bool {
	drawSet := make(map[int]bool)
	for _, troopix := range troopixs {
		drawSet[troopix] = true
	}
	return drawSet
}
func createDrawSet() map[int]bool {
	drawSet := make(map[int]bool)
	for i := 1; i <= cards.NOTroop; i++ {
		drawSet[i] = true
	}
	return drawSet
}
func removeFromDeck(troopixs []int, deckSet map[int]bool) {
	for _, troopix := range troopixs {
		delete(deckSet, troopix)
	}
}
