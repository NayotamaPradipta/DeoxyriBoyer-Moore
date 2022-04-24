package algorithm

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BuildLast(disease string) [128]int {
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

func StartBoyerMoore(dna string, disease string) bool {
	// Boyer-Moore Algorithm
	// I.S. Input DNA valid (tidak ada huruf kecil, tidak ada huruf selain AGCT, tidak ada spasi)
	// F.S. Boolean true or false

	// Deklarasi Variabel
	var lastOccurence [128]int
	var dnaLength int
	var diseaseLength int
	var i int

	// Simpan last occurence untuk masing-masing karakter pada disease
	lastOccurence = BuildLast(disease)

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
			var lo int = lastOccurence[dna[i]]
			i = i + diseaseLength - Min(j, lo+1)
			j = diseaseLength - 1
		}
	}
	return false
}
