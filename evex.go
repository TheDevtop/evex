package main

/*
	Evalute Expression
	Version: 1.1
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

// Output evexMap content
func output() error {
	if buf, err := json.MarshalIndent(evexMap, tokenEmpty, tokenBigspace); err != nil {
		return err
	} else {
		fmt.Printf("%s\n", string(buf))
	}
	return nil
}

// Map and evaluate strings to a set
func evalSet(setStr []string) ([]float64, error) {
	var set []float64
	for _, str := range setStr {
		if num, err := strconv.ParseFloat(str, sizeFloat); err != nil {
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

// Reduction
func fnReduce(set []float64) float64 {
	var total float64 = 00.00
	for _, num := range set {
		total += num
	}
	return total
}

// Evaluate length
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

// Map string to function, or error
func evalFunc(fnStr string) (func([]float64) float64, error) {
	switch fnStr {
	case tokenReduce:
		return fnReduce, nil
	case tokenCount:
		return fnCount, nil
	case tokenHigh:
		return fnHigh, nil
	case tokenLow:
		return fnLow, nil
	default:
		return nil, fmt.Errorf(errEvalFunction, fnStr)
	}
}

// Parse, evaluate, and apply expression
func parse(file string) error {
	file = strings.TrimSuffix(file, tokenNewline)
	for index, line := range strings.Split(file, tokenNewline) {
		if strings.HasPrefix(line, tokenHash) || line == tokenEmpty {
			continue
		}
		tokens := strings.Fields(line)
		if len(tokens) < sizeMinTokens {
			return fmt.Errorf(errMinTokens, index)
		}
		if fn, err := evalFunc(tokens[indcFunc]); err != nil {
			return err
		} else if set, err := evalSet(tokens[indcSet:]); err != nil {
			return err
		} else {
			evexMap[tokens[indcVal]] = fn(set)
		}
	}
	return nil
}

// Detect arguments and read from input
func input() ([]byte, error) {
	if len(os.Args) < 2 {
		if buf, err := io.ReadAll(os.Stdin); err != nil {
			return nil, err
		} else {
			return buf, nil
		}
	} else {
		if buf, err := os.ReadFile(os.Args[1]); err != nil {
			return nil, err
		} else {
			return buf, err
		}
	}
}

// Program entrypoint
func main() {
	evexMap = make(map[string]float64)
	if buf, err := input(); err != nil {
		fmt.Println(err)
		os.Exit(exitErr)
	} else {
		if err := parse(string(buf)); err != nil {
			fmt.Println(err)
			os.Exit(exitErr)
		}
	}
	if err := output(); err != nil {
		fmt.Println(err)
		os.Exit(exitErr)
	}
	os.Exit(exitDef)
}
