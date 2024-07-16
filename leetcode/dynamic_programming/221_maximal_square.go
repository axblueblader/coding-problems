package leetcode

// https://leetcode.com/problems/maximal-square/description/

// we can break it down into smaller cases
// [1] -> this is a square
// [1 1]
// [1 1]
// how do we know this is a square -> top left = 1, left = 1, top = 1
// what if we store the square size result in (i,j), the matrix becomes
// [1 1]
// [1 2]
// consider what happens as we extend to square of 3
// [1 1 1]    [1 1 1]
// [1 1 1] => [1 2 2]
// [1 1 1]    [1 2 3]
// consider this case when there's a hole
// [1 1 1]    [1 1 1]
// [0 1 1] => [0 1 2]
// [1 1 1]    [1 1 2]
// and this one
// [0 1 1]    [0 1 1]
// [1 1 1] => [1 1 2]
// [1 1 1]    [1 2 2]
// => we need to get the min(left, top left, top) + 1 to get the correct area
func maximalSquare(matrix [][]byte) int {
	var max byte
	max = 0
	a := matrix
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			a[i][j] = a[i][j] - 48 // becuz input is string
			// a[i][j] == 0 -> skip
			// a[i][j] == 1 && upper surround all 1 -> min(upper surround) + 1
			if i-1 >= 0 && j-1 >= 0 {
				if a[i][j] >= 1 && a[i-1][j] >= 1 && a[i][j-1] >= 1 && a[i-1][j-1] >= 1 {
					a[i][j] = minNum(a[i-1][j], a[i][j-1], a[i-1][j-1]) + 1
				}
			}
			if a[i][j] > max {
				max = a[i][j]
			}
		}
	}
	return int(max * max)
}

func minNum(arr ...byte) byte {
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}
