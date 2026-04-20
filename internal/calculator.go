package internal

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrEmptyStack = errors.New("empty stack")
)

type Calculator interface {
	Process(input string) (float64, error)
}

type internalCalculator struct {
	stack Stack[float64]
}

func NewCalculator() Calculator {
	return &internalCalculator{
		stack: NewStack[float64](),
	}
}

func (c *internalCalculator) Process(input string) (float64, error) {
	tokens := strings.Fields(input)
	for _, token := range tokens {
		switch token {
		case "+":
			if err := c.doAdd(); err != nil {
				return 0, err
			}
		case "-":
			if err := c.doSubtract(); err != nil {
				return 0, err
			}
		case "*":
			if err := c.doMultiply(); err != nil {
				return 0, err
			}
		case "/":
			if err := c.doDivide(); err != nil {
				return 0, err
			}
		default:
			val, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			c.stack.Push(val)
		}
	}
	result, ok := c.stack.Pop()
	if !ok {
		return 0, ErrEmptyStack
	}
	return result, nil
}

func (c* internalCalculator) getTopPair() (float64, float64, error) {
	a, ok := c.stack.Pop()
	if !ok {
		return 0, 0, ErrEmptyStack
	}
	b, ok := c.stack.Pop()
	if !ok {
		return 0, 0, ErrEmptyStack
	}
	return a, b, nil
}

func (c* internalCalculator) doMath(fn func(float64, float64) float64) error {
	a, b, err := c.getTopPair()
	if err != nil {
		return err
	}
	result := fn(a, b)
	c.stack.Push(result)
	return nil
}	

func (c* internalCalculator) doAdd() error {
	return c.doMath(func(a, b float64) float64 {
		return b + a
	})
}

func (c* internalCalculator) doSubtract() error {
	return c.doMath(func(a, b float64) float64 {
		return b - a
	})
}

func (c* internalCalculator) doMultiply() error {
	return c.doMath(func(a, b float64) float64 {
		return b * a
	})
}

func (c* internalCalculator) doDivide() error {
	return c.doMath(func(a, b float64) float64 {
		return b / a
	})	
}