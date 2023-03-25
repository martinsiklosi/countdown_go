// Martin Siklosi

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
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
	var output []Exp
	for _, e1 := range v1 {
		for _, e2 := range v2 {
			if e1.con&e2.con != 0 {
				continue
			}
			for _, comb := range UsefulCombs(e1, e2) {
				exp_id := CreateID(comb, n_vars)
				_, is_in := id_set[exp_id]
				if is_in {
					continue
				}
				output = append(output, comb)
				id_set[exp_id] = true
			}
		}
	}
	return output
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

// Return distance between an expression and the target.
func Dist(e Exp, target int) int {
	diff := e.val - target
	if diff >= 0 {
		return diff
	}
	return -diff
}

// Generate string of the best expression given numbers and target.
func RunNumbers(nums []int, target int) string {
	// Generate base expressions
	n_vars := len(nums)
	exp_sets := make([][]Exp, n_vars)
	for i, num := range nums {
		exp_sets[0] = append(
			exp_sets[0],
			*NewExp(fmt.Sprint(num), num, (1 << i)),
		)
	}
	id_set := make(map[int]bool)
	for _, e := range exp_sets[0] {
		e_id := CreateID(e, n_vars)
		id_set[e_id] = true
	}

	// Find all useful combinations
	for i := 0; i < n_vars; i++ {
		for j := 0; j < i; j++ {
			exp_sets[i] = append(
				exp_sets[i],
				Perms(
					exp_sets[j],
					exp_sets[i-j-1],
					id_set,
					n_vars,
				)...)
		}
	}

	// Print the best expression
	var exps []Exp
	for _, v := range exp_sets {
		exps = append(exps, v...)
	}
	sort.SliceStable(exps, func(i, j int) bool {
		return Dist(exps[i], target) < Dist(exps[j], target)
	})
	best := exps[0]
	return fmt.Sprintf("%v = %v", best.str, best.val)
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
	ans := RunNumbers(nums, target)
	fmt.Print(ans)

}
