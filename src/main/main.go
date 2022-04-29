package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/NayotamaPradipta/DeoxyriBoyer-Moore/src/algorithm"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/*
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
		// Manual input buat testing backend
		INSERT INTO hasilprediksi VALUES ("10 August 2020", "Ken Kaneki", "Ghoul", "True");
		INSERT INTO hasilprediksi VALUES ("27 October 2021", "Chizuru Ichinose", "Kawaiism", "True");
		INSERT INTO hasilprediksi VALUES ("27 October 2021", "Gabimaru", "Klinefelter", "False");
		riwayat := todb.SELECTRIWAYAT("Klinefelter 11 April 2020", "")
		fmt.Println(riwayat)
	*/
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/tesdna")
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	type jenispenyakit struct {
		Id_penyakit   int    `json:"id_penyakit"`
		Nama_penyakit string `json:"nama_penyakit"`
		Rantai_dna    string `json:"rantai_dna"`
	}

	type hasilprediksi struct {
		Tanggal_prediksi     string `json:"tanggal_prediksi"`
		Nama_pasien          string `json:"nama_pasien"`
		Penyakit_terprediksi string `json:"penyakit_terprediksi"`
		Status_terprediksi   string `json:"status_terprediksi"`
	}

	router.GET("/disease/:penyakit_terprediksi", func(c *gin.Context) {
		var (
			hp     hasilprediksi
			hPs    []hasilprediksi
			result gin.H
		)
		penyakit_terprediksi := c.Param("penyakit_terprediksi")
		if algorithm.IsValidSearchDiseaseOnly(penyakit_terprediksi) {
			rows, err := db.Query("SELECT tanggal_prediksi, nama_pasien, penyakit_terprediksi, status_terprediksi FROM hasilprediksi WHERE penyakit_terprediksi = ?;", penyakit_terprediksi)
			if err != nil {
				result = gin.H{
					"data":  "null",
					"Error": err.Error(),
				}
			} else {
				for rows.Next() {
					err = rows.Scan(&hp.Tanggal_prediksi, &hp.Nama_pasien, &hp.Penyakit_terprediksi, &hp.Status_terprediksi)
					hPs = append(hPs, hp)
					if err != nil {
						result = gin.H{
							"data":  "null",
							"Error": err.Error(),
						}
					} else {
						result = gin.H{
							"data": hPs,
						}
					}

				}
				defer rows.Close()
			}
			c.IndentedJSON(http.StatusOK, result)
		} else {
			result = gin.H{
				"data": "Invalid",
			}
			c.IndentedJSON(http.StatusOK, result)
		}

	})
	router.GET("/date/:tanggal_prediksi", func(c *gin.Context) {
		var (
			hp     hasilprediksi
			hPs    []hasilprediksi
			result gin.H
		)
		tanggal_prediksi := c.Param("tanggal_prediksi")
		tanggal_prediksi = tanggal_prediksi[:2] + " " + tanggal_prediksi[2:]
		tanggal_prediksi = tanggal_prediksi[:len(tanggal_prediksi)-4] + " " + tanggal_prediksi[len(tanggal_prediksi)-4:]
		if algorithm.IsValidSearchDateOnly(tanggal_prediksi) {
			rows, err := db.Query("SELECT tanggal_prediksi, nama_pasien, penyakit_terprediksi, status_terprediksi FROM hasilprediksi WHERE tanggal_prediksi = ?;", tanggal_prediksi)
			if err != nil {
				result = gin.H{
					"data":  "null",
					"Error": err.Error(),
				}
			} else {
				for rows.Next() {
					err = rows.Scan(&hp.Tanggal_prediksi, &hp.Nama_pasien, &hp.Penyakit_terprediksi, &hp.Status_terprediksi)
					hPs = append(hPs, hp)
					if err != nil {
						result = gin.H{
							"data":  "null",
							"Error": err.Error(),
						}
					} else {
						result = gin.H{
							"data": hPs,
						}
					}
				}
				defer rows.Close()
			}
			c.IndentedJSON(http.StatusOK, result)
		} else {
			result = gin.H{
				"data": "Invalid",
			}
			c.IndentedJSON(http.StatusOK, result)
		}
	})
	router.GET("/dnd/:penyakit_terprediksi/:tanggal_prediksi", func(c *gin.Context) {
		var (
			hp     hasilprediksi
			hPs    []hasilprediksi
			result gin.H
		)
		penyakit_terprediksi := c.Param("penyakit_terprediksi")
		tanggal_prediksi := c.Param("tanggal_prediksi")
		tanggal_prediksi = tanggal_prediksi[:2] + " " + tanggal_prediksi[2:]
		tanggal_prediksi = tanggal_prediksi[:len(tanggal_prediksi)-4] + " " + tanggal_prediksi[len(tanggal_prediksi)-4:]
		combine := tanggal_prediksi + " " + penyakit_terprediksi
		if algorithm.IsValidSearchDateAndDisease(combine) {
			rows, err := db.Query("SELECT tanggal_prediksi, nama_pasien, penyakit_terprediksi, status_terprediksi FROM hasilprediksi WHERE tanggal_prediksi = ? AND penyakit_terprediksi = ?;", tanggal_prediksi, penyakit_terprediksi)
			if err != nil {
				result = gin.H{
					"data":  "null",
					"Error": err.Error(),
				}
			} else {
				for rows.Next() {
					err = rows.Scan(&hp.Tanggal_prediksi, &hp.Nama_pasien, &hp.Penyakit_terprediksi, &hp.Status_terprediksi)
					hPs = append(hPs, hp)
					if err != nil {
						result = gin.H{
							"data":  "null",
							"Error": err.Error(),
						}
					} else {
						result = gin.H{
							"data": hPs,
						}
					}
				}
				defer rows.Close()
			}
			c.IndentedJSON(http.StatusOK, result)
		} else {
			result = gin.H{
				"data": "Invalid",
				"test": tanggal_prediksi,
			}
			c.IndentedJSON(http.StatusOK, result)
		}

	})

	router.POST("/insertDisease", func(c *gin.Context) {
		id_penyakit, _ := strconv.Atoi(c.PostForm("id_penyakit"))
		nama_penyakit := c.PostForm("nama_penyakit")
		rantai_dna := c.PostForm("rantai_dna")
		if algorithm.IsValidString(rantai_dna) {
			query := fmt.Sprintf(`INSERT INTO tesdna.jenispenyakit (id_penyakit, nama_penyakit, rantai_dna) SELECT * FROM (SELECT '%d', '%s', '%s') AS tmp WHERE NOT EXISTS (SELECT nama_penyakit FROM tesdna.jenispenyakit WHERE nama_penyakit='%s') LIMIT 1`, id_penyakit, nama_penyakit, rantai_dna, nama_penyakit)
			insert, err := db.Prepare(query)

			if err != nil {
				fmt.Print(err.Error())
			}
			_, err = insert.Exec(id_penyakit, nama_penyakit, rantai_dna)
			if err != nil {
				c.IndentedJSON(http.StatusOK, gin.H{
					"Message": "Insertion of new disease is unsuccessful",
				})
			}
			c.IndentedJSON(http.StatusOK, gin.H{
				"Message": "Insertion of new disease is successful",
			})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"Message": "Invalid Rantai DNA"})
		}
	})
	router.POST("/insertDNA", func(c *gin.Context) {
		nama_penyakit := c.PostForm("nama_penyakit")
		query := fmt.Sprintf(`SELECT rantai_dna FROM jenispenyakit WHERE nama_penyakit = "%s"`, nama_penyakit)
		var rdna string
		err = db.QueryRow(query).Scan(&rdna)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
		} else {
			nama_pasien := c.PostForm("nama_pasien")
			rantai_dna := c.PostForm("rantai_dna")
			if algorithm.IsValidString(rantai_dna) { // Input valid dan disease yang dicari ada di database
				now := time.Now()
				y, m, d := now.Date()
				date := strconv.Itoa(d) + " " + m.String() + " " + strconv.Itoa(y)
				var diagnosis string
				if algorithm.StartBoyerMoore(rantai_dna, rdna) && algorithm.StartKMP(rantai_dna, rdna) {
					diagnosis = "True"
				} else {
					diagnosis = "False"
				}
				query := fmt.Sprintf(`INSERT IGNORE INTO hasilprediksi VALUES ('%s', '%s', '%s', '%s')`, date, nama_pasien, nama_penyakit, diagnosis)
				insert, err := db.Query(query)
				if err != nil {
					c.IndentedJSON(http.StatusOK, gin.H{
						"Message": "Invalid Input file!",
					})
				} else {
					var hP hasilprediksi
					hP.Tanggal_prediksi = date
					hP.Nama_pasien = nama_pasien
					hP.Penyakit_terprediksi = nama_penyakit
					hP.Status_terprediksi = diagnosis
					c.IndentedJSON(http.StatusOK, gin.H{
						"HasilPrediksi": hP,
					})

				}
				defer insert.Close()
			} else {
				c.IndentedJSON(http.StatusOK, gin.H{
					"Message": "Invalid Input file!",
				})
			}
		}
	})
	router.Run(":8080")
}
