package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

// ********* Advent of Code 2024 *********
// --- Day 13: Claw Contraption --- Puzzle 1
// https://adventofcode.com/2024/day/13

const INPUT_FILE_PATH string = "data.txt"
const OUTPUT_FILE_PATH string = "output.txt"

// Структура представляет точку на двумерной плоскости
type Point struct {
	x int
	y int
}

// Структура представляет игровой автомат, имеющий кнопки A и B, клешню Claw и приз Prize
// При нажатии кнопки А: происходит смещение клешни на AX по оси X и на AY по оси Y
// При нажатии кнопки B: происходит смещение клешни на BX по оси X и на BY по оси Y
type Machine struct {
	AX, AY int
	BX, BY int
	Claw   Point
	Prize  Point
}

func (m *Machine) PushButtonA() {
	m.Claw.x += m.AX
	m.Claw.y += m.AY
}

func (m *Machine) PushButtonB() {
	m.Claw.x += m.BX
	m.Claw.y += m.BY
}

func (m Machine) Display() {
	fmt.Printf("A(%v,%v) B(%v,%v) Claw (%v,%v) Prize(%v,%v)",
		m.AX,
		m.AY,
		m.BX,
		m.BY,
		m.Claw.x,
		m.Claw.y,
		m.Prize.x,
		m.Prize.y)
}

func main() {
	start := time.Now()

	input := GetData(INPUT_FILE_PATH)
	answer := day13_1(input)
	fmt.Println(answer)

	fmt.Printf("%s \n", time.Since(start)) // время выполнения функции
}

func day13_1(machines []Machine) (result int) {

	var bruteforce int
	var cramer int

	for _, m := range machines {
		bruteforce += calcCheapestWinCombination_bruteForce(m)
		cramer += calcCheapestWinCombination_Cramer(m)
	}

	fmt.Println("BruteForce: ", bruteforce)
	fmt.Println("Cramer:     ", cramer)

	return cramer
}

// Функция вычисляет самую дешевую комбинацию нажатия кнопок A в B на игровом автомате, требуемую для перемещение клешни в координаты приза
// Если комбинации не существут - функция возвращает 0
// Реализована перебором всех комбинаций
func calcCheapestWinCombination_bruteForce(machine Machine) (result int) {

	const BUTTON_A_COST = 3 // Стоимость нажатия кнопки A
	const BUTTON_B_COST = 1 // Стоимость нажатия кнопки B

	var success [][]int // успешные комбинации нажатия кнопок [A,B]

	for A := 0; A <= 100; A++ {
		for B := 0; B <= 100; B++ {

			// Копия машины, чтобы не изменять исходную
			tempClaw := Point{machine.Claw.x, machine.Claw.y}
			tempPrize := Point{machine.Prize.x, machine.Prize.y}
			tempMachine := Machine{machine.AX, machine.AY, machine.BX, machine.BY, tempClaw, tempPrize}

			for range A {
				tempMachine.PushButtonA()
			}

			for range B {
				tempMachine.PushButtonB()
			}

			if tempMachine.Claw.x == tempMachine.Prize.x && tempMachine.Claw.y == tempMachine.Prize.y { // попали в приз
				success = append(success, []int{A, B})
				break
			}

			if tempMachine.Claw.x > tempMachine.Prize.x || tempMachine.Claw.y > tempMachine.Prize.y { // проскочили приз
				break
			}
		}
	}

	// поиск самого дешевого решения
	if len(success) != 0 {
		for _, combination := range success {

			costOfWin := combination[0]*BUTTON_A_COST + combination[1]*BUTTON_B_COST

			if result == 0 {
				result = costOfWin
			} else if costOfWin < result {
				result = costOfWin
			}
		}
	}

	return result
}

// Функция вычисляет самую дешевую комбинацию нажатия кнопок A в B на игровом автомате, требуемую для перемещение клешни в координаты приза
// Если комбинации не существут - функция возвращает 0
// Реализована с использованием правила Крамера - способа решения систем линейных уравнений с числом уравнений равным числу неизвестных
func calcCheapestWinCombination_Cramer(machine Machine) (result int) {

	const BUTTON_A_COST = 3 // Стоимость нажатия кнопки A
	const BUTTON_B_COST = 1 // Стоимость нажатия кнопки B

	var A, B float64 // количество нажатий кнопок A и B (их требуется найти)

	Ax, Ay := machine.AX, machine.AY           // смещение по осям x и y при нажатии кнопки A
	Bx, By := machine.BX, machine.BY           // смещение по осям x и y при нажатии кнопки B
	Px, Py := machine.Prize.x, machine.Prize.y // координаты приза по осям x и y

	// Задача может быть выражена следующей системой уравнений:
	// Ax*A + Bx*B = Px
	// Ay*A + By*B = Py
	// Система решается с использованием метода Крамера

	D := (Ax * By) - (Ay * Bx)   // Главная детерминанта
	D_a := (Px * By) - (Py * Bx) // Детерминанта для A
	D_b := (Py * Ax) - (Px * Ay) // Детерминанта для B

	A = float64(D_a) / float64(D)
	B = float64(D_b) / float64(D)

	// Если A и B - целые, то можно выиграть приз для данной игровой машины
	if math.Floor(A) == A && math.Floor(B) == B {
		costOfWin := A*BUTTON_A_COST + B*BUTTON_B_COST
		result = int(costOfWin)
	}

	return result
}

// Функция извлекает из файла filename набор исходных данных
func GetData(filename string) (machines []Machine) {

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var acc []string

	for scanner.Scan() {

		patternPrize := regexp.MustCompile(`Prize:`) // содержит ли строка "Prize:", если содержит - то это последняя строка, описывающая автомат
		strPrize := patternPrize.FindAllString(scanner.Text(), -1)

		patternNums := regexp.MustCompile(`\d+`) // выхватывает из строки числа
		numsInLine := patternNums.FindAllString(scanner.Text(), -1)

		acc = append(acc, numsInLine...)

		if len(strPrize) != 0 { // это последняя строка в блоке, описывающем отдельный автомат

			var nums []int

			for _, s := range acc {
				num, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				nums = append(nums, num)
			}

			ax := nums[0]
			ay := nums[1]
			bx := nums[2]
			by := nums[3]
			prizeX := nums[4]
			prizeY := nums[5]

			claw := Point{0, 0}
			prize := Point{prizeX, prizeY}
			machine := Machine{ax, ay, bx, by, claw, prize}

			machines = append(machines, machine)

			acc = []string{} // сбрасываем накопитель значений для следующей итерации
		}
	}

	file.Close()

	return machines
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
