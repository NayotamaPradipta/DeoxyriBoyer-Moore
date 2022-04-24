package algorithm
// Border Function / fail
func b(disease string) []int{
	var diseaseLength = len(disease)
	b := make([]int, diseaseLength)
	var prevLength int = 0
	var i int = 1 
	b[0] = 0
	for ok := true;ok; ok = (i < diseaseLength){
		if (disease[i] == disease[prevLength]){
			prevLength++
			b[i] = prevLength
			i++
		} else {
			if (prevLength != 0){
				prevLength = b[prevLength - 1]
			} else {
				b[i] = prevLength
				i++
			}
		}
	}
	return b
}



func StartKMP(dna string, disease string) bool{
	var dnaLength int
	var diseaseLength int
	dnaLength = len(dna)
	diseaseLength = len(disease)
	var b = b(disease)
	var i int = 0 // Index disease 
	var j int = 0 // Index dna
	// Looping hingga disease ditemukan atau sudah search semua dna
	for ok := true; ok; ok = (j < dnaLength) {
		if disease[i] == dna[j]{
			i++
			j++
		}
		if i == diseaseLength {
			return true
		} else if j < dnaLength && disease[i] != dna[j] {
			if (i != 0){
				i = b[i-1]
			} else {
				j++
			}
		}
	} 
	return false
}
