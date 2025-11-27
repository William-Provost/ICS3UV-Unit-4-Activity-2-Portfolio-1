// Author: William
// Version: 1.0.0
// Date: 2025-11-27
// Fileoverview: This program calculates base^exponent using a while-style loop.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// variable declarations
	var baseString string
	var exponentString string
	var base int
	var exponent int
	var count int
	var result int = 1

	reader := bufio.NewReader(os.Stdin)

	// get base from user
	fmt.Print("Enter the base: ")
	baseString, _ = reader.ReadString('\n')
	baseString = strings.TrimSpace(baseString)
	base, _ = strconv.Atoi(baseString)

	// get exponent from user
	fmt.Print("Enter the exponent: ")
	exponentString, _ = reader.ReadString('\n')
	exponentString = strings.TrimSpace(exponentString)
	exponent, _ = strconv.Atoi(exponentString)

	// while loop to multiply base exponent times
	count = 0
	for count < exponent {
		result = result * base
		count++
	}

	// print result
	fmt.Printf("%d raised to the power of %d is: %d\n", base, exponent, result)

	fmt.Println("\nDone.")
}
