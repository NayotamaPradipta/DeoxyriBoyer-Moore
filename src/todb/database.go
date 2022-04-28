/*
package todb

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/NayotamaPradipta/DeoxyriBoyer-Moore/src/algorithm"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type jenispenyakit struct {
	id_penyakit               int       `json:"id_penyakit"`
	nama_penyakit             string    `json:"nama_penyakit"`
	rantai_dna                string    `json:"rantai_dna"`
}

type hasilprediksi struct {
	tanggal_prediksi          string    `json:"tanggal_prediksi"`
	nama_pasien               string    `json:"nama_pasien"`
	penyakit_terprediksi      string    `json:"penyakit_terprediksi"`
	status_terprediksi        string    `json:"status_terprediksi"`
}

func InsertNewDisease(diseaseName string, dna string, password string) { 
	// I.S. dna sudah disaring menggunakan regex
	// Capture connection properties.
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/tesdna")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	// Hitung jumlah row sekarang untuk inisialisasi id_penyakit yang ingin ditambahkan
	row, err := db.Query("SELECT COUNT(*) FROM tesdna.jenispenyakit")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var count int
	for row.Next(){
		if err := row.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}
	// Insert to database
	query := fmt.Sprintf(`INSERT INTO tesdna.jenispenyakit (id_penyakit, nama_penyakit, rantai_dna) SELECT * FROM (SELECT '%d', '%s', '%s') AS tmp WHERE NOT EXISTS (SELECT nama_penyakit FROM tesdna.jenispenyakit WHERE nama_penyakit='%s') LIMIT 1`, count + 1, diseaseName, dna, diseaseName)
	insert, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("INSERT SUCCESS")
	}
	defer insert.Close()
}

func InsertNewPrediction(tanggal_prediksi string, nama_pasien string, penyakit_terprediksi string, status_terprediksi string, password string) {
	var err error
	db, err = sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/tesdna")
	if err != nil {
		log.Fatal(err)     
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	query := fmt.Sprintf(`INSERT IGNORE INTO hasilprediksi VALUES ('%s', '%s', '%s', '%s')`, tanggal_prediksi, nama_pasien, penyakit_terprediksi, status_terprediksi)
	insert, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("INSERT SUCCESS")
	}
	defer insert.Close()
}



func SELECTDNA(diseaseName string, password string) (string) {
	// Capture connection properties.
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/tesdna")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	var ddna string
	query := fmt.Sprintf(`SELECT rantai_dna FROM tesdna.jenispenyakit WHERE nama_penyakit = "%s"`, diseaseName)
	// if table == "jenispenyakit" {
	err = db.QueryRow(query).Scan(&ddna)
	if err.Error() == "sql: no rows in result set" {
		// fmt.Println("Disease " + diseaseName + " is not in database")
		ddna = ""
	} else if err != nil {
		panic(err.Error())
	}
	return ddna
}

func SELECTRIWAYAT(input string, password string) []hasilprediksi{
	// I.S. Input string harus dipastikan benar menurut regex
	// Capture connection properties.
	// Get a database handle.
	var hp []hasilprediksi
	if algorithm.IsValidSearchDateAndDisease(input){
		splitInput := strings.Fields(input)
		if _, err := strconv.Atoi(splitInput[0]); err == nil { // Tanggal duluan
			date := splitInput[0] + " " + splitInput[1] + " " + splitInput[2]
			disease := splitInput[3]
			query := fmt.Sprintf(`SELECT * FROM tesdna.hasilprediksi WHERE tanggal_prediksi = '%s' AND penyakit_terprediksi = '%s'`, date, disease)
			rows, err := db.Query(query)

			if err == nil {
				defer rows.Close()
				for rows.Next(){
					var tp, np, pt, st string
					err := rows.Scan(&tp, &np, &pt, &st)
					if err != nil {
						log.Fatal(err)
					}
					hp = append(hp, hasilprediksi{tanggal_prediksi: tp, nama_pasien: np, penyakit_terprediksi: pt, status_terprediksi: st})
				}
				if err := rows.Err(); err != nil {
					log.Fatal(err)
				} 
				return hp
			} else if err.Error() == "sql: no rows in result set" {
				fmt.Println("No result found")
			} else {
				log.Fatal(err)
			}

		} else {
			date := splitInput[1] + " " + splitInput[2] + " " + splitInput[3]
			disease := splitInput[0]
			query := fmt.Sprintf(`SELECT * FROM tesdna.hasilprediksi WHERE tanggal_prediksi = '%s' AND penyakit_terprediksi = '%s'`, date, disease)
			rows, err := db.Query(query)
			if err == nil {
				defer rows.Close()
				for rows.Next(){
					var tp, np, pt, st string
					err := rows.Scan(&tp, &np, &pt, &st)
					if err != nil {
						log.Fatal(err)
					}
					hp = append(hp, hasilprediksi{tanggal_prediksi: tp, nama_pasien: np, penyakit_terprediksi: pt, status_terprediksi: st})
				}
				if err := rows.Err(); err != nil {
					log.Fatal(err)
				} 
				return hp
			} else if err.Error() == "sql: no rows in result set" {
				fmt.Println("No result found")
			} else {
				log.Fatal(err)
			}
		}

	} else if algorithm.IsValidSearchDateOnly(input){
		query := fmt.Sprintf(`SELECT * FROM tesdna.hasilprediksi WHERE tanggal_prediksi = '%s'`, input)
		rows, err := db.Query(query)
		if err == nil {
			defer rows.Close()
			for rows.Next(){
				var tp, np, pt, st string
				err := rows.Scan(&tp, &np, &pt, &st)
				if err != nil {
					log.Fatal(err)
				}
				hp = append(hp, hasilprediksi{tanggal_prediksi: tp, nama_pasien: np, penyakit_terprediksi: pt, status_terprediksi: st})
			}
			if err := rows.Err(); err != nil {
				log.Fatal(err)
			} 
			return hp
		} else if err.Error() == "sql: no rows in result set" {
			fmt.Println("No result found")
		} else {
			log.Fatal(err)
		}
	} else { // Search hanya dengan nama penyakit
		query := fmt.Sprintf(`SELECT * FROM tesdna.hasilprediksi WHERE penyakit_terprediksi = '%s'`, input)
		rows, err := db.Query(query)
		if err == nil {
			defer rows.Close()
			for rows.Next(){
				var tp, np, pt, st string
				err := rows.Scan(&tp, &np, &pt, &st)
				if err != nil {
					log.Fatal(err)
				}
				hp = append(hp, hasilprediksi{tanggal_prediksi: tp, nama_pasien: np, penyakit_terprediksi: pt, status_terprediksi: st})
			}
			if err := rows.Err(); err != nil {
				log.Fatal(err)
			} 
			return hp
		} else if err.Error() == "sql: no rows in result set" {
			fmt.Println("No result found")
		} else {
			log.Fatal(err)
		}
	} 
	return nil
}
*/