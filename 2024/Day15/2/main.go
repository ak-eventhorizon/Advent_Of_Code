package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
// layout - содержимое матрицы: @ - робот, # - стена, O - перемещаемый ящик, . - пустая клетка
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

		stackToMove = append(stackToMove, []int{f.rX, f.rY})
		allMoveElements := moveUpRecur(stackToMove, f.layout)

		fmt.Println("All blocks to move: ", allMoveElements)

		// for nextY := f.rY; nextY >= 0; nextY-- { // от робота проверяем все точки вверх, собирая ящики, которые можно передвинуть до первого свободного поля или до стены

		// 	if f.layout[nextY][f.rX] == "@" || f.layout[nextY][f.rX] == "O" {
		// 		stackToMove = append(stackToMove, []int{f.rX, nextY})
		// 	} else if f.layout[nextY][f.rX] == "#" { // уперлись в стену, нет места для перемещения
		// 		break
		// 	} else if f.layout[nextY][f.rX] == "." { // нашлось свободное место для перемещения вверх
		// 		slices.Reverse(stackToMove) // переворачиваем стек и начинаем двигать каждый элемент на свободную ячейку вверх

		// 		for _, elem := range stackToMove {
		// 			cX := elem[0]
		// 			cY := elem[1]
		// 			value := f.layout[cY][cX]
		// 			nX = elem[0]
		// 			nY = elem[1] - 1

		// 			f.layout[cY][cX] = "."
		// 			f.layout[nY][nX] = value

		// 			// если двигаем робота - обновляем его координаты
		// 			if value == "@" {
		// 				f.rX = nX
		// 				f.rY = nY
		// 			}
		// 		}
		// 		break
		// 	}
		// }

	case "v":
		for nextY := f.rY; nextY < len(f.layout)-1; nextY++ {
			if f.layout[nextY][f.rX] == "@" || f.layout[nextY][f.rX] == "O" {
				stackToMove = append(stackToMove, []int{f.rX, nextY})
			} else if f.layout[nextY][f.rX] == "#" {
				break
			} else if f.layout[nextY][f.rX] == "." {
				slices.Reverse(stackToMove) // переворачиваем стек и начинаем двигать каждый элемент на свободную ячейку вниз

				for _, elem := range stackToMove {
					cX := elem[0]
					cY := elem[1]
					value := f.layout[cY][cX]
					nX = elem[0]
					nY = elem[1] + 1

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

// TODO
func moveUpRecur(stackToMove [][]int, layout [][]string) (result [][]int) {

	var currentStack [][]int

	for _, cell := range stackToMove {

		x := cell[0]
		y := cell[1]

		if layout[y][x] == "@" && layout[y-1][x] == "[" {
			currentStack = append(currentStack, []int{x, y - 1})     // []
			currentStack = append(currentStack, []int{x + 1, y - 1}) // @
		} else if layout[y][x] == "@" && layout[y-1][x] == "]" {
			currentStack = append(currentStack, []int{x, y - 1})     // []
			currentStack = append(currentStack, []int{x - 1, y - 1}) //  @
		} else if layout[y][x] == "[" && layout[y-1][x] == "[" { // блок сверху вровень с текущим блоком
			currentStack = append(currentStack, []int{x, y - 1})     // []
			currentStack = append(currentStack, []int{x + 1, y - 1}) // [
		} else if layout[y][x] == "]" && layout[y-1][x] == "[" { // блок сверху наполовину смещен вправо
			currentStack = append(currentStack, []int{x, y - 1})     //  []
			currentStack = append(currentStack, []int{x + 1, y - 1}) //  ]
		} else if layout[y][x] == "[" && layout[y-1][x] == "]" { // блок сверху наполовину смещен влево
			currentStack = append(currentStack, []int{x, y - 1})     // []
			currentStack = append(currentStack, []int{x - 1, y - 1}) //  [
		}
	}

	if len(currentStack) == 0 {
		fmt.Println("len = 0") // DEBUG
		return currentStack
	} else {
		fmt.Println("rec call - ", currentStack) // DEBUG
		result = append(result, moveUpRecur(currentStack, layout)...)
		return result
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

	// for _, direction := range moveSet {
	// 	field.MoveRobot(direction)
	// }

	field.SaveToFile(OUTPUT_FILE_PATH)

	// вычисление суммы GPS всех ящиков по условию задачи
	// for y, line := range field.layout {
	// 	for x, char := range line {
	// 		if char == "O" {
	// 			result += (100*y + x)
	// 		}
	// 	}
	// }

	// примитивный контроллер движения робота по полю кнопками w-a-s-d (q-выход)
	isLoopActive := true
	for isLoopActive {
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()

		if err != nil {
			panic(err)
		}

		switch char {
		case 'w':
			field.MoveRobot("^")
			field.SaveToFile(OUTPUT_FILE_PATH)
		case 'a':
			field.MoveRobot("<")
			field.SaveToFile(OUTPUT_FILE_PATH)
		case 's':
			field.MoveRobot("v")
			field.SaveToFile(OUTPUT_FILE_PATH)
		case 'd':
			field.MoveRobot(">")
			field.SaveToFile(OUTPUT_FILE_PATH)
		case 'q':
			fmt.Println("QUIT")
			isLoopActive = false
		}
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
