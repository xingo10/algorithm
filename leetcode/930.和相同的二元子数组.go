/*
 * @lc app=leetcode.cn id=930 lang=golang
 *
 * [930] 和相同的二元子数组
 */

// @lc code=start
func numSubarraysWithSum(nums []int, goal int) int {
	/*
		解题思路：
			子数组和为goal，使用前缀和
	*/
	ans := 0
	sum := 0
	count := make(map[int]int)
	for i := range nums {
		count[sum]++
		sum += nums[i]
		ans += count[sum-goal]
	}
	return ans
}

// @lc code=end

