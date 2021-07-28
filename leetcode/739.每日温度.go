/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	/*
		解题思路：
			该题是求当前天数之后，下一个更高温度的距离，也可以用单调递增栈实现。

	*/
	// 单调递增栈
	ans := make([]int, len(temperatures), len(temperatures))
	// 栈中存储温度下标，方便计算下标距离
	stack := make([]int, 0)

	for i := range temperatures {
		// 温度相同的话，认为没有更高温度，因此>=需要变为>
		for len(stack) != 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			si := stack[len(stack)-1]
			ans[si] = i - si
			// pop
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return ans
}

// @lc code=end

