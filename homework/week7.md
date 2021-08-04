# 第7周作业

## [冗余连接](https://leetcode-cn.com/problems/redundant-connection/)

```go
func findRedundantConnection(edges [][]int) []int {
    /*
    解题思路：
        在一棵树中，边的数量比节点的数量少1。如果一棵树有 N 个节点，则这棵树有 N-1 条边。
        1. 如果两个顶点属于不同的连通分量，则说明在遍历到当前的边之前，这两个顶点之间不连通，
        因此当前的边不会导致环出现，合并这两个顶点的连通分量。
        2. 如果两个顶点属于相同的连通分量，则说明在遍历到当前的边之前，这两个顶点之间已经连通，
        因此当前的边导致环出现，为附加的边，将当前的边作为答案返回。
    */
    // parent[i] = j表示的是第i个节点前驱节点为j, 连通树中根节点k满足parent[k]==k
    var parent = make([]int, len(edges)+1)
    for i := 1; i <= len(edges); i++ {
        parent[i] = i
    }

    for i := 0; i < len(edges); i++ {
        a := find(edges[i][0], parent)
        b := find(edges[i][1], parent)
        if a == b {
            return []int{edges[i][0], edges[i][1]}
        } else {
            // 合并
            union(a, b, parent)
        }
    }
    return nil
}

func find(x int, p []int) int {
    if p[x] != x {
        p[x] = find(p[x], p)
    }

    return p[x]
}

func union(x, y int, p []int) {
    x_root := find(x, p)
    y_root := find(y, p)
    if x_root != y_root {
        p[x_root] = y_root
    }
}
```

## [岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)

```go
// 参考了其他的解法
var total int
func numIslands(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])
	mmap := initMap(grid, n, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				if i-1 >= 0 && grid[i-1][j] == '1' {
					union(i*m+j, (i-1)*m+j, mmap)
				}
				if i+1 < n && grid[i+1][j] == '1' {
					union(i*m+j, (i+1)*m+j, mmap)
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					union(i*m+j, i*m+j-1, mmap)
				}
				if j+1 < m && grid[i][j+1] == '1' {
					union(i*m+j, i*m+j+1, mmap)
				}
			}
		}
	}
	res := total
	total = 0
	return res
}

func initMap(grid [][]byte, n, m int) map[int]int {
    res := make(map[int]int)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				res[i*m+j] = i*m+j
				total++
			}
		}
	}
    return res
}

func find(x int, m map[int]int) int {
	if m[x] != x {
		m[x] = find(m[x], m)
	}
	return m[x]
}

func union(x, y int, m map[int]int) {
	xp := find(x, m)
	yp := find(y, m)
	if xp == yp {
		return
	}
	m[xp] = yp
	total--
}
```
