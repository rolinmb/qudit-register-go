package main

import (
    "fmt"
    "math"
    "math/cmplx"
)

type Qudit struct {
    Dimension int
    Amplitudes []complex128
}

func singletonQudit(cargs []complex128) (*Qudit, error) {
    norm := 0.0
    for _, c := range cargs {
        norm += math.Pow(cmplx.Abs(c), 2)
    }
    if math.Abs(norm - 1.0) < EPSILON {
        qd := &Qudit {Dimension: len(cargs), Amplitudes: cargs }
        return qd, nil
    }
    return nil, fmt.Errorf("src/qudit.go : singletonQudit() :: ERROR ::: Invalid complex number arguments for a valid qudit.")
}

func (qd *Qudit) measure(observation float64) int {
    cuumProb := 0.0
    for i, c := range qd.Amplitudes {
	    cuumProb += math.Pow(cmplx.Abs(c), 2)
	    if observation < cuumProb {
            return i
	    }
    }
    return len(qd.Amplitudes) - 1 
}

func (qd *Qudit) applyGate(gate [][]complex128) error {
    if qd == nil {
        return fmt.Errorf("src/qudit.go : applyGate() :: ERROR ::: Qudit is nil, cannot apply gate.")
    }
    if len(gate) != qd.Dimension {
        return fmt.Errorf("src/qudit.go : applyGate() :: ERROR ::: Gate dimension (%d) does not match qudit dimension (%d).", len(gate), qd.Dimension)
    }
    for i := range gate {
        if len(gate[i]) != qd.Dimension {
            return fmt.Errorf("src/register.go : applyGate() :: ERROR ::: Gate is not square.")
        }
    }
    newAmplitudes := make([]complex128, qd.Dimension)
    for i := 0; i < qd.Dimension; i++ {
        for j := 0; j < qd.Dimension; j++ {
            newAmplitudes[i] += gate[i][j] * qd.Amplitudes[j]
        }
    }
    qd.Amplitudes = newAmplitudes
    return nil
}