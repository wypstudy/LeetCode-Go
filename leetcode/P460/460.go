package main

import (
	"fmt"
)

func NewLFU(c int) LFUCache {
	return Constructor(c)
}

type LFUCache struct {
	kv         map[int]int
	head, tail *Node
	capacity   int
	size       int
}

type Node struct {
	pre, next *Node
	key, freq int
}

func Constructor(capacity int) LFUCache {
	head := &Node{}
	tail := &Node{}
	head.pre = nil
	head.next = tail
	tail.pre = head
	tail.next = nil
	return LFUCache{
		kv:       make(map[int]int),
		head:     head,
		tail:     tail,
		capacity: capacity,
		size:     0,
	}
}

func (lfu *LFUCache) Get(key int) int {
	var (
		v  int
		ok bool
	)
	if v, ok = lfu.kv[key]; !ok {
		return -1
	}
	lfu.refresh(key)
	return v
}

func (lfu *LFUCache) Put(key int, value int) {
	if _, ok := lfu.kv[key]; ok {
		lfu.kv[key] = value
		lfu.refresh(key)
		return
	}
	lfu.kv[key] = value
	n := &Node{
		key:  key,
		freq: 1,
	}
	lfu.size++
	if lfu.size > lfu.capacity {
		lfu.eliminate()
	}
	lfu.addNode(n)
}

func (lfu *LFUCache) eliminate() {
	// fmt.Printf("eliminate %d\n", lfu.tail.pre.key)
	if lfu.size == 0 {
		return
	}
	delete(lfu.kv, lfu.tail.pre.key)
	lfu.removeNode(lfu.tail.pre)
	lfu.size--
}

func (lfu *LFUCache) refresh(key int) {
	nowPtr := lfu.findNode(key)
	if nowPtr == nil {
		return
	}
	nowPtr.freq++
	// fmt.Printf("refresh %d freq %d\n", key, nowPtr.freq)
	lfu.removeNode(nowPtr)
	lfu.addNode(nowPtr)
}

func (lfu *LFUCache) findNode(key int) *Node {
	ptr := lfu.head.next
	for ptr.key != key && ptr != lfu.tail {
		ptr = ptr.next
	}
	if ptr == lfu.tail {
		return nil
	}
	return ptr
}

func (lfu *LFUCache) removeNode(n *Node) {
	nowPre := n.pre
	nowNext := n.next
	nowPre.next = nowNext
	nowNext.pre = nowPre
}

func (lfu *LFUCache) addNode(n *Node) {
	// find new position
	newPtr := lfu.head
	for newPtr.next != lfu.tail && newPtr.next.freq > n.freq {
		newPtr = newPtr.next
	}
	// add after new position
	newNext := newPtr.next
	newPtr.next = n
	n.pre = newPtr
	n.next = newNext
	newNext.pre = n
}

func (lfu *LFUCache) Print() {
	fmt.Printf("kv=%v\n", lfu.kv)
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	lfu := NewLFU(2)
	lfu.Put(1, 1) // cache=[1,_], cnt(1)=1
	lfu.Print()
	lfu.Put(2, 2) // cache=[2,1], cnt(2)=1, cnt(1)=1
	lfu.Print()
	fmt.Println(lfu.Get(1)) // 返回 1
	lfu.Print()
	// cache=[1,2], cnt(2)=1, cnt(1)=2
	lfu.Put(3, 3) // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
	lfu.Print()
	// cache=[3,1], cnt(3)=1, cnt(1)=2
	fmt.Println(lfu.Get(2)) // 返回 -1（未找到）
	lfu.Print()
	fmt.Println(lfu.Get(3)) // 返回 3
	lfu.Print()
	// cache=[3,1], cnt(3)=2, cnt(1)=2
	lfu.Put(4, 4) // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
	lfu.Print()
	// cache=[4,3], cnt(4)=1, cnt(3)=2
	fmt.Println(lfu.Get(1)) // 返回 -1（未找到）
	lfu.Print()
	fmt.Println(lfu.Get(3)) // 返回 3
	lfu.Print()
	// cache=[3,4], cnt(4)=1, cnt(3)=3
	fmt.Println(lfu.Get(4)) // 返回 4
	lfu.Print()
	// cache=[3,4], cnt(4)=2, cnt(3)=3
}
