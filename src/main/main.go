package main

import (
	"fmt"
	"github.com/NayotamaPradipta/DeoxyriBoyer-Moore/src/algorithm"
	"github.com/NayotamaPradipta/DeoxyriBoyer-Moore/src/todb"
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
		fmt.Println("Invalid DNA String!")
	}
	klinefelter := algorithm.GetDNASequenceFromFile("Klinefelter.txt")
	if algorithm.IsValidString(klinefelter){
		todb.InsertNewDisease("Klinefelter", klinefelter, "")
	} else {
		fmt.Println("Invalid Disease DNA!")
	}
	fmt.Println(todb.SELECTDNA("x", ""))
	// Testing searching 
	if algorithm.IsValidSearchDiseaseOnly("testDisease") {
		todb.SELECTDNA("testDisease", "")
	} 
	if algorithm.IsValidSearchDateAndDisease("HIV 20 September 1999") {
		fmt.Println("Valid!")
	}
	if algorithm.IsValidSearchDateOnly("11 April 2020"){
		fmt.Println("Valid!")
	}
	if algorithm.IsValidSearchDateAndDisease("19 October 2020 Klinefelter"){
		fmt.Println("Valid!")
	}
	if !algorithm.IsValidSearchDateAndDisease("32 December 2020 HIV"){
		fmt.Println("Invalid!")
	}
	todb.SELECTRIWAYAT("32 December 2020 HIV", "")
	todb.InsertNewPrediction("11 April 2020", "Kaori Miyazono", "Klinefelter", "False", "")
	riwayat := todb.SELECTRIWAYAT("Klinefelter 11 April 2020", "")
	fmt.Println(riwayat)
}
