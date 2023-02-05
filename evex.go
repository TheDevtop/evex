package main

/*
	Evalute Expression
	Version: 0.9
*/

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var evexMap map[string]float64

func fnFold(set []float64) float64 {
	var total float64 = 00.00
	for _, num := range set {
		total += num
	}
	return total
}

func fnCount(set []float64) float64 {
	return float64(len(set))
}

// Evaluate highest
func fnHigh(set []float64) float64 {
	max := set[0]
	for _, num := range set {
		if num > max {
			max = num
		}
	}
	return max
}

// Evaluate lowest
func fnLow(set []float64) float64 {
	min := set[0]
	for _, num := range set {
		if num < min {
			min = num
		}
	}
	return min
}

// Parse and evaluate expression
func parse(file string) error {
	file = strings.TrimSuffix(file, tokenNewline)
	for index, line := range strings.Split(file, tokenNewline) {
		if strings.HasPrefix(line, tokenHash) {
			continue
		}
		tokens := strings.Fields(line)
		if len(tokens) < sizeMinTokens {
			return fmt.Errorf(errMinTokens, index)
		}
		if fn, err := evalFunction(tokens[indcOperation]); err != nil {
			return err
		} else if set, err := evalSet(tokens[indcSet:]); err != nil {
			return err
		} else {
			evalForm(evexForm{
				value:    tokens[indcValue],
				function: fn,
				set:      set,
			})
		}
	}
	return nil
}

// Evaluate and apply form
func evalForm(form evexForm) {
	evexMap[form.value] = form.function(form.set)
}

// Map and evaluate strings to a set
func evalSet(setStr []string) ([]float64, error) {
	var set []float64
	for _, str := range setStr {
		if num, err := strconv.ParseFloat(str, sizeFloat64); err != nil {
			if num, ok := evexMap[str]; !ok {
				return set, err
			} else {
				set = append(set, num)
				continue
			}
		} else {
			set = append(set, num)
			continue
		}
	}
	return set, nil
}

// Map string to function, or error
func evalFunction(opStr string) (func([]float64) float64, error) {
	switch opStr {
	case tokenFold:
		return fnFold, nil
	case tokenCount:
		return fnCount, nil
	case tokenHigh:
		return fnHigh, nil
	case tokenLow:
		return fnLow, nil
	default:
		return nil, fmt.Errorf(errEvalFunction, opStr)
	}
}

// Output evexMap content
func output() error {
	if buf, err := json.MarshalIndent(evexMap, tokenEmpty, tokenBigspace); err != nil {
		return err
	} else {
		fmt.Printf("%s\n", string(buf))
	}
	return nil
}

// Program entrypoint
func main() {
	var (
		buf []byte
		err error
	)
	if len(os.Args) < 2 {
		if buf, err = io.ReadAll(os.Stdin); err != nil {
			fmt.Println(err)
			os.Exit(exitErr)
		}
	} else {
		if buf, err = os.ReadFile(os.Args[1]); err != nil {
			fmt.Println(err)
			os.Exit(exitErr)
		}
	}
	evexMap = make(map[string]float64)
	parse(string(buf))
	if err = output(); err != nil {
		fmt.Println(err)
		os.Exit(exitErr)
	}
	os.Exit(exitDef)
}
