-- MariaDB dump 10.19  Distrib 10.6.5-MariaDB, for Win64 (AMD64)
--
-- Host: localhost    Database: tesdna
-- ------------------------------------------------------
-- Server version	10.6.5-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `tesdna`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `tesdna` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `tesdna`;

--
-- Table structure for table `hasilprediksi`
--

DROP TABLE IF EXISTS `hasilprediksi`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `hasilprediksi` (
  `tanggal_prediksi` varchar(255) NOT NULL,
  `nama_pasien` varchar(255) NOT NULL,
  `penyakit_terprediksi` varchar(255) NOT NULL,
  `status_terprediksi` varchar(255) NOT NULL,
  PRIMARY KEY (`tanggal_prediksi`,`nama_pasien`,`penyakit_terprediksi`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `hasilprediksi`
--

LOCK TABLES `hasilprediksi` WRITE;
/*!40000 ALTER TABLE `hasilprediksi` DISABLE KEYS */;
INSERT INTO `hasilprediksi` (`tanggal_prediksi`, `nama_pasien`, `penyakit_terprediksi`, `status_terprediksi`) VALUES ('15 December 2020','Seryuu Ubiquitous','Trisomy','True'),('17 August 2010','Malty Melromarc','Klinefelter','True'),('24 October 2021','Chizuru Ichinose','Klinefelter','False'),('24 October 2021','Danzo Shimura','Neurofibromatosis','True'),('29 April 2022','apaaja','Klinefelter','True');
/*!40000 ALTER TABLE `hasilprediksi` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jenispenyakit`
--

DROP TABLE IF EXISTS `jenispenyakit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `jenispenyakit` (
  `id_penyakit` int(11) NOT NULL AUTO_INCREMENT,
  `nama_penyakit` varchar(255) NOT NULL,
  `rantai_dna` varchar(255) NOT NULL,
  PRIMARY KEY (`id_penyakit`,`nama_penyakit`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jenispenyakit`
--

LOCK TABLES `jenispenyakit` WRITE;
/*!40000 ALTER TABLE `jenispenyakit` DISABLE KEYS */;
INSERT INTO `jenispenyakit` (`id_penyakit`, `nama_penyakit`, `rantai_dna`) VALUES (1,'Klinefelter','CTAACCTTCATCCTT'),(2,'Trisomy','TCGACTACGACTACGTACG'),(3,'Neurofibromatosis','GACTGTAGCTACAGCT'),(4,'test','ACTGTACGATCGATCGATGCA');
/*!40000 ALTER TABLE `jenispenyakit` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-30  0:06:43
