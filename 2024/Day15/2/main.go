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
// --- Day 15: Warehouse Woes --- Puzzle 2
// https://adventofcode.com/2024/day/15

const INPUT_FILE_PATH string = "data.txt"
const OUTPUT_FILE_PATH string = "output.txt"

// Структура представляет поле (двумерную матрицу)
// rX, rY - координаты робота @
// layout - содержимое матрицы: @ - робот, # - стена, [] - перемещаемый ящик, . - пустая клетка
type Field struct {
	rX     int
	rY     int
	layout [][]string
}

func (f *Field) SaveToFile(filename string) {

	os.Truncate(filename, 0) // очистка файла

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	for _, line := range f.layout {
		s := strings.Join(line, "")
		_, err2 := file.WriteString(s + "\n")
		if err2 != nil {
			panic(err2)
		}
	}
}

func (f *Field) MoveRobot(direction string) {

	var nX, nY int          // координаты точки, в которую произойдет движение
	var stackToMove [][]int // набор точек [x,y] который нужно сдвинуть

	switch direction {

	case "^":
		nX = f.rX
		nY = f.rY - 1

		if f.layout[nY][nX] == "." { // шаг робота в свободую клетку
			f.layout[f.rY][f.rX] = "."
			f.rX = nX
			f.rY = nY
			f.layout[f.rY][f.rX] = "@"
		} else if f.layout[nY][nX] == "#" { // шаг робота в стену
			break
		} else if f.layout[nY][nX] == "[" || f.layout[nY][nX] == "]" { // шаг робота в перемещаемый ящик
			stackToMove = append(stackToMove, []int{f.rX, f.rY})
			allBoxesToMove := moveUpRecur(stackToMove, f.layout)
			stackToMove = append(stackToMove, allBoxesToMove...) // все ящики цепочки и сам робот

			// проверка, что ни один из объектов цепочки не упирается в стену, которая мешает перемещению всей цепочки
			isStackMovable := true
			for _, elem := range stackToMove {
				x := elem[0]
				y := elem[1]
				if f.layout[y-1][x] == "#" {
					isStackMovable = false
					break
				}
			}

			// если цепочку можно сдвинуть вверх - двигаем каждый ее элемент, начиная с конца
			if isStackMovable {
				slices.Reverse(stackToMove) // переворачиваем стек и начинаем двигать каждый элемент на свободную ячейку вверх
				for _, elem := range stackToMove {
					x := elem[0]
					y := elem[1]
					value := f.layout[y][x]
					nX = elem[0]
					nY = elem[1] - 1

					f.layout[y][x] = "."
					f.layout[nY][nX] = value

					// если двигаем робота - обновляем его координаты
					if value == "@" {
						f.rX = nX
						f.rY = nY
					}
				}
			}
		}

	case "v":
		nX = f.rX
		nY = f.rY + 1

		if f.layout[nY][nX] == "." { // шаг робота в свободую клетку
			f.layout[f.rY][f.rX] = "."
			f.rX = nX
			f.rY = nY
			f.layout[f.rY][f.rX] = "@"
		} else if f.layout[nY][nX] == "#" { // шаг робота в стену
			break
		} else if f.layout[nY][nX] == "[" || f.layout[nY][nX] == "]" { // шаг робота в перемещаемый ящик
			stackToMove = append(stackToMove, []int{f.rX, f.rY})
			allBoxesToMove := moveDownRecur(stackToMove, f.layout)
			stackToMove = append(stackToMove, allBoxesToMove...) // все ящики цепочки и сам робот

			// проверка, что ни один из объектов цепочки не упирается в стену, которая мешает перемещению всей цепочки
			isStackMovable := true
			for _, elem := range stackToMove {
				x := elem[0]
				y := elem[1]
				if f.layout[y+1][x] == "#" {
					isStackMovable = false
					break
				}
			}

			// если цепочку можно сдвинуть вверх - двигаем каждый ее элемент, начиная с конца
			if isStackMovable {
				slices.Reverse(stackToMove) // переворачиваем стек и начинаем двигать каждый элемент на свободную ячейку вверх
				for _, elem := range stackToMove {
					x := elem[0]
					y := elem[1]
					value := f.layout[y][x]
					nX = elem[0]
					nY = elem[1] + 1

					f.layout[y][x] = "."
					f.layout[nY][nX] = value

					// если двигаем робота - обновляем его координаты
					if value == "@" {
						f.rX = nX
						f.rY = nY
					}
				}
			}
		}

	case "<":
		for nextX := f.rX; nextX >= 0; nextX-- {
			if f.layout[f.rY][nextX] == "@" || f.layout[f.rY][nextX] == "[" || f.layout[f.rY][nextX] == "]" {
				stackToMove = append(stackToMove, []int{nextX, f.rY})
			} else if f.layout[f.rY][nextX] == "#" {
				break
			} else if f.layout[f.rY][nextX] == "." {
				slices.Reverse(stackToMove) // переворачиваем стек и начинаем двигать каждый элемент на свободную ячейку влево

				for _, elem := range stackToMove {
					cX := elem[0]
					cY := elem[1]
					value := f.layout[cY][cX]
					nX = elem[0] - 1
					nY = elem[1]

					f.layout[cY][cX] = "."
					f.layout[nY][nX] = value

					// если двигаем робота - обновляем его координаты
					if value == "@" {
						f.rX = nX
						f.rY = nY
					}
				}
				break
			}
		}

	case ">":
		for nextX := f.rX; nextX < len(f.layout[0])-1; nextX++ {
			if f.layout[f.rY][nextX] == "@" || f.layout[f.rY][nextX] == "[" || f.layout[f.rY][nextX] == "]" {
				stackToMove = append(stackToMove, []int{nextX, f.rY})
			} else if f.layout[f.rY][nextX] == "#" {
				break
			} else if f.layout[f.rY][nextX] == "." {
				slices.Reverse(stackToMove) // переворачиваем стек и начинаем двигать каждый элемент на свободную ячейку вправо

				for _, elem := range stackToMove {
					cX := elem[0]
					cY := elem[1]
					value := f.layout[cY][cX]
					nX = elem[0] + 1
					nY = elem[1]

					f.layout[cY][cX] = "."
					f.layout[nY][nX] = value

					// если двигаем робота - обновляем его координаты
					if value == "@" {
						f.rX = nX
						f.rY = nY
					}
				}
				break
			}
		}
	}
}

