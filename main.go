package main

import (
	"fmt"
	"strconv"
	"strings"
)

const op = "+-*/"
const rome = "XIV"

var numRome = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

func parsSign(a string) string {
	if strings.ContainsAny(a, op) {
		return a
	}
	panic("На вход поступил неверный арифметический знак")
}
func getNums(x, y, s string) interface{} {
	if strings.ContainsAny(rome, x) && strings.ContainsAny(rome, y) {
		return parsRome(x, y, s)
	}
	first, err1 := strconv.Atoi(x)
	second, err2 := strconv.Atoi(y)
	if err1 != nil || err2 != nil {
		panic("Ожидаются только арбские либо только рисмские числа которые не больше 10")
	}

	return getOp(first, second, s)
}

func parsRome(x, y, s string) string {
	l, ok1 := numRome[x]
	r, ok2 := numRome[y]
	if !ok1 || !ok2 {
		panic("Ожидаются только арбские либо только рисмские числа которые не больше 10")
	} else if (l - r) < 0 {
		panic("В римской системе нет отрицательных чисел.")
	}
	num := getOp(l, r, s)
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	res := strings.Builder{}
	for i := 0; i < len(symbols); i++ {
		for num >= values[i] {
			res.WriteString(symbols[i])
			num -= values[i]
		}
	}
	return res.String()

}

func getOp(x, y int, s string) int {
	var res int
	switch s {
	case "+":
		res = x + y
	case "-":
		res = x - y
	case "/":
		res = x / y
	case "*":
		res = x * y
	}
	return res
}

func main() {
	var x, s, y string
	_, err := fmt.Scan(&x, &s, &y)
	if err != nil {
		return
	}
	sign := parsSign(s)
	fmt.Println(getNums(x, y, sign))
}
