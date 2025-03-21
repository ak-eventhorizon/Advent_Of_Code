package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

// ********* Advent of Code 2024 *********
// --- Day 9: Disk Fragmenter --- Puzzle 1
// https://adventofcode.com/2024/day/9

func main() {

	diskMap := GetData("data.txt")
	answer := day9_1(diskMap)

	fmt.Println(answer)
}

func day9_1(diskMap string) (result int) {

	expandedDiskMap := diskMapExpand(diskMap)
	defrag(expandedDiskMap)

	return result
}

// Функция получает произвольную карту диска: 0..111....22222
// Возвращает дефрагментированную карту, полученную из исходной:
// 0..111....22222
// 02.111....2222.
// 022111....222..
// 0221112...22...
// 02211122..2....
// 022111222......
func defrag(input []string) (result []string) {

	fmt.Println("Input:   ", input) // DEBUG

	emptySpaceCounter := 0 // количество пропусков, подлежажих дефрагментации
	for _, v := range input {
		if v == "." {
			emptySpaceCounter++
		}
	}

	fmt.Println("toFill:  ", emptySpaceCounter) // DEBUG

	// копия исходного
	tmp := slices.Clone(input)
	// удалить из копии "."
	tmp = slices.DeleteFunc(tmp, func(n string) bool {
		return n == "."
	})
	// перевернуть копию
	slices.Reverse(tmp)
	fmt.Println("Fillers: ", tmp) // DEBUG

	// отрезать нужное количество заполнителя по количеству свободного места для дефрагментации
	tmp = tmp[:emptySpaceCounter] // ТУТ НЕПРАВИЛЬНО, на самом деле нужно 12
	fmt.Println("NeedFill:", tmp) // DEBUG

	// поразрядное наполнение результата
	for _, v := range input {

		if v == "." {
			result = append(result, tmp[0])
			tmp = slices.Delete(tmp, 0, 1)
		} else {
			result = append(result, v)
		}

		if len(tmp) == 0 {
			break
		}

	}
	fmt.Println("Result:  ", result) // DEBUG

	return result
}

// Функция получает карту диска в свернутом виде: "12345"
// И разворачивает ее по правилам, описанным в задаче: [0 . . 1 1 1 . . . . 2 2 2 2 2]
func diskMapExpand(diskMap string) (result []string) {

	fileID := 0

	for i, char := range diskMap {

		lenElem, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}

		if i%2 == 0 { // каждый четный элемент карты - это размер файла
			for range lenElem {
				result = append(result, fmt.Sprint(fileID))
			}
			fileID++
		} else { // каждый нечетный элемент карты - это размер пустого блока
			for range lenElem {
				result = append(result, ".")
			}
		}
	}

	return result
}

// Функция извлекает из файла filename строку с условием задачи.
func GetData(filename string) (s string) {

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s = scanner.Text()
	}

	file.Close()

	return s
}

// Функция сохраняет строку в файл filename.
func SaveData(s string, filename string) {

	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	file.WriteString(s)

	file.Close()
}
