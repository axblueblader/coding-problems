package leetcode

// https://leetcode.com/problems/implement-trie-prefix-tree/

// In a Trie, every node stores a value to indicate that this node is the end of a word
// and it's children is a map of character that comes after it
type Node struct {
	children map[rune]*Node
	isWord   bool
}
type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{
			children: make(map[rune]*Node),
			isWord:   false,
		},
	}
}

func (trie *Trie) Insert(word string) {
	curNode := trie.root
	for _, ch := range word {
		_, ok := curNode.children[ch]
		if !ok {
			curNode.children[ch] = &Node{
				children: make(map[rune]*Node),
			}
		}
		curNode = curNode.children[ch]
	}
	curNode.isWord = true
}

func (trie *Trie) Search(word string) bool {
	curNode := trie.root
	for _, ch := range word {
		nextNode, ok := curNode.children[ch]
		if ok {
			curNode = nextNode
		} else {
			return false
		}
	}
	return curNode.isWord
}

func (trie *Trie) StartsWith(prefix string) bool {
	curNode := trie.root
	for _, ch := range prefix {
		nextNode, ok := curNode.children[ch]
		if ok {
			curNode = nextNode
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
