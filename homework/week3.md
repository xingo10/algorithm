# 第三周

## [从中序与后序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    n := len(inorder)
    if n == 0 {
        return nil
    }
    root := &TreeNode{postorder[n-1], nil, nil}
    i := 0
    for ; i < n; i++ {
        if inorder[i] == postorder[n-1] {
            break
        }
    }
    root.Left = buildTree(inorder[:i], postorder[:i])
    root.Right = buildTree(inorder[i+1:], postorder[i:n-1])
    return root
}
```

## [210. 课程表 II](https://leetcode-cn.com/problems/course-schedule-ii/)

```go
func findOrder(numCourses int, prerequisites [][]int) []int {
    var (
        edges = make([][]int, numCourses)
        indeg = make([]int, numCourses)
        learned = make([]int, 0)
    )

    // 加边
    for _, p := range prerequisites {
        ai := p[0]
        bi := p[1]
        // 因为ai依赖bi，因此图的有向边是bi->ai
        edges[bi] = append(edges[bi], ai)
        // 记录入边
        indeg[ai]++
    }

    queue := make([]int, 0)
    for i := 0; i < numCourses; i++ {
        // 如果某个课程入边是0，说明学习该课程不依赖其他课程，放入queue中进行学习
        if indeg[i] == 0 {
            queue = append(queue, i)
        }
    }

    // BFS
    for len(queue) > 0 {
        // 取队列头
        q := queue[0]
        queue = queue[1:]
        learned = append(learned, q)
        for _, v := range edges[q] {
            indeg[v]--
            if indeg[v] == 0 {
                queue = append(queue, v)
            }
        }
    }
    if len(learned) == numCourses {
        return learned
    }
    return nil
}
```
