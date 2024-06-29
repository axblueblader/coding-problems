package leetcode

// https://leetcode.com/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/description/

// Naive approach
//
// We can construct a graph from the edges, storing nodes to children array
// Have another array to store the results
// Iterate each node and depth first search into all it's children, adding itself as ancestor
// After full search, remove the duplicates and sort the result array
//
// Notice we need the ancestors, so it's possible to have store nodes to ancestors instead
// Then we can apply the same DFS but upwards, to parents of parents
//
// For the result, it needs to be a 'sorted set', there are data structures to help with this
// But since the limit is known, we can use an array of bool mapping node index to visit state
// Since the index is already sorted we can simply traverse the array and append it as an ancestors if it had been visited from the DFS
func getAncestors(n int, edges [][]int) [][]int {
	anc := make([][]int, n)
	for _, edge := range edges {
		anc[edge[1]] = append(anc[edge[1]], edge[0])
	}

	var dfs func(int, []bool)
	dfs = func(parent int, visited []bool) {
		if visited[parent] {
			return
		}
		visited[parent] = false
		for _, granParent := range anc[parent] {
			dfs(granParent, visited)
		}
	}

	res := make([][]int, n)
	for node, parents := range anc {
		visited := make([]bool, n)
		for _, parent := range parents {
			dfs(parent, visited)
		}
		for i, isAncestor := range visited {
			if isAncestor {
				res[node] = append(res[node], i)
			}
		}
	}
	return res
}
