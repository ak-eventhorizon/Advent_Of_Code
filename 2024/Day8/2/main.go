package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// ********* Advent of Code 2024 *********
// --- Day 8: Resonant Collinearity --- Puzzle 2
// https://adventofcode.com/2024/day/8

func main() {

	startLayout := GetData("data.txt")
	answer := day8_2(startLayout)
	fmt.Println(answer)
}

func day8_2(layout [][]string) (result int) {

	resultField := slices.Clone(layout) // копия исходного поля, для установки на него вычисленных точек

	antennas := make(map[string][][]int) // структура для хранения всех видов антенн с координатами - map[A:[[8 1] [5 2] [7 3]] B:[[6 5] [9 9]]]

	// обход поля для составления списка всех антенн с их координатами
	for y, line := range layout {
		for x, name := range line {
			if name != "." {
				// проверка - существует ли уже в перечне антенна такого вида
				if _, isPresent := antennas[name]; isPresent { // добавить координаты текущей к существующему набору
					antennas[name] = append(antennas[name], []int{x, y})
				} else { // создать набор антенн такого вида и добавить к нему координаты текущей
					antennas[name] = make([][]int, 0)
					antennas[name] = append(antennas[name], []int{x, y})
				}
			}
		}
	}

	// обход всех видов антенн. по каждому виду построение всех антинод для каждой антенны относительно каждой другой
	for _, v := range antennas {

		antinodes := findAllAntiNodes(v, len(layout[0]), len(layout))
		// fmt.Println(k, " --> ", v, "--antinodes-->", antinodes) // DEBUG print

		// нарисовать антиноды # на resultField
		for _, antinode := range antinodes {

			x := antinode[0]
			y := antinode[1]

			resultField[y][x] = "#" // установка антиноды на поле (антинода может находиться на одной точке с антенной)

		}
	}

	// вычисление сколько всего антинод установлено на поле
	for _, line := range resultField {
		for _, char := range line {
			if char == "#" {
				result++
			}
		}
	}

	SaveData(resultField, "output.txt")
	return result
}

// Функция получает набор антенн [[x1 y1] [x2 y2] [x3 y3]...] и размерность поля sizeX, sizeY
// Возвращает набор антинод для входного набора антенн, вычисленных по условиям задачи
func findAllAntiNodes(antennas [][]int, sizeX, sizeY int) (antinodes [][]int) {

	for i := range antennas {

		tmp := slices.Clone(antennas)
		currentAntenna := antennas[i]
		restAntennas := slices.Delete(tmp, i, i+1) // удаление текущей антенны из небора

		for _, secondAntenna := range restAntennas {
			antinodes = append(antinodes, reflectVector(currentAntenna, secondAntenna, sizeX, sizeY)...)
			antinodes = append(antinodes, reflectVector(secondAntenna, currentAntenna, sizeX, sizeY)...)
		}

	}

	return antinodes
}

// Функция вычисляет координаты всей цепочки антинод, которые можно построить внутри поля (размерностью sizeX на sizeY) на основании точек p1 и p2 относительно точки p1
//
// ....................   ->   ....................
// .......1............   ->   ....................
// ....................   ->   ....................
// .........2..........   ->   .........#..........
// ....................   ->   ....................
// ....................   ->   ...........#........
// ....................   ->   ....................
// ....................   ->   .............#......
// ....................   ->   ....................
// ....................   ->   ...............#....
func reflectVector(p1, p2 []int, sizeX, sizeY int) (antinodes [][]int) {

	antinodes = make([][]int, 0)

	x1, y1 := p1[0], p1[1]
	x2, y2 := p2[0], p2[1]

	deltaX := x1 - x2
	deltaY := y1 - y2

	// построение цепочки отражений относительно точки p1 по вычисленному вектору
	antinodes = append(antinodes, []int{x2, y2}) // первая антинода для точки p1 - это всегда точка p2
	nextX := x1 + deltaX
	nextY := y1 + deltaY
	for {
		if nextX < 0 || nextX >= sizeX || nextY < 0 || nextY >= sizeY { // выход следующей точки за пределы обозначенного поля
			break
		}
		antinodes = append(antinodes, []int{nextX, nextY})
		nextX += deltaX
		nextY += deltaY
	}

	return antinodes
}

// Функция извлекает из текстового файла карту задачи.
func GetData(filename string) (layout [][]string) {

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
		layout = append(layout, line)
	}

	file.Close()

	return layout
}

// Функция сохраняет полученную карту layout в файл filename.
func SaveData(layout [][]string, filename string) {

	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	for i, line := range layout {
		var s string
		for _, char := range line {
			s += char
		}
		file.WriteString(s)

		if i != len(layout)-1 {
			file.WriteString("\n")
		}
	}

	file.Close()
}
