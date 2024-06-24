package trie

import "errors"

type Node[T any] struct {
	value T
	next  map[string]*Node[T]
	isEnd bool
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		value: value,
		next:  make(map[string]*Node[T]),
	}
}

type Trie[T any] struct {
	root *Node[T]
}

func NewTrie[T any]() Trie[T] {
	return Trie[T]{
		root: &Node[T]{
			next: make(map[string]*Node[T]),
		},
	}
}

func (t *Trie[T]) Insert(keys []string, value T) *Node[T] {
	current := t.root

	for _, key := range keys {
		if len(key) > 1 && key[0] == '{' && key[len(key)-1] == '}' {
			key = "*"
		}

		if _, exist := current.next[key]; !exist {
			current.next[key] = NewNode[T](value)
		}
		current = current.next[key]
	}

	current.isEnd = true

	return current
}

func (t *Trie[T]) Search(keys []string) (T, error) {
	var (
		current   = t.root
		zeroValue T
	)

	for _, key := range keys {
		if next, exist := current.next[key]; exist {
			current = next
		} else if next, exist := current.next["*"]; exist {
			current = next
		} else {
			return zeroValue, errors.New("key does not exist")
		}
	}

	if current.isEnd {
		return current.value, nil
	}

	return zeroValue, errors.New("key does not exist")
}
