package binarysearchtree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_binarySearchTree_Delete(t1 *testing.T) {
	t1.Run("delete root", func(t *testing.T) {
		a := assert.New(t)
		y := NewBinarySearchTree(10)

		// delete
		r := y.Search(10)
		y.Delete(r)

		a.Equal([]int64{}, y.InorderTreeWalk())
		y.Insert(10)
		a.Equal([]int64{10}, y.InorderTreeWalk())
	})
	t1.Run("delete node having no child", func(t *testing.T) {
		a := assert.New(t)
		y := NewBinarySearchTree(10)
		y.Insert(15)

		// delete
		r := y.Search(15)
		y.Delete(r)

		a.Equal([]int64{10}, y.InorderTreeWalk())
		y.Insert(15)
		a.Equal([]int64{10, 15}, y.InorderTreeWalk())
	})

	t1.Run("delete node having left child", func(t *testing.T) {
		a := assert.New(t)
		y := NewBinarySearchTree(10)
		y.Insert(15)
		y.Insert(12)

		// delete
		r := y.Search(15)
		y.Delete(r)

		a.Equal([]int64{10, 12}, y.InorderTreeWalk())
		y.Insert(15)
		a.Equal([]int64{10, 12, 15}, y.InorderTreeWalk())
	})

	t1.Run("delete node having right child", func(t *testing.T) {
		a := assert.New(t)
		y := NewBinarySearchTree(10)
		y.Insert(15)
		y.Insert(17)

		r := y.Search(15)
		y.Delete(r)

		a.Equal([]int64{10, 17}, y.InorderTreeWalk())
		y.Insert(15)
		a.Equal([]int64{10, 15, 17}, y.InorderTreeWalk())
	})

	t1.Run("delete node having both side", func(t *testing.T) {
		a := assert.New(t)
		y := NewBinarySearchTree(10)
		y.Insert(15)
		y.Insert(12)
		y.Insert(17)

		r := y.Search(15)
		y.Delete(r)

		a.Equal([]int64{10, 12, 17}, y.InorderTreeWalk())
		y.Insert(15)
		a.Equal([]int64{10, 12, 15, 17}, y.InorderTreeWalk())
	})
}

func Test_binarySearchTree_InorderTreeWalk(t1 *testing.T) {
	a := assert.New(t1)
	y := NewBinarySearchTree(10)
	y.Insert(15)
	y.Insert(12)
	y.Insert(17)
	a.Equal([]int64{10, 12, 15, 17}, y.InorderTreeWalk())

	y.Insert(11)
	a.Equal([]int64{10, 11, 12, 15, 17}, y.InorderTreeWalk())
	y.Insert(1)
	a.Equal([]int64{1, 10, 11, 12, 15, 17}, y.InorderTreeWalk())
}

func Test_binarySearchTree_Insert(t1 *testing.T) {
	a := assert.New(t1)
	y := NewBinarySearchTree(10)
	a.Equal([]int64{10}, y.InorderTreeWalk())
	y.Insert(15)
	a.Equal([]int64{10, 15}, y.InorderTreeWalk())
}

func Test_binarySearchTree_Maximum(t1 *testing.T) {
	a := assert.New(t1)
	y := NewBinarySearchTree(10)
	a.Equal([]int64{10}, y.InorderTreeWalk())
	y.Insert(15)

	a.Equal(15, y.Maximum().Value)
	y.Insert(14)
	a.Equal(15, y.Maximum().Value)
	y.Insert(7)
	a.Equal(15, y.Maximum().Value)
	y.Insert(16)
	a.Equal(16, y.Maximum().Value)
	y.Insert(100)
	a.Equal(100, y.Maximum().Value)
}

func Test_binarySearchTree_Minimum(t1 *testing.T) {
	a := assert.New(t1)
	y := NewBinarySearchTree(10)
	a.Equal([]int64{10}, y.InorderTreeWalk())
	y.Insert(15)

	a.Equal(10, y.Minimum().Value)
	y.Insert(14)
	a.Equal(10, y.Minimum().Value)
	y.Insert(7)
	a.Equal(7, y.Minimum().Value)
	y.Insert(16)
	a.Equal(7, y.Minimum().Value)
	y.Insert(100)
	a.Equal(7, y.Minimum().Value)
}

func Test_binarySearchTree_Predecessor(t1 *testing.T) {
	a := assert.New(t1)
	y := NewBinarySearchTree(10)
	y.Insert(5)
	a.Equal(5, y.Predecessor().Value)
	y.Insert(7)
	a.Equal(7, y.Predecessor().Value)
	y.Insert(9)
	a.Equal(9, y.Predecessor().Value)
	y.Insert(8)
	a.Equal(9, y.Predecessor().Value)
	y.Insert(1)
	a.Equal(9, y.Predecessor().Value)
	y.Insert(6)
	a.Equal(9, y.Predecessor().Value)
	y.Insert(13)
	a.Equal(9, y.Predecessor().Value)

	z := y.Search(9)
	a.Equal(8, z.Predecessor().Value)
	z = y.Search(8)
	a.Equal(7, z.Predecessor().Value)
	z = y.Search(7)
	a.Equal(6, z.Predecessor().Value)
	z = y.Search(6)
	a.Equal(5, z.Predecessor().Value)
	z = y.Search(5)
	a.Equal(1, z.Predecessor().Value)
	z = y.Search(1)
	a.Equal(nil, z.Predecessor())
}

func Test_binarySearchTree_Search(t1 *testing.T) {
	a := assert.New(t1)
	y := NewBinarySearchTree(10)
	elms := []int64{5, 7, 1, 6, 8, 11, 12}
	for _, e := range elms {
		y.Insert(e)
	}

	for _, e := range elms {
		z := y.Search(e)
		a.NotNil(z)
		a.Equal(e, z.Value)
	}
}

func Test_binarySearchTree_Successor(t1 *testing.T) {
	a := assert.New(t1)
	y := NewBinarySearchTree(10)
	y.Insert(12)
	a.Equal(12, y.Successor().Value)
	y.Insert(5)
	a.Equal(12, y.Successor().Value)
	y.Insert(7)
	a.Equal(12, y.Successor().Value)
	y.Insert(9)
	a.Equal(12, y.Successor().Value)
	y.Insert(8)
	a.Equal(12, y.Successor().Value)
	y.Insert(1)
	a.Equal(12, y.Successor().Value)
	y.Insert(6)
	a.Equal(12, y.Successor().Value)
	y.Insert(13)
	a.Equal(12, y.Successor().Value)
	y.Insert(11)
	a.Equal(11, y.Successor().Value)

	elms := []int64{1, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	for ind, e := range elms {
		z := y.Search(e)
		if ind == len(elms)-1 {
			a.Nil(z)
		} else {
			a.NotNil(z)
			a.Equal(elms[ind+1], z.Value)
		}
	}
}
