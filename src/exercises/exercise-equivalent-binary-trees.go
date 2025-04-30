/* Exercise: Equivalent Binary Trees
1. Implement the Walk function.

2. Test the Walk function.

The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.

Create a new channel ch and kick off the walker:

go Walk(tree.New(1), ch)
Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

4. Test the Same function.

Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

The documentation for Tree can be found here. */

package exercises

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walker func(t *tree.Tree)

	walker = func(t *tree.Tree) {
	if t == nil {
		return
	}
	//visit left node
	walker(t.Left)
	//visit current node
	ch <- t.Value
	//visit right node
	walker(t.Right)
	}
	walker(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 && !ok2 {
			return true
		}
		if !ok1 || !ok2 {
			return false
		}
		if v1 != v2 {
			return false
		}
	}
}

func Test() {
	same := Same(tree.New(1), tree.New(2))
	fmt.Println(same)
}
