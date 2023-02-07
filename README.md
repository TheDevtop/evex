# Evaluate expression

Evex is an evaluation engine (interpreter),
and behaves similair to functional languages.

### Usage:

1. Clone this repository with: `$ git clone`
2. Build the program with: `$ go build`
3. Write some evex code into a file
4. Execute the code with: `$ evex < file.vx`

### Syntax:

```
fa := 3.0
fx := fa fa fa
```
In this example _fa_ is evaluated as 3.0.
By applying this to _fx_ we get _fx_ = (3.0+3.0+3.0) = 9.

### Operators:

- Reduction **:=**
- Count/Length **#=**
- Selection (highest) **|=**
- Selection (lowest) **&=**
