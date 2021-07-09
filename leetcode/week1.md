# 第一周实战

## [合并有序数组](https://leetcode-cn.com/problems/merge-sorted-array/)

```go
func merge(nums1 []int, m int, nums2 []int, n int)  {
    /*
    解题思路:
        不开新的内存空间
        1. 为避免两个数组从0开始比较，导致覆盖nums1中后面的元素，因此采用倒序，数组由后向前比较
        2. 因为nums1的空间大小等于m+n，所以nums1的最后一位下标为m+n-1，从这个下标开始插入数据
    */
    tail := m+n-1
    // 需要先把m和n的长度减一，变为数组下标
    m = m-1
    n = n-1
    for m >= 0 || n >= 0 {
        // 如果n<0，说明nums2已经遍历完，无需再继续比较，直接把nums1中剩余元素放入数组即可
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
    /*
    解题思路：
        因为是原地修改数组，因此考虑用非重复项覆盖数组前面的重复项
        所以需要记录一个全局的index下标，表示当前非重复项的下标位置
        1. 如果数组长度为0或1，那肯定是没有重复元素的，直接返回数组长度即可
        2. 如果数组长度大于1，从1开始遍历数组，比对当前下标为i元素和i-1的元素是否相等
        3. 相等说明重复，跳过
        4. 如果不相等，把i元素覆盖index的位置，同时index向后移动一
    */
    if len(nums) == 0 {
        return 0
    }
    index := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            continue
        }
        nums[index] = nums[i]
        index++
    }
    return index
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

## [反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var dummy *ListNode
    for head != nil {
        temp := head.Next
        head.Next = dummy
        dummy = head
        head = temp
    }
    return dummy
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
