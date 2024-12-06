package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadfile(t *testing.T) {

	test_updates := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}

	// key must be printed before value
	test_rules := map[int][]int{
		47: {53, 13, 61, 29},
		97: {13, 61, 47, 29, 53, 75},
		75: {29, 53, 47, 61, 13},
		61: {13, 53, 29},
		29: {13},
		53: {29, 13},
	}

	rules, updates := readfile("./Input_test.txt")
	assert.Equal(t, test_rules, rules, "Rules do not match")
	assert.Equal(t, test_updates, updates, "Updates do not match")
}

func TestDay5Part1(t *testing.T) {

	assert.Equal(t, 143, Day5Part1("./Input_test.txt"), "Inputs do not match")
}

func TestDay5Part2(t *testing.T) {

	assert.Equal(t, 9, Day5Part2("./Input_test.txt"), "Inputs do not match")
}
