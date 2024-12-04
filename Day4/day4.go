package day4

import (
	"os"
	"strings"
)

const word_length = 3

type cord struct {
	x int //north south
	y int //east west
}

func Day4Part1(path string) (wordcount int) {

	wordsearch := readfile(path)

	for row := range wordsearch {
		for column := range wordsearch[row] {

			if wordsearch[row][column] == 'X' {

				if check_s(wordsearch, cord{row, column}) {
					wordcount++
				}

				if check_n(wordsearch, cord{row, column}) {
					wordcount++
				}

				if check_w(wordsearch, cord{row, column}) {
					wordcount++
				}

				if check_e(wordsearch, cord{row, column}) {
					wordcount++
				}

				if check_nw(wordsearch, cord{row, column}) {
					wordcount++
				}

				if check_ne(wordsearch, cord{row, column}) {
					wordcount++
				}

				if check_sw(wordsearch, cord{row, column}) {
					wordcount++
				}

				if check_se(wordsearch, cord{row, column}) {
					wordcount++
				}

			}
		}
	}

	return
}

func Day4Part2(path string) (xmas_count int) {
	wordsearch := readfile(path)
	for row := 1; row < len(wordsearch)-1; row++ {
		for column := 1; column < len(wordsearch[row])-1; column++ {
			if wordsearch[row][column] == 'A' {

				pt1 := (wordsearch[row-1][column-1] == 'M' && wordsearch[row+1][column+1] == 'S')
				pt1 = pt1 || (wordsearch[row-1][column-1] == 'S' && wordsearch[row+1][column+1] == 'M')

				pt2 := (wordsearch[row+1][column-1] == 'M' && wordsearch[row-1][column+1] == 'S')
				pt2 = pt2 || (wordsearch[row+1][column-1] == 'S' && wordsearch[row-1][column+1] == 'M')

				valid := pt1 && pt2

				if valid {

					xmas_count++
				}
			}
		}
	}
	return
}

func readfile(path string) (wordsearch [][]rune) {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file_stat, _ := file.Stat()

	b := make([]byte, file_stat.Size())
	file.Read(b)

	rows := strings.Split(string(b), "\r\n")

	wordsearch = make([][]rune, len(rows))

	for row := range rows {
		wordsearch[row] = make([]rune, len(rows[row]))
		for letter := 0; letter < len(rows[row]); letter++ {
			wordsearch[row][letter] = rune(rows[row][letter])
		}
	}
	return
}

func check_w(wordsearch [][]rune, index cord) bool {
	if index.y-word_length < 0 {
		return false
	}

	if wordsearch[index.x][index.y-1] == 'M' && wordsearch[index.x][index.y-2] == 'A' && wordsearch[index.x][index.y-3] == 'S' {
		return true
	}

	return false
}

func check_e(wordsearch [][]rune, index cord) bool {
	if index.y+word_length >= len(wordsearch[index.x]) {
		return false
	}

	if wordsearch[index.x][index.y+1] == 'M' && wordsearch[index.x][index.y+2] == 'A' && wordsearch[index.x][index.y+3] == 'S' {
		return true
	}

	return false
}

func check_n(wordsearch [][]rune, index cord) bool {
	if index.x-word_length < 0 {
		return false
	}

	if wordsearch[index.x-1][index.y] == 'M' && wordsearch[index.x-2][index.y] == 'A' && wordsearch[index.x-3][index.y] == 'S' {
		return true
	}

	return false
}

func check_s(wordsearch [][]rune, index cord) bool {
	if index.x+word_length >= len(wordsearch) {
		return false
	}

	if wordsearch[index.x+1][index.y] == 'M' && wordsearch[index.x+2][index.y] == 'A' && wordsearch[index.x+3][index.y] == 'S' {
		return true
	}

	return false
}

func check_nw(wordsearch [][]rune, index cord) bool {
	if index.x-word_length < 0 || index.y-word_length < 0 {
		return false
	}

	if wordsearch[index.x-1][index.y-1] == 'M' && wordsearch[index.x-2][index.y-2] == 'A' && wordsearch[index.x-3][index.y-3] == 'S' {
		return true
	}

	return false
}

func check_ne(wordsearch [][]rune, index cord) bool {
	if index.x-word_length < 0 || index.y+word_length >= len(wordsearch[index.x]) {
		return false
	}

	if wordsearch[index.x-1][index.y+1] == 'M' && wordsearch[index.x-2][index.y+2] == 'A' && wordsearch[index.x-3][index.y+3] == 'S' {
		return true
	}

	return false
}

func check_sw(wordsearch [][]rune, index cord) bool {
	if index.x+word_length >= len(wordsearch) || index.y-word_length < 0 {
		return false
	}

	if wordsearch[index.x+1][index.y-1] == 'M' && wordsearch[index.x+2][index.y-2] == 'A' && wordsearch[index.x+3][index.y-3] == 'S' {
		return true
	}

	return false
}

func check_se(wordsearch [][]rune, index cord) bool {
	if index.x+word_length >= len(wordsearch) || index.y+word_length >= len(wordsearch[index.x]) {
		return false
	}

	if wordsearch[index.x+1][index.y+1] == 'M' && wordsearch[index.x+2][index.y+2] == 'A' && wordsearch[index.x+3][index.y+3] == 'S' {
		return true
	}

	return false
}
