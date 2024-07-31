package leetcode

// https://leetcode.com/problems/bus-routes/

// 1 stop can have multiple bus passes by
// 1 section can be part of multiple routes e.g (1,2,3) (1,2,4) (2,4,8)
// intuition: directed graph with edge value being the bus route number
// set difference?
// node 1 has route set (0,1); 2: (0,1,2) 3: (0) 4: (1,2) 8: (2)
// start = 1, end = 8
// at node 1, try going to 2,3,4
// go to node 2 -> option becomes set2 - set1 = (1,2,3,4,8) - (1,2,3,4) = (8)
// go to node 8

// (1,2) (2,3,4) (3,5) (5,6) (6,8) (4,8)
// for shortest path it's better to use BFS
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	// stop to route
	// stop 1 -> [0,1]
	busGraph := make(map[int]map[int]bool)
	for routeNo, stops := range routes {
		for _, stop := range stops {
			if busGraph[stop] == nil {
				busGraph[stop] = make(map[int]bool)
			}
			busGraph[stop][routeNo] = true
		}
	}

	// Use BFS so we can find the shortest path
	routeQ := []int{}
	visited := make(map[int]bool)
	for routeNo := range busGraph[source] {
		routeQ = append(routeQ, routeNo)
		visited[routeNo] = true
	}

	count := 1
	for len(routeQ) > 0 {
		routesOfThisStop := len(routeQ)
		// check every bus/route that goes through this stop
		for i := 0; i < routesOfThisStop; i++ {
			route := routeQ[0]
			routeQ = routeQ[1:]

			for _, stop := range routes[route] {
				if stop == target {
					return count
				}

				for routeNo := range busGraph[stop] {
					if !visited[routeNo] {
						visited[routeNo] = true
						routeQ = append(routeQ, routeNo)
					}
				}
			}
		}
		count++
	}

	return -1
}
