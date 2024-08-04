package leetcode

// https://leetcode.com/problems/flatten-nested-list-iterator/

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct{}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool { return true }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int { return 1 }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger { return nil }

type NestedIterator struct {
	Res []int
	Idx int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	stack := []*NestedInteger{}
	for i := len(nestedList) - 1; i >= 0; i-- {
		if nestedList[i].IsInteger() || len(nestedList[i].GetList()) > 0 {
			stack = append(stack, nestedList[i])
		}
	}

	res := []int{}
	for len(stack) > 0 {
		nextVal := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for nextVal != nil && !nextVal.IsInteger() {
			nextList := nextVal.GetList()
			for i := len(nextList) - 1; i >= 0; i-- {
				if nextList[i].IsInteger() || len(nextList[i].GetList()) > 0 {
					stack = append(stack, nextList[i])
				}
			}
			if len(stack) > 0 {
				nextVal = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			} else {
				nextVal = nil
			}
		}
		if nextVal != nil {
			res = append(res, nextVal.GetInteger())
		}
	}

	return &NestedIterator{
		Res: res,
		Idx: 0,
	}
}

func (this *NestedIterator) Next() int {
	val := this.Res[this.Idx]
	this.Idx += 1
	return val
}

func (this *NestedIterator) HasNext() bool {
	return this.Idx < len(this.Res)
}
