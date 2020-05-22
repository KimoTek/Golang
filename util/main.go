package util

import "math"

func Square(x float32) float32 {
    return x * x
}

func Circle(r float32) float32 {
    return math.Pi * r * r
}

func cube(x float32) float32 {
    return x * x * x
}
