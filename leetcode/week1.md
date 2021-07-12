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

## [K 个一组翻转链表](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    /*
    解题思路：
        1. 分多组遍历整个链表，重新组织链表
        2. 反转每组链表(每组移动k-1步)
        3. 上一组链表末尾指向反转后的当前组链表头
        4. 反转后当前组链表末尾指向未反转时该组链表末尾的next
        5. 需要记录上一组链表尾和该组原有链表末尾的next
        6. 为了避免处理表头为null的情况，在表头增加一个保护点，方便进行反转处理
    */
    protect := &ListNode{0, head}
    // 上一组的末尾
    last := protect
    for head != nil {
        // 获取每组尾部，如果end为nil，说明到了链表末尾或者是剩余长度不足k
        end := getEnd(head, k)
        if end == nil {
            break
        }

        // 保留下一组的头节点，用于下一次遍历head
        nextGroupHead := end.Next

        // 反转后，end是该组头节点，head是末尾
        reverse(head, end)
        
        // 上一组末尾和该组头关联
        last.Next = end
        // head的next指向下一组，是使得last的next指向下一组
        head.Next = nextGroupHead
        
        // 分组遍历
        // head是该组末尾，last是上一组末尾，因此把当前组末尾赋值给last
        last = head
        // head到下一组链表头
        head = nextGroupHead
    }
    return protect.Next
}

func getEnd(head *ListNode, k int) *ListNode {
    for head != nil {
        k--
        if k == 0 {
            break
        }
        head = head.Next
    }
    return head
}

/*
  dummy              end
    3   -> 4    ->    5     
         start

         dummy       end
    3   -> 4    ->    5
                    start
*/
func reverse(start, end *ListNode) {
    if start == end {
        return
    }
    dummy := start
    start = start.Next
    for start != end {
        temp := start.Next
        start.Next = dummy
        dummy = start
        start = temp
    }
    // 反转后，end成组的头节点，end的next需要指向dummy，因为dummy在start的前一个节点
    end.Next = dummy
    return
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
