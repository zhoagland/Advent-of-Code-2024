package day2

import (
	"encoding/csv"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day2Part1(path string) (safe_reports int) {

	result := readfile(path)
	safe_reports = 0

	for _, reports := range result {

		safe, _ := calculate_safety(reports)

		if safe {
			safe_reports += 1
		}
	}

	return safe_reports
}

func Day2Part2(path string) (safe_reports int) {

	result := readfile(path)
	safe_reports = 0

	for idx := 0; idx < len(result); idx++ {
		reports := result[idx]

		safe, _ := calculate_safety(reports)

		if !safe {
			for removed_idx := 0; removed_idx < len(reports); removed_idx++ {

				left_half := make([]int, removed_idx)
				copy(left_half, reports[:removed_idx])

				right_half := make([]int, len(reports)-removed_idx-1)
				copy(right_half, reports[removed_idx+1:])

				arr_missing_one := append(left_half, right_half...)
				value, _ := calculate_safety(arr_missing_one)

				if value {
					safe = true
					break
				}

			}
		}

		if safe {
			safe_reports += 1
		}

	}

	return safe_reports
}

func readfile(path string) (level [][]int) {
	var temp_arr []int

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
		output := strings.Split(record[0], " ")

		for _, values := range output {
			intval, _ := strconv.Atoi(values)

			temp_arr = append(temp_arr, intval)
		}
		level = append(level, temp_arr)
		temp_arr = make([]int, 0)
	}

	return level
}

func calculate_safety(reports []int) (safe bool, fail_idx int) {

	previous_delta_slope := 1

	for level := 1; level < len(reports); level++ {

		delta := reports[level] - reports[level-1]

		if (math.Abs(float64(delta)) < 1) || (math.Abs(float64(delta)) > 3) {
			return false, level
		}

		delta_slope := delta / int(math.Abs(float64(delta)))

		if (delta_slope != previous_delta_slope) && (level != 1) {
			return false, level
		}

		previous_delta_slope = delta_slope

	}

	return true, 0
}
