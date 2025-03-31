package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// ********* Advent of Code 2024 *********
// --- Day 11: Plutonian Pebbles --- Puzzle 1
// https://adventofcode.com/2024/day/11

func main() {

	input := GetData("data.txt")
	answer := day11_1(input, 25)

	fmt.Println(answer)
}

func day11_1(input []string, blinks int) (result int) {

	tmp := slices.Clone(input)
	// fmt.Println(tmp) // Debug

	for range blinks {
		tmp = blink(tmp)
		// fmt.Println(tmp) // Debug
	}

	return len(tmp)
}

// Функция производит одну итерацию манипуляций с набором данных по следующим правилам:
// 1) Если элемент = "0", он заменяется на "1"
// 2) Если количество знаков элемента четное "2000" - оно разбивается на две равные части "20" и "00", лидирующие нули убираются -> "20" и "0"
// 3) В осталных случаях - значение элемента умножается на 2024
func blink(input []string) (result []string) {

	for _, elem := range input {
		if elem == "0" {
			result = append(result, "1")
		} else if len(elem)%2 == 0 {
			s := strings.Split(elem, "")

			part1 := s[:len(elem)/2]
			p1 := strings.Join(part1, "")

			// избавляемся от лидирующих нулей конвертацией Atoi -> Itoa
			part2 := s[len(elem)/2:]
			val, err := strconv.Atoi(strings.Join(part2, ""))
			if err != nil {
				panic(err)
			}
			p2 := strconv.Itoa(val)

			result = append(result, p1, p2)
		} else {
			val, err := strconv.Atoi(elem)
			if err != nil {
				panic(err)
			}
			val *= 2024
			s := strconv.Itoa(val)
			result = append(result, s)
		}
	}

	return result
}

// Функция извлекает из файла filename исходняе данные для задачи
func GetData(filename string) (result []string) {

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		line := strings.Split(scanner.Text(), " ")
		result = append(result, line...)
	}

	file.Close()

	return result
}

// Функция сохраняет полученную строку input в файл filename
func SaveData(input []string, filename string) {

	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	line := strings.Join(input, " ")
	file.WriteString(line)

	file.Close()
}
