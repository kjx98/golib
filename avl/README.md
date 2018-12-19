### AVL - AVL tree
#### forked from Yawning Angel (yawning at schwanenlied dot me)

[![GoDoc](https://godoc.org/github.com/kjx98/golib/avl.git?status.svg)](https://godoc.org/github.com/kjx98/golib/avl.git)

A generic Go AVL tree implementation, derived from [Eric Biggers' C code][1],
in the spirt of [the runtime library's containers][2].
Replace cmpFunc with Itemer interface, slower 10% than cmpFunc

Features:

 * Size
 * Insertion
 * Deletion
 * Search
 * In-order traversal (forward only) with an iterator or callback.
 * Non-recursive.

Note:

 * The package itself is free from external dependencies, the unit tests use
   [testify][3].

[1]: https://github.com/ebiggers/avl_tree
[2]: https://golang.org/pkg/container
