/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	/*
		解题思路：
			题目是找出连续子数组，可以想到使用前缀和
			1. 求前缀和数组
			2. 因子数组和需要满足>=target，这里采用二分查找来实现，复杂度为O(logN)
	*/
	sums := make([]int, len(nums)+1, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	if sums[len(sums)-1] < target {
		return 0
	}

	ans := len(nums) + 1
	for i := 1; i < len(sums); i++ {
		low := i
		high := len(sums) - 1
		// fmt.Println("====start====", low, i, high)
		for low <= high {
			mid := (high-low)/2 + low
			if sums[mid] >= sums[i-1]+target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		if low <= len(nums) {
			// fmt.Println("====end====", low, i, ans)
			ans = min(ans, low-i+1)
		}
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

