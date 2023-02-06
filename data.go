package main

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
	tokenReduce   = ":="
	tokenCount    = "#="
	tokenHigh     = "|="
	tokenLow      = "&="
)

// Sizes and indeces
const (
	sizeMinTokens = 3  // Minimal token size
	sizeFloat     = 64 // 64-bit floating point
	indcVal       = 0  // Index of value token
	indcFunc      = 1  // Index of function token
	indcSet       = 2  // Index of set token
)

// Error strings
const (
	errMinTokens    = "length error, index %d"
	errEvalFunction = "token error, can't evaluate (%s)"
)
