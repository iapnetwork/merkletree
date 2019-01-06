package merkletree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildEmptyTree(t *testing.T) {
	tree := Tree{}
	leaves := []*Node{}

	assert.Error(t, tree.BuildTree(leaves))
}

func TestBuildTreeSingleNode(t *testing.T) {
	tree := Tree{}
	leaves := []*Node{
		NewNode([]byte("1")),
	}

	assert.NoError(t, tree.BuildTree(leaves))

	assert.NotNil(t, tree.Root)
	assert.Equal(t, tree.Root.Hash, leaves[0].Hash)
}

func TestBuildTree(t *testing.T) {
	tree := Tree{}
	leaves := []*Node{
		NewNode([]byte("1")),
		NewNode([]byte("2")),
	}

	rootHash := computeHash(append(leaves[0].Hash, leaves[1].Hash...))

	assert.NoError(t, tree.BuildTree(leaves))
	assert.Equal(t, tree.Root.Hash, rootHash)
}

func TestBuildTreeOddNode(t *testing.T) {
	tree := Tree{}
	leaves := []*Node{
		NewNode([]byte("1")),
		NewNode([]byte("2")),
		NewNode([]byte("3")),
	}

	firstRootHash := computeHash(append(leaves[0].Hash, leaves[1].Hash...))
	rootHash := computeHash(append(firstRootHash, leaves[2].Hash...))

	assert.NoError(t, tree.BuildTree(leaves))
	assert.Equal(t, tree.Root.Hash, rootHash)
}

func TestAppendLeaf(t *testing.T) {
	tree := Tree{}
	tree.AppendLeaf(NewNode([]byte("1")))
	tree.AppendLeaf(NewNode([]byte("2")))

	for i, leaf := range tree.Leaves {
		assert.Equal(t, tree.Leaves[i], leaf)
	}
}

func TestFindLeaf(t *testing.T) {
	target := NewNode([]byte("2"))

	tree := Tree{}
	tree.AppendLeaf(NewNode([]byte("1")))
	tree.AppendLeaf(target)
	leaf := tree.FindLeaf(target.Hash)

	assert.Equal(t, leaf.Hash, target.Hash)
}

func TestFindLeafNotFound(t *testing.T) {
	target := computeHash([]byte("3"))

	tree := Tree{}
	tree.AppendLeaf(NewNode([]byte("1")))
	tree.AppendLeaf(NewNode([]byte("2")))

	leaf := tree.FindLeaf(target)

	assert.Nil(t, leaf)
}