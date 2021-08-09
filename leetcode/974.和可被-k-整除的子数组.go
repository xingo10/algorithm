/*
 * @lc app=leetcode.cn id=974 lang=golang
 *
 * [974] 和可被 K 整除的子数组
 */

// @lc code=start
func subarraysDivByK(nums []int, k int) int {
	/*
		解题思路：
			根据题意，连续子数组和可被K整除
			=> A[i]~A[j]的前缀和可被K整除
			=> (preSum(j)-preSum(i-1))%K == 0
			=> preSum(j)%K == preSum(i-1)%K
		前缀和求解的话，需要循环两次，需要进行优化。
		我们用map来存放前缀和modK后的结果作为key，value为出现这个key的次数

		补充：前缀和 为负数 的情况
		举例 K = 4，求出某个前缀和为 -1，-1 % K 应该为 3，但有的语言中：-1 % K = -1，这个 -1，要加上 K，转成正数的 3。
		K=4的情况下，为什么 preSum 值为 -1 和 3 属于同一类？算出 -1 也要算进来。因为：
		-1 和 3 分别模 4 的结果看似不相等，但前缀和之差：3-(-1) 为 4。4 % K = 0，所形成的子数组满足元素和被 4 整除。所以前缀和 -1 和 3 其实是等价的。
	*/
	ans := 0
	preSum := 0
	m := map[int]int{0: 1}
	for i := range nums {
		preSum = (preSum + nums[i]) % k
		if preSum < 0 {
			preSum += k
		}
		if v, ok := m[preSum]; ok {
			ans += v
			m[preSum]++
		} else {
			m[preSum] = 1
		}
	}

	return ans
}

// @lc code=end

