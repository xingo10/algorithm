# 前缀和

## [寻找数组的中心下标](https://leetcode-cn.com/problems/find-pivot-index/)

- 普通解法

```go
func pivotIndex(nums []int) int {
    /*
    解题思路：
        根据题意，左侧元素和+中心元素+右侧元素和=总和，即leftSum+pivot+RightSum=sum
        因为中心下标左右两侧元素和相等，所以2*leftSum+pivot=sum
        既而推导出：pivot=sum-2*leftSum
        根据上面推导出的结果来进行求解
        1. 计算数组所有元素总和sum
        2. 遍历数组，每遍历一个元素i，比对nums[i]=sum-2*leftSum
        3. 如果不相等，把元素累加到leftSum
        4. 如果相等直接返回下标
    */
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
    }

    leftSum := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == sum - 2*leftSum {
            return i
        }
        leftSum += nums[i]
    }
    return -1
}
```

- 前缀和解法

```go
func pivotIndex(nums []int) int {
    /*
    解题思路：
        根据题意，相当于计算0~i-1和i+1~length这两个区间的和是否相等
        是区间和的问题，可以考虑使用前缀和求解

        1. 计算nums的前缀和数组sums
        2. 遍历前缀和数组
        3. 每遍历一个下标元素，判断左侧区间和是否等于右侧区间和
            左侧区间和是从0到i-1，即sums[i-1]-sums[0]，因为sums[0]始终为0，所以是sums[i-1]
            右侧区间和是从i+1到length，即sums[length]-sums[i]
        4. 如果相等返回下标i-1，否则继续遍历
    */
    // 计算前缀和数组
    sums := make([]int, len(nums)+1, len(nums)+1)
    for i := 1; i <= len(nums); i++ {
        sums[i] = sums[i-1] + nums[i-1]
    }

    sum := sums[len(sums)-1]
    for i := 1; i < len(sums); i++ {
        if sums[i-1] == sum - sums[i] {
            return i-1
        }
    }

    return -1
}
```
