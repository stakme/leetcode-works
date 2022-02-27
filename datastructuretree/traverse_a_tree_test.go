package datastructuretree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PreorderTraversal(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    *TreeNode
		expected []int
	}{
		{
			name: "sample test",
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "sample test2",
			input: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			expected: []int{3, 1, 2},
		},
		{
			name:     "empty",
			input:    nil,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, preorderTraversal(tt.input))
		})
	}
}

func TestInorderTraversal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *TreeNode
		expected []int
	}{
		{
			name: "sample test",
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: []int{1, 3, 2},
		},
		{
			name: "sample test2",
			input: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			expected: []int{1, 3, 2},
		},
		{
			name:     "empty",
			input:    nil,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, inorderTraversal(tt.input))
		})
	}
}

func TestPostorderTraversal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *TreeNode
		expected []int
	}{
		{
			name: "sample test",
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: []int{3, 2, 1},
		},
		{
			name: "sample test2",
			input: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			expected: []int{1, 2, 3},
		},
		{
			name:     "empty",
			input:    nil,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, postorderTraversal(tt.input))
		})
	}
}

func TestLevelOrder(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *TreeNode
		expected [][]int
	}{
		{
			name: "Example 1",
			input: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name:     "Example 2",
			input:    &TreeNode{Val: 1},
			expected: [][]int{{1}},
		},
		{
			name:     "empty",
			input:    nil,
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, levelOrder(tt.input))
		})
	}
}
