package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Farshad-Davari/bmi/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (float64, float64) {

	fmt.Print(info.WeigthPrompt)
	weight, _ := reader.ReadString('\n')

	fmt.Print(info.HeightPrompt)
	height, _ := reader.ReadString('\n')

	userWeight, _ := strconv.ParseFloat(strings.TrimSpace(weight), 64)
	userHeight, _ := strconv.ParseFloat(strings.TrimSpace(height), 64)

	return userWeight, userHeight
}
