package algorithm

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
	"os"
)

func GetDNASequenceFromFile(filename string) string{
	mydir, _ := os.Getwd()
	absPath, _ := filepath.Abs(filepath.Dir(filepath.Dir(mydir)) + "/test/" + filename)
	b, err := ioutil.ReadFile(absPath)
	var str = ""
	if (err == nil){
		str := string(b)
		return str
	} 
	return str
}

func IsValidString(dna string) bool {
	// Regex Parsing
	match, err := regexp.MatchString(`^[ACGT]+$`, dna)
	if err == nil {
		return match
	} else {
		return false
	}
}