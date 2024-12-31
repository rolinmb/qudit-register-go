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

func pauliXGate(dim int) [][]complex128 { // Bit flip operation
    gate := make([][]complex128, dim)
    for i := 0; i < dim; i++ {
        gate[i] = make([]complex128, dim)
        gate[i][(i+1)%dim] = 1
    }
    return gate
}

func pauliYGate(dim int) [][]complex128 { // Bit flip + phase flip
    gate := make([][]complex128, dim)
    for i := 0; i < dim; i++ {
        gate[i] = make([]complex128, dim)
        gate[i][(i+1)%dim] = complex(0, 1) // Imaginary component
        gate[(i+1)%dim][i] = complex(0, -1)
    }
    return gate
}

func pauliZGate(dim int) [][]complex128 { // Phase flip operation
    gate := make([][]complex128, dim)
    for i := 0; i < dim; i++ {
        gate[i] = make([]complex128, dim)
        if i%2 == 0 {
            gate[i][i] = 1
        } else {
            gate[i][i] = -1
        }
    }
    return gate
}

func phaseShiftGate(dim int, theta float64) [][]complex128 {
    gate := identityGate(dim)
    for i := 0; i < dim; i++ {
        gate[i][i] = cmplx.Exp(complex(0, theta*float64(i))) // e^(i*theta)
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
// TODO: Test the two-qudit gates below here
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
}

func cZGate(dim int) [][]complex128 {
    gate := make([][]complex128, dim*dim)
    for i := 0; i < dim*dim; i++ {
        gate[i] = make([]complex128, dim*dim)
    }
    for i := 0; i < dim; i++ {
        gate[i][i] = 1
    }
    for i := dim; i < dim*dim; i++ {
        gate[i][i] = -1
    }
    return gate
}

func swapGate(dim int) [][]complex128 {
    gate := identityGate(dim * dim)
    for i := 0; i < dim; i++ {
        for j := 0; j < dim; j++ {
            gate[i*dim+j][j*dim+i] = 1
            gate[j*dim+i][i*dim+j] = 1
            gate[i*dim+j][i*dim+j] = 0
        }
    }
    return gate
}
// TODO: test 3-qudit gates below here
// Toffoli Gate: third qudit is flipped if the first two qudits are in state 1
func toffoliGate(dim int) [][]complex128 {
    gate := identityGate(dim * dim * dim)
    for i := 0; i < dim; i++ {
        for j := 0; j < dim; j++ {
            if i == 1 && j == 1 {
                gate[dim*dim+i][dim*dim+j] = 1
            }
        }
    }
    return gate
}
