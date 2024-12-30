package main

import (
    "time"
    "math/rand"
)

const (
    ITERS = 1000
    EPSILON = 1e-5
)

func getObservation() float64 {
    time.Sleep(1*time.Nanosecond)
    rand.Seed(time.Now().UnixNano())
    return rand.Float64()
}