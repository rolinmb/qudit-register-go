package main

import (
    "fmt"
)

func tensorProduct(a0,a1 []complex128) []complex128 {
    result := make([]complex128, len(a0)*len(a1))
    for i, amp0 := range a0 {
        for j, amp1 := range a1 {
            result[i*len(a1) + j] = amp0 * amp1
        }
    }
    return result
}

type QuantumRegister struct {
    Qudits []*Qudit
}

func (qr *QuantumRegister) TensorProduct() *Qudit {
    if len(qr.Qudits) == 0 {
        return nil
    }
    jointAmplitudes := qr.Qudits[0].Amplitudes
    for i := 1; i < len(qr.Qudits); i++ {
        jointAmplitudes = tensorProduct(jointAmplitudes, qr.Qudits[i].Amplitudes)
    }
    return &Qudit{Dimension: len(jointAmplitudes), Amplitudes: jointAmplitudes}
}

func (qr *QuantumRegister) measure(observation float64) []int {
    jointState := qr.TensorProduct()
    if jointState == nil {
        fmt.Println("src/register.go : measure() :: ERROR ::: No Qudits in Quantum Register to take Tensor Product from.")
        return nil
    }
    result := jointState.measure(observation)
    numQudits := len(qr.Qudits) // Decode the joint measurement result into individual qudit states
    results := make([]int, numQudits)
    base := len(qr.Qudits[0].Amplitudes)
    for i := numQudits - 1; i >= 0; i-- {
        results[i] = result % base
        result /= base
    }
    return results
}

func (qr *QuantumRegister) applyGateToQudit(index int, gate [][]complex128) error {
    if index < 0 || index >= len(qr.Qudits) {
        return fmt.Errorf("src/register.go : applyGateToQudit() :: ERROR ::: Index out of range.")
    }
    return qr.Qudits[index].applyGate(gate)
}

func (qr *QuantumRegister) applyGateToQudits(gate [][]complex128, controlIndex, targetIndex int) error {
    if controlIndex < 0 || controlIndex >= len(register) || targetIndex < 0 || targetIndex >= len(register) {
        return fmt.Errorf("src/register.go : applyGateToQudits() :: ERROR ::: Invalid qudit indices.")
    }
    if controlIndex == targetIndex {
        return fmt.Errorf("src/register.go : applyGateToQudits() :: ERROR ::: Control and target indices cannot be the same.")
    }
    // Get dimensions of the two qudits
    controlDim := qr.Qudits[controlIndex].Dimension
    targetDim := qr.Qudits[targetIndex].Dimension
    // Check gate dimensions
    if len(gate) != controlDim*targetDim || len(gate[0]) != controlDim*targetDim {
        return fmt.Errorf("src/qudit.go : applyGateToQudits() :: ERROR ::: Gate dimensions do not match control-target qudit dimensions.")
    }
    // Calculate the new state for the combined register
    totalDim := controlDim * targetDim
    newAmplitudes := make([]complex128, totalDim)
    combinedAmplitudes := make([]complex128, totalDim)
    // Combine the amplitudes of the control and target qudits
    for c := 0; c < controlDim; c++ {
        for t := 0; t < targetDim; t++ {
            index := c*targetDim + t
            combinedAmplitudes[index] = qr.Qudits[controlIndex].Amplitudes[c] * qr.Qudits[targetIndex].Amplitudes[t]
        }
    }
    // Apply the gate
    for i := 0; i < totalDim; i++ {
        for j := 0; j < totalDim; j++ {
            newAmplitudes[i] += gate[i][j] * combinedAmplitudes[j]
        }
    }
    // Update the states of the control and target qudits
    for c := 0; c < controlDim; c++ {
        for t := 0; t < targetDim; t++ {
            index := c*targetDim + t
            qr.Qudits[controlIndex].Amplitudes[c] = newAmplitudes[index]
            qr.Qudits[targetIndex].Amplitudes[t] = newAmplitudes[index]
        }
    }
    return nil
}
