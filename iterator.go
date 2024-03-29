// Copyright 2015, Hu Keping. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rbtree

// Iterator is the function of iteration entity which would be
// used by those functions like `Ascend`, `Dscend`, etc.
//
// A typical Iterator with Print :
// func loop_with_print(item rbtree.Item) bool {
//         i, ok := item.(XXX)
//         if !ok {
//                 return false
//         }
//         fmt.Printf("%+v\n", i)
//         return true
// }
type Iterator func(i Item) bool

// Ascend will call iterator once for each element greater or equal than pivot
// in ascending order. It will stop whenever the iterator returns false.
func (t *Rbtree) Ascend(pivot Item, iterator Iterator) {
	t.ascend(t.root, pivot, iterator)
}

func (t *Rbtree) ascend(x *Node, pivot Item, iterator Iterator) bool {
	if x == t.NIL {
		return true
	}

	if !less(x.Item, pivot) {
		if !t.ascend(x.Left, pivot, iterator) {
			return false
		}
		if !iterator(x.Item) {
			return false
		}
	}

	return t.ascend(x.Right, pivot, iterator)
}

// Descend will call iterator once for each element less or equal than pivot
// in descending order. It will stop whenever the iterator returns false.
func (t *Rbtree) Descend(pivot Item, iterator Iterator) {
	t.descend(t.root, pivot, iterator)
}

func (t *Rbtree) descend(x *Node, pivot Item, iterator Iterator) bool {
	if x == t.NIL {
		return true
	}

	if !less(pivot, x.Item) {
		if !t.descend(x.Right, pivot, iterator) {
			return false
		}
		if !iterator(x.Item) {
			return false
		}
	}

	return t.descend(x.Left, pivot, iterator)
}

// AscendRange will call iterator once for elements greater or equal than @ge
// and less than @lt, which means the range would be [ge, lt).
// It will stop whenever the iterator returns false.
func (t *Rbtree) AscendRange(ge, lt Item, iterator Iterator) {
	t.ascendRange(t.root, ge, lt, iterator)
}

func (t *Rbtree) ascendRange(x *Node, inf, sup Item, iterator Iterator) bool {
	if x == t.NIL {
		return true
	}

	if !less(x.Item, sup) {
		return t.ascendRange(x.Left, inf, sup, iterator)
	}
	if less(x.Item, inf) {
		return t.ascendRange(x.Right, inf, sup, iterator)
	}

	if !t.ascendRange(x.Left, inf, sup, iterator) {
		return false
	}
	if !iterator(x.Item) {
		return false
	}
	return t.ascendRange(x.Right, inf, sup, iterator)
}

// SliceAscend will recursively go through Nodes and return a slice of Nodes by ascending order.
func (t *Rbtree) SliceAscend() []*Node {
	result := make([]*Node, t.count)
	count := 0

	t.dfsLeft(t.root, &count, result)
	return result
}

func (t *Rbtree) dfsLeft(x *Node, count *int, result []*Node) {
	if x == t.NIL {
		return
	}
	t.dfsLeft(x.Left, count, result)
	result[*count] = x
	*count++
	t.dfsLeft(x.Right, count, result)
}

// SliceDescend will recursively go through Nodes and return a slice of Nodes by descending order.
func (t *Rbtree) SliceDescend() []*Node {
	result := make([]*Node, t.count)
	count := 0

	t.dfsRight(t.root, &count, result)
	return result
}

func (t *Rbtree) dfsRight(x *Node, count *int, result []*Node) {
	if x == t.NIL {
		return
	}
	result[*count] = x
	t.dfsRight(x.Right, count, result)
	result[*count] = x
	*count++
	t.dfsRight(x.Left, count, result)
}

// SliceAscendFirstN will recursively go through the first N Nodes and return a slice of Nodes by ascending order.
func (t *Rbtree) SliceAscendFirstN(length int) []*Node {
	n := uint(length)
	if n > t.count {
		n = t.count
	}
	result := make([]*Node, n)
	count := 0

	t.dfsLeftFirstN(t.root, &count, length, result)
	return result
}

func (t *Rbtree) dfsLeftFirstN(x *Node, count *int, length int, result []*Node) {
	if x == t.NIL {
		return
	}
	t.dfsLeftFirstN(x.Left, count, length, result)
	if *count == length {
		return
	}
	result[*count] = x
	*count++
	if *count == length {
		return
	}
	t.dfsLeftFirstN(x.Right, count, length, result)
}

// SliceDescend will recursively go through the first N Nodes and return a slice of Nodes by descending order.
func (t *Rbtree) SliceDescendFirstN(length int) []*Node {
	n := uint(length)
	if n > t.count {
		n = t.count
	}
	result := make([]*Node, n)
	count := 0

	t.dfsRightFirstN(t.root, &count, length, result)
	return result
}

func (t *Rbtree) dfsRightFirstN(x *Node, count *int, length int, result []*Node) {
	if x == t.NIL {
		return
	}
	result[*count] = x
	t.dfsRightFirstN(x.Right, count, length, result)
	if *count == length {
		return
	}
	result[*count] = x
	*count++
	if *count == length {
		return
	}
	t.dfsRightFirstN(x.Left, count, length, result)
}
