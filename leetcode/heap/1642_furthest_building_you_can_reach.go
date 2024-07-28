package leetcode

// https://leetcode.com/problems/furthest-building-you-can-reach/

import "container/heap"

// greedy: we should use ladder for higher jumps or when there's not enough bricks left
// we can try to use all ladders first,
// when we meet a tall building, we can swap used ladder

// which ladder to swap with?
// bricks used
// 6 5 4
// now i need
// if 7 => swap to get 6
// if 3 => swap to get 4
// swap with the minimum to fulfill this jump => need to use a heap for quickest access to the minimum

func furthestBuilding(heights []int, bricks int, ladders int) int {
	prevH := heights[0]
	brickCost := &Heap{}
	i := 1
	for i = 1; i < len(heights); i++ {

		diff := heights[i] - prevH
		prevH = heights[i]
		if diff <= 0 {
			continue
		}

		if ladders > 0 {
			ladders--
			heap.Push(brickCost, diff)
			continue
		}

		heap.Push(brickCost, diff)
		minCost := heap.Pop(brickCost).(int)

		bricks -= minCost
		if bricks < 0 {
			break
		}
	}

	return i - 1
}

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old) - 1
	v := old[n]
	*h = old[:n]
	return v
}
