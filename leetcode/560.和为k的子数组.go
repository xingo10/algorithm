/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	/*
		解题思路：
			找到和为k的连续子数组 -> 前缀和
			求sum[j]-sum[i]=k -> sum[i]=sum[j]-k
			求sum中，有多少个i满足上面的公式
			使用map存储每个sum[i]，判断sum-k在map中存在的数量
	*/
	ans := 0
	sum := 0
	count := map[int]int{0: 1}
	for _, n := range nums {
		sum += n
		if v, ok := count[sum-k]; ok {
			ans += v
		}
		count[sum]++
	}

	return ans
}

// @lc code=end

