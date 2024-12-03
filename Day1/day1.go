package day1

import (
	"encoding/csv"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day1Part1(path string) (difference int64) {

	array1, array2 := readfile(path)

	slices.Sort(array1)
	slices.Sort(array2)

	for idx := 0; idx < len(array1); idx++ {
		difference += int64(math.Abs(float64(array1[idx]) - float64(array2[idx])))
	}

	return difference

}

func Day1Part2(path string) (similarity int) {

	// Could optimize by adding hashmap and doing memoization
	similarity = 0

	count := 0

	array1, array2 := readfile(path)

	slices.Sort(array1)
	slices.Sort(array2)

	for out := 0; out < len(array1); out++ {
		for in := 0; in < len(array2); in++ {

			if array2[in] > array1[out] {
				break
			}

			if array1[out] == array2[in] {
				count++
			}
		}

		similarity += (array1[out]) * (count)
		count = 0
	}

	return similarity
}

func readfile(path string) (array1 []int, array2 []int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, record := range records {
		output := strings.Split(record[0], "   ")

		intrep1, _ := strconv.Atoi(output[0])
		intrep2, _ := strconv.Atoi(output[1])

		array1 = append(array1, intrep1)
		array2 = append(array2, intrep2)
	}

	return
}
