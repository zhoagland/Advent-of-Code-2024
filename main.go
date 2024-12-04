package main

import (
	day1 "AdventOfCode2024/v2/Day1"
	day2 "AdventOfCode2024/v2/Day2"
	day3 "AdventOfCode2024/v2/Day3"
	day4 "AdventOfCode2024/v2/Day4"
	"fmt"
)

func main() {

	fmt.Print("Day 1 Part 1 Answer: ")
	fmt.Println(day1.Day1Part1("./Day1/Input.txt"))
	fmt.Print("Day 1 Part 2 Answer: ")
	fmt.Println(day1.Day1Part2("./Day1/Input.txt"))
	fmt.Print("Day 2 Part 1 Answer: ")
	fmt.Println(day2.Day2Part1("./Day2/Input.txt"))
	fmt.Print("Day 2 Part 2 Answer: ")
	fmt.Println(day2.Day2Part2("./Day2/Input.txt"))
	fmt.Print("Day 3 Part 1 Answer: ")
	fmt.Println(day3.Day3Part1("./Day3/Input.txt"))
	fmt.Print("Day 3 Part 2 Answer: ")
	fmt.Println(day3.Day3Part2("./Day3/Input.txt"))
	fmt.Print("Day 4 Part 1 Answer: ")
	fmt.Println(day4.Day4Part1("./Day4/Input.txt"))
	fmt.Print("Day 4 Part 2 Answer: ")
	fmt.Println(day4.Day4Part2("./Day4/Input.txt"))
}
