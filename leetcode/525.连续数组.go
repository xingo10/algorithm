/*
 * @lc app=leetcode.cn id=525 lang=golang
 *
 * [525] 连续数组
 */

// @lc code=start
func findMaxLength(nums []int) int {
	/*
		解题思路:
			根据题意，找到含有相同数量的0和1的连续子数组，假设0都变为-1
			相当于找到一个和为0的最长子数组

	*/
	m := map[int]int{0: -1}
	res := 0
	sum := 0
	for i := range nums {
		if nums[i] == 0 {
			sum--
		} else {
			sum++
		}
		if v, ok := m[sum]; ok {
			res = max(res, i-v)
		} else {
			m[sum] = i
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=en

