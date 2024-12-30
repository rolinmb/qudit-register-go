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
    qudit1, err := singletonQudit([]complex128{0, -1})
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
    fmt.Println("\nsrc/main.go : main() :: Tensor Product of Quantum Register:", jointState.Amplitudes)
    qrCounts := make(map[string]int)
    for i := 0; i < ITERS; i++ {
        measurement := register.measure(getObservation())
        stateKey := fmt.Sprint(measurement)
        qrCounts[stateKey]++
    }
    fmt.Printf("\nsrc/main.go : main() :: Quantum Register results after %d iterations :::\n", ITERS)
    for state, count := range qrCounts {
        fmt.Printf("\t -> Quantum Register State %s: %d occurrences\n", state, count)
    }
}
