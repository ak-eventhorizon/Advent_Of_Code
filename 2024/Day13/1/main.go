package main

import (
	"bufio"
	"fmt"
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
	fmt.Printf("A(%v,%v) B(%v,%v) Claw (%v,%v) Prize(%v,%v)\n",
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

	// DEBUG - тестируем на одной машине из списка
	// result = calcCheapestWinCombination(machines[0])

	for _, m := range machines {
		result += calcCheapestWinCombination(m)
	}

	return result
}

// Функция вычисляет самую дешевую комбинацию нажатия кнопок A в B на игровом автомате, требуемую для перемещение клешни в координаты приза
// Одно нажатие кнопки A стоит 3 монеты, одно нажатие кнопки B стоит 1 монету
// Если комбинации не существут - функция возвращает 0
func calcCheapestWinCombination(machine Machine) (result int) {

	// успешные комбинации нажатия кнопок [A,B]
	var success [][]int

	// пробуем все комбинации нажатий кнопок от "ни разу" до лимита в 100 раз (лимит из условия задачи), сохраняем те, которые приводят к победе
	for countA := 0; countA <= 100; countA++ {

		if countA != 0 {
			machine.PushButtonA()
		}

		if machine.Claw.x == machine.Prize.x && machine.Claw.y == machine.Prize.y { // попали в приз нажатием только кнопки А
			success = append(success, []int{countA, 0})
			break
		}

		if machine.Claw.x > machine.Prize.x || machine.Claw.y > machine.Prize.y { // проскочили приз нажатием только кнопки А
			break
		}

		savedClawX := machine.Claw.x // сохраняем координаты клешни, для ее возвращения в исходную позицию после отработки вложенного цикла
		savedClawY := machine.Claw.y

		for countB := 1; countB <= 100; countB++ {

			machine.PushButtonB()

			// DEBUG
			// var ding string
			// if machine.Claw.x == machine.Prize.x && machine.Claw.y == machine.Prize.y {
			// 	ding = "<<<<<<< DING!!!"
			// }
			// s := fmt.Sprintf("A(%v) B(%v) Claw:(%v %v) -- Prize:(%v %v) %v", countA, countB, machine.Claw.x, machine.Claw.y, machine.Prize.x, machine.Prize.y, ding)
			// SaveToFile(s, OUTPUT_FILE_PATH)

			if machine.Claw.x == machine.Prize.x && machine.Claw.y == machine.Prize.y { // попали в приз
				success = append(success, []int{countA, countB})
				machine.Claw.x = savedClawX // возвращаем клешню на состояние до внутреннего цикла
				machine.Claw.y = savedClawY
				break
			}

			if machine.Claw.x > machine.Prize.x || machine.Claw.y > machine.Prize.y { // проскочили приз
				machine.Claw.x = savedClawX // возвращаем клешню на состояние до внутреннего цикла
				machine.Claw.y = savedClawY
				break
			}
		}
	}

	// поиск самого дешевого решения
	if len(success) != 0 {

		// fmt.Println("Win combinations: ", success) // DEBUG

		for _, combination := range success {

			costOfWin := combination[0]*3 + combination[1]*1

			// fmt.Print("For A = ", combination[0], " B = ", combination[1], " ") // DEBUG
			// fmt.Println("Cost of Win =", costOfWin)                             // DEBUG

			if result == 0 {
				result = costOfWin
			} else if costOfWin < result {
				result = costOfWin
			}
		}
	}

	return result
}

//TODO -- сделать вторую функцию calcCheapestWinCombination, вычисляющуу результат по методу Крамера. Сравнить результаты функций

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
