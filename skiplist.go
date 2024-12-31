package skiplist

import (
	"math"
	"math/rand/v2"

	"golang.org/x/exp/constraints"
)


const (
	maxLevel = 16
)

type Skiplist[T constraints.Ordered] struct {
	maxLevel int
	head     *Node[T]
	level    int
}

type Node[T constraints.Ordered] struct {
	nexts []*Node[T]
	val   T
}

func NewNode[T constraints.Ordered](val T, level int) *Node[T] {
	return &Node[T]{
		val:   val,
		nexts: make([]*Node[T], level),
	}
}

func Constructor[T constraints.Ordered]() Skiplist[T] {

	return Skiplist[T]{
		maxLevel: maxLevel,
		head:     NewNode(MinValue[T](), maxLevel),
		level:    1,
	}
}

func (list *Skiplist[T]) Search(target T) bool {
	curr := list.head
	for level := list.level - 1; level >= 0; level-- {
		curr = list.findNext(curr, level, target)
		if curr.nexts[level] != nil && curr.nexts[level].val == target {
			return true
		}
	}
	return false
}

func (list *Skiplist[T]) Add(key T) {
	newLevel := list.randLevel()
	if newLevel > list.level {
		list.level = newLevel
	}
	node := NewNode(key, newLevel)
	curr := list.head
	for level := list.level - 1; level >= 0; level-- {
		curr = list.findNext(curr, level, key)
		if level < newLevel {
			node.nexts[level] = curr.nexts[level]
			curr.nexts[level] = node
		}
	}
}

func (list *Skiplist[T]) Erase(key T) bool {
	ok := false
	curr := list.head
	for level := list.level - 1; level >= 0; level-- {
		curr = list.findNext(curr, level, key)
		if curr.nexts[level] != nil && curr.nexts[level].val == key {
			curr.nexts[level] = curr.nexts[level].nexts[level]
			ok = true
		}
	}
	for list.level > 1 && list.head.nexts[list.level-1] == nil {
		list.level--
	}
	return ok
}

func (list *Skiplist[T]) randLevel() int {
	const prob = 1 << 30
	i := 1

	for ; i < list.maxLevel; i++ {
		if rand.Int32() < prob {
			break
		}
	}
	return i
}

func (list *Skiplist[T]) findNext(node *Node[T], level int, target T) *Node[T] {
	curr := node
	for curr.nexts[level] != nil && curr.nexts[level].val < target {
		curr = curr.nexts[level]
	}
	return curr
}

func MinValue[T constraints.Ordered]() T {
	var zero T
	switch any(zero).(type) {
	case int:
		return any(math.MinInt).(T)
	case int8:
		return any(math.MinInt8).(T)
	case int16:
		return any(math.MinInt16).(T)
	case int32:
		return any(math.MinInt32).(T)
	case int64:
		return any(math.MinInt64).(T)
	case uint:
		return any(0).(T)
	case uint8:
		return any(0).(T)
	case uint16:
		return any(0).(T)
	case uint32:
		return any(0).(T)
	case uint64:
		return any(0).(T)
	case float32:
		return any(-math.MaxFloat32).(T)
	case float64:
		return any(-math.MaxFloat64).(T)
	case string:
		return any("").(T)
	default:
		panic("unsupported type for MinValue")
	}
}
