package todb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type jenispenyakit struct {
	id_penyakit               int
	nama_penyakit, rantai_dna string
}

type hasilprediksi struct {
	tanggal_prediksi, nama_pasien, penyakit_terprediksi, status_terprediksi string
}

func insert(query string, password string) {
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
	insert, err := db.Query("INSERT INTO tesdna.jenispenyakit values ('1', 'tes', 'ABCDEFGH')")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("INSERT SUCCESS")
	}
	defer insert.Close()
	var p jenispenyakit
	err = db.QueryRow("SELECT * FROM tesdna.jenispenyakit where id_penyakit = 1").Scan(&p.id_penyakit, &p.nama_penyakit, &p.rantai_dna)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("id = %d, nama = %s, dna = %s\n", p.id_penyakit, p.nama_penyakit, p.rantai_dna)
	//print(insert)
	fmt.Println("Connected!")
}

func SELECT(query string, password string, table string) {
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
	if table == "jenispenyakit" {
		var p jenispenyakit
		err = db.QueryRow(query).Scan(&p.id_penyakit, &p.nama_penyakit, &p.rantai_dna)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("id = %d, nama = %s, dna = %s\n", p.id_penyakit, p.nama_penyakit, p.rantai_dna)

	} else if table == "hasilprediksi" {
		var p hasilprediksi
		err = db.QueryRow(query).Scan(&p.tanggal_prediksi, &p.nama_pasien, &p.penyakit_terprediksi, &p.status_terprediksi)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("tanggal = %s, nama = %s, penyakit = %s, status = %s\n", p.tanggal_prediksi, p.nama_pasien, p.penyakit_terprediksi, p.status_terprediksi)
	}
	fmt.Println("Connected!")
}
