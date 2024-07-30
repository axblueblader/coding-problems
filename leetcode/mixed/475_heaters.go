package leetcode

// https://leetcode.com/problems/heaters/

import (
	"fmt"
	"math"
	"sort"
)

// heater radius
// warmed if
// housePosition <= heaterBefore + radius
// housePosition >= heaterAfter - radius

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	fmt.Println(houses)
	fmt.Println(heaters)

	maxDist := 0
	if len(heaters) == 1 {
		for _, housePos := range houses {
			maxDist = max(maxDist, int(math.Abs(float64(housePos-heaters[0]))))
		}
		return maxDist
	}
	h := 0
	i := 0

	for i < len(heaters)-1 && h < len(houses) {
		// fmt.Println(i, h, "heater:", heaters[i], " house:", houses[h], maxDist)
		heaterPos := heaters[i]
		if houses[h] <= heaterPos {
			maxDist = max(maxDist, heaterPos-houses[h])
			h++
			continue
		}
		if heaterPos+maxDist >= houses[h] {
			h++
			continue
		}
		if houses[h]-heaterPos < heaters[i+1]-houses[h] {
			maxDist = houses[h] - heaterPos
			h++
			continue
		}
		i++
	}

	if h < len(houses) {
		maxDist = max(maxDist, houses[len(houses)-1]-heaters[i])
	}

	return maxDist
}
