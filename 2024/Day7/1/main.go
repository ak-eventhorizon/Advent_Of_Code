package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ********* Advent of Code 2024 *********
// --- Day 7: Bridge Repair --- Puzzle 1
// https://adventofcode.com/2024/day/7

type Case struct {
	value   int
	numbers []int
}

func main() {

	data := GetData("data.txt")

	answer := day7_1(data)
	fmt.Println(answer)
}

func day7_1(cases []Case) (result int) {

	for _, v := range cases {
		fmt.Println(v.value, v.numbers)
	}

	fmt.Println(GetCombinations([]string{"A", "B"}, 3))

	return result
}

// Функция возвращает все комбинации строк длиной n, которые можно составить из символов symbols
func GetCombinations(symbols []string, n int) (list []string) {

	combinations := math.Pow(float64(len(symbols)), float64(n)) // количество комбинаций = количество_символов**n
	list = make([]string, int(combinations))

	for i := 0; i < int(combinations); i++ {

	}

	for i := 0; i < n; i++ {
		for _, symbol := range symbols {
			list[i] += symbol
		}
	}
	return list
}

// Функция извлекает из текстового файла все условия задачи.
func GetData(filename string) (cases []Case) {

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		var str []string
		var nums []int

		part1 := strings.Split(scanner.Text(), ":")[0]
		part2 := strings.Split(strings.Trim(strings.Split(scanner.Text(), ":")[1], " "), " ")

		str = append(str, part1)
		str = append(str, part2...)

		for _, v := range str {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			nums = append(nums, num)
		}

		var currentCase = new(Case)
		currentCase.value = nums[0]
		currentCase.numbers = nums[1:]

		cases = append(cases, *currentCase)
	}

	file.Close()

	return cases
}
