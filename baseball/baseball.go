package baseball

import "fmt"

type Game interface {
}

type gameNumbers struct {
	nums string
}

func CreateGame(nums string) (Game, error) {
	if len(nums) < 3 || len(nums) > 4 {
		return nil, fmt.Errorf("invaild nums: %s", nums)
	}
	for i := 0; i < len(nums); i++ {
		if !(nums[i] >= '0' && nums[i] <= '9') {
			return nil, fmt.Errorf("invalid nums: %s", nums)
		}
	}
	used := make(map[uint8]bool)
	for i := 0; i < len(nums); i++ {
		_, found := used[nums[i]]
		if found {
			return nil, fmt.Errorf("invalid nums: %s", nums)
		}
		used[nums[i]] = true
	}
	return &gameNumbers{nums: nums}, nil
}
