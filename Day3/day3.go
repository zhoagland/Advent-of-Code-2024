package day3

import (
	"os"
	"strconv"
	"strings"
)

func Day3Part1(path string) (sum int) {

	var mul_param_1 int
	var mul_param_2 int

	valid_chars := map[rune]bool{'0': true, '1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true}

	memory := readfile(path)
	start_of_command := false
	digits_in_param := 0
	delimiter := false
	end_of_command := false

	for idx := 0; idx < len(memory); idx++ {
		if idx+4 < len(memory) && strings.Compare(memory[idx:idx+4], "mul(") == 0 {
			start_of_command = true
			delimiter = false
			digits_in_param = 0
			end_of_command = false
			idx += 4
		}
		if _, exists := valid_chars[rune(memory[idx])]; exists && start_of_command {
			digits_in_param++
		}
		if (memory[idx] == ',') && ((digits_in_param > 0) && (digits_in_param < 4)) && start_of_command {
			mul_param_1, _ = strconv.Atoi(memory[idx-digits_in_param : idx])
			digits_in_param = 0
			delimiter = true
		}
		if (memory[idx] == ')') && ((digits_in_param > 0) && (digits_in_param < 4)) && start_of_command && delimiter {
			mul_param_2, _ = strconv.Atoi(memory[idx-digits_in_param : idx])
			digits_in_param = 0
			end_of_command = true
		}

		if end_of_command {
			sum += (mul_param_1 * mul_param_2)
			start_of_command = false
			delimiter = false
			digits_in_param = 0
			end_of_command = false
		}
	}

	return sum
}

func Day3Part2(path string) (sum int) {

	var mul_param_1 int
	var mul_param_2 int

	valid_chars := map[rune]bool{'0': true, '1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true}

	dont_flag := false
	memory := readfile(path)
	start_of_command := false
	digits_in_param := 0
	delimiter := false
	end_of_command := false

	for idx := 0; idx < len(memory); idx++ {

		if idx+4 < len(memory) && strings.Compare(memory[idx:idx+4], "do()") == 0 {
			idx += 4
			dont_flag = false
		}
		if idx+7 < len(memory) && strings.Compare(memory[idx:idx+7], "don't()") == 0 {
			idx += 7
			dont_flag = true
		}

		if !dont_flag {
			if idx+4 < len(memory) && strings.Compare(memory[idx:idx+4], "mul(") == 0 {
				start_of_command = true
				delimiter = false
				digits_in_param = 0
				end_of_command = false
				idx += 4
			}
			if _, exists := valid_chars[rune(memory[idx])]; exists && start_of_command {
				digits_in_param++
			}
			if (memory[idx] == ',') && ((digits_in_param > 0) && (digits_in_param < 4)) && start_of_command {
				mul_param_1, _ = strconv.Atoi(memory[idx-digits_in_param : idx])
				digits_in_param = 0
				delimiter = true
			}
			if (memory[idx] == ')') && ((digits_in_param > 0) && (digits_in_param < 4)) && start_of_command && delimiter {
				mul_param_2, _ = strconv.Atoi(memory[idx-digits_in_param : idx])
				digits_in_param = 0
				end_of_command = true
			}

			if end_of_command {
				sum += (mul_param_1 * mul_param_2)
				start_of_command = false
				delimiter = false
				digits_in_param = 0
				end_of_command = false
			}
		}
	}

	return sum
}

func readfile(path string) (memory string) {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file_stat, _ := file.Stat()

	b := make([]byte, file_stat.Size())
	file.Read(b)

	return string(b)
}