func main() {
	start := time.Now()

	field, moveSet := GetData(INPUT_FILE_PATH)
	answer := day15_2(field, moveSet)
	fmt.Println(answer)

	fmt.Printf("%s \n", time.Since(start)) // время выполнения функции
}

func day15_2(field Field, moveSet []string) (result int) {

	for _, direction := range moveSet {
		field.MoveRobot(direction)
	}

	field.SaveToFile(OUTPUT_FILE_PATH)

	// вычисление суммы GPS всех ящиков по условию задачи
	for y, line := range field.layout {
		for x, char := range line {
			if char == "[" {
				result += (100*y + x)
			}
		}
	}

	// // примитивный контроллер движения робота по полю кнопками w-a-s-d (q-выход)
	// isLoopActive := true
	// for isLoopActive {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	char, _, err := reader.ReadRune()

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	switch char {
	// 	case 'w':
	// 		field.MoveRobot("^")
	// 		field.SaveToFile(OUTPUT_FILE_PATH)
	// 	case 'a':
	// 		field.MoveRobot("<")
	// 		field.SaveToFile(OUTPUT_FILE_PATH)
	// 	case 's':
	// 		field.MoveRobot("v")
	// 		field.SaveToFile(OUTPUT_FILE_PATH)
	// 	case 'd':
	// 		field.MoveRobot(">")
	// 		field.SaveToFile(OUTPUT_FILE_PATH)
	// 	case 'q':
	// 		fmt.Println("QUIT")
	// 		isLoopActive = false
	// 	}
	// }

	return result
}

// Функция рекурсивно строит цепочку из всех ящиков поля layout, которые можно сдвинуть вверх
func moveUpRecur(stackToMove [][]int, layout [][]string) (result [][]int) {

	var nextIter [][]int // набор элементов для перемещения в следующей итерации

	for _, cell := range stackToMove {

		x := cell[0]
		y := cell[1]

		if layout[y][x] == "@" && layout[y-1][x] == "[" {
			nextIter = append(nextIter, []int{x, y - 1})     // []
			nextIter = append(nextIter, []int{x + 1, y - 1}) // @
		} else if layout[y][x] == "@" && layout[y-1][x] == "]" {
			nextIter = append(nextIter, []int{x, y - 1})     // []
			nextIter = append(nextIter, []int{x - 1, y - 1}) //  @
		} else if layout[y][x] == "]" && layout[y][x+1] == "[" && layout[y-1][x] == "[" {
			nextIter = append(nextIter, []int{x, y - 1})     //  []
			nextIter = append(nextIter, []int{x + 1, y - 1}) // >][
		} else if layout[y][x] == "[" && layout[y][x-1] == "]" && layout[y-1][x] == "]" {
			nextIter = append(nextIter, []int{x, y - 1})     // []
			nextIter = append(nextIter, []int{x - 1, y - 1}) // ][<
		} else if layout[y][x] == "[" && layout[y-1][x] == "[" { // блок сверху вровень с текущим блоком
			nextIter = append(nextIter, []int{x, y - 1})     // []
			nextIter = append(nextIter, []int{x + 1, y - 1}) // [
		} else if layout[y][x] == "]" && layout[y-1][x] == "[" { // блок сверху наполовину смещен вправо
			nextIter = append(nextIter, []int{x, y - 1})     //  []
			nextIter = append(nextIter, []int{x + 1, y - 1}) //  ]
		} else if layout[y][x] == "[" && layout[y-1][x] == "]" { // блок сверху наполовину смещен влево
			nextIter = append(nextIter, []int{x, y - 1})     // []
			nextIter = append(nextIter, []int{x - 1, y - 1}) //  [
		}
	}

	nextIter = removeDuplicates(nextIter)

	if len(nextIter) == 0 {
		return nextIter
	}

	// fmt.Println("rec call - ", nextIter) // DEBUG
	nextIter = append(nextIter, moveUpRecur(nextIter, layout)...)
	return nextIter
}

// Функция рекурсивно строит цепочку из всех ящиков поля layout, которые можно сдвинуть вниз
func moveDownRecur(stackToMove [][]int, layout [][]string) (result [][]int) {

	var nextIter [][]int // набор элементов для перемещения в следующей итерации

	for _, cell := range stackToMove {

		x := cell[0]
		y := cell[1]

		if layout[y][x] == "@" && layout[y+1][x] == "[" {
			nextIter = append(nextIter, []int{x, y + 1})     // @
			nextIter = append(nextIter, []int{x + 1, y + 1}) // []
		} else if layout[y][x] == "@" && layout[y+1][x] == "]" {
			nextIter = append(nextIter, []int{x, y + 1})     //  @
			nextIter = append(nextIter, []int{x - 1, y + 1}) // []
		} else if layout[y][x] == "]" && layout[y][x+1] == "[" && layout[y+1][x] == "[" {
			nextIter = append(nextIter, []int{x, y + 1})     // >][
			nextIter = append(nextIter, []int{x + 1, y + 1}) //  []
		} else if layout[y][x] == "[" && layout[y][x-1] == "]" && layout[y+1][x] == "]" {
			nextIter = append(nextIter, []int{x, y + 1})     // ][<
			nextIter = append(nextIter, []int{x - 1, y + 1}) // []
		} else if layout[y][x] == "[" && layout[y+1][x] == "[" { // блок снизу вровень с текущим блоком
			nextIter = append(nextIter, []int{x, y + 1})     // [
			nextIter = append(nextIter, []int{x + 1, y + 1}) // []
		} else if layout[y][x] == "]" && layout[y+1][x] == "[" { // блок снизу наполовину смещен вправо
			nextIter = append(nextIter, []int{x, y + 1})     //  ]
			nextIter = append(nextIter, []int{x + 1, y + 1}) //  []
		} else if layout[y][x] == "[" && layout[y+1][x] == "]" { // блок снизу наполовину смещен влево
			nextIter = append(nextIter, []int{x, y + 1})     //  [
			nextIter = append(nextIter, []int{x - 1, y + 1}) // []
		}
	}

	nextIter = removeDuplicates(nextIter)

	if len(nextIter) == 0 {
		return nextIter
	}

	// fmt.Println("rec call - ", nextIter) // DEBUG
	nextIter = append(nextIter, moveDownRecur(nextIter, layout)...)
	return nextIter
}

// Функция удаляет дубли из произвольного слайса слайсов целых чисел
func removeDuplicates(source [][]int) (result [][]int) {

	// превращение слайса слайсов интов в слайс строк, который можно отсортировать и удалить дубли в сортированном
	strSlice := []string{}
	for _, elem := range source {
		strElem := strings.Trim(strings.Replace(fmt.Sprint(elem), " ", "-", -1), "[]") // [3 4] --> "3-4"
		strSlice = append(strSlice, strElem)
	}

	slices.Sort(strSlice)               // сортировка
	strSlice = slices.Compact(strSlice) // удаление дубликатов

	// превращение слайса строк обратно в слайс слайсов интов
	for _, s := range strSlice {
		elements := strings.Split(s, "-")
		element := []int{}
		for _, v := range elements {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			element = append(element, num)
		}
		result = append(result, element)
	}

	return result
}

// Функция извлекает из файла filename набор исходных данных
func GetData(filename string) (field Field, moveSet []string) {

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var isMatrixFinished bool // признак того, что матрицу с игровым полем полностью извлекли из файла и начинается вторая часть файла

	for scanner.Scan() {

		if scanner.Text() == "" {
			isMatrixFinished = true
			continue
		}

		if !isMatrixFinished {
			var transformedLine []string
			line := strings.Split(scanner.Text(), "")

			// По условиям задачи 15_2 - исходная матрица подлежит трансформации по набору правил
			for _, char := range line {
				if char == "#" {
					transformedLine = append(transformedLine, "#", "#")
				} else if char == "O" {
					transformedLine = append(transformedLine, "[", "]")
				} else if char == "." {
					transformedLine = append(transformedLine, ".", ".")
				} else if char == "@" {
					transformedLine = append(transformedLine, "@", ".")
				}
			}

			field.layout = append(field.layout, transformedLine)
		} else {
			line := strings.Split(scanner.Text(), "")
			moveSet = append(moveSet, line...)
		}
	}

	file.Close()

	// Поиск координат робота на поле
	for y, line := range field.layout {
		for x, char := range line {
			if char == "@" {
				field.rX = x
				field.rY = y
			}
		}
	}

	return field, moveSet
}

// Функция дописывает строку line в файл filename
func SaveToFile(line string, filename string) {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err2 := file.WriteString(line + "\n")

	if err2 != nil {
		panic(err2)
	}
}
