package data_structure

import "fmt"

type bst_node struct {
	value int16
	right *bst_node
	left  *bst_node
}

type bst interface {
	Insert(value int16)           // Insert new value to Binary Search Tree
	Search(value int16) *bst_node // Finding value from Binary Search Tree
	_insert_recursive(parent *bst_node, newNode *bst_node)
	_search_recursive(parent *bst_node, value int16) *bst_node
}

type bstImpl struct {
	root *bst_node
}

func newBST() bst {
	return &bstImpl{}
}

func (b *bstImpl) Insert(value int16) {
	newNode := &bst_node{
		value: value,
	}
	if b.root == nil {
		fmt.Println(value, "Inserted as Root")
		b.root = newNode
	} else {
		b._insert_recursive(b.root, newNode)
	}
}
func (b *bstImpl) _insert_recursive(parent *bst_node, newNode *bst_node) {
	switch {
	case newNode.value > parent.value:
		if parent.right == nil {
			fmt.Println(newNode.value, "Was inserted to the right of", parent.value)
			parent.right = newNode
			return
		}
		b._insert_recursive(parent.right, newNode)
	case newNode.value < parent.value:
		fmt.Println(newNode.value, "Was inserted to the left of", parent.value)
		if parent.left == nil {
			parent.left = newNode
			return
		}
		b._insert_recursive(parent.left, newNode)
	}
}

func (b *bstImpl) Search(value int16) *bst_node {
	if b.root.value == value {
		return b.root
	} else {
		return b._search_recursive(b.root, value)
	}
}
func (b *bstImpl) _search_recursive(parent *bst_node, value int16) *bst_node {
	if parent == nil {
		return parent
	}
	if parent.value == value {
		return parent
	}
	if value < parent.value {
		fmt.Printf("%v is lesser than %v, searching to the left\n", value, parent.value)
		return b._search_recursive(parent.left, value)
	}
	fmt.Printf("%v is greater than %v, searching to the right\n", value, parent.value)
	return b._search_recursive(parent.right, value)
}

/*Binary Search Tree*/
func BST() {
	myBST := newBST()
	fmt.Println("Inserting [15, 10, 20, 25, 16]")
	myBST.Insert(15)
	myBST.Insert(10)
	myBST.Insert(20)
	myBST.Insert(25)
	myBST.Insert(16)

	/* The Tree will look like this:
	      15
	     /  \
	   10   20
	       /  \
	      16  25
	*/

	fmt.Println("")
	x := myBST.Search(16)
	if x != nil {
		fmt.Println(x.value)
	} else {
		fmt.Println("not found!")
	}
}
