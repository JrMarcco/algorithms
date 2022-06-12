package leet_code

import "testing"

type element struct {
	Val  int
	Prev *element
	Next *element
}

type MyCircularDeque struct {
	cap   int
	count int
	Head  *element
	Tail  *element
}

func constructor(k int) MyCircularDeque {
	head := &element{}
	tail := &element{}

	head.Next = tail
	tail.Prev = head

	return MyCircularDeque{
		cap:   k,
		count: 0,
		Head:  head,
		Tail:  tail,
	}
}

func (mcd *MyCircularDeque) InsertFront(value int) bool {
	if mcd.IsFull() {
		return false
	}

	newElement := &element{
		Val:  value,
		Prev: mcd.Head,
		Next: mcd.Head.Next,
	}

	mcd.Head.Next.Prev = newElement
	mcd.Head.Next = newElement

	mcd.count++
	return true
}

func (mcd *MyCircularDeque) InsertLast(value int) bool {
	if mcd.IsFull() {
		return false
	}

	newElement := &element{
		Val:  value,
		Prev: mcd.Tail.Prev,
		Next: mcd.Tail,
	}

	mcd.Tail.Prev.Next = newElement
	mcd.Tail.Prev = newElement
	mcd.count++

	return true
}

func (mcd *MyCircularDeque) DeleteFront() bool {
	if mcd.IsEmpty() {
		return false
	}

	mcd.Head.Next.Next.Prev = mcd.Head
	mcd.Head.Next = mcd.Head.Next.Next

	mcd.count--
	return true
}

func (mcd *MyCircularDeque) DeleteLast() bool {
	if mcd.IsEmpty() {
		return false
	}

	mcd.Tail.Prev.Prev.Next = mcd.Tail
	mcd.Tail.Prev = mcd.Tail.Prev.Prev

	mcd.count--
	return true
}

func (mcd *MyCircularDeque) GetFront() int {
	if mcd.IsEmpty() {
		return -1
	}

	return mcd.Head.Next.Val
}

func (mcd *MyCircularDeque) GetRear() int {
	if mcd.IsEmpty() {
		return -1
	}

	return mcd.Tail.Prev.Val
}

func (mcd *MyCircularDeque) IsEmpty() bool {
	return mcd.count == 0
}

func (mcd *MyCircularDeque) IsFull() bool {
	return mcd.cap == mcd.count
}

func TestDeque(t *testing.T) {
	circularDeque := constructor(3)

	t.Log(circularDeque.InsertLast(1))
	t.Log(circularDeque.InsertLast(2))
	t.Log(circularDeque.InsertFront(3))
	t.Log(circularDeque.InsertFront(4))
	t.Log(circularDeque.GetRear())
	t.Log(circularDeque.IsFull())
	t.Log(circularDeque.DeleteLast())
	t.Log(circularDeque.InsertFront(4))
	t.Log(circularDeque.GetFront())
}
