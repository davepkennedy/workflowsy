package internal_test

import (
	"testing"

	"workflowsy/internal"

	"github.com/stretchr/testify/assert"
)

func TestNewStackIsEmpty(t *testing.T) {
	stack := internal.NewStack[int]()
	assert.Zero(t, stack.Size())
}

func TestItemCanBeAdded(t *testing.T) {
	stack := internal.NewStack[int]()
	stack.Push(42)
	
	val, ok := stack.Pop()
	assert.True(t, ok)
	assert.Equal(t, 42, val)
	assert.Zero(t, stack.Size())
}

func TestItemCanBeRemoved(t *testing.T) {
	stack := internal.NewStack[int]()
	stack.Push(1)
	stack.Pop()
	assert.Zero(t, stack.Size())
}

func TestPopFromEmptyStack(t *testing.T) {
	stack := internal.NewStack[int]()
	val, ok := stack.Pop()
	assert.False(t, ok)
	assert.Zero(t, val)
}

func TestItemsAreInLifoOrder(t *testing.T) {
	stack := internal.NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	val, ok := stack.Pop()
	assert.True(t, ok)
	assert.Equal(t, 3, val)
	val, ok = stack.Pop()
	assert.True(t, ok)
	assert.Equal(t, 2, val)
	val, ok = stack.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, val)
}