package services

import (
	"fmt"
	"sync"
)

type Queue struct {
	list chan int
	wg   sync.WaitGroup
	Lock sync.Mutex
}

func NewQueue(size int64) *Queue {

	return &Queue{
		list: make(chan int, size),
	}
}

func (q *Queue) InputQueue(value int) {
	q.wg.Add(1)

	q.list <- value

}

func (q *Queue) outputQueue() (int, *int) {
	q.wg.Add(1)
	n := <-q.list
	return n, &n
}

func (q *Queue) PrintQueue() {
	q.wg.Add(1)
	val, v := q.outputQueue()
	fmt.Printf("v: %v , value : %d \n", v, val)
	q.wg.Done()
}

func (q *Queue) WaitQueue() {

	q.wg.Wait()
	close(q.list)
}
