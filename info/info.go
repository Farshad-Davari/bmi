package info

import "fmt"

const mainTitle = "BMI Calculator"
const separator = "------------------"
const WeigthPrompt = "Please enter your weight(kg): "
const HeightPrompt = "Please enter your height(m): "

func PrintWelcome() {

	fmt.Println(mainTitle)
	fmt.Println(separator)
}
