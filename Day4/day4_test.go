package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadfile(t *testing.T) {

	input := [][]rune{
		{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
		{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
		{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
		{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
		{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
		{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
		{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
		{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
		{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'}}

	assert.Equal(t, input, readfile("./Input_test.txt"), "Inputs do not match")
}

func TestDay4Part1(t *testing.T) {

	assert.Equal(t, 18, Day4Part1("./Input_test.txt"), "Inputs do not match")
}

func TestDay4Part2(t *testing.T) {

	assert.Equal(t, 9, Day4Part2("./Input_test.txt"), "Inputs do not match")
}
