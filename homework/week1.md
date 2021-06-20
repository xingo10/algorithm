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
