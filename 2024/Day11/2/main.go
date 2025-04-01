package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// ********* Advent of Code 2024 *********
// --- Day 11: Plutonian Pebbles --- Puzzle 2
// https://adventofcode.com/2024/day/11

// вводим в программу кэш, который будет хранить количество камней, вычисленное для комбинации "исходный_камень-количество_итераций"
// например для четырех итераций исходного камня "42" - кэш будет хранить значения:
// "42-1" --> 2 		["4" "2"]
// "42-2" --> 2 		["8096" "4048"]
// "42-3" --> 4 		["80" "96" "40" "48"]
// "42-4" --> 8 		["8" "0" "9" "6" "4" "0" "4" "8"]
var cache map[string]int

func main() {

	cache = make(map[string]int) // инициализация кэша

	input := GetData("data.txt")
	answer := day11_2(input, 75)
	fmt.Println(answer)
}

func day11_2(stones []string, blinks int) (result int) {

	start := time.Now()

	for _, stone := range stones {
		result += count_stone_blinks(stone, blinks)
	}

	fmt.Printf("%s \n", time.Since(start)) // замер времени выполнения
	return result
}

// Функция вычисляет во сколько камней превратится текущий камень stone через depth итерций --
// Рекурсия с кэшированием результата
func count_stone_blinks(stone string, depth int) (result int) {

	// проверка - присутствуют ли данные в кэше
	key := stone + "-" + strconv.Itoa(depth)
	v, isCacheContainKey := cache[key]

	if isCacheContainKey {
		return v
	} else {
		left_stone, right_stone := single_blink_stone(stone)

		// условие выхода из рекурсии
		if depth == 1 {
			if right_stone == "" {
				return 1
			} else {
				return 2
			}
		} else {
			result = count_stone_blinks(left_stone, depth-1)
			if right_stone != "" {
				result += count_stone_blinks(right_stone, depth-1)
			}
		}

		cache[key] = result
		return result
	}
}

// Функция производит один шаг превращения камня stone и возвращает камни (leftStone, rightStone), в которые он превратился
// Второго камня (rightStone) может не быть, в этом случае вместо него возвращается пустая строка
// Правила превращения:
// 1) Если элемент = "0", он заменяется на "1"
// 2) Если количество знаков элемента четное "2000" - он разбивается на две равные части "20" и "00", лидирующие нули убираются -> "20" и "0"
// 3) В осталных случаях - значение элемента умножается на 2024
func single_blink_stone(stone string) (leftStone, rightStone string) {

	if stone == "0" {
		leftStone = "1"
		rightStone = ""
	} else if len(stone)%2 == 0 {
		s := strings.Split(stone, "")

		part1 := s[:len(stone)/2]
		leftStone = strings.Join(part1, "")

		// избавляемся от лидирующих нулей конвертацией Atoi <-> Itoa
		part2 := s[len(stone)/2:]
		val, err := strconv.Atoi(strings.Join(part2, ""))
		if err != nil {
			panic(err)
		}

		rightStone = strconv.Itoa(val)

	} else {
		val, err := strconv.Atoi(stone)

		if err != nil {
			panic(err)
		}

		val *= 2024

		leftStone = strconv.Itoa(val)
		rightStone = ""
	}

	return leftStone, rightStone
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
