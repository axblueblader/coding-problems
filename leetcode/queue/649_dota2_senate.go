package leetcode

// https://leetcode.com/problems/dota2-senate/description/

// the best strategy is to ban next opposition until there's a clear winner
func predictPartyVictory(senate string) string {
	rVoters := 0
	dVoters := 0
	queue := []rune(senate)
	for _, ch := range queue {
		if ch == 'R' {
			rVoters++
		} else {
			dVoters++
		}
	}

	// right to band
	rBans := 0 // Radiant can ban X number of dires next
	dBans := 0
	for rVoters > 0 && dVoters > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == 'R' {
			if dBans > 0 {
				// i am banned
				rVoters--
				dBans--
				continue
			}
			// i can ban the next opposition voters
			rBans++
		}
		if v == 'D' {
			if rBans > 0 {
				dVoters--
				rBans--
				continue
			}
			dBans++
		}
		// i'm not ban so i can continue to vote
		// move me to last of queue, meaning i'm going to next round
		queue = append(queue, v)
	}

	if dVoters == 0 {
		return "Radiant"
	}

	return "Dire"
}
