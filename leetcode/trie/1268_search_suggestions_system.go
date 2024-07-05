package leetcode

// https://leetcode.com/problems/search-suggestions-system/description/
import "sort"

// While constructing the trie, stores a reference/index to the word itself at each prefix
// insert "mobile" - index 0
// stores at node 'm' stores [0], at node 'mo' stores [0] ... at node "mobile" stores[0]
// insert "mousepad" - index 4
// at node 'm' stores [0,4], at node 'mo' stores [0,4], at node 'mou' stores [4]
// For the sorted requirements, we can sort the original product array
// so later we don't need to sort it for every result
func suggestedProducts(products []string, searchWord string) [][]string {
	type node struct {
		children     map[rune]*node
		productIndex []int
	}

	root := &node{
		children:     make(map[rune]*node),
		productIndex: []int{},
	}
	sort.Strings(products)
	for i, word := range products {
		curNode := root
		for _, ch := range word {
			_, ok := curNode.children[ch]
			if !ok {
				curNode.children[ch] = &node{
					children:     make(map[rune]*node),
					productIndex: []int{},
				}
			}
			curNode = curNode.children[ch]
			curNode.productIndex = append(curNode.productIndex, i)
		}
	}

	res := make([][]string, len(searchWord))
	curNode := root
	for i, ch := range searchWord {
		next, ok := curNode.children[ch]
		if ok {
			curNode = next
			recommends := []string{}
			for _, val := range curNode.productIndex {
				recommends = append(recommends, products[val])
			}
			res[i] = recommends[:min(len(recommends), 3)]
		} else {
			// if we don't break this case will fail
			// ["baby"]
			// "bacb"
			break
		}
	}

	return res
}

// Use Trie is useful but overkill in the context of this problem because the input changes
// on every test case, hence caching have no effect
// In a different context where we do multiple search on the same products
// then Trie will be more efficient
// For this however,
// It's simpler to sort the array and use binary search on the characters
func suggestedProductsBinarySearch(products []string, searchWord string) [][]string {
	res := make([][]string, len(searchWord))
	l := 0
	r := len(products) - 1

	sort.Slice(products, func(a, b int) bool {
		return products[a] < products[b]
	})
	// fmt.Println(products)
	for i := range searchWord {
		// fmt.Println(i,l,r,products[l][i],products[r][i],searchWord[i],len(products[l]),i)
		for l <= r && (len(products[l]) <= i || products[l][i] != searchWord[i]) {
			// fmt.Println("l++")
			l++
		}
		for l <= r && (len(products[r]) <= i || products[r][i] != searchWord[i]) {
			r--
		}
		res[i] = make([]string, 0)
		num := min(3, r-l+1)
		for j := 0; j < num; j++ {
			res[i] = append(res[i], products[l+j])
		}
	}
	return res
}
