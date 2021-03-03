package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func parse(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error: format: int '<operator>' int")
		os.Exit(1)
	}
	return value
}

type calc struct {
	intA     int
	intB     int
	operator string
}

func (self calc) operate() int {
	total := 0

	switch self.operator {
	case "+":
		total = self.intA + self.intB
	case "-":
		total = self.intA - self.intB
	case "*":
		total = self.intA * self.intB
	case "/":
		total = self.intA / self.intB
	default:
		total = 0
	}

	return total
}
