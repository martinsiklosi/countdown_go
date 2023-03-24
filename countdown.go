// Martin Siklosi

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Expression type
type Exp struct {
	str string
	val int
	con int
}

func NewExp(str string, val int, con int) *Exp {
	return &Exp{str, val, con}
}

// Add two expressions and return result.
func Add(e1, e2 Exp) (Exp, error) {
	str := fmt.Sprintf("(%s)+(%s)", e1.str, e2.str)
	val := e1.val + e2.val
	con := e1.con + e2.con
	return *NewExp(str, val, con), nil
}

// Multiply two expressions and return result.
func Mult(e1, e2 Exp) (Exp, error) {
	if e1.val == 1 || e2.val == 1 {
		return *new(Exp), errors.New("Unnecessary multiplication.")
	}
	str := fmt.Sprintf("%s*%s", e1.str, e2.str)
	val := e1.val * e2.val
	con := e1.con + e2.con
	return *NewExp(str, val, con), nil
}

// Subtract two expressions and return result.
func Sub(e1, e2 Exp) (Exp, error) {
	if e1.val <= e2.val {
		return *new(Exp), errors.New("Unnecessary subtraction.")
	}
	str := fmt.Sprintf("(%s)-(%s)", e1.str, e2.str)
	val := e1.val - e2.val
	con := e1.con + e2.con
	return *NewExp(str, val, con), nil
}

// Divide two expressions and return result.
func Div(e1, e2 Exp) (Exp, error) {
	if e2.val == 1 || e1.val%e2.val != 0 {
		return *new(Exp), errors.New("Unnecessary or invalid division.")
	}
	str := fmt.Sprintf("%s/(%s)", e1.str, e2.str)
	val := e1.val / e2.val
	con := e1.con + e2.con
	return *NewExp(str, val, con), nil
}

// Return the useful combinations of two expressions.
func UsefulCombs(e1, e2 Exp) []Exp {
	methods := []func(Exp, Exp) (Exp, error){Add, Mult, Sub, Div}
	var output []Exp
	for _, method := range methods {
		comb, err := method(e1, e2)
		if err == nil {
			output = append(output, comb)
		}
	}
	return output
}

// Return permutations of expressions in v2 and v2.
func Perms(v1, v2 []Exp, id_set map[int]bool, n_vars int) []Exp {
	return *new([]Exp)
}

// Generate id of expression
func CreateID(e Exp, n_vars int) int {
	return (e.val << n_vars) + e.con
}

// Convert string of integers serperated by spaces to list of integers.
func STIs(s string) ([]int, error) {
	var nums []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err != nil {
			return *new([]int), err
		}
		nums = append(nums, i)
	}
	return nums, nil
}

func main() {
	// Take input
	fmt.Print("Enter numbers: ")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	nums, err := STIs(string(char))
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter target: ")
	char2, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	target, err := strconv.Atoi(string(char2))
	if err != nil {
		panic(err)
	}

	fmt.Println(nums)
	fmt.Println(target)
	// Generate base expressions
	// Find all useful combinations
	// Print the best expression
}
