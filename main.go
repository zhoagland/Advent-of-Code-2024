package main

import (
	day1 "AdventOfCode2024/v2/Day1"
	day2 "AdventOfCode2024/v2/Day2"
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
}
