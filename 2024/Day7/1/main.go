package main

import (
	"bufio"
	"fmt"
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

	// for _, v := range cases {
	// 	fmt.Println(v.value, v.numbers)
	// }

	GetAllCombinations([]string{"A", "B"}, 3)

	return result
}

// Функция возвращает все комбинации длиной k, которые можно составить из набора символов set
func GetAllCombinations(set []string, k int) (list []string) {
	n := len(set)
	printCombinationRec(set, "", n, k)

	return list
}

func printCombinationRec(set []string, prefix string, n int, k int) {

	// базовый случай - возвращает префикс
	if k == 0 {
		fmt.Println(prefix)
		return
	}

	for i := range n {
		newPrefix := prefix + set[i]
		printCombinationRec(set, newPrefix, n, k-1)

	}

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
