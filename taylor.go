package main

import (
	"fmt"
	"math"
)

// Taylor is the representation of a taylor series.
type Taylor struct {
	coeff []float64
}

// String Method added to be used by default when printing.
func (T Taylor) String() string {
	str := "[ "
	for _, v := range T.coeff {
		str += " " + fmt.Sprintf("%.10f", v)
	}
	str += " ]"
	return str
}

// NewTaylor creates a new empty instance of Taylor.
func NewTaylor(N int) Taylor {
	return Taylor{make([]float64, N, N)}
}

// NewVariable creates a new instacnce with one variable.
func NewVariable(x float64, N int) Taylor {
	f := NewTaylor(N)
	f.coeff[0] = x
	f.coeff[1] = 1
	return f
}

// NewConstant creates a new instacnce with one constant.
func NewConstant(c float64, N int) Taylor {
	f := NewTaylor(N)
	f.coeff[0] = c
	return f
}

// Add adds two taylor series.
func Add(A Taylor, B Taylor) Taylor {
	C := NewTaylor(len(A.coeff))
	for i := 0; i < len(A.coeff); i++ {
		C.coeff[i] = A.coeff[i] + B.coeff[i]
	}
	return C
}

// Subtract subs two taylor series.
func Subtract(A Taylor, B Taylor) Taylor {
	C := NewTaylor(len(A.coeff))
	for i := 0; i < len(A.coeff); i++ {
		C.coeff[i] = A.coeff[i] - B.coeff[i]
	}
	return C
}

// Multiply takes two series and multiplies them.
func Multiply(A Taylor, B Taylor) Taylor {
	C := NewTaylor(len(A.coeff))
	for i := 0; i < len(A.coeff); i++ {
		for j := 0; j <= i; j++ {
			C.coeff[i] += A.coeff[j] * B.coeff[i-j]
		}
	}
	return C
}

// Divide divides two taylor series.
func Divide(A Taylor, B Taylor) Taylor {
	C := NewTaylor(len(A.coeff))
	C.coeff[0] = A.coeff[0] / B.coeff[0]
	for i := 1; i < len(A.coeff); i++ {
		var sum = A.coeff[i]
		for j := 0; j < i; j++ {
			sum -= C.coeff[j] * B.coeff[i-j]
		}
		C.coeff[i] = sum / B.coeff[0]
	}
	return C
}

// Exp powers a taylor series.
func Exp(A Taylor) Taylor {
	C := NewTaylor(len(A.coeff))
	C.coeff[0] = math.Exp(A.coeff[0])
	for i := 1; i < len(A.coeff); i++ {
		var sum float64 = 0
		for j := 1; j <= i; j++ {
			sum += float64(j) * A.coeff[j] * C.coeff[i-j]
		}
		C.coeff[i] = sum / float64(i)
	}
	return C
}

// Log takes a taylor series and returns its log.
func Log(A Taylor) Taylor {
	C := NewTaylor(len(A.coeff))
	C.coeff[0] = math.Log(A.coeff[0])
	for i := 1; i < len(A.coeff); i++ {
		var sum float64
		for j := 1; j < i; j++ {
			sum += float64(j) * C.coeff[j] * A.coeff[i-j]
		}
		C.coeff[i] = (A.coeff[i] - sum/float64(i)) / A.coeff[0]
	}
	return C
}

/* Example */
func main() {
	A := NewTaylor(3)

	A.coeff[0] = 1.5
	A.coeff[1] = 1

	C := Log(A)
	fmt.Println(C)
}
