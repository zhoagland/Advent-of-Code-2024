package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadfile(t *testing.T) {

	test_string := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	assert.Equal(t, test_string, readfile("./Input_test.txt"), "File string doesn't match test case")
}

func TestDay3Part1(t *testing.T) {

	assert.Equal(t, 161, Day3Part1("./Input_test.txt"), "Sum doesn't match test case")
}

func TestDay3Part2(t *testing.T) {

	assert.Equal(t, 48, Day3Part2("./Input_test2.txt"), "Sum doesn't match test case")
}
