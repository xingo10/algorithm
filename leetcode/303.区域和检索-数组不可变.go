/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	return NumArray{sums: sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	// 题意是要包含left和right，所以应该是left-1 ~ right+1 的和，
	// 而前缀和本身left就已经是-1了，所以这里只要right+1就可以。
	return this.sums[right+1] - this.sums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end

