# 第一周作业

## [合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/)

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }

    prehead := &ListNode{}
    res := prehead
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            res.Next = l1
            l1 = l1.Next
        } else {
            res.Next = l2
            l2 = l2.Next
        }
        res = res.Next
    }

    if l1 != nil {
        res.Next = l1
    }
    if l2 != nil {
        res.Next = l2
    }
    return prehead.Next
}
```

## [和为K的子数组](https://leetcode-cn.com/problems/subarray-sum-equals-k/)

```go
func subarraySum(nums []int, k int) int {
    res := 0
    count := map[int]int{0:1}
    preSum := 0
    for i := 0; i < len(nums); i++ {
        preSum += nums[i]
        if v, ok := count[preSum-k]; ok {
            res += v
        }
        count[preSum]++
    }
    
    return res
}
```
