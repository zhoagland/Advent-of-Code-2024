package day1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Part1(t *testing.T) {

	result := Day1Part1("./Input_test.txt")
	assert.Equal(t, 11, result, "Needs to meet test case.")

	fmt.Println(result)
}

func TestDay1Part2(t *testing.T) {

	result := Day1Part2("./Input_test.txt")
	assert.Equal(t, 31, result, "Needs to meet test case.")

	fmt.Println(result)
}

func TestReadfile(t *testing.T) {
	return_arr1, return_arr2 := readfile("./Input_test.txt")

	arr1 := []int{3, 4, 2, 1, 3, 3}
	arr2 := []int{4, 3, 5, 3, 9, 3}

	assert.Equal(t, return_arr1, arr1, "Array 1 doesn't match test case")
	assert.Equal(t, return_arr2, arr2, "Array 2 doesn't match test case")
}
