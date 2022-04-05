package list

type singleLinkNode[T any] struct {
	data T
	next *singleLinkNode[T]
}

type LinkedList[T any] struct {
	count int
	head  *singleLinkNode[T]
	tail  *singleLinkNode[T]
}

func NewLinkedList[T any](elements ...T) *LinkedList[T] {
	ll := &LinkedList[T]{}
	for _, element := range elements {
		ll.Add(element)
	}
	return ll
}

func (l *LinkedList[T]) Add(element T) {
	n := &singleLinkNode[T]{data: element}
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.count++
}

func (l *LinkedList[T]) Count() int {
	return l.count
}

func (l *LinkedList[T]) Iterator() *iterator[T] {
	return &iterator[T]{next: l.head}
}

type iterator[T any] struct {
	next *singleLinkNode[T]
}

func (iter *iterator[T]) Next() (bool, T) {
	curr := iter.next
	if curr == nil {
		var empty T
		return false, empty
	}
	iter.next = curr.next
	return true, curr.data
}
