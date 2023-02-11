package stack_test

import (
	"testing"
	"udemy/go/unit_test/stack"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	s := stack.New()
	t.Run("A stack is empty on construction", func(t *testing.T) {
		assert.True(t, s.IsEmpty())
	})

	t.Run("A stack has size 0 on construction", func(t *testing.T) {
		assert.True(t, s.Size() == 0)
	})
}

func TestInsert(t *testing.T) {
	t.Run("after n push to empty stack (n>0), the stack is non empty and it size equal n", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)

		assert.False(t, s.IsEmpty())
		assert.Equal(t, s.Size(), 2)
	})
	t.Run("if one pushes x then pops, the value popped is x, and the size is one less than old size", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Push(6)
		assert.Equal(t, 3, s.Size())
		val, _ := s.Pop()
		assert.Equal(t, val, 6)
		assert.Equal(t, s.Size(), 2)
	})

	t.Run("if one pushes x then peeks, the value returned is x, but the siz stays the same", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Push(6)
		assert.Equal(t, 3, s.Size())
		val, _ := s.Peek()
		assert.Equal(t, val, 6)
		assert.Equal(t, s.Size(), 3)
	})
}

func TestError(t *testing.T) {
	s := stack.New()

	t.Run("popping from empty stack return error: errnosuchelement", func(t *testing.T) {
		_, err := s.Pop()

		if err == nil {
			t.Fail()
			t.Logf("Expect error is not nil, but got %v", err)
		}

		assert.Equal(t, stack.ErrorNoSuchElement, err)
	})
	t.Run("peeking from empty stack return error: errnosuchelement", func(t *testing.T) {
		_, err := s.Peek()

		if err == nil {
			t.Fail()
			t.Logf("Expect error is not nil, but got %v", err)
		}

		assert.Equal(t, stack.ErrorNoSuchElement, err)
	})
}
