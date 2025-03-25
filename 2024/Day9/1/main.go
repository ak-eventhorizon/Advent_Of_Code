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
// --- Day 9: Disk Fragmenter --- Puzzle 1
// https://adventofcode.com/2024/day/9

func main() {

	diskMap := GetData("data.txt")
	answer := day9_1(diskMap)

	fmt.Println(answer)
}

func day9_1(diskMap string) (result int) {

	expandedDiskMap := diskMapExpand(diskMap)
	defragedDiskMap := defrag(expandedDiskMap)

	SaveData(strings.Join(defragedDiskMap, ""), "output.txt")

	result = checksum(defragedDiskMap)

	return result
}

// Функция излекает контрольную сумму из карты диска
func checksum(input []string) (result int) {

	for i, v := range input {

		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		result += num * i
	}

	return result
}

// Функция получает произвольную развернутую карту диска: 0..111....22222
// Возвращает дефрагментированную карту, полученную из исходной:
// 0..111....22222
// 02.111....2222.
// 022111....222..
// 0221112...22...
// 02211122..2....
// 022111222......
// 022111222
func defrag(input []string) (result []string) {

	// fmt.Println(input) // DEBUG

	for i := 0; i < len(input); i++ { // поиск очередного свободного блока слева-направо >->->->->->->

		if input[i] == "." { // свободный блок найден
			for j := len(input) - 1; j >= 0; j-- { // поиск очередного значения для перемещения в свободный блок справа-налево <-<-<-<-<-<

				if i == j { // если прямой и обратный обходы встретились >->->->i=j<-<-<-<
					break
				}

				if input[j] != "." { // нашли значение для перемещения
					input[i] = input[j] // перемещение значения в свободный блок
					input[j] = "."      // освобождение значения
					// fmt.Println(input)  // DEBUG
					break
				}
			}
		}
	}

	// Удаление пустых ячеек после дефрагментации [022111222......] --> [022111222]
	result = slices.DeleteFunc(input, func(s string) bool {
		return s == "."
	})

	// fmt.Println(result) // DEBUG

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
