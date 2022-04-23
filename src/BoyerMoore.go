package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
)

func getDNASequenceFromFile(filename string) string{
	absPath, _ := filepath.Abs("DeoxyriBoyer-Moore/test/" + filename)
	b, err := ioutil.ReadFile(absPath)
	var str = ""
	if (err == nil){
		str := string(b)
		return str
	} 
	return str

}	

func isValidString(dna string) bool {
	// Regex Parsing
	match, err := regexp.MatchString(`^[ACGT]+$`, dna)
	if err == nil {
		return match
	} else {
		return false
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func buildLast(disease string) [128]int {
	// Mengembalikan sebuah array of integer berukuran 128 (Banyaknya karakter ASCII).
	// Jika index ke-i bernilai -1, maka tidak ada karakter tersebut di pattern
	// Jika tidak bernilai -1, maka nilai tersebut menunjukkan index terakhir karakter di pattern
	var lastOccurence [128]int
	for i := 0; i < 128; i++ {
		lastOccurence[i] = -1
	}
	for i := 0; i < len(disease); i++ {
		lastOccurence[disease[i]] = i
	}
	return lastOccurence
}

func BoyerMoore(dna string, disease string) bool {
	// Boyer-Moore Algorithm
	// I.S. Input DNA valid (tidak ada huruf kecil, tidak ada huruf selain AGCT, tidak ada spasi)
	// F.S. Boolean true or false

	// Deklarasi Variabel
	var lastOccurence [128]int
	var dnaLength int
	var diseaseLength int
	var i int

	// Simpan last occurence untuk masing-masing karakter pada disease
	lastOccurence = buildLast(disease)

	dnaLength = len(dna)
	diseaseLength = len(disease)
	i = diseaseLength - 1
	if i > dnaLength-1 {
		return false
	}
	var j int
	j = diseaseLength - 1
	for ok := true; ok; ok = (i <= dnaLength) {
		if disease[j] == dna[i] {
			if j == 0 {
				return true
			} else {
				i--
				j--
			}
		} else {
			var lo int
			lo = lastOccurence[dna[i]]
			i = i + diseaseLength - min(j, lo+1)
			j = diseaseLength - 1
		}
	}
	return false
}

func main() { // Buat testing 
	// Asumsi sequence DNA pengguna > sequence penyakit
	dnaToTest := getDNASequenceFromFile("dnaInput.txt")
	disease := "GATC"
	fmt.Println(dnaToTest)
	if isValidString(dnaToTest) {
		if !BoyerMoore(dnaToTest, disease) {
			fmt.Println("Disease not detected!")
		} else {
			fmt.Println("Disease detected!")
		}
	} else {
		fmt.Println("Invalid DNA!")
	}
}
