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
-- Database: `go_restapi_gereja`
--

-- --------------------------------------------------------

--
-- Table structure for table `keluargas`
--

CREATE TABLE `keluargas` (
  `id` int(11) NOT NULL,
  `no_kk` varchar(255) DEFAULT NULL,
  `nama_keluarga` varchar(255) DEFAULT NULL,
  `alamat` varchar(255) DEFAULT NULL,
  `status` enum('Aktif','Pindah','Disabled') DEFAULT NULL,
  `tgl_nikah` datetime DEFAULT NULL,
  `lampiran` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `keluargas`
--

INSERT INTO `keluargas` (`id`, `no_kk`, `nama_keluarga`, `alamat`, `status`, `tgl_nikah`, `lampiran`) VALUES
(1, '236571211', 'jUANN pm', 'belgia, tambunan', 'Aktif', '2020-03-12 00:00:00', '5fcf53204e904.jpg'),
(2, '1219182', 'veri mutiha', 'pasar siantar', 'Disabled', '2022-04-30 00:00:00', 'Douwes.jpg');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `keluargas`
--
ALTER TABLE `keluargas`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `keluargas`
--
ALTER TABLE `keluargas`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
