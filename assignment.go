package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse(text string) (number1 float64, number2 float64, operator string, err error) {
	splits := strings.Split(text, " ")

	if len(splits) != 3 {
		err = errors.New("Wrong arguments")
		return
	}

	number1, err = strconv.ParseFloat(splits[0], 64)

	if err != nil {
		err = errors.New("Parse float error")
		return
	}

	number2, err = strconv.ParseFloat(splits[2], 64)
	operator = splits[1]

	if err != nil {
		err = errors.New("Parse float error")
		return
	}

	return
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover with err: ", err)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		number1, number2, operator, err := Parse(text)

		if err != nil {
			panic(err)
		}

		switch operator {
		case "+":
			fmt.Println(number1 + number2)
		case "-":
			fmt.Println(number1 - number2)
		case "*":
			fmt.Println(number1 * number2)
		case "/":
			if number2 == 0 {
				panic("Can not divide to 0")
			}
			fmt.Println(number1 / number2)
		default:
			panic("Your input operator has not been supported")
		}

	}
}
