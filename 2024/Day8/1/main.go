package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// ********* Advent of Code 2024 *********
// --- Day 8: Resonant Collinearity --- Puzzle 1
// https://adventofcode.com/2024/day/8

// Поле, с расстановкой объектов на нем
type Field struct {
	state [][]string // карта поля, на котором располагается объекты
}

func main() {

	startLayout := GetData("data.txt")
	answer := day8_1(startLayout)
	fmt.Println(answer)
}

func day8_1(layout [][]string) (result int) {

	// resField := slices.Clone(layout) // копия исходного поля, для установки на него вычисленных точек

	antennas := make(map[string][][]int) // структура для хранения всех видов антенн с координатами - map[A:[[8 1] [5 2] [7 3]] B:[[6 5] [9 9]]]

	// обход поля для составления списка всех антенн с их координатами
	for y, line := range layout {
		for x, name := range line {
			if name != "." {
				// проверка - существует ли уже в перечне антенна такого вида
				if _, isPresent := antennas[name]; isPresent { // добавить текущую к существующему набору
					antennas[name] = append(antennas[name], []int{x, y})
				} else { // создать набор антенн такого вида и добавить к нему текущую
					antennas[name] = make([][]int, 0)
					antennas[name] = append(antennas[name], []int{x, y})
				}
			}
		}
	}

	// обход всех видов антенн. по каждому виду построение отражения (антиноды) для каждой антенны от каждой другой
	for k, v := range antennas {
		antinodes := findAllAntiNodes(v)
		fmt.Println(k, " --> ", v, "--antinodes-->", antinodes) // DEBUG print
		// TODO нарисовать антиноды на resField ----------> #
	}

	return result
}

// Функция получает набор антенн [x,y], возвращает набор антинод для входного набора антенн, вычисленных по условиям задачи
func findAllAntiNodes(antennas [][]int) (antinodes [][]int) {

	for i := range antennas {

		tmp := slices.Clone(antennas)
		currentAntenna := antennas[i]
		restAntens := slices.Delete(tmp, i, i+1)

		for _, secondAntenna := range restAntens {
			antinodes = append(antinodes, reflect(currentAntenna, secondAntenna))
			antinodes = append(antinodes, reflect(secondAntenna, currentAntenna))
		}

	}

	return antinodes
}

// Функция отражает точку p2[x2,y2] относительно точки p1[x1,y1] и возвращает координаты отраженной точки [xr,yr]
func reflect(p1, p2 []int) (r []int) {

	x1 := p1[0]
	y1 := p1[1]

	x2 := p2[0]
	y2 := p2[1]

	deltaX := x1 - x2
	deltaY := y1 - y2

	r = make([]int, 2)

	r[0] = x1 + deltaX
	r[1] = y1 + deltaY

	return r
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
