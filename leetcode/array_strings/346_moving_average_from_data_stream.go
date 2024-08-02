package leetcode

// https://leetcode.com/problems/moving-average-from-data-stream/

type MovingAverage struct {
	Size int
	Sum  int
	Vals []int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		Size: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.Vals) < this.Size {
		this.Sum += val
		this.Vals = append(this.Vals, val)
		return float64(this.Sum) / float64(len(this.Vals))
	}
	this.Sum = this.Sum - this.Vals[0] + val
	this.Vals = this.Vals[1:]
	this.Vals = append(this.Vals, val)
	return float64(this.Sum) / float64(len(this.Vals))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
