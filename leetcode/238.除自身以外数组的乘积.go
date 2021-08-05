/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	/*
		解题思路：
		1、计算所有元素总乘积，除以nums[i]即可（题目不允许使用除法）
		2、暴力法 O(N^2)：
			for i in 0~n
				for j in 0~n
					if j != i: p *= nums[j]
		3、类似前缀和方式：
			其实求除i之外其余元素的乘积，可以按i的左侧和右侧分开计算，即
			p = p(左)*p(右) -> p = 0~i的乘积 * i+1~n的乘积
			左侧可以使用类似前缀和的概念，p[i]=p[i-1]*nums[i]
			右侧的话，因为乘法两侧数字顺序无所谓，所以可以变成从数组尾到0开始算前缀乘积
			最后的乘积p=l[i]*r[i]
	*/
	left := make([]int, len(nums)+1)
	right := make([]int, len(nums)+1)
	// 索引为0的元素左侧没有元素，因此前缀积的话，left[0]=1,可以保证1乘任何数都是本身
	left[0] = 1
	// nums数组最后右侧没有元素，因此需要设置右侧前缀积的0下标为1
	right[len(right)-1] = 1

	for i := 1; i <= len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	for j := len(nums) - 1; j >= 0; j-- {
		right[j] = right[j+1] * nums[j]
	}
	// fmt.Println(left, right)
	ans := make([]int, len(nums))
	for i := range nums {
		ans[i] = left[i] * right[i+1]
	}
	return ans
}

// @lc code=end

