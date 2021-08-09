/*
 * @lc app=leetcode.cn id=528 lang=golang
 *
 * [528] 按权重随机选择
 */

// @lc code=start
/*
	解题思路：
		把权重映射到数组中，使得下标和权重建立对映关系，使用前缀和数组
		找到第一个大于等于随机数的下标，返回
		关于权重算法实践的介绍：https://www.codeleading.com/article/26355762048/
*/
type Solution struct {
	sum    int
	prefix []int
}

func Constructor(w []int) Solution {
	p := make([]int, len(w))
	p[0] = w[0]
	for i := 1; i < len(w); i++ {
		p[i] = p[i-1] + w[i]
	}
	return Solution{p[len(p)-1], p}
}

func (this *Solution) PickIndex() int {
	target := rand.Intn(this.sum)
	left, right := 0, len(this.prefix)-1
	for left < right {
		mid := (right-left)/2 + left
		if target >= this.prefix[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
// @lc code=end

