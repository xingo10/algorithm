# [模拟行走机器人](https://leetcode-cn.com/problems/walking-robot-simulation/)

```go
func robotSim(commands []int, obstacles [][]int) int {
    // 把障碍物存放在哈希表中
    // 每次移动一格，判断下一步的(nextx,nexty)是否在障碍物哈希表中
    // 如果有的话，就执行下一个命令
    obstacleMap := make(map[string]int)
    for i := range obstacles {
        obstacleMap[calcHash(obstacles[i][0], obstacles[i][1])]++
    }

    // 定义方向数组
    //          N  E  S  W
    dx := []int{0, 1, 0, -1}
    dy := []int{1, 0, -1, 0}
    direction := 0
    x, y, ans := 0, 0, 0
    
    for _, cmd := range commands {
        if cmd > 0 {
            for i := 0; i < cmd; i++ {
                nextx := x + dx[direction]
                nexty := y + dy[direction]
                if _, ok := obstacleMap[calcHash(nextx, nexty)]; ok {
                    continue
                }
                x = nextx
                y = nexty
                ans = max(ans, x*x + y*y)
            }
        } else if cmd == -1 {
            direction = (direction + 1) % 4
        } else if cmd == -2 {
            // 左转相当于右转三次，所以这里使用+3， 避免-1出现下标为负数的情况
            direction = (direction + 3) % 4 
        }
    }
    return ans
}

func calcHash(x, y int) string {
    return fmt.Sprintf("%d,%d", x, y)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```

# 字母异位词分组

排序字符串中字母，采用哈希表实现

```go
func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)
    for i := range strs {
        b := []byte(strs[i])
        sort.Slice(b, func(i, j int) bool {
            return b[i] < b[j]
        })

        if _, ok := m[string(b)]; ok {
            m[string(b)] = append(m[string(b)], strs[i])
        } else {
            m[string(b)] = []string{strs[i]}
        }
    }

    res := make([][]string, 0)
    for _, v := range m {
        res = append(res, v)
    }
    return res
}
```
