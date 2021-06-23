# 第一周实战

## [合并有序数组](https://leetcode-cn.com/problems/merge-sorted-array/)

```go
func merge(nums1 []int, m int, nums2 []int, n int)  {
    tail := m+n-1
    // 最后元素的下标为长度-1
    m--
    n--
    for m >= 0 || n >= 0 {
        if n < 0 || (m >= 0 && nums1[m] > nums2[n]) {
            nums1[tail] = nums1[m]
            m--
        } else {
            nums1[tail] = nums2[n]
            n--
        }
        tail--
    }
}
```

## [去重](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/submissions/)

```go
func removeDuplicates(nums []int) int {
    if len(nums) <= 0 {
        return 0
    }
    n := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] {
            nums[n] = nums[i]
            n++
        }
    }
    return n
}
```

## [移动零](https://leetcode-cn.com/problems/move-zeroes/)

```go
func moveZeroes(nums []int)  {
    n := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[n] = nums[i]
            n++
        }
    }
    for j := n; j < len(nums); j++ {
        nums[j] = 0
    }
}
```

## [两数之和 II](https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/)

```go
func twoSum(numbers []int, target int) []int {
    i, j := 0, len(numbers)-1
    for i < j {
        sum := numbers[i] + numbers[j]
        if sum == target {
            return []int{i+1, j+1}
        } else if sum < target {
            i++
        } else {
            j--
        }
    }
    return []int{}
}
```
