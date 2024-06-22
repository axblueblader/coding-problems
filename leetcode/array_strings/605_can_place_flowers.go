package leetcode

// Given we're at position i-th, what is the condition to put a flower there
// f[i] == 0
// f[i-1] == 0
// f[i+1] == 0
// We can iterate and keep count of number of flowers planted then return true once it's larger or equal to n
//
// Refactor:
// Checking for the array limits adds a few ugly 'if' checks
// Notice at f[0], f[i-1] is essentially 0
// If we store f[i] into a variable 'prev' for previous we can reduce one if check
//
// Small optimization: exit immediately once we can place more than 'n'
func canPlaceFlowers(flowerbed []int, n int) bool {
	canPlace := 0
	prev := 0
	bedLen := len(flowerbed)
	for i := 0; i < bedLen; i++ {
		if prev == 0 && flowerbed[i] == 0 {
			if i+1 >= bedLen || flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				canPlace++
				if canPlace >= n {
					return true
				}
			}
		}
		prev = flowerbed[i]
	}
	return canPlace >= n
}
