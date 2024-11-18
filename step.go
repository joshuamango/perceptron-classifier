package main

import (
	"fmt"
)

func main() {
	//x := 0.582182
	//k := 20.0
	learning_rate := 0.1
	inputs := []float64{1.0, 0.0, 1.0, 1.0, 0.0}
	weights := []float64{0.5, 0.5, 0.5, 0.5, 0.5}
	outputs := make([]float64, len(inputs))
	for range 100 {
		for i := range len(inputs) {
			weights[i] = weights[i] + learning_rate*(inputs[i]-outputs[i])*inputs[i]
			outputs[i] = inputs[i] * weights[i]
		}
	}
	fmt.Printf("weights: %g\n", weights)
	fmt.Println(outputs)
	//fmt.Printf("%.7g\n", math.Round(1.0/2.0 + 1.0/2.0 * math.Tanh(k * x)))
}
