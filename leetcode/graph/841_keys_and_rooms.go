package leetcode

// https://leetcode.com/problems/keys-and-rooms/

// the list keys in a room is similar to list of connected vertices in a graph
// visit room 0, get all the keys, visit each room that the key unlocks and repeat
// until all rooms are visisted
// we can either use recursion or a queue to
func canVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)
	// queue that starts with room 0
	queue := []int{0}
	qIndex := 0
	for qIndex < len(queue) {
		roomNumber := queue[qIndex]
		qIndex++
		if visited[roomNumber] {
			continue
		}
		visited[roomNumber] = true
		for _, key := range rooms[roomNumber] {
			if !visited[key] {
				queue = append(queue, key)
			}
		}
	}
	return len(rooms) == len(visited)
}
