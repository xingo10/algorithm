/*
 * @lc app=leetcode.cn id=1004 lang=golang
 *
 * [1004] 最大连续1的个数 III
 */

// @lc code=start
func longestOnes(nums []int, k int) int {
	/*
		解题思路:
			根据题意，是求 一段区间内0的个数 <= k
			把0和1互换，即：0变为1，1变为0，
			这样就等价于求一段区间内元素和 <=k的最长子数组
			p[j] - p[i] <= k, 即p[i] >= p[j] - k
			因为数组是非0即1，所以是个递增正整数数组
			使用二分查找，找到第一个满足上式的下标i
	*/
	sums := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		// 1-nums[i-1]是为了把0变为1，1变为0
		sums[i] = sums[i-1] + (1 - nums[i-1])
	}
	// fmt.Println(sums)
	ans := 0
	for i := range sums {
		left := sort.SearchInts(sums, sums[i]-k)
		ans = max(ans, i-left)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

