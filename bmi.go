package main

import "github.com/Farshad-Davari/bmi/info"

func main() {

	// output the informations
	info.PrintWelcome()

	// user inputs
	weight, height := getUserMetrics()

	// calculate the BMI
	bmi := calculateBMI(weight, height)

	// output the BMI
	printBMI(bmi)
}

func calculateBMI(weight float64, height float64) float64 {

	return weight / (height * height)
}
