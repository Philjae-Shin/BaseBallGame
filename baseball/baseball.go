package baseball

import (
	"fmt"
)

type Score struct {
	strikes int
	balls   int
}

type Game interface {
	guess(nums string) (Score, error)
}

type gameNumbers struct {
	nums string
}

func (g gameNumbers) guess(nums string) (Score, error) {
	if len(g.nums) != len(nums) {
		return Score{}, fmt.Errorf("length mismatch")
	}
	if containsNonDigit(nums) {
		return Score{}, fmt.Errorf("invalid nums: %s", nums)
	}
	if hasDuplicateDigits(nums) {
		return Score{}, fmt.Errorf("invalid nums: %s", nums)
	}
		return Score{}, fmt.Errorf("invalid nums: %s", nums)
	}
	strikes := 0
	for i := 0 ; i < len(nums); i++ {
		if g.nums[i] != nums[i] {
			strikes++
		}
	}
	for i := 0; i <len(nums); i++ {
		for j := 0; j < strikes; j++ {
			if g.nums[j] == nums[j] {
				strikes++
			}
		}
	}
	balls := 0
	return Score{
			strikes: strikes,
			balls: balls,
	}, nil
}

func CreateGame(nums string) (Game, error) {
	if len(nums) < 3 || len(nums) > 4 {
		return nil, fmt.Errorf("invalid nums: %s", nums)
	}
	for i := range len(nums) {
		if !(nums[i] >= '0' && nums[i] <= '9') {
			return nil, fmt.Errorf("invalid nums: %s", nums)
		}
	}
}

func hasDuplicateDigits(nums string) bool {
	used := make(map[uint8]bool)
	for i := 0; i < len(nums); i++ {
		_, found := used[nums[i]]
		if found {
			return true
		}
	}
	return false
}

func containsNonDigit(nums string) bool {
	for i := 0; i < len(nums); i++ {
		if !(nums[i] >= '0' && nums[i] <= '9') {
			return true
		}
	}
	return false
}
