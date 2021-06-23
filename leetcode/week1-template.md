# 前缀和、数组计数模板

```go
func numberOfSubarrays(nums []int, k int) int {
    n := len(nums)
    s := make([]int, n+1, n+1)
    for i := 1; i <= n; i++ {
        s[i] = s[i-1] + nums[i-1]%2
    }

    // for i 1~n
    //   for j 0~i-1
    //     if s[i]-s[j]==k {
    //        ans++
    //     }
    // 固定外层循环，内层可以看成
    // 每个i，有多少个j满足s[i]-s[j]=k  ->   每个i，有多少个j满足s[j]=s[i]-k
    // 相当于计算数组中，有多少个元素满足s[i]-k
    // 使用map统计数组s中每个元素值的个数
    
    count := make(map[int]int)
    for i := 0; i < len(s); i++ {
        count[s[i]]++
    }

    // 遍历数组s，查看每个s[i]-k的数量
    ans := 0
    for i := 1; i <= n; i++ {
        ans += count[s[i]-k]
    }
    return ans
}
```
