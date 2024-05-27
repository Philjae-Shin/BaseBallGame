package baseball

import (
	"testing"
)

func TestGuessInvaildNums(t *testing.T) {
	game, _ := CreatGame("123")
	_, err := game.guess("12")
	if err == nil {
		t.Fatalf("error must be returned: %s", "12")
	}
}
