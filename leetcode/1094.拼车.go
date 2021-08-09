/*
 * @lc app=leetcode.cn id=1094 lang=golang
 *
 * [1094] 拼车
 */

// @lc code=start
func carPooling(trips [][]int, capacity int) bool {
	/*
		解题思路：
			差分+前缀和还原
			只要判断前缀和数组中是否有元素大于cap即可。
	*/
	diff := make([]int, 1000+1)
	for _, trip := range trips {
		diff[trip[1]] += trip[0]
		diff[trip[2]] -= trip[0]
	}

	sum := 0
	for i := range diff {
		sum += diff[i]
		if sum > capacity {
			return false
		}
	}

	return true
}

// @lc code=end

