/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	/*
		解题思路：
			下一个更大元素，可以采用单调递增栈
			先计算出nums2中每个下标对应的下一个更大元素，存放在ans
			然后遍历nums1，根据nums1中元素在ans的位置，筛选出合适的答案
	*/
	ans := make([]int, len(nums2))
	// 初始化所有的ans都是-1
	for i := range ans {
		ans[i] = -1
	}
	// 放置nums2下标
	stack := make([]int, 0)
	// 存放nums2中每个元素的下标，方便遍历nums1的时候，找到ans里的答案
	m := make(map[int]int)

	// 把nums2下标的下一个更大元素都计算出来
	for i := range nums2 {
		for len(stack) != 0 && nums2[i] >= nums2[stack[len(stack)-1]] {
			// 栈顶元素的下标
			si := stack[len(stack)-1]
			// 距离下一个更大元素的距离
			// ans[si] = i - si
			ans[si] = nums2[i]
			// 单调递增栈，所以如果当前元素 >= 栈顶元素，pop栈顶元素
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		m[nums2[i]] = i
	}
	res := make([]int, 0)
	for _, n := range nums1 {
		res = append(res, ans[m[n]])
	}

	return res
}

// @lc code=end

