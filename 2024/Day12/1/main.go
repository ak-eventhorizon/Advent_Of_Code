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
// --- Day 12: Garden Groups --- Puzzle 1
// https://adventofcode.com/2024/day/12

func main() {
	start := time.Now()

	// input := GetData("data.txt")
	// answer := day12_1(input)
	// fmt.Println(answer)

	tmp := [][]string{}
	tmp = append(tmp, []string{"O", "O", "O", "O", "O"})
	tmp = append(tmp, []string{"O", "X", "O", "X", "O"})
	tmp = append(tmp, []string{"O", "O", "O", "O", "O"})
	tmp = append(tmp, []string{"O", "X", "O", "X", "O"})
	tmp = append(tmp, []string{"O", "O", "O", "O", "O"})
	tmp2 := extractRegions(tmp)
	fmt.Println(tmp2)

	fmt.Printf("%s \n", time.Since(start)) // измерение времени выполнения функции
}

func day12_1(input [][]string) (result int) {

	os.Truncate("output.txt", 0) //очистка файла

	gardens := extractIdenticalGardens(input)

	for _, garden := range gardens {
		transformed := transformToWeights(garden)

		for i := 0; i < len(garden); i++ {
			line := strings.Join(garden[i], "") + "  -->  " + strings.Join(transformed[i], "")
			SaveToFile(line, "output.txt")
		}

		price := calcFencePrice(transformed)
		result += price

		SaveToFile("Price: "+strconv.Itoa(price)+"\n", "output.txt")
	}

	return result
}

// Функция вычисляет стоимость забора для полученной матрицы
// Стоимость = количество_клеток_сада * количество_секций_забора
// +-+-+-+-+
// |3 2 2 3|	4 * ( 3 + 2 + 2 + 3) = 40
// +-+-+-+-+
func calcFencePrice(matrix [][]string) (result int) {

	cellCount := 0 // количество клеток, занятых садом
	fenceSum := 0  // количество секций забора, требуемых для огораживания сада

	for _, line := range matrix {
		for _, char := range line {
			if char != "." {
				val, err := strconv.Atoi(char)
				if err != nil {
					panic(err)
				}

				cellCount++
				fenceSum += val
			}
		}
	}

	return cellCount * fenceSum
}

// Функция получает карту сада и возвращает карту, в которой значения клеток сада заменеты на их вес (количество ребер, для которых требуется забор)
//
//	AAAA		+-+-+-+-+		3223
//	....		|A A A A|		....
//	....		+-+-+-+-+		....
//	....						....
func transformToWeights(matrix [][]string) (result [][]string) {

	// делаем копию исходного поля
	for _, line := range matrix {
		result = append(result, slices.Clone(line))
	}

	lenY := len(matrix)
	lenX := len(matrix[0])

	for y, line := range matrix {
		for x, char := range line {

			if char == "." {
				continue
			}

			weight := 0

			// нужен ли забор сверху
			if y == 0 || matrix[y-1][x] != char {
				weight++
			}

			// нужен ли забор снизу
			if y == lenY-1 || matrix[y+1][x] != char {
				weight++
			}

			// нужен ли забор слева
			if x == 0 || matrix[y][x-1] != char {
				weight++
			}

			// нужен ли забор справа
			if x == lenX-1 || matrix[y][x+1] != char {
				weight++
			}

			result[y][x] = strconv.Itoa(weight)
		}
	}

	return result
}

