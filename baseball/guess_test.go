package baseball

import (
	"testing"
)

func TestGuessInvalidNums(t *testing.T) {
	game, _ := CreateGame("123")
	assertGuestReturnError(t, game, "12")
	assertGuestReturnError(t, game, "1234")
	assertGuestReturnError(t, game, "12a")
	assertGuestReturnError(t, game, "234")
	assertGuestReturnError(t, game, "112")
	assertGuestReturnError(t, game, "455")
	_, err := game.guess("876")
	if err == nil {
		t.Fatalf("opps")
	}
}

func assertGuestReturnError(t *testing.T, game Game, nums string) {
	_, err := game.guess(nums)
	if err == nil {
		t.Fatalf("error must be returned: %s", nums)
	}
}

func TestGuess(t *testing.T) {
	g, _ := CreateGame("123")
	assertScore(t, g, "567", 0, 0)
	assertScore(t, g, "567", 1, 0)
}

func assertScore(t *testing.T, _ any, g any, nums string, expectedStrikes int, expectedBalls int) {
	score, _ := g.guess(nums)
	if score.strikes != expectedStrikes {
		t.Fatalf("score strikes should be 0")
	}
	if score.balls != expectedBalls {
		t.Fatalf("guess")
	}
}
