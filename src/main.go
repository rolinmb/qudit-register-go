package main

import (
    "fmt"
)

func main() {
    qudit0, err := singletonQudit([]complex128{1, 0, 0})
    if err != nil {
        fmt.Println("src/main.go : main() :: ERROR ::: Failed to initialize qudit0")
        return
    }
    qudit1, err := singletonQudit([]complex128{0, 0, -1i})
    if err != nil {
        fmt.Println("src/main.go : main() :: ERROR ::: Failed to initialize qudit1")
        return
    }
    register := QuantumRegister{
        Qudits: []*Qudit{qudit0, qudit1},
    }
    jointState := register.TensorProduct()
    if jointState == nil {
        fmt.Println("src/main.go : main() :: ERROR ::: Tensor product failed. Quantum register is empty.")
        return
    }
    fmt.Println("\nsrc/main.go : main() :: Tensor Product of Quantum Register before Hadamard Gate:", jointState.Amplitudes)
    qrCounts := make(map[string]int)
    for i := 0; i < ITERS; i++ {
        measurement := register.measure(getObservation())
        stateKey := fmt.Sprint(measurement)
        qrCounts[stateKey]++
    }
    fmt.Printf("\nsrc/main.go : main() :: Quantum Register before Hadamard Gate results after %d iterations :::\n", ITERS)
    for state, count := range qrCounts {
        fmt.Printf("\t -> Quantum Register State %s: %d occurrences\n", state, count)
    }
    fmt.Println("\nsrc/main.go : main() :: qudit0 before Hadamard Gate:", register.Qudits[0].Amplitudes)
    hgate := hadamardGate(3)
    if err := register.applyGateToQudit(0, hgate); err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("src/main.go : main() :: qudit0 after Hadamard Gate:", register.Qudits[0].Amplitudes)
    jointState = register.TensorProduct()
    if jointState == nil {
        fmt.Println("src/main.go : main() :: ERROR ::: Tensor product failed. Quantum register is empty.")
        return
    }
    fmt.Println("\nsrc/main.go : main() :: Tensor Product of Quantum Register after Hadamard Gate:", jointState.Amplitudes)
    qrCounts = make(map[string]int)
    for i := 0; i < ITERS; i++ {
        measurement := register.measure(getObservation())
        stateKey := fmt.Sprint(measurement)
        qrCounts[stateKey]++
    }
    fmt.Printf("\nsrc/main.go : main() :: Quantum Register after Hadamard Gate results after %d iterations :::\n", ITERS)
    for state, count := range qrCounts {
        fmt.Printf("\t -> Quantum Register State %s: %d occurrences\n", state, count)
    }
}