// Функция извлекает обособленные регионы из исходной матрицы
//
// .....    ...   result["O-0"]    result["X-0"]    result["X-1"]
// OOOOO    -->      OOOOO            .....            .....
// OXXOO    -->      O..OO            .XX..            .....
// OOOOO    -->      OOOOO            .....            .....
// OOOXO    -->      OOO.O            .....            ...X.
// OOOOO    -->      OOOOO            .....            .....
func extractRegions(matrix [][]string) (result map[string][][]string) {

	result = map[string][][]string{}

	lenY := len(matrix)
	lenX := len(matrix[0])

	//анализ каждой клетки исходного поля
	for y, line := range matrix {
		for x, char := range line {

			// пропускаем клетку, если она пустая
			if char == "." {
				continue
			}

			// создаем пустую матрицу (заполненную ".") по размеру исходной, на которую будем наносить каждый регион и сохранять его отдельно
			var blankMatrix [][]string
			for range lenY {
				line := []string{}
				for range lenX {
					line = append(line, ".")
				}
				blankMatrix = append(blankMatrix, line)
			}

			// поиском в ширину вычиляются все координаты региона
			start := []int{x, y}
			currentCells := [][]int{}
			currentCells = append(currentCells, start)
			visitedCells := [][]int{}
			visitedCells = append(visitedCells, start)

			for i := 1; i <= 9; i++ { // TODO цикл изменить на бесконечный с условием выхода
				currentCells = stepBFS(currentCells, matrix)
				visitedCells = append(visitedCells, currentCells...)
			}

			// сохраняем регион записью в мап

			// в исходной поле очищаем все клетки найденного региона на "." чтобы больше их не анализировать

		}
	}

	return result
}

func stepBFS(currentCells [][]int, matrix [][]string) (nextCells [][]int) {

	return nextCells
}

// На основе координат исходной клетки x,y вычисляется весь регион клеток с таким же значением, соединенный по вертикали и горизонтали
func calcNextStep(px, py int, matrix [][]string) (result [][]int) {

	regionName := matrix[py][px]
	lenY := len(matrix)
	lenX := len(matrix[0])
	currentStep := [][]int{}
	currentStep = append(currentStep, []int{px, py})
	nextStep := [][]int{}

	for _, c := range currentStep {
		x := c[0]
		y := c[1]

		// проверка соседа сверху
		if y-1 >= 0 && matrix[y-1][x] == regionName {
			nextStep = append(nextStep, []int{x, y - 1})
		}
		// проверка соседа снизу
		if y+1 < lenY && matrix[y+1][x] == regionName {
			nextStep = append(nextStep, []int{x, y + 1})
		}
		// проверка соседа слева
		if x-1 >= 0 && matrix[y][x-1] == regionName {
			nextStep = append(nextStep, []int{x - 1, y})
		}
		// проверка соседа справа
		if x+1 < lenX && matrix[y][x+1] == regionName {
			nextStep = append(nextStep, []int{x + 1, y})
		}
	}

	return result
}

// TODO -- Удалить после завершения функции extractRegions
// Функция разбирает исходную матрицу на структуру нескольких матриц, каждая из которых содержит карту только своего вида
func extractIdenticalGardens(matrix [][]string) (result map[string][][]string) {

	result = map[string][][]string{}

	lenY := len(matrix)
	lenX := len(matrix[0])

	//анализ каждой клетки исходного поля
	for y, line := range matrix {
		for x, char := range line {

			_, isKeyExistInMap := result[char]

			// если такого ключа в мапе еще нет - его нужно создать и заполнить пустой матрицей
			if !isKeyExistInMap {
				// создаем пустую матрицу (заполненную ".") по размеру исходной, на которую будем наносить каждый вид сада и сохранять его отдельно
				var blankMatrix [][]string
				for range lenY {
					line := []string{}
					for range lenX {
						line = append(line, ".")
					}
					blankMatrix = append(blankMatrix, line)
				}
				result[char] = slices.Clone(blankMatrix)
			}

			result[char][y][x] = char
		}
	}

	return result
}

// Функция извлекает из файла filename матрицу исходных данных
func GetData(filename string) (matrix [][]string) {

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var line []string
		for _, char := range scanner.Text() {
			line = append(line, string(char))
		}
		matrix = append(matrix, line)
	}

	file.Close()

	return matrix
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
