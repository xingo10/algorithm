# 第五周作业

## [在 D 天内送达包裹的能力](https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/)

```go
func shipWithinDays(weights []int, days int) int {
    /*
    解题思路：
        假设船的运载能力为m的时候可以在指定的days内送达所有包裹，
        那么运载能力在最低和最大运载能力之间肯定可以找到一个临界值，满足在days内送达所有包裹。
        这样的话可以转换为：二分查找，一个正整数升序排列的数组中，找到一个值（可以满足指定的条件）。

        根据上面分析，接下来需要判断值是否满足条件，这个条件怎么来判定。
        必须要按顺序来送包裹，所以遍历weights，累加每个包裹的重量，重量超过运载能力，就第二天送；
        最后得到在运载能力为n的时候，需要几天送达，从而判断该运载能力是否满足条件。
        
        最低运载能力，一定是大于等于单个包裹的最大重量，否则无法完全运送完包裹
        而最大的运载能力是小于等于所有包裹的总和

    */
    left, right := getMaxAndSumWeights(weights)
    // 二分查找
    for left < right {
        mid := left + (right-left)/2
        // 如果能在days完成，说明运载能力较大，可以减少right
        if check(weights, mid, days) {
            right = mid
        } else {
            left = mid+1
        }
    }
    return left
}

func check(weights []int, cap, days int) bool {
    // sum是当天包裹重量总和
    sum := 0
    // 送达需要的天数
    day := 1
    for _, w := range weights {
        if day > days {
            return false
        }
        // 判断如果当天包裹总重量sum > 运载能力cap，就需要第二天再运送
        if sum + w > cap {
            day++
            sum = 0
        }
        sum += w
    }
    return day <= days
}

func getMaxAndSumWeights(weights []int) (max int, sum int) {
    for _, w := range weights {
        if w > max {
            max = w
        }
        sum += w
    }
    return
}
```

## [爱吃香蕉的珂珂](https://leetcode-cn.com/problems/koko-eating-bananas/)

```go
func minEatingSpeed(piles []int, h int) int {
    /*
    解题思路：
        分析题目，暴力的方式是从1开始，每次递增1，计算每次吃完香蕉的时间。
        随着吃的速度增加，耗费的时间随之减少，可以看出来是单调的。
        与lc.1011这道题类似，都可以采用二分查找的方式。

        二分查找需要找到两个左右边界
        左：从最小的1开始
        右：因为每小时最多吃掉某一堆香蕉，所以吃香蕉的速度不会每堆香蕉的最大值
    */
    left, right := 1, getMax(piles)
    for left < right {
        mid := left + (right-left)/2
        if check(piles, h, mid) {
            right = mid
        } else {
            left = mid+1
        }
    }
    return left
}

func check(piles []int, h, speed int) bool {
    hour := 0
    for _, p := range piles {
        if hour > h {
            return false
        }
        hour += p/speed
        // 如果无法整除，说明该堆香蕉剩下不到speed根，还需要再花1小时
        if p%speed != 0 {
            hour++
        }
    }
    return hour <= h
}

func getMax(piles []int) (max int) {
    for _, v := range piles {
        if v > max {
            max = v
        }
    }
    return 
}
```