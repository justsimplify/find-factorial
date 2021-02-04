package main

import (
	"exercie/modules/calculation"
	"fmt"
	"os"
)

func main() {
	fmt.Println(calculation.Calculate(os.Args[1]))
}