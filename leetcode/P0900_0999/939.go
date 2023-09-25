package main

type RecentCounter struct {
	pingQueue []int
}

func Constructor() RecentCounter {
	return RecentCounter{pingQueue: make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.pingQueue = append(this.pingQueue, t)
	i, j := 0, len(this.pingQueue)-1
	for i < j && this.pingQueue[j]-this.pingQueue[i] > 3000 {
		i++
	}
	this.pingQueue = this.pingQueue[i:]
	return len(this.pingQueue)
}
