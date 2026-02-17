package fn

import "slices"

// MaxConsecutive checks that consecutive slot lengths do not exceed the given limit
func MaxConsecutive(slots []int, limit int) bool {
	if len(slots) == 0 {
		return true
	}
	slices.Sort(slots)
	group := []int{slots[0]}
	prev := slots[0]
	for _, slot := range slots[1:] {
		if prev+1 == slot {
			// Consecutive: add slot to group
			group = append(group, slot)
		} else {
			// Not consecutive: check group size and reset group
			if len(group) > limit {
				return false
			}
			group = []int{slot}
		}
		prev = slot
	}
	// Check last group
	return len(group) <= limit
}
