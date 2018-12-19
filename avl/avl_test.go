// avl_test.go - AVL tree tests.
//
// To the extent possible under law, Yawning Angel has waived all copyright
// and related or neighboring rights to avl, using the Creative
// Commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package avl

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

type itemT	int
func (k itemT) Less(v interface {}) bool {
	return k < v.(itemT)
}

func TestAVLTree(t *testing.T) {
	Equal := func(a, b interface{}, ss string, args ...interface{}) {
		if !reflect.DeepEqual(a, b) {
			t.Errorf(ss, args...)
		}
	}
	Nil := func(a interface{}, ss string, args ...interface{}) {
		if !reflect.ValueOf(a).IsNil() {
			t.Errorf(ss, args...)
		}
	}

	tree := New()
	Equal(0, tree.Len(), "Len(): empty")
	Nil(tree.First(), "First(): empty")
	Nil(tree.Last(), "Last(): empty")

	iter := tree.Iterator()
	Nil(iter.First(), "Iterator: First(), empty")
	Nil(iter.Next(), "Iterator: Next(), empty")

	// Test insertion.
	const nrEntries = 1024
	insertedMap := make(map[itemT]*Node)
	for len(insertedMap) != nrEntries {
		v := itemT(rand.Int())
		if insertedMap[v] != nil {
			continue
		}
		insertedMap[v] = tree.Insert(v)
		tree.validate(t)
	}
	Equal(nrEntries, tree.Len(), "Len(): After insertion")
	tree.validate(t)

	// Ensure that all entries can be found.
	for k, v := range insertedMap {
		Equal(v, tree.Find(k), "Find(): %v", k)
		Equal(k, v.Value, "Find(): %v Value", k)
	}

	// Test the iterators.
	fwdInOrder := make([]int, 0, nrEntries)
	for k := range insertedMap {
		fwdInOrder = append(fwdInOrder, int(k))
	}
	sort.Ints(fwdInOrder)
	Equal(itemT(fwdInOrder[0]), tree.First().Value, "First(), full")
	Equal(itemT(fwdInOrder[nrEntries-1]), tree.Last().Value, "Last(), full")

	iter = tree.Iterator()
	visited := 0
	for node := iter.First(); node != nil; node = iter.Next() {
		v, idx := node.Value, visited
		Equal(itemT(fwdInOrder[visited]), v, "Iterator: Forward[%v]", idx)
		Equal(node, iter.Get(), "Iterator: Forward[%v]: Get()", idx)
		visited++
	}
	Equal(nrEntries, visited, "Iterator: Forward: Visited")

	// Test the ForEach.
	forEachValues := make([]int, 0, nrEntries)
	forEachFn := func(n *Node) bool {
		forEachValues = append(forEachValues, int(n.Value.(itemT)))
		return true
	}
	tree.ForEach(forEachFn)
	Equal(fwdInOrder, forEachValues, "ForEach: Forward")

	// Test removal.
	for i, idx := range rand.Perm(nrEntries) { // In random order.
		v := itemT(fwdInOrder[idx])
		node := tree.Find(v)
		Equal(v, node.Value, "Find(): %v (Pre-remove)", v)

		tree.Remove(node)
		Equal(nrEntries-(i+1), tree.Len(), "Len(): %v (Post-remove)", v)
		tree.validate(t)

		node = tree.Find(v)
		Nil(node, "Find(): %v (Post-remove)", v)
	}
	Equal(0, tree.Len(), "Len(): After removal")
	Nil(tree.First(), "First(): After removal")
	Nil(tree.Last(), "Last(): After removal")

	// Refill the tree.
	for _, v := range fwdInOrder {
		tree.Insert(itemT(v))
	}

	// Test that removing the node doesn't break the iterator.
	iter = tree.Iterator()
	visited = 0
	for node := iter.Get(); node != nil; node = iter.Next() { // Omit calling First().
		v, idx := node.Value, visited
		Equal(itemT(fwdInOrder[idx]), v, "Iterator: Forward[%v] (Pre-Remove)", idx)
		Equal(itemT(fwdInOrder[idx]), tree.First().Value, "First() (Iterator, remove)")
		visited++

		tree.Remove(node)
		tree.validate(t)
	}
	Equal(0, tree.Len(), "Len(): After iterating removal")
}

func (t *Tree) validate(te *testing.T) {
	checkInvariants(te, t.root, nil)
}

func checkInvariants(te *testing.T, node, parent *Node) int {
	Equal := func(a, b interface{}) {
		if !reflect.DeepEqual(a, b) {
			te.Error(a, "notEqual", b)
		}
	}
	if node == nil {
		return 0
	}

	// Validate the parent pointer.
	Equal(parent, node.parent)

	// Validate that the balance factor is -1, 0, 1.
	switch node.balance {
	case -1, 0, 1:
	default:
		te.Error(node.balance)
	}

	// Recursively derive the height of the left and right sub-trees.
	lHeight := checkInvariants(te, node.left, node)
	rHeight := checkInvariants(te, node.right, node)

	// Validate the AVL invariant and the balance factor.
	Equal(int(node.balance), rHeight-lHeight)
	if lHeight > rHeight {
		return lHeight + 1
	}
	return rHeight + 1
}

func BenchmarkAVLInsert(b *testing.B) {
	b.StopTimer()
	tree := New()
	for i := 0; i < 1e6; i++ {
		tree.Insert(itemT(i))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v := (rand.Int() % 1e6) + 2e6
		tree.Insert(itemT(v))
	}
}

func BenchmarkAVLFind(b *testing.B) {
	b.StopTimer()
	tree := New()
	for i := 0; i < 1e6; i++ {
		tree.Insert(itemT(i))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v := (rand.Int() % 1e6)
		tree.Find(itemT(v))
	}
}
