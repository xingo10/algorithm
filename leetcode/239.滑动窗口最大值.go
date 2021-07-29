/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	/*
		解题思路：
			滑动窗口的最大值，使用单调递减队列
	*/
	var ans, queue []int
	for i := range nums {
		// 去掉队首，保证窗口大小合法性
		for len(queue) > 0 && queue[0] <= i-k {
			queue = queue[1:]
		}
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if i >= k-1 {
			ans = append(ans, nums[queue[0]])
		}
	}

	return ans
}

// @lc code=end

