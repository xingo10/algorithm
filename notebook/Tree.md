# 树

## 二叉树的四种遍历模版

### 前序遍历

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorder(root *TreeNode) []int {
	var ans []int
	var order func(root *TreeNode)
	order = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		// 二叉树
		order(root.Left)
		order(root.Right)
		/** N叉树
		for _, child := range root.Children {
			order(child)
		}
		*/
	}
	order(root)
	return ans
}
```

### 中序遍历

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorder(root *TreeNode) []int {
	var ans []int
	var order func(root *TreeNode)
	order = func(root *TreeNode) {
		if root == nil {
			return
		}
		order(root.Left)
		ans = append(ans, root.Val)
		order(root.Right)
	}
	order(root)
	return ans
}
```

### 后序遍历

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorder(root *TreeNode) []int {
	var ans []int
	var order func(root *TreeNode)
	order = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 二叉树
		order(root.Left)
		order(root.Right)
		/** N叉树
		for _, child := range root.Children {
			order(child)
		}
		*/
		ans = append(ans, root.Val)
	}
	order(root)
	return ans
}
```

### 层序遍历

二叉树

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**输入：
    3
   / \
  9  20
    /  \
   15   7
解题思路：
	1. 初始化ans记录结果
	2. 记录每层数据到数组nodes中，只要数据中有值，就不断遍历，直到树的最底部没有节点为止
	3. 每层遍历的时候，记录本层的节点值以及下一层的节点
	4. 每层遍历结束后，把下一层节点覆盖到nodes中，进行下一层的遍历
*/
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    ans := make([][]int, 0)
    nodes := []*TreeNode{root}
    for len(nodes) != 0 {
        // 记录每层的节点值
        levelNodes := make([]int, 0)
        nextNodes := make([]*TreeNode, 0)
        // 遍历每层数据，并记录本层的下一层节点
        for i := range nodes {
            n := nodes[i]
            // 本层数据记录下来
            levelNodes = append(levelNodes, n.Val)
            if n.Left != nil {
                nextNodes = append(nextNodes, n.Left)
            }
            if n.Right != nil {
                nextNodes = append(nextNodes, n.Right)
            }
        }
        ans = append(ans, levelNodes)
        nodes = nextNodes
    }
    return ans
}
```

N叉树

```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    if root == nil {
        return nil
    }
    ans := make([][]int, 0)
    nodes := []*Node{root}
    for len(nodes) != 0 {
        levelNodes := make([]int, 0)
        nextNodes := make([]*Node, 0)
        for i := range nodes {
            n := nodes[i]
            levelNodes = append(levelNodes, n.Val)
            for _, child := range n.Children {
                nextNodes = append(nextNodes, child)
            }
        }
        ans = append(ans, levelNodes)
        nodes = nextNodes
    }

    return ans
}
```

### 树的最大深度

- 二叉树的最大深度

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := maxDepth(root.Left)
    right := maxDepth(root.Right)
    if left > right {
        return left + 1
    }
    return right + 1
}
```

- N叉树的最大深度

```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }

    max := 0
		// 计算每个子节点的深度，保存最大深度
    for _, child := range root.Children {
        depth := maxDepth(child)
        if depth > max {
            max = depth
        }
    }
    return max + 1
}
```

## 构造二叉树

**构造二叉树，首先是找到根节点，一般是前序或者后序是最容易找到根节点的，然后通过另外的数组来找到根节点所在下标位置，分割数组，进行左右子树的处理。**

有前序的话，通过前序找到根节点，然后在后序（或中序）中找到根节点所在的下标位置，进行左右子树分割

中序和后序的情况下，通过后序更容易找到根节点。因此在中序中找到根节点下标位置，进行左右子树分割。

### 前序和中序构造二叉树

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &TreeNode{preorder[0], nil, nil}
    index := 0
    // 查找中序中父节点位置
    for i := range inorder {
        if inorder[i] == preorder[0] {
            index = i
            break   
        }
    }
    // 左子树为：从前序中下标为1的元素 ~ 中序父节点位置之前的元素长度+1
    root.Left = buildTree(preorder[1:index+1], inorder[:index])
    root.Right = buildTree(preorder[index+1:], inorder[index+1:])
    return root
}
```

### 前序和后序构造二叉树

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(pre []int, post []int) *TreeNode {
    if len(pre) == 0 {
        return nil
    }
    // 前序数组长度
    n := len(pre)
    // 前序数组第一个元素为根节点
    root := &TreeNode{pre[0], nil, nil}
    if n == 1 {
        return root
    }

    i := 0
    for ; i < n; i++ {
        if post[i] == pre[1] {
            break
        }
    }
    // i+2 是因为i为根节点的下标，需要计算下标为i的数组的长度+1
    root.Left = constructFromPrePost(pre[1:i+2], post[:i+1])
    root.Right = constructFromPrePost(pre[i+2:], post[i+1:n-1])

    return root
}
```

### 中序和后序构造二叉树

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

### 序列化和反序列化构造二叉树

[https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/)

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    var s strings.Builder
    var preorder func(root *TreeNode)
    preorder = func(root *TreeNode) {
        if root == nil {
            s.WriteString("null ")
            return
        }
        s.WriteString(fmt.Sprintf("%d ", root.Val))
        preorder(root.Left)
        preorder(root.Right)
    }
    preorder(root)
    return s.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    sp := strings.Split(data, " ")
    var build func() *TreeNode
    build = func() *TreeNode {
        first := sp[0]
        sp = sp[1:]
        if first == "null" {
            return nil
        }
        val, _ := strconv.Atoi(first)
        root := &TreeNode{Val: val}
        root.Left = build()
        root.Right = build()
        return root
    }
    return build()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
```