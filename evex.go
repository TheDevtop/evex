package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type form struct {
	val string
	set []float64
}

const (
	tokenNewline = "\n"
	tokenEval    = ":="
	lineMinSize  = 3  // Minimal line size
	exitDef      = 0  // Exit without errors
	exitErr      = 1  // Exit with error
	indexVal     = 0  // Index position of value
	indexSet     = 2  // Index position of set
	bitsize      = 64 // 64-bit floating point
	f64Default   = 00.00
)

var evexMap map[string]float64

// Parse file contents
func parse(file string) error {
	file = strings.TrimSuffix(file, tokenNewline)
	lines := strings.Split(file, tokenNewline)

	// For each line, split the line into tokens,
	// convert the tail-end tokens to float64,
	// or error out.
	for index, line := range lines {
		tokens := strings.Fields(line)
		if len(tokens) < lineMinSize {
			return fmt.Errorf("main.parse: error at index %d", index)
		}
		if nset, err := substitute(tokens[indexSet:]); err != nil {
			return err
		} else {
			compute(form{val: tokens[indexVal], set: nset})
		}
	}
	return nil
}

// Substitute strings for float64
func substitute(set []string) ([]float64, error) {
	var nset = make([]float64, len(set))

	// For each token, attempt to convert to float64,
	// or substitute them with known values,
	// or error out.
	for index, item := range set {
		if f, err := strconv.ParseFloat(item, bitsize); err != nil {
			if fm, ok := evexMap[item]; !ok {
				return nil, err
			} else {
				nset[index] = fm
				continue
			}
		} else {
			nset[index] = f
			continue
		}
	}
	return nset, nil
}

// Compute and assign expression
func compute(f form) {
	var v float64 = f64Default
	for _, s := range f.set {
		v += s
	}
	evexMap[f.val] = v
}

// Output evexMap content
func output() error {
	if buf, err := json.Marshal(evexMap); err != nil {
		return err
	} else {
		fmt.Printf("%s\n", string(buf))
	}
	return nil
}

// Program entrypoint
func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: evex [file]")
		os.Exit(exitErr)
	}
	if buf, err := os.ReadFile(os.Args[1]); err != nil {
		fmt.Println(err)
		os.Exit(exitErr)
	} else {
		evexMap = make(map[string]float64)
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
