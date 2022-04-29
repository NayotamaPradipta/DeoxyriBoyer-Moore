CREATE TABLE jenispenyakit
(id_penyakit int NOT NULL AUTO_INCREMENT,
nama_penyakit varchar(255) NOT NULL,
rantai_dna varchar(255) NOT NULL,
PRIMARY KEY (id_penyakit, nama_penyakit)
);

CREATE TABLE hasilprediksi 
(tanggal_prediksi varchar(255) NOT NULL,
nama_pasien varchar(255) NOT NULL,
penyakit_terprediksi varchar(255) NOT NULL,
status_terprediksi varchar(255) NOT NULL
PRIMARY KEY (tanggal_prediksi, nama_pasien, penyakit_terprediksi)
);