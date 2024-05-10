-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 10, 2024 at 04:43 AM
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
-- Database: `go_restapi_jemaatkeluarga`
--

-- --------------------------------------------------------

--
-- Table structure for table `jemaat_keluargas`
--

CREATE TABLE `jemaat_keluargas` (
  `id` int(11) NOT NULL,
  `nik` varchar(255) DEFAULT NULL,
  `no_kk` varchar(255) DEFAULT NULL,
  `status` enum('Suami','Istri','Anak','Menikah') DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `jemaat_keluargas`
--

INSERT INTO `jemaat_keluargas` (`id`, `nik`, `no_kk`, `status`) VALUES
(6, '2182178721100', 'kk1267215121rg', 'Suami'),
(7, 'A2182178721100', 'kk771216571', 'Anak'),
(8, '1234567890123456', '12345678', 'Suami'),
(9, '2345678901234567', '23456789', 'Istri'),
(10, '3456789012345678', '34567890', 'Anak'),
(11, '4567890123456789', '45678901', 'Menikah'),
(12, '5678901234567890', '56789012', 'Suami'),
(13, '6789012345678901', '67890123', 'Istri'),
(14, '7890123456789012', '78901234', 'Anak'),
(15, '8901234567890123', '89012345', 'Menikah'),
(16, '9012345678901234', '90123456', 'Anak'),
(17, '0123456789012345', '01234567', 'Anak');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `jemaat_keluargas`
--
ALTER TABLE `jemaat_keluargas`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `jemaat_keluargas`
--
ALTER TABLE `jemaat_keluargas`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
