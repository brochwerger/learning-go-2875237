package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	float1, _ := getInputFromUser(reader, "Value 1: ", true)
	float2, _ := getInputFromUser(reader, "Value 2: ", true)
	_, operation := getInputFromUser(reader, "Operation: ", false)

	var result float64

	switch operation {
		case "+": result = float1 + float2
		case "-": result = float1 - float2
		case "*": result = float1 * float2
		case "/": result = float1 / float2
		default:
			panic("Operation must be one of: +, -, *, /")
	
	}

	result = math.Round(result*100) / 100
	fmt.Printf("Result of %v %v %v is %v\n\n", float1, operation, float2, result)

}

func getInputFromUser(reader *bufio.Reader, msg string, isNumber bool) (float64, string) {
	fmt.Print(msg)
	input, _ := reader.ReadString('\n')
	if isNumber {

		number, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err != nil {
			fmt.Println(err)
			panic("Value must be a number")
		}
		return number, ""
	
	} else {
		return 0, strings.TrimSpace(input)
	}

}