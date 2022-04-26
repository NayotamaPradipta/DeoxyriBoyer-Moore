package algorithm
// File untuk pengecekan regex
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

func IsValidSearchDiseaseOnly(search string) bool {
	// Regex Parsing 
	match, err := regexp.MatchString(`([A-Z][a-z]+|[A-Z]+)`, search)
	if err == nil {
		if match {
			return true
		} else {
			return false
		}
	} else {
		panic(err.Error())
	}
} 

func IsValidSearchDateOnly(search string) bool {
	match, err := regexp.MatchString(`^(3[01]|[12][0-9]|0?[1-9])\s(Jan(uary)?|Feb(ruary)?|Mar(ch)?|Apr(il)?|May|Jun(e)?|Jul(y)?|Aug(ust)?|Sep(tember)?|Oct(ober)?|Nov(ember)?|Dec(ember)?)\s\d{4}$`, search)
	if err == nil {
		if match {
			return true
		} else {
			return false
		}
	} else {
		panic(err.Error())
	}
}

func IsValidSearchDateAndDisease(search string) bool {
	match, err := regexp.MatchString(`((3[01]|[12][0-9]|0?[1-9])\s(Jan(uary)?|Feb(ruary)?|Mar(ch)?|Apr(il)?|May|Jun(e)?|Jul(y)?|Aug(ust)?|Sep(tember)?|Oct(ober)?|Nov(ember)?|Dec(ember)?)\s\d{4}\s([A-Z][a-z]+|[A-Z]+))|(([A-Z][a-z]+|[A-Z]+)\s(3[01]|[12][0-9]|0?[1-9])\s(Jan(uary)?|Feb(ruary)?|Mar(ch)?|Apr(il)?|May|Jun(e)?|Jul(y)?|Aug(ust)?|Sep(tember)?|Oct(ober)?|Nov(ember)?|Dec(ember)?)\s\d{4})`, search)
	if err == nil {
		if match {
			return true
		} else {
			return false
		}
	} else {
		panic(err.Error())
	}
}
