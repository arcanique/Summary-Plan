/*
 * @lc app=leetcode.cn id=622 lang=golang
 *
 * [622] 设计循环队列
 */

// @lc code=start
type MyCircularQueue struct {
	cap  uint
	head uint
	tail uint
	data []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	if k <= 0 {
		k = 0
	}
	return MyCircularQueue{
		cap:  uint(k),
		head: 0,
		tail: 0,
		data: make([]int, k),
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.head-this.tail == this.cap {
		return false
	}

	wrIdx := this.head % this.cap
	this.data[wrIdx] = value
	this.head++

	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if 0 == this.head-this.tail {
		return false
	}

	this.tail++

	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if 0 == this.head-this.tail {
		return -1
	}
	rearIdx := this.tail % this.cap

	return this.data[rearIdx]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if 0 == this.head-this.tail {
		return -1
	}
	frontIdx := (this.head - 1) % this.cap

	return this.data[frontIdx]

}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if 0 == this.head-this.tail {
		return true
	}

	return false
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if this.head-this.tail == this.cap {
		return true
	}

	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end
