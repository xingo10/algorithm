/*
 * @lc app=leetcode.cn id=523 lang=golang
 *
 * [523] 连续的子数组和
 */

// @lc code=start
func checkSubarraySum(nums []int, k int) bool {
	/*
		解题思路：
			根据题意，求子数组元素总和为k的倍数
			按暴力法破解的话：
			for i in 0~n
				for j in 0~n
					(sum[j] - sum[i])%6 == 0
			也就是说sum[j]=sum[i]%6，即判断在sum里，有没有j满足上面的条件
			使用map来存储每个sum[i]%6的结果
	*/
	if len(nums) < 2 {
		return false
	}
	sum := 0
	m := map[int]int{
		0: -1,
	}
	for i := range nums {
		sum = (sum + nums[i]) % k
		if v, ok := m[sum]; ok {
			if i-v >= 2 {
				return true
			}
		} else {
			m[sum] = i
		}
	}
	return false
}

// @lc code=end

