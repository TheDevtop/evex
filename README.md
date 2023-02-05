# Evaluate expression

Evex is an evaluation engine which interpets syntax,
and behaves similair to lambda calculus.
Every evaluation is solely based on addition.

### Usage:

1. Clone this repository with: `$ git clone`
2. Build the program with: `$ go build`
3. Write some evex code into a file
4. Execute the code with: `$ evex [file]`

### Syntax:

```
fa := 3.0
fx := fa fa fa
```
In this example _fa_ is evaluated as 3.0.
By applying this to _fx_ we get _fx = (3.0+3.0+3.0) = 9_.
