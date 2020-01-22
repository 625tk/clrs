package binarysearchtree

type BinarySearchTree struct {
	Parent *BinarySearchTree
	Value  int64
	Left   *BinarySearchTree
	Right  *BinarySearchTree
}

func NewBinarySearchTree(v int64) *BinarySearchTree {
	return &BinarySearchTree{
		Parent: nil,
		Value:  v,
		Left:   nil,
		Right:  nil,
	}

}

func (t *BinarySearchTree) InorderTreeWalk() (sortedValues []int64) {
	// implement me
	return []int64{}
}

func (t *BinarySearchTree) Search(x int64) *BinarySearchTree {
	// implement me
	return nil
}

func (t *BinarySearchTree) Minimum() *BinarySearchTree {
	// implement me
	return nil
}

func (t *BinarySearchTree) Maximum() *BinarySearchTree {
	// implement me
	return nil
}

func (t *BinarySearchTree) Successor() *BinarySearchTree {
	// implement me
	// 1つ後の値なければnil
	return nil
}

func (t *BinarySearchTree) Predecessor() *BinarySearchTree {
	// implement me
	// 1つ前の値なければnil
	return nil
}

func (t *BinarySearchTree) Insert(x int64) {
	// implement me
	return
}

func (t *BinarySearchTree) Delete(x *BinarySearchTree) {
	// implement me
	return
}
