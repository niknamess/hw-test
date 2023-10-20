package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	List // Remove me after realization.
	// Place your code here.
	size int
	head *ListItem
	tail *ListItem
}

func (l list) Len() int {
	return l.size
}

func (l list) Front() *ListItem {
	return l.head
}

func (l list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.head == nil {
		l.head = &ListItem{v, nil, nil}
		l.size++
		l.tail = l.head
	} else {
		prevHead := l.head
		l.head = &ListItem{v, prevHead, nil}
		prevHead.Prev = l.head
		l.size++
	}
	return l.head
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.tail == nil {
		l.tail = &ListItem{v, nil, nil}
		l.size++
		l.head = l.tail
	} else {
		prevTail := l.tail
		l.tail = &ListItem{v, nil, prevTail}
		prevTail.Next = l.tail
		l.size++
	}
	return l.tail
}

func (l *list) Remove(i *ListItem) {
	prevItem := i.Prev
	nextItem := i.Next
	switch {
	case prevItem == nil && nextItem == nil:
		l.head = nil
		l.tail = nil
	case prevItem == nil:
		l.head = nextItem
		nextItem.Prev = nil
	case nextItem == nil:
		l.tail = prevItem
		prevItem.Next = nil
	default:
		prevItem.Next = nextItem
		nextItem.Prev = prevItem
	}
	l.size--
}

func (l *list) MoveToFront(item *ListItem) {
	l.Remove(item)
	if l.head == nil {
		l.head = item
		item.Next = nil
		item.Prev = nil
		l.tail = l.head
		l.size++
	} else {
		prevHead := l.head
		l.head = item
		item.Next = prevHead
		item.Prev = nil
		prevHead.Prev = l.head
		l.size++
	}
}

func NewList() List {
	return new(list)
}
