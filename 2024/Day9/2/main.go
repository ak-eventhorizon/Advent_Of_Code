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
// --- Day 9: Disk Fragmenter --- Puzzle 2
// https://adventofcode.com/2024/day/9

func main() {

	diskMap := getData("data.txt")
	answer := day9_1(diskMap)

	fmt.Println(answer)
}

func day9_1(diskMap string) (result int) {

	expandedDiskMap := expand(diskMap)
	defragedDiskMap := defrag(expandedDiskMap)

	saveData(strings.Join(defragedDiskMap, ""), "output.txt")

	result = checksum(defragedDiskMap)

	return result
}

// Функция излекает контрольеую сумму из карты диска
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

// Функция получает произвольную развернутую карту диска: 0..111....22233
// Возвращает дефрагментированную карту, полученную из исходной, перенося файлы целиком:
// 0..111....22233
// 033111....222..
// 033111222......
// 033111222
func defrag(input []string) (result []string) {

	fmt.Println("Input: ", input) // DEBUG

	blocks := [][]string{}
	block := []string{}

	files := [][]string{}
	file := []string{}

	// разбираем ввод на однородные блоки [0 0 . . . 1 1 1 . . . 2] --> [ [0 0] [. . .] [1 1 1] [. . .] [2] ]
	for i := 0; i < len(input); i++ {

		block = append(block, input[i])

		if (i+1 < len(input)) && (input[i+1] == input[i]) { // следующий элемент является продолжением текущего
			continue
		} else { // этот элемент последний в блоке
			blocks = append(blocks, block)
			block = []string{}
		}
	}

	// Поиск всех файлов справа-налево <-<-<-<-<-<
	for i := len(input) - 1; i >= 0; i-- {

		if input[i] != "." { // если элемент не пустой - значит начался файл
			file = append(file, input[i])

			if (i-1 >= 0) && (input[i-1] == input[i]) { // следующий элемент является продолжением файла
				continue
			} else { // этот элемент последний в данном файле, копируем файл в список файлов, сбрасываем переменную для следующего
				files = append(files, file)
				file = []string{}
			}
		}
	}

	fmt.Println("Blocks:", blocks) // DEBUG
	fmt.Println("Files: ", files)  // DEBUG

	// TODO rebuildDisk()

	return result
}

// Функуция получает карту диска и пытается переместить каждый файл с конца в ближайшее слева свободное место, подходящее по размеру
// При успешном перемещении файла он стирается в исходном месте расположения
func rebuildDisk(initialMap [][]string) (rebuildMap [][]string) {

	rebuildMap = slices.Clone(initialMap) // копия исходной карты для помещения в нее перестроенных файлов

	for i := len(initialMap) - 1; i >= 0; i-- { // поиск следующего файла справа-налево <-<-<-<-<
		if initialMap[i][0] != "." { // блок содержит файл

			fileLen := len(initialMap[i])

			for j := 0; j < len(rebuildMap); j++ { // поиск следующего пустого блока слева направо >->->->->
				if rebuildMap[j][0] == "." { // блок содержит пустое место
					// ...[. . . . .]... + [2 2 2] -> ...[2 2 2] [. .]...
					// ...[. . .]... + [2 2 2] -> ...[2 2 2]...
				}
			}

		}
	}

	return rebuildMap
}

// Функция получает карту диска в свернутом виде: "12345"
// И разворачивает ее по правилам, описанным в задаче: [0 . . 1 1 1 . . . . 2 2 2 2 2]
func expand(diskMap string) (result []string) {

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
func getData(filename string) (s string) {

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
func saveData(s string, filename string) {

	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	file.WriteString(s)

	file.Close()
}
