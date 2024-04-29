package baseball

import (
	"testing"
)

func TestInvalidNums(t *testing.T) {
	//Only numbers
	//No same numbers
	assertError(t, "01")
	assertError(t, "01456")
	assertError(t, "abc")
	assertError(t, "14z")
	assertError(t, "a87a4")
	assertError(t, "011")
	assertError(t, "4400")
}

func assertError(t *testing.T, nums string) {
	_, err := CreateGame(nums)
	if err == nil {
		t.Fatalf("error must be returned: %s", nums)
	}
}

func TestCreateGame(t *testing.T) {
	game, err := CreateGame("123")
	if err != nil || game == nil {
		t.Fatalf("Game must be returned")
	}
	game2, err := CreateGame("1234")
	if err != nil || game2 == nil {
		t.Fatalf("Game must be returned")
	}
}

//10:31
