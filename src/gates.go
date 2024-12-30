package main

import (
    "math"
    "math/cmplx"
)

func identityGate(dim int) [][]complex128 {
    gate := make([][]complex128, dim)
    for i := 0; i < dim; i++ {
        gate[i] = make([]complex128, dim)
        gate[i][i] = 1
    }
    return gate
}

func pauliXGate(dim int) [][]complex128 {
    gate := make([][]complex128, dim)
    for i := 0; i < dim; i++ {
        gate[i] = make([]complex128, dim)
        gate[i][(i+1)%dim] = 1
    }
    return gate
}

func fourierTransformGate(dim int) [][]complex128 {
    gate := make([][]complex128, dim)
    omega := cmplx.Exp(2 * math.Pi * 1i / complex(float64(dim), 0))
    for i := 0; i < dim; i++ {
        gate[i] = make([]complex128, dim)
        for j := 0; j < dim; j++ {
            gate[i][j] = cmplx.Pow(omega, complex(float64(i*j), 0)) / complex(math.Sqrt(float64(dim)), 0)
        }
    }
    return gate
} 