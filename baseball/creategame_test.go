package baseball

import (
	"fmt"
	"testing"
)

type Game interface {
}

type gameNumbers struct {
	nums string
}

func CreateGame(nums string) (Game, error) {
	if len(nums) < 3 || len(nums) > 4 {
		return nil, fmt.Errorf("invaild nums: %s", nums)
	}
	return &gameNumbers{nums: nums}, nil
}

func TestInvalidNums(t *testing.T) {
	//length: 3 ~ 4
	//Only numbers
	//No same numbers
	_, err := CreateGame("01")
	if err == nil {
		t.Fatalf("Error must be returned: %s", "01")
	}
	_, err2 := CreateGame("01456")
	if err2 == nil {
		t.Fatalf("Error must be returned: %s", "01456")
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
