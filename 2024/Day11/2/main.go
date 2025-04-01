package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

// ********* Advent of Code 2024 *********
// --- Day 11: Plutonian Pebbles --- Puzzle 2
// https://adventofcode.com/2024/day/11

func main() {

	input := GetData("data.txt")
	answer := day11_2(input, 25)

	fmt.Println(answer)
}

func day11_2(input []string, blinks int) (result int) {

	// вводим кэш, который будет хранить для каждого значение - его преобразование по указанным в задаче правилам
	cache := make(map[string][]string)

	tmp := slices.Clone(input)

	for i := 0; i < blinks; i++ {
		start := time.Now()

		tmp = blink(tmp, cache)

		duration := time.Since(start)
		fmt.Printf("%v -- %s \n", i, duration) // фиксация времени выполнения итерации
	}

	return len(tmp)
}

// Функция производит одну итерацию манипуляций с набором данных input по следующим правилам:
// 1) Если элемент = "0", он заменяется на "1"
// 2) Если количество знаков элемента четное "2000" - оно разбивается на две равные части "20" и "00", лидирующие нули убираются -> "20" и "0"
// 3) В осталных случаях - значение элемента умножается на 2024
// cache - общая между всеми итерациями кеширующая структура, сохраняющая рпезультаты каждого преобразования
func blink(input []string, cache map[string][]string) (result []string) {

	for _, elem := range input {

		if cache[elem] == nil { // пишем результат преобразования в кэш, если его еще там нет

			if elem == "0" {
				cache[elem] = []string{"1"}
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

				cache[elem] = []string{p1, p2}
			} else {
				val, err := strconv.Atoi(elem)

				if err != nil {
					// fmt.Println("ding", elem) // DEBUG
					panic(err)
				}

				val *= 2024
				s := strconv.Itoa(val)

				cache[elem] = []string{s}
			}
		}

		result = append(result, cache[elem]...)
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
