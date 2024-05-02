-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 02, 2024 at 11:04 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_restapi_jemaat`
--

-- --------------------------------------------------------

--
-- Table structure for table `jemaats`
--

CREATE TABLE `jemaats` (
  `id` int(11) NOT NULL,
  `nik` varchar(20) DEFAULT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `jenis_kelamin` enum('Laki-laki','Perempuan') DEFAULT NULL,
  `alamat` varchar(255) DEFAULT NULL,
  `tempat_lahir` varchar(255) DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `gambar_profile` varchar(255) DEFAULT NULL,
  `status_gereja` enum('Aktif','Pindah','Meninggal','Nonaktif') DEFAULT NULL,
  `status_baptis` tinyint(1) DEFAULT NULL,
  `status_sidi` tinyint(1) DEFAULT NULL,
  `status_nikah` tinyint(1) DEFAULT NULL,
  `id_sektor` int(11) DEFAULT NULL,
  `no_telepon` varchar(20) DEFAULT NULL,
  `no_kk` varchar(20) DEFAULT NULL,
  `username` varchar(10) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `sektor_role` enum('Penanggung Jawab','Anggota') DEFAULT NULL,
  `lampiran` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `jemaats`
--

INSERT INTO `jemaats` (`id`, `nik`, `nama`, `jenis_kelamin`, `alamat`, `tempat_lahir`, `tanggal_lahir`, `gambar_profile`, `status_gereja`, `status_baptis`, `status_sidi`, `status_nikah`, `id_sektor`, `no_telepon`, `no_kk`, `username`, `password`, `sektor_role`, `lampiran`) VALUES
(1, '12345678901234567890', 'John Doe', 'Laki-laki', 'Jl. Contoh No. 123', 'jakarta', '1990-05-15', '330px-AlanJacksonApr10.jpg', 'Aktif', 1, 1, 0, 123, '081234567890', '1234567890', NULL, NULL, NULL, NULL),
(2, '17827287323', 'Silaen Huhu', 'Laki-laki', 'Jl. Orgorg, Sumatera Utara', 'tangerang', '1990-05-17', '330px-AlanJacksonApr10.jpg', 'Aktif', 1, 1, 0, 123, '081234567890', '1234567890', NULL, NULL, NULL, NULL),
(3, '17827287323', 'John Doe', 'Laki-laki', 'Jalan Contoh No. 123', 'Jakarta barat', '1990-05-20', '330px-AlanJacksonApr10.jpg', 'Aktif', 0, 0, 0, 1, '081234567890', NULL, 'johndoe', 'password123', 'Anggota', '');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `jemaats`
--
ALTER TABLE `jemaats`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `jemaats`
--
ALTER TABLE `jemaats`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
