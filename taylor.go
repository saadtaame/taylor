package main

import (
	"math"
)

type Taylor struct {
	coeff []float64
}

func NewTaylor(N int) Taylor {
	return Taylor{make([]float64, N, N)}
}

func NewVariable(x float64, N int) Taylor {
	f := NewTaylor(N)
	f.coeff[0] = x
	f.coeff[1] = 1
	return f
}

func NewConstant(c float64, N int) Taylor {
	f := NewTaylor(N)
	f.coeff[0] = c
	return f
}

func Print(A Taylor) {
	print("[");
	for _, v := range A.coeff {
		print(" ", v)
	}
	print(" ]")
}

func Add(A Taylor, B Taylor) Taylor {
	C := NewTaylor(len(A.coeff))
	for i := 0; i < len(A.coeff); i++ {
		C.coeff[i] = A.coeff[i] + B.coeff[i];
	}
	return C
}

func Subtract(A Taylor, B Taylor) Taylor {
	C := NewTaylor(len(A.coeff))
	for i:= 0; i < len(A.coeff); i++ {
		C.coeff[i] = A.coeff[i] - B.coeff[i]
	}
	return C
}

func Multiply(A Taylor, B Taylor) Taylor {
	C := NewTaylor(len(A.coeff))
	for i := 0; i < len(A.coeff); i++ {
		for j := 0; j <= i; j++ {
			C.coeff[i] += A.coeff[j] * B.coeff[i - j]
		}
	}
	return C
}

func Divide(A Taylor, B Taylor) Taylor {
	C := NewTaylor(len(A.coeff))
	C.coeff[0] = A.coeff[0] / B.coeff[0]
	for i := 1; i < len(A.coeff); i++ {
		var sum = A.coeff[i]
		for j := 0; j < i; j++ {
			sum -= C.coeff[j] * B.coeff[i - j]
		}
		C.coeff[i] = sum / B.coeff[0]
	}
	return C
}

func Exp(A Taylor) Taylor {
	C := NewTaylor(len(A.coeff))
	C.coeff[0] = math.Exp(A.coeff[0])
	for i := 1; i < len(A.coeff); i++ {
		var sum float64 = 0
		for j := 1; j <= i; j++ {
			sum += float64(j) * A.coeff[j] * C.coeff[i - j]
		}
		C.coeff[i] = sum / float64(i)
	}
	return C
}

func Log(A Taylor) Taylor {
	C := NewTaylor(len(A.coeff))
	C.coeff[0] = math.Log(A.coeff[0])
	for i := 1; i < len(A.coeff); i++ {
		var sum float64 = 0
		for j := 1; j < i; j++ {
			sum += float64(j) * C.coeff[j] * A.coeff[i - j]
		}
		C.coeff[i] = (A.coeff[i] - sum / float64(i)) / A.coeff[0]
	}
	return C
}

/* Example */
func main() {
	A := NewTaylor(3)

	A.coeff[0] = 1.5
	A.coeff[1] = 1

	C := Log(A);
	Print(C)
}
