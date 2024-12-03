package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadfile(t *testing.T) {

	return_arr := readfile("./Input_test.txt")

	arr1 := [][]int{{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9}}

	assert.Equal(t, arr1, return_arr, "2d Array doesn't match test case")

}

func TestDay2Part1(t *testing.T) {

	safe_report_count := Day2Part1("./Input_test.txt")

	assert.Equal(t, 2, safe_report_count, "Safe report count doesn't match test case")
}

func TestDay2Part2(t *testing.T) {

	dapened_safe_report_count := Day2Part2("./Input_test.txt")

	assert.Equal(t, 4, dapened_safe_report_count, "Safe report count doesn't match test case")
}
