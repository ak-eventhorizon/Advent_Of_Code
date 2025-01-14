package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ********* Advent of Code 2024 *********
// --- Day 1: Historian Hysteria --- Puzzle 2
// https://adventofcode.com/2024/day/1

func main() {

	var slice1, slice2 []int = GetDataFromFile("data.txt")

	answer := day1_2(slice1, slice2)
	fmt.Printf("%v\n", answer)

}

// Такой вариант реализует описание алгоритма решения, указанного в постановке задачи
func day1_2(list1 []int, list2 []int) int {

	var result int = 0

	for _, v1 := range list1 {

		count := 0

		for _, v2 := range list2 {
			if v1 == v2 {
				count++
			}
		}

		result += v1 * count

	}

	return result
}

// Функция разбирает текстовый файл с исходными данными на два слайса
func GetDataFromFile(filename string) ([]int, []int) {

	var slice1, slice2 []int

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		values := strings.Split(line, "   ")

		num1, err := strconv.Atoi(values[0])
		if err != nil {
			fmt.Println(err)
		}
		slice1 = append(slice1, num1)

		num2, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Println(err)
		}
		slice2 = append(slice2, num2)
	}

	file.Close()

	return slice1, slice2
}
