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
