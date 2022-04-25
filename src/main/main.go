package main

import (
	"fmt"

	"github.com/NayotamaPradipta/DeoxyriBoyer-Moore/src/algorithm"
)

func main() {
	// Asumsi sequence DNA pengguna > sequence penyakit
	dnaToTest := algorithm.GetDNASequenceFromFile("dnaInput.txt")
	disease := "GATC"
	fmt.Println(dnaToTest)
	if algorithm.IsValidString(dnaToTest) {
		if !algorithm.StartBoyerMoore(dnaToTest, disease) {
			fmt.Println("Disease not detected with Boyer-Moore!")
		} else {
			fmt.Println("Disease detected with Boyer-Moore!")
		}
		if !algorithm.StartKMP(dnaToTest, disease) {
			fmt.Println("Disease not detected with Boyer-Moore!")
		} else {
			fmt.Println("Disease detected with KMP!")
		}
	} else {
		fmt.Println("Invalid DNA!")
	}
}
