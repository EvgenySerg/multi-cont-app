package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	ress := Sum(1, 2)
	assert.Equal(t, 3, ress)
}

func TestSecond(t *testing.T) {
	ress := Sum(3, 5)
	assert.Equal(t, 8, ress)
}

func TestThird(t *testing.T) {
	ress := Sum(3, 5)
	assert.Equal(t, 8, ress)
}
