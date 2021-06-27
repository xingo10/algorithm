# [LRU](https://leetcode-cn.com/problems/lru-cache/)

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
