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

// Функция излекает контрольную сумму из карты диска
func checksum(input []string) (result int) {

	for i, v := range input {

		if v == "." {
			continue
		}

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
func defrag(input []string) (result []string) {

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

	// извлечение всех файлов справа-налево <-<-<-<-<-<
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

	// файл за файлом проводим дефрагментацию
	for _, file := range files {
		blocks = moveFile(blocks, file[0])
	}

	// собираем данные в формат результата
	for _, v := range blocks {
		result = append(result, v...)
	}

	// fmt.Println(input)  // DEBUG
	// fmt.Println(result) // DEBUG

	return result
}

// Функуция получает карту диска и ID файла, пробует переместить его в свободное место слева
// Если это удается - удаляет исходное расположение файла. Возвращает слайс (измененный или исходный в зависимости от результата перемещения файла)
func moveFile(initMap [][]string, fileID string) (resMap [][]string) {

	resMap = slices.Clone(initMap)

	// поиск файла справа-налево <-<-<-<-<
	for i := len(initMap) - 1; i >= 0; i-- {
		if initMap[i][0] == fileID { // файл найден

			file := initMap[i]
			fileIndex := i
			fileLen := len(initMap[i])

			// поиск пустого блока от начала до текущего файла
			for j := 0; j < i; j++ {
				if initMap[j][0] == "." { // пустой блок найден

					spaceLen := len(initMap[j])
					spaceBlock := make([]string, fileLen)
					for k := range spaceBlock {
						spaceBlock[k] = "."
					}

					if fileLen == spaceLen { // файл укладывается в свободное место точно по размеру - [. . .] + [2 2 2] -> [2 2 2]
						resMap[j] = file               // перемещение файла в пустой блок
						resMap[fileIndex] = spaceBlock // замещение исходной позиции файла пустым местом
						break

					} else if fileLen < spaceLen { // файл меньше свободного блока - [. . . . .] + [2 2 2] -> [2 2 2] + [. .]
						partOne := file                             // первая часть - файл
						partTwo := make([]string, spaceLen-fileLen) // вторая часть - пустое место, оставшееся после размещения файла
						for k := range spaceLen - fileLen {
							partTwo[k] = "."
						}
						resMap[fileIndex] = spaceBlock                      // замещение исходной позиции файла пустым местом
						resMap = slices.Delete(resMap, j, j+1)              // удаление свободного блока, вместо которого будет установлен файл и добор
						resMap = slices.Insert(resMap, j, partOne, partTwo) // вставка файла и остаточного пустого блока
						break

					}
				}
			}
		}
	}

	return resMap
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
