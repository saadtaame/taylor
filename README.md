# taylor
Routines for manipulating Taylor series. These are useful for computing higher order derivatives for example. The tool is written in Google's Go programming language.

# API
* Creating constants and variables
* Creating series
* Arithmetic operators / functions on taylor series

# Supported functions
* Add, Subtract, Multiply, Divide
* Exp, Log

# Example

In the following example we compute the first three terms in the Taylor expansion (centered at 0) of `log(x + 1.5)`.

<pre>
func main() {
    A := NewTaylor(3)
    A.coeff[0] = 1.5
    A.coeff[1] = 1

    C := Log(A);

    Print(C)
}
</pre>

The program outputs:

<pre>
[ +4.054651e-001 +6.666667e-001 -2.222222e-001 ]
</pre>

This is very useful if you want to check your calculus homework. This tool will be extended by adding support for new functions. A little exercise to get yourself familiar with the tool is to write a small program to compute the `n'th` order derivative of some function (expressible using the supported functions) evaluated at 0.
