package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var operators = map[rune]bool{
	'+': true,
	'-': true,
	'*': true,
	'/': true,
}

var romanNumbers = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}

var numToRoman = map[int]string{
	100: "C",
	90:  "XC",
	50:  "L",
	40:  "XL",
	10:  "X",
	9:   "IX",
	8:   "VIII",
	7:   "VII",
	6:   "VI",
	5:   "V",
	4:   "IV",
	3:   "III",
	2:   "II",
	1:   "I",
}

func strToInt(operandStr string) (int, string) {
	number, isRoman := romanNumbers[operandStr]
	numberFormat := "roman"
	if !isRoman {
		numberFormat = "arabic"
		var err error
		number, err = strconv.Atoi(operandStr)
		if err != nil {
			panic("Операнд не является числом в арабской форме записи: " + operandStr)
		}
	}

	if !(number >= 1 && number <= 10) {
		panic("Операнд находится вне диапазона значений от 1 до 10: " + operandStr)
	}

	return number, numberFormat
}


func getOperator(s string) string {
	var operators = []string{"+", "-", "*", "/"}
	var operatorList []string
	var operator string
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(operators); j++ {
			if string(s[i]) == operators[j] {
				operatorList = append(operatorList, operators[j])

			}

		}
	}
	operator = operatorList[0]
	if len(operatorList) != 1 {
		panic("Ошибка ввода операции!!!")
	}
	return operator
}

func getOperands(s string, operator string) (operands []int, format string) {
	operandsList := strings.Split(s, operator)
	if len(operandsList) != 2 {
		panic(fmt.Sprintf("Количество операндов должно быть равно 2"))
	}

	operands = make([]int, 0, 2)
	for _, operandStr := range operandsList {
		number, numberFormat := strToInt(operandStr)
		if format == "" {
			format = numberFormat
		} else if numberFormat != format {
			panic("Форма записи чисел у операндов не совпадает: " + numberFormat + " != " + format)
		}

		operands = append(operands, number)
	}

	return operands, format
}

func printResult(number int, numberFormat string) {
	if numberFormat == "arabic" {
		fmt.Println(number)
	} else if numberFormat == "roman" {
		if number < 1 {
			panic("Результат должен быть положительным числом большим 1")
		}
		fmt.Println(toRomanStr(number))
	}
}

func calculate(s string) {
	operator := getOperator(s)
	operands, numberFormat := getOperands(s, operator)

	var result int
	x, y := operands[0], operands[1]

	switch operator {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		result = x / y
	default:
		panic("Задана неверная операция")
	}

	printResult(result, numberFormat)
}

var romans []int = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}

func toRomanStr(n int) string {
	romanStr := ""

	for i := len(romans) - 1; n > 0; i -= 1 {
		div := n / romans[i]
		n = n % romans[i]

		for ; div > 0; div -= 1 {
			romanStr += numToRoman[romans[i]]
		}
	}

	return romanStr
}

func main() {
	fmt.Println("Введите два числа от 1 до 10 и одну простейшую математическую операцию между ними")
	reader := bufio.NewReader(os.Stdin)
	console, _ := reader.ReadString('\n')
	s := strings.ReplaceAll(console, " ", "")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	calculate(strings.ToUpper(s))

}
