-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 09, 2024 at 05:58 PM
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
-- Database: `go_restapi_pelayangereja`
--

-- --------------------------------------------------------

--
-- Table structure for table `pelayan_gerejas`
--

CREATE TABLE `pelayan_gerejas` (
  `id` int(11) NOT NULL,
  `nik` varchar(255) DEFAULT NULL,
  `peran` enum('Pendeta','Penatua','PHJ','Pelayan_Ibadah','Tata Usaha') DEFAULT NULL,
  `tgl_terima_jabatan` datetime DEFAULT NULL,
  `tgl_akhir_jabatan` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `pelayan_gerejas`
--

INSERT INTO `pelayan_gerejas` (`id`, `nik`, `peran`, `tgl_terima_jabatan`, `tgl_akhir_jabatan`) VALUES
(1, '121210010702001002', 'Tata Usaha', '2024-05-04 00:00:00', '2026-05-04 00:00:00'),
(3, '1212100107020003', 'PHJ', '2022-05-27 00:00:00', '2026-05-27 00:00:00'),
(4, '1212100107020004', 'Penatua', '2022-05-22 00:00:00', '2022-05-22 00:00:00'),
(5, '12345678901234567', 'Pendeta', '2013-11-24 00:00:00', '2017-12-24 00:00:00'),
(6, '3671982801027', 'Tata Usaha', '2017-01-24 00:00:00', '2020-08-14 00:00:00');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `pelayan_gerejas`
--
ALTER TABLE `pelayan_gerejas`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `pelayan_gerejas`
--
ALTER TABLE `pelayan_gerejas`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
