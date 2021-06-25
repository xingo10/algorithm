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

# [串联所有单词的子串](https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/)

```go
func findSubstring(s string, words []string) []int {
    /* 解题思路：
        需要判断字符串s中子串和words完全匹配，那么就说明子串长度应该是words中所有元素的长度
        因为words中每个元素长度相同，所以words中所有元素长度为len(words)*len(words[0])(即：第一个元素的字符串长度)
        for循环类似滑动获取子串
    
        子串想和words完全匹配，可以把子串按len(words[0])进行分割
        因为匹配没有顺序要求，所以只需要保证分割后，每个串的出现次数和words中出现次数一致即可
        通过哈希表存储分割后的每个串和对应的次数
        和words中元素出现次数做对比，完全一致说明匹配成功
    */

    ans := []int{}
    // 根据上面思路，最后需要对比子串中和words中元素出现次数，所以先对words中元素出现次数做统计
    countWordsMap := countWords(words)

    // 接着获取s的子串
    // 先决定子串的长度是多少
    wordsLength := len(words) * len(words[0])
    // start + wordsLength <= len(s) 是保证随着start的下标后移，子串长度如果小于words的总长度，说明子串不可能完全匹配words，可以忽略后面的子串，循环结束
    end := wordsLength
    for start := 0; start + wordsLength <= len(s); start++ {
        subString := s[start:end]
        countWordsOfSubString := getWordsOfSubString(subString, len(words[0]))
        if reflect.DeepEqual(countWordsOfSubString, countWordsMap) {
            ans = append(ans, start)
        }
        end++
    }
    return ans
}

func getWordsOfSubString(s string, length int) map[string]int {
    res := make(map[string]int)
    for i := 0; i + length <= len(s); i += length {
        res[s[i:i+length]]++
    }
    return res
}

func countWords(words []string) map[string]int {
    res := make(map[string]int)
    for i := range words {
        res[words[i]]++
    }
    return res
}
```
