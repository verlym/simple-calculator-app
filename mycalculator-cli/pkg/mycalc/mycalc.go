package mycalc

import (
	"fmt"
	"strconv"
)

func Add(args []string) (sum int) {
	for _, i := range args {
		realInt, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println(err)
		}
		sum += realInt
	}
	return sum
}

func Subtract(args []string) (diff int) {
	num1, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err)
	}

	num2, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
	}

	return num1 - num2
}

func Multiply(args []string) (product int) {
	product = 1
	for _, i := range args {
		realInt, err := strconv.Atoi(i)

		if err != nil {
			fmt.Println(err)
		}
		product *= realInt
	}

	return product
}

func Divide(args []string) (dividend int) {
	num1, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err)
	}

	num2, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
	}

	return num1 / num2
}