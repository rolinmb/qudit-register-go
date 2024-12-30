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

func hadamardGate(dim int) [][]complex128 { // same as fourier transform gate; superimposes or undoes superposition
    gate := make([][]complex128, dim)
    omega := cmplx.Exp(2 * math.Pi * 1i / complex(float64(dim), 0)) // primitive root of unity
    for i := 0; i < dim; i++ {
        gate[i] = make([]complex128, dim)
        for j := 0; j < dim; j++ {
            gate[i][j] = cmplx.Pow(omega, complex(float64(i*j), 0)) / complex(math.Sqrt(float64(dim)), 0)
        }
    }
    return gate
}

func cNotGate(dim int) [][]complex128 {
    gate := make([][]complex128, dim*dim) // Create a 2D matrix of size d^2 for the control-target qudit pair
    // Initialize the matrix elements to 0
    for i := 0; i < dim*dim; i++ {
        gate[i] = make([]complex128, dim*dim)
    }
    // Identity operation when control qudit is in state |0>
    for i := 0; i < dim; i++ {
        gate[i][i] = 1
    }
    // Cyclic permutation when control qudit is in state |1>
    for i := 0; i < dim-1; i++ {
        gate[dim+i][dim+i+1] = 1
    }
    gate[2*dim-1][dim] = 1 // Circular permutation for last state
    return gate
} // TODO: Test this function