package leetcode

// https://leetcode.com/problems/asteroid-collision/description/

// the main idea is to traverse from the left
// if it's an asteroids moving to the right (value > 0), then we add it to the answer stack
// when we see an opposite direction, we need to continue popping the stack until all colision is resolved
// if opposite > rightmost asteroids => it's destroyed, pop the stack
// if opposite == current => both are destroyed, pop the stack
// if current is < 0, we can stop and add to stack
// if by the end, there is no greater, add to stack
func asteroidCollision(asteroids []int) []int {
	answer := []int{}

	for _, val := range asteroids {
		// fmt.Println(answer)
		if val > 0 {
			answer = append(answer, val)
		} else {
			collided := false
			for len(answer) > 0 && answer[len(answer)-1] > 0 {
				prev := answer[len(answer)-1]

				if prev > -val {
					collided = true
					break
				}

				if prev == -val {
					answer = answer[:len(answer)-1]
					collided = true
					break
				}

				answer = answer[:len(answer)-1]
			}

			if !collided {
				answer = append(answer, val)
			}
		}
	}
	return answer
}
