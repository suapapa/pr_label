package main

import "errors"

const (
	DefaultOrderQLen = 100
)

type OrderQ struct {
	q               []*Order
	pushIdx, popIdx int
}

func NewOrderQ(qSize int) *OrderQ {
	return &OrderQ{
		q: make([]*Order, qSize),
	}
}

func (q *OrderQ) IsEmpty() bool {
	return q.pushIdx == q.popIdx
}

func (q *OrderQ) IsFull() bool {
	if q.popIdx < q.pushIdx {
		return (q.pushIdx - q.popIdx) == len(q.q)-1
	} else if q.popIdx > q.pushIdx {
		return q.pushIdx+1 == q.popIdx
	}
	return false // empty
}

func (q *OrderQ) Push(o *Order) error {
	if q.IsFull() {
		return errors.New("full Q")
	}

	q.q[q.pushIdx] = o
	q.pushIdx += 1
	if q.pushIdx >= len(q.q) {
		q.pushIdx = 0
	}

	return nil
}

func (q *OrderQ) Pop() (*Order, error) {
	if q.IsEmpty() {
		return nil, errors.New("empty Q")
	}

	ret := q.q[q.popIdx]
	q.popIdx -= 1
	if q.popIdx < 0 {
		q.popIdx = len(q.q) - 1
	}

	return ret, nil
}
