package algo

import "fmt"

func BinarySearch(slice []int64, target int64) (int64, error) {
	highest := len(slice) - 1
	lowest := 0

	for lowest <= highest {
		middle := (lowest + highest) / 2
		guess := slice[middle]
		if guess == target {
			return target, nil
		}
		if guess > target {
			highest = middle - 1
		} else {
			lowest = middle + 1
		}
	}

	return 0, fmt.Errorf("not found")
}
