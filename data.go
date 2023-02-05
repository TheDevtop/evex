package main

type evexForm struct {
	value    string
	function func([]float64) float64
	set      []float64
}

// Exit codes
const (
	exitDef = 0 // Exit without errors
	exitErr = 1 // Exit with error
)

// Tokens
const (
	tokenNewline  = "\n"
	tokenEmpty    = ""
	tokenBigspace = "    "
	tokenHash     = "#"
	tokenFold     = ":="
	tokenCount    = "#="
	tokenHigh     = "|="
	tokenLow      = "&="
)

// Sizes and indeces
const (
	sizeMinTokens = 3  // Minimal token size
	sizeFloat64   = 64 // 64-bit floating point
	indcValue     = 0
	indcOperation = 1
	indcSet       = 2
)

// Error strings
const (
	errMinTokens    = "token size error at index %d"
	errEvalFunction = "can't evaluate token (%s)"
)
