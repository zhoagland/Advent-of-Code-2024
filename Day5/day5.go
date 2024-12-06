package day5

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day5Part1(path string) (middle_letter_sum int) {
	rules, updates := readfile(path)

	for update := range updates {
		good_update := true
		for idx := range updates[update] {
			good_update = good_update && values_not_in_other_slice(updates[update][:idx], rules[updates[update][idx]])
			if !good_update {
				break
			}
		}
		if good_update {
			index_to_add := ((len(updates[update]) - 1) / 2)
			middle_letter_sum += updates[update][index_to_add]
		}
	}

	return
}

func Day5Part2(path string) (xmas_count int) {

	return
}

func readfile(path string) (rules map[int][]int, updates [][]int) {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file_stat, _ := file.Stat()

	b := make([]byte, file_stat.Size())
	file.Read(b)

	rows := strings.Split(string(b), "\r\n")
	spacerow := 0
	rules = make(map[int][]int, 1)

	for row := range rows {
		if strings.Contains(rows[row], "|") {
			kv := strings.Split(rows[row], "|")
			k, _ := strconv.Atoi(kv[0])
			v, _ := strconv.Atoi(kv[1])
			rules[k] = append(rules[k], v)
		} else if strings.Compare(rows[row], "") == 0 {
			updates = make([][]int, len(rows)-(row+1))
			spacerow = row + 1
		} else {
			vals := strings.Split(rows[row], ",")
			updates[row-spacerow] = make([]int, len(vals))
			for val := range vals {
				v, _ := strconv.Atoi(vals[val])
				updates[row-spacerow][val] = v
			}
		}

	}
	return
}

func values_not_in_other_slice(values_to_check []int, bad_values []int) (passed bool) {

	passed = true
	if len(bad_values) != 0 {
		for check := range values_to_check {
			passed = passed && !slices.Contains(bad_values, values_to_check[check])
		}
	}
	return
}
