# [LRU Cache](https://leetcode-cn.com/problems/lru-cache/)

模拟golang里的"container/list"包实现双向链表

```go
type Entry struct {
    key int
    value int
}

type Element struct {
    Value Entry
    prev *Element
    next *Element
}

type List struct {
    root Element
    length int
}

// 建立root的保护点，减少判断nil的情况
func newList() *List {
    l := new(List)
    l.root.prev = &l.root
    l.root.next = &l.root
    l.length = 0
    return l
}

func (l *List) move(e, at *Element) *Element {
    e.prev.next = e.next
    e.next.prev = e.prev

    e.next = at.next
    e.prev = at

    e.next.prev = e
    e.prev.next = e

    return e
}

func (l *List) insert(e, at *Element) *Element {
    e.next = at.next
    e.prev = at

    e.next.prev = e
    e.prev.next = e

    l.length++

    return e
}

func (l *List) MoveToFront(e *Element) {
    // 如果root的下一个就是e，说明e就在链表头
    if l.root.next == e {
        return
    }
    l.move(e, &l.root)
}

func (l *List) PushToFront(e Entry) *Element {
    return l.insert(&Element{Value: e}, &l.root)
}

// Back 返回最后一个元素
func (l *List) Back() *Element {
    if l.length == 0 {
        return nil
    }
    return l.root.prev
}

func (l *List) Remove(e *Element) Entry {
    e.prev.next = e.next
    e.next.prev = e.prev
    
    e.prev = nil
    e.next = nil

    l.length--

    return e.Value
}

// cache存储key，value是元素，查找可以实现O(1)的复杂度
// dl是个双向链表，实现O(1)的插入和删除
type LRUCache struct {
    cache map[int]*Element
    cap int
    dl *List
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        cache: make(map[int]*Element),
        cap: capacity,
        dl: newList(),
    }
}


func (this *LRUCache) Get(key int) int {
    if e, ok := this.cache[key]; ok {
        this.dl.MoveToFront(e)
        return e.Value.value
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if e, ok := this.cache[key]; ok {
        // 如果key存在，更新双向链表里该项的数据
        e.Value = Entry{key, value}
        this.dl.MoveToFront(e)
        return
    }
    // 如果不存在，把数据插入到链表头
    this.cache[key] = this.dl.PushToFront(Entry{key, value})
    // 如果缓存cache的长度大于cap，需要删除掉链表最后一个元素
    if len(this.cache) > this.cap {
        v := this.dl.Remove(this.dl.Back())
        delete(this.cache, v.key)
    }

}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
```

# [子域名访问计数](https://leetcode-cn.com/problems/subdomain-visit-count/)

```go
import "strings"

func subdomainVisits(cpdomains []string) []string {
    /*
    解题思路：
        用map存放每个域名和对应的访问次数。其中域名作为key存入map，访问次数作为value
        把域名通过"."进行分割，不断修改域名去掉前缀，并使用map记录相应的访问次数
    */
    domainMap := make(map[string]int)
    for _, value := range cpdomains {
        // 数组中每个元素格式是"次数 域名"，所以这里使用" "分割
        // 第一个为访问次数，第二个为域名
        ss := strings.Split(value, " ")
        count, _ := strconv.Atoi(ss[0])
        // 对于每个域名，通过"."继续分割，这里使用递归的方式
        // 传入domainMap是为了记录分割后的域名及访问次数
        visits(ss[1], count, domainMap)
    }

    ans := make([]string, 0)
    for k, v := range domainMap {
        ans = append(ans, fmt.Sprintf("%d %s", v, k))
    }
    return ans
}

func visits(domain string, count int, domainMap map[string]int) {
    domainMap[domain] += count
    dotIndex := strings.Index(domain, ".")
    if dotIndex < 0 {
        return
    }
    visits(domain[dotIndex+1:], count, domainMap)
}
```
