# 前缀和、差分

## 前缀和

### [寻找数组的中心下标](https://leetcode-cn.com/problems/find-pivot-index/)

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

### [最大子序和](https://leetcode-cn.com/problems/maximum-subarray/)

- 解法一

```go
func maxSubArray(nums []int) int {
    /*
    解题思路：
        连续子数组，并且不修改原数组，因此考虑前缀和

        计算前缀和，同时记录最小的前缀和元素
        每次计算前缀和的时候，计算当前sum-最小前缀和是多少，把差值大的sum赋给结果用于返回
        每次计算前缀和的时候，进行最小的前缀和和当前前缀和的比较，把小的赋给minSum
    */
    sum := 0
    minSum := 0
    ans := nums[0]
    for _, n := range nums {
        sum += n
        ans = max(ans, sum-minSum)
        minSum = min(minSum, sum)
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
```

- 解法二

```go
func maxSubArray(nums []int) int {
    sums := make([]int, len(nums)+1, len(nums)+1)
    for i := 1; i <= len(nums); i++ {
        sums[i] = sums[i-1] + nums[i-1]
    }
    
    ans := nums[0]
    minSum := 0
    for i := 1; i <= len(nums); i++ {
        ans = max(ans, sums[i]-minSum)
        minSum = min(sums[i], minSum)
    }

    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
```

## 差分

### [航班预订统计](https://leetcode-cn.com/problems/corporate-flight-bookings/)

```go
func corpFlightBookings(bookings [][]int, n int) []int {
    /*
    解题思路：
        分析题目可知，需要频繁的对bookings中first~last区间的元素进行修改，因此可以使用差分
    */
    // 构建差分数组diff: diff[l] += d, diff[r+1] -= d
    // 因为这道题下标是从1开始计算，所以构建差分数组的时候，要下标-1
    diff := make([]int, n)
    for _, b := range bookings {
        first := b[0]
        last := b[1]
        seat := b[2]
        diff[first-1] += seat
        if last < len(diff) {
            diff[last] -= seat
        }
    }
    
    // 构建差分数组的前缀和，即为原数组
    res := make([]int, n)
    res[0] = diff[0]
    for i := 1; i < n; i++ {
        res[i] = res[i-1] + diff[i]
    }
    return res


}
```