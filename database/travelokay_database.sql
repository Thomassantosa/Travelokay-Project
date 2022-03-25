-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 25, 2022 at 04:27 PM
-- Server version: 10.4.20-MariaDB
-- PHP Version: 8.0.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `travelokay_database`
--

-- --------------------------------------------------------

--
-- Table structure for table `airlines`
--

CREATE TABLE `airlines` (
  `airline_id` int(10) NOT NULL,
  `airline_name` varchar(20) NOT NULL,
  `airline_contact` varchar(15) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `airlines`
--

INSERT INTO `airlines` (`airline_id`, `airline_name`, `airline_contact`) VALUES
(1, 'Air Asia', '0908615237'),
(2, 'Lion Air', '090861575'),
(3, 'Qatar Airways', '0568615237'),
(4, 'Sriwijaya Air', '0908645646');

-- --------------------------------------------------------

--
-- Table structure for table `airplanes`
--

CREATE TABLE `airplanes` (
  `airplane_id` int(10) NOT NULL,
  `airline_id` int(10) NOT NULL,
  `airplane_model` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `airplanes`
--

INSERT INTO `airplanes` (`airplane_id`, `airline_id`, `airplane_model`) VALUES
(1, 1, 'Boeing 101'),
(2, 1, 'Boeing 102'),
(3, 2, 'Boeing 103'),
(4, 2, 'Boeing 104'),
(5, 3, 'Airbus 101'),
(6, 3, 'Airbus 102'),
(7, 4, 'Airbus 103'),
(8, 4, 'Airbus 104');

-- --------------------------------------------------------

--
-- Table structure for table `airports`
--

CREATE TABLE `airports` (
  `airport_id` int(10) NOT NULL,
  `airport_code` varchar(5) NOT NULL,
  `airport_name` varchar(255) NOT NULL,
  `airport_city` varchar(50) NOT NULL,
  `airport_country` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `airports`
--

INSERT INTO `airports` (`airport_id`, `airport_code`, `airport_name`, `airport_city`, `airport_country`) VALUES
(1, 'CGK', 'Soekarno Hatta Airport', 'Tangerang', 'Indonesia'),
(2, 'BDO', 'Husein Sastranegara Airport', 'Bandung', 'Indonesia'),
(3, 'SIN', 'Changi Airport', 'Singapore', 'Singapore'),
(4, 'DPS', 'Ngurah Rai Airport', 'Bali', 'Indonesia');

-- --------------------------------------------------------

--
-- Table structure for table `buscompanies`
--

CREATE TABLE `buscompanies` (
  `buscompany_id` int(10) NOT NULL,
  `buscompany_name` varchar(20) NOT NULL,
  `buscompany_contact` varchar(15) NOT NULL,
  `buscompany_address` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `buscompanies`
--

INSERT INTO `buscompanies` (`buscompany_id`, `buscompany_name`, `buscompany_contact`, `buscompany_address`) VALUES
(1, 'Budiman', '08162534212', 'Jl. Sastrojaya'),
(2, 'Damri', '08162534122', 'Jl. Soekarno Hatta'),
(3, 'Harapan Jaya', '08162534213', 'Jl. Ali Sastro'),
(4, 'NPM', '08162534267', 'Jl. Agnezmo');

-- --------------------------------------------------------

--
-- Table structure for table `buses`
--

CREATE TABLE `buses` (
  `bus_id` int(10) NOT NULL,
  `buscompany_id` int(10) NOT NULL,
  `bus_model` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `buses`
--

INSERT INTO `buses` (`bus_id`, `buscompany_id`, `bus_model`) VALUES
(1, 1, 'HD'),
(2, 1, 'SHD'),
(3, 1, 'HDD'),
(4, 2, 'HD'),
(5, 2, 'SHD'),
(6, 2, 'HDD'),
(7, 3, 'HD'),
(8, 3, 'SHD'),
(9, 3, 'HDD'),
(10, 4, 'HD'),
(11, 4, 'SHD'),
(12, 4, 'HDD');

-- --------------------------------------------------------

--
-- Table structure for table `busstations`
--

CREATE TABLE `busstations` (
  `busstation_id` int(10) NOT NULL,
  `busstation_code` varchar(5) NOT NULL,
  `busstation_name` varchar(255) NOT NULL,
  `busstation_city` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `busstations`
--

INSERT INTO `busstations` (`busstation_id`, `busstation_code`, `busstation_name`, `busstation_city`) VALUES
(1, 'JKT-1', 'Terminal Lebak Bulus', 'Jakarta'),
(2, 'JKT-2', 'Terminal Grogol', 'Jakarta'),
(3, 'JKT-3', 'Terminal Blok M', 'Jakarta'),
(4, 'BDG-1', 'Terminal Cicaheum', 'Bandung'),
(5, 'BDG-2', 'Terminal Ledeng', 'Bandung'),
(6, 'BDG-3', 'Terminal Ciroyom', 'Bandung'),
(7, 'SBY-1', 'Terminal Osowilangun', 'Surabaya'),
(8, 'SMG-1', 'Terminal Mangkang', 'Semarang'),
(9, 'SMG-2', 'Terminal Pengaron', 'Semarang');

-- --------------------------------------------------------

--
-- Table structure for table `bustrips`
--

CREATE TABLE `bustrips` (
  `bustrip_id` int(10) NOT NULL,
  `bus_id` int(10) NOT NULL,
  `departure_busstation` int(10) NOT NULL,
  `destination_busstation` int(10) NOT NULL,
  `bustrip_number` varchar(20) NOT NULL,
  `departure_time` time NOT NULL,
  `arrival_time` time NOT NULL,
  `departure_date` date NOT NULL,
  `arrival_date` date NOT NULL,
  `travel_time` int(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `bustrips`
--

INSERT INTO `bustrips` (`bustrip_id`, `bus_id`, `departure_busstation`, `destination_busstation`, `bustrip_number`, `departure_time`, `arrival_time`, `departure_date`, `arrival_date`, `travel_time`) VALUES
(1, 1, 1, 4, 'B-1', '10:00:00', '14:00:00', '2020-11-29', '2020-11-29', 4),
(2, 1, 4, 1, 'B-2', '15:00:00', '19:00:00', '2020-11-28', '2020-11-28', 4),
(3, 1, 6, 7, 'B-3', '15:00:00', '20:00:00', '2020-11-28', '2020-11-28', 5),
(4, 1, 7, 4, 'B-4', '15:00:00', '20:00:00', '2020-11-29', '2020-11-29', 5),
(5, 2, 4, 7, 'B-5', '15:00:00', '20:00:00', '2020-11-29', '2020-11-29', 5),
(6, 2, 8, 2, 'B-6', '15:00:00', '20:00:00', '2020-11-29', '2020-11-29', 5),
(7, 2, 2, 8, 'B-7', '20:00:00', '01:00:00', '2021-11-28', '2020-11-29', 5),
(8, 3, 5, 1, 'B-8', '07:00:00', '10:00:00', '2020-11-29', '2020-11-29', 3),
(9, 3, 1, 5, 'B-9', '07:00:00', '10:00:00', '2020-11-29', '2020-11-29', 3),
(10, 3, 5, 9, 'B-10', '07:00:00', '12:00:00', '2020-11-29', '2020-11-29', 5);

-- --------------------------------------------------------

--
-- Table structure for table `flights`
--

CREATE TABLE `flights` (
  `flight_id` int(10) NOT NULL,
  `airplane_id` int(10) NOT NULL,
  `departure_airport` int(10) NOT NULL,
  `destination_airport` int(10) NOT NULL,
  `flight_type` varchar(15) NOT NULL,
  `flight_number` varchar(20) NOT NULL,
  `departure_time` time NOT NULL,
  `arrival_time` time NOT NULL,
  `departure_date` date NOT NULL,
  `arrival_date` date NOT NULL,
  `travel_time` int(2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `flights`
--

INSERT INTO `flights` (`flight_id`, `airplane_id`, `departure_airport`, `destination_airport`, `flight_type`, `flight_number`, `departure_time`, `arrival_time`, `departure_date`, `arrival_date`, `travel_time`) VALUES
(1, 1, 1, 3, 'International', 'FL-1', '11:20:00', '12:20:00', '2021-11-28', '2021-11-28', 1),
(2, 1, 1, 4, 'Local', 'FL-2', '15:35:00', '16:35:00', '2021-11-28', '2021-11-28', 1),
(3, 1, 1, 4, 'Local', 'FL-3', '19:35:00', '20:35:00', '2021-11-29', '2021-11-29', 1),
(4, 2, 2, 3, 'International', 'FL-4', '12:10:00', '13:10:00', '2021-11-28', '2021-11-28', 1),
(5, 2, 2, 4, 'Local', 'FL-5', '12:10:00', '13:10:00', '2021-11-29', '2021-11-29', 1),
(6, 2, 2, 3, 'International', 'FL-6', '20:10:00', '21:10:00', '2021-11-29', '2021-11-29', 1),
(7, 3, 3, 4, 'International', 'FL-7', '20:10:00', '21:10:00', '2021-11-28', '2021-11-28', 1),
(8, 3, 4, 3, 'International', 'FL-8', '20:10:00', '21:10:00', '2021-11-29', '2021-11-29', 1),
(9, 3, 4, 3, 'International', 'FL-9', '11:10:00', '12:10:00', '2021-11-29', '2021-11-29', 1),
(10, 4, 4, 1, 'Local', 'FL-10', '11:10:00', '13:10:00', '2021-11-29', '2021-11-29', 2),
(11, 3, 1, 3, 'International', 'FL_011', '21:00:00', '23:00:00', '2021-11-28', '2021-11-28', 2),
(12, 2, 1, 3, 'International', 'FL_022', '21:00:00', '22:00:00', '2021-11-28', '2021-11-28', 1),
(16, 7, 2, 3, 'International', 'FL_555', '21:00:00', '02:00:00', '2021-11-28', '2021-11-29', 5);

-- --------------------------------------------------------

--
-- Table structure for table `hotels`
--

CREATE TABLE `hotels` (
  `hotel_id` int(10) NOT NULL,
  `hotel_name` varchar(100) NOT NULL,
  `hotel_star` int(1) NOT NULL,
  `hotel_rating` decimal(2,1) NOT NULL DEFAULT 0.0,
  `hotel_review` int(10) NOT NULL DEFAULT 0,
  `hotel_facility` text NOT NULL,
  `hotel_address` text NOT NULL,
  `hotel_city` varchar(50) NOT NULL,
  `hotel_country` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `hotels`
--

INSERT INTO `hotels` (`hotel_id`, `hotel_name`, `hotel_star`, `hotel_rating`, `hotel_review`, `hotel_facility`, `hotel_address`, `hotel_city`, `hotel_country`) VALUES
(1, 'Aston Martin Bandung', 4, '4.6', 4, 'Gym, Swimming Pool, Spa, Gift Shop', 'Jl. Empang No.21', 'Bandung', 'Indonesia'),
(2, 'Santika Bali', 4, '4.8', 5, 'Gym, Swimming Pool, Spa', 'Jl. Ali Sastro No.25', 'Bali', 'Indonesia'),
(3, 'Sheraton Jakarta', 5, '4.8', 6, 'Gym, Swimming Pool, Spa, Water Purification System', 'Jl. Blackpink No.4', 'Jakarta', 'Indonesia'),
(4, 'Marina Bay Sands Singapore', 5, '5.0', 2, 'Gym, Swimming Pool, Spa, Jacuzi', 'Jl. Singapore No.50', 'Singapore', 'Singapore'),
(5, 'Grand Mercure Bali', 4, '4.8', 4, 'Gym, Swimming Pool, Spa, Gift Shop', 'Jl. Gajah Mada No.21', 'Bali', 'Indonesia'),
(6, 'Amaris Jakarta', 3, '4.2', 5, 'Gym, Swimming Pool, Spa, Gift Shop', 'Jl. Airlangga No.21', 'Jakarta', 'Indonesia'),
(7, 'Amaris Bandung', 3, '4.0', 6, 'Gym, Swimming Pool,', 'Jl. Gunung Nungkep No.69', 'Bandung', 'Indonesia'),
(8, 'Amaris Bali', 3, '3.8', 7, 'Gym, Swimming Pool, Spa, Jacuzi', 'Jl. PBO No.91', 'Bali', 'Indonesia'),
(9, 'Grand Mercure Bali', 4, '4.8', 8, 'Gym, Swimming Pool, Spa', 'Jl. Airlangga Sucipto No82', 'Bali', 'Indonesia'),
(10, 'Grand Mercure Singapore', 5, '4.9', 7, 'Gym, Swimming Pool, Spa', 'Jl. Majulah Singapore No.90', 'Singapore', 'Singapore'),
(11, 'Aston Martin Singapore', 4, '4.6', 3, 'Gym, Swimming Pool, Spa, Gift Shop', 'Jl. Empang No.25', 'Bandung', 'Indonesia');

-- --------------------------------------------------------

--
-- Table structure for table `imagelist`
--

CREATE TABLE `imagelist` (
  `image_id` int(10) NOT NULL,
  `hotel_id` int(10) NOT NULL,
  `room_id` int(10) NOT NULL,
  `tour_id` int(10) NOT NULL,
  `image_path` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `order_id` int(10) NOT NULL,
  `user_id` int(10) NOT NULL,
  `flight_id` int(10) DEFAULT NULL,
  `traintrip_id` int(10) DEFAULT NULL,
  `bustrip_id` int(10) DEFAULT NULL,
  `room_id` int(10) DEFAULT NULL,
  `tour_id` int(10) DEFAULT NULL,
  `order_date` date NOT NULL,
  `person_name` varchar(255) NOT NULL,
  `phone_number` varchar(20) NOT NULL,
  `email` varchar(255) NOT NULL,
  `transaction_type` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`order_id`, `user_id`, `flight_id`, `traintrip_id`, `bustrip_id`, `room_id`, `tour_id`, `order_date`, `person_name`, `phone_number`, `email`, `transaction_type`) VALUES
(1, 1, 1, NULL, NULL, NULL, NULL, '2021-11-22', 'Thomas Budi', '085634567', 'thomas@gmail.com', 'Gopay'),
(2, 5, NULL, 1, NULL, NULL, NULL, '2021-11-23', 'Thomas Budi Santosa', '23456123', 'thomassantosa07@gmail.com', 'BCA'),
(3, 5, 11, NULL, NULL, NULL, NULL, '2021-11-24', 'Farrel bin Bapaknya Farrel', '6666666666', 'farrelgeming@gmail.com', 'Bitcoin'),
(4, 5, 11, NULL, NULL, NULL, NULL, '2021-11-24', 'Renald bin Bapaknya Renald', '999999', 'renaldgeming@gmail.com', 'BRI'),
(5, 5, 4, NULL, NULL, NULL, NULL, '2021-11-24', 'Renald bin Bapaknya Renald', '01234567', 'renaldgeming@gmail.com', 'BRI'),
(6, 5, 12, NULL, NULL, NULL, NULL, '2021-11-24', 'Chris bin Bapaknya Chris', '1234567', 'chrisnggageming@gmail.com', 'Gopay'),
(8, 5, NULL, NULL, NULL, 12, NULL, '2021-11-24', 'Thomas Anak Baik', '123451234', 'thomasbaik@gmail.com', 'Bitcoin'),
(9, 14, NULL, NULL, NULL, 9, NULL, '2021-11-24', 'Budi Anak Pak Budi', '12345678', 'budibudiman@gmail.com', 'BCA'),
(10, 17, NULL, NULL, NULL, 14, NULL, '2021-11-24', 'Farrel Buaya', '1234567', 'farrel@gmail.com', 'OVO'),
(11, 14, NULL, NULL, NULL, 12, NULL, '2021-11-24', 'Thomas Orang Setia', '12345678', 'thomasbaik@gmail.com', 'Bitcoin'),
(12, 5, 1, NULL, NULL, NULL, NULL, '2021-11-24', 'Farrel Udah Punya Pacar', '12345678', 'farrel@gmail.com', 'Gopay'),
(13, 5, NULL, 1, NULL, NULL, NULL, '2021-11-24', 'Chris Sayang Istri', '123456789', 'chrisss@gmail.com', 'Gopay'),
(14, 5, NULL, NULL, 3, NULL, NULL, '2021-11-24', 'Renald Duta Grabfood', '12345678', 'renald123@gmail.com', 'Bitcoin'),
(15, 19, 12, NULL, NULL, NULL, NULL, '2021-11-25', 'Thomas Budi Santosa', '086234567', 'thomassantosa07@gmail.com', 'Bitcoin'),
(16, 19, NULL, 1, NULL, NULL, NULL, '2021-11-25', 'Farrel', '123456', 'ggGeming', 'BCA'),
(17, 19, NULL, NULL, NULL, 17, NULL, '2021-11-25', 'Renald', '12345678', 'nggaGGgemning@Gmail.com', 'BRI');

-- --------------------------------------------------------

--
-- Table structure for table `rooms`
--

CREATE TABLE `rooms` (
  `room_id` int(10) NOT NULL,
  `hotel_id` int(10) NOT NULL,
  `room_name` varchar(100) NOT NULL,
  `room_type` varchar(20) NOT NULL,
  `room_price` int(10) NOT NULL,
  `room_facility` text NOT NULL,
  `room_capacity` int(2) NOT NULL,
  `room_status` tinyint(1) NOT NULL DEFAULT 0,
  `checkin` date DEFAULT NULL,
  `checkout` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `rooms`
--

INSERT INTO `rooms` (`room_id`, `hotel_id`, `room_name`, `room_type`, `room_price`, `room_facility`, `room_capacity`, `room_status`, `checkin`, `checkout`) VALUES
(1, 1, 'Superior Twin Room', 'Superior Twin', 400000, 'AC, Kamar Mandi, Dapur, Kopi, Teh', 2, 0, NULL, NULL),
(2, 1, 'Deluxe King Room', 'Deluxe King', 750000, 'AC, Kamar Mandi, Shower, Dapur, Kopi, Teh', 4, 0, NULL, NULL),
(3, 2, 'Deluxe Twin Room', 'Deluxe King', 650000, 'AC, Kamar Mandi, Shower, Dapur, Kopi, Teh', 2, 0, NULL, NULL),
(4, 2, 'Premiere King Room', 'Premiere King', 800000, 'AC, Kamar Mandi, Shower,Tv, Bathtub, Dapur, Kopi, Teh', 4, 0, NULL, NULL),
(5, 3, 'Premiere Twin Room', 'Premiere Twin', 750000, 'AC, Kamar Mandi, Shower,Tv, Bathtub, Dapur, Kopi, Teh', 2, 0, NULL, NULL),
(6, 3, 'Junior Suite Room', 'Junior Suite', 450000, 'AC, Kamar Mandi, Shower,Tv, Kopi, Teh', 2, 0, NULL, NULL),
(7, 4, 'Family Room', 'Family', 1000000, 'AC, Kamar Mandi, Shower,Tv, Bathtub, Dapur, Kopi, Teh', 4, 0, NULL, NULL),
(8, 4, 'Tjokro Room', 'Tjokro', 1000000, 'AC, Kamar Mandi, Shower,Tv, Bathtub, Dapur, Kopi, Teh', 4, 0, NULL, NULL),
(9, 5, 'Superior Twin Room', 'Superior Twin', 400000, 'AC, Kamar Mandi, Dapur, Kopi, Teh', 2, 0, NULL, NULL),
(10, 5, 'Deluxe King Room', 'Deluxe King', 750000, 'AC, Kamar Mandi, Shower, Dapur, Kopi, Teh', 4, 0, NULL, NULL),
(11, 6, 'Deluxe Twin Room', 'Deluxe King', 650000, 'AC, Kamar Mandi, Shower, Dapur, Kopi, Teh', 2, 0, NULL, NULL),
(12, 6, 'Premiere King Room', 'Premiere King', 800000, 'AC, Kamar Mandi, Shower,Tv, Bathtub, Dapur, Kopi, Teh', 4, 0, NULL, NULL),
(13, 7, 'Premiere Twin Room', 'Premiere Twin', 750000, 'AC, Kamar Mandi, Shower,Tv, Bathtub, Dapur, Kopi, Teh', 2, 0, NULL, NULL),
(14, 7, 'Junior Suite Room', 'Junior Suite', 450000, 'AC, Kamar Mandi, Shower,Tv, Kopi, Teh', 2, 0, NULL, NULL),
(15, 8, 'Family Room', 'Family', 1000000, 'AC, Kamar Mandi, Shower,Tv, Bathtub, Dapur, Kopi, Teh', 4, 0, NULL, NULL),
(16, 8, 'Tjokro Room', 'Tjokro', 1000000, 'AC, Kamar Mandi, Shower,Tv, Bathtub, Dapur, Kopi, Teh', 4, 0, NULL, NULL),
(17, 9, 'Deluxe King Room', 'Deluxe King', 750000, 'AC, Kamar Mandi, Shower, Dapur, Kopi, Teh', 4, 0, NULL, NULL),
(18, 9, 'Deluxe Twin Room', 'Deluxe King', 650000, 'AC, Kamar Mandi, Shower, Dapur, Kopi, Teh', 2, 0, NULL, NULL),
(19, 10, 'Premiere King Room', 'Premiere King', 800000, 'AC, Kamar Mandi, Shower,Tv, Bathtub, Dapur, Kopi, Teh', 4, 0, NULL, NULL),
(20, 10, 'Premiere Twin Room', 'Premiere Twin', 750000, 'AC, Kamar Mandi, Shower,Tv, Bathtub, Dapur, Kopi, Teh', 2, 0, NULL, NULL),
(21, 11, 'Junior Suite Room', 'Junior Suite', 450000, 'AC, Kamar Mandi, Shower,Tv, Kopi, Teh', 2, 0, NULL, NULL),
(22, 11, 'Family Room', 'Family', 1000000, 'AC, Kamar Mandi, Shower,Tv, Bathtub, Dapur, Kopi, Teh', 4, 0, NULL, NULL),
(23, 11, 'Tjokro Room', 'Tjokro', 1000000, 'AC, Kamar Mandi, Shower,Tv, Bathtub, Dapur, Kopi, Teh', 4, 0, NULL, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `roomstatus`
--

CREATE TABLE `roomstatus` (
  `roomstatus_id` int(10) NOT NULL,
  `room_id` int(10) NOT NULL,
  `checkin` date NOT NULL,
  `checkout` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `schedules`
--

CREATE TABLE `schedules` (
  `schedule_id` int(10) NOT NULL,
  `tour_id` int(10) NOT NULL,
  `schedule_day` varchar(10) NOT NULL,
  `opentime` time NOT NULL,
  `closetime` time NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `seats`
--

CREATE TABLE `seats` (
  `seat_id` int(10) NOT NULL,
  `airplane_id` int(10) DEFAULT NULL,
  `train_id` int(10) DEFAULT NULL,
  `bus_id` int(10) DEFAULT NULL,
  `seat_type` varchar(10) NOT NULL,
  `seat_avaliable` int(3) NOT NULL,
  `baggage_capacity` int(3) NOT NULL,
  `seat_price` int(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `seats`
--

INSERT INTO `seats` (`seat_id`, `airplane_id`, `train_id`, `bus_id`, `seat_type`, `seat_avaliable`, `baggage_capacity`, `seat_price`) VALUES
(1, 1, NULL, NULL, 'Economy', 50, 40, 150000),
(2, 1, NULL, NULL, 'Bussiness', 25, 50, 300000),
(3, 1, NULL, NULL, 'First Clas', 25, 75, 500000),
(4, 2, NULL, NULL, 'Economy', 50, 40, 150000),
(5, 2, NULL, NULL, 'Bussiness', 25, 50, 300000),
(6, 2, NULL, NULL, 'First Clas', 25, 75, 500000),
(7, 3, NULL, NULL, 'Economy', 50, 40, 150000),
(8, 3, NULL, NULL, 'Bussiness', 25, 50, 300000),
(9, 3, NULL, NULL, 'First Clas', 25, 75, 500000),
(10, 4, NULL, NULL, 'Economy', 50, 40, 150000),
(11, 4, NULL, NULL, 'Bussiness', 25, 50, 300000),
(12, 4, NULL, NULL, 'First Clas', 25, 75, 500000),
(13, 5, NULL, NULL, 'Economy', 50, 40, 150000),
(14, 5, NULL, NULL, 'Bussiness', 25, 50, 300000),
(15, 5, NULL, NULL, 'First Clas', 25, 75, 500000),
(16, 6, NULL, NULL, 'Economy', 50, 40, 150000),
(17, 6, NULL, NULL, 'Bussiness', 25, 50, 300000),
(18, 6, NULL, NULL, 'First Clas', 25, 75, 500000),
(19, 7, NULL, NULL, 'Economy', 50, 40, 150000),
(20, 7, NULL, NULL, 'Bussiness', 25, 50, 300000),
(21, 7, NULL, NULL, 'First Clas', 25, 75, 500000),
(22, 8, NULL, NULL, 'Economy', 50, 40, 150000),
(23, 8, NULL, NULL, 'Bussiness', 25, 50, 300000),
(24, 8, NULL, NULL, 'First Clas', 25, 75, 500000),
(25, NULL, 1, NULL, 'Economy', 50, 40, 100000),
(26, NULL, 1, NULL, 'Bussiness', 25, 50, 300000),
(27, NULL, 1, NULL, 'First Clas', 25, 75, 400000),
(28, NULL, 2, NULL, 'Economy', 50, 40, 100000),
(29, NULL, 2, NULL, 'Bussiness', 25, 50, 300000),
(30, NULL, 2, NULL, 'First Clas', 25, 75, 400000),
(31, NULL, 3, NULL, 'Economy', 50, 40, 100000),
(32, NULL, 3, NULL, 'Bussiness', 25, 50, 300000),
(33, NULL, 3, NULL, 'First Clas', 25, 75, 400000),
(34, NULL, 4, NULL, 'Economy', 50, 40, 100000),
(35, NULL, 4, NULL, 'Bussiness', 25, 50, 300000),
(36, NULL, 4, NULL, 'First Clas', 25, 75, 400000),
(37, NULL, NULL, 1, 'Economy', 50, 40, 75000),
(38, NULL, NULL, 1, 'Bussiness', 25, 50, 150000),
(39, NULL, NULL, 1, 'First Clas', 25, 75, 20000),
(40, NULL, NULL, 2, 'Economy', 50, 40, 75000),
(41, NULL, NULL, 2, 'Bussiness', 25, 50, 150000),
(42, NULL, NULL, 2, 'First Clas', 25, 75, 200000),
(43, NULL, NULL, 3, 'Economy', 50, 40, 75000),
(44, NULL, NULL, 3, 'Bussiness', 25, 50, 150000),
(45, NULL, NULL, 3, 'First Clas', 25, 75, 200000),
(47, NULL, NULL, 4, 'Economy', 50, 40, 75000),
(48, NULL, NULL, 4, 'Bussiness', 25, 50, 150000),
(49, NULL, NULL, 4, 'First Clas', 25, 75, 200000);

-- --------------------------------------------------------

--
-- Table structure for table `stations`
--

CREATE TABLE `stations` (
  `station_id` int(10) NOT NULL,
  `station_code` varchar(50) NOT NULL,
  `station_name` varchar(255) NOT NULL,
  `station_city` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `stations`
--

INSERT INTO `stations` (`station_id`, `station_code`, `station_name`, `station_city`) VALUES
(1, 'BD', 'Stasion Bandung Kota', 'Bandung'),
(2, 'SB', 'Stasion Surabaya Kota', 'Surabaya'),
(3, 'JAKK', 'Stasion Jakarta Kota', 'Jakarta'),
(4, 'YK', 'Stasion Yogyakarta Kota', 'Yogyakarta'),
(5, 'SMT', 'Stasion Semarang Kota', 'Semarang');

-- --------------------------------------------------------

--
-- Table structure for table `tours`
--

CREATE TABLE `tours` (
  `tour_id` int(10) NOT NULL,
  `tour_name` varchar(100) NOT NULL,
  `tour_type` varchar(50) NOT NULL,
  `tour_rating` decimal(2,1) NOT NULL DEFAULT 0.0,
  `tour_review` int(10) NOT NULL,
  `tour_highlight` text NOT NULL,
  `tour_desc` text NOT NULL,
  `tour_facility` text NOT NULL,
  `tour_address` text NOT NULL,
  `tour_city` varchar(50) NOT NULL,
  `tour_province` varchar(50) NOT NULL,
  `tour_country` varchar(50) NOT NULL,
  `tour_price` int(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `trains`
--

CREATE TABLE `trains` (
  `train_id` int(10) NOT NULL,
  `train_model` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `trains`
--

INSERT INTO `trains` (`train_id`, `train_model`) VALUES
(1, 'KL01'),
(2, 'KL02'),
(3, 'KM01'),
(4, 'KM02');

-- --------------------------------------------------------

--
-- Table structure for table `traintrips`
--

CREATE TABLE `traintrips` (
  `traintrip_id` int(10) NOT NULL,
  `train_id` int(10) NOT NULL,
  `departure_station` int(10) NOT NULL,
  `destination_station` int(10) NOT NULL,
  `trainTrip_number` varchar(20) NOT NULL,
  `departure_time` time NOT NULL,
  `arrival_time` time NOT NULL,
  `departure_date` date NOT NULL,
  `arrival_date` date NOT NULL,
  `travel_time` int(2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `traintrips`
--

INSERT INTO `traintrips` (`traintrip_id`, `train_id`, `departure_station`, `destination_station`, `trainTrip_number`, `departure_time`, `arrival_time`, `departure_date`, `arrival_date`, `travel_time`) VALUES
(1, 1, 1, 2, 'TR-1', '09:00:00', '15:00:00', '2021-11-28', '2021-11-28', 6),
(2, 1, 2, 1, 'TR-1', '16:00:00', '22:00:00', '2021-11-28', '2021-11-28', 6),
(3, 4, 3, 2, 'TR-3', '16:00:00', '22:00:00', '2021-11-28', '2021-11-28', 6),
(4, 2, 2, 3, 'TR-4', '10:00:00', '16:00:00', '2021-11-28', '2021-11-28', 6),
(5, 3, 1, 4, 'TR-5', '10:00:00', '14:00:00', '2021-11-29', '2021-11-29', 6),
(6, 3, 3, 4, 'TR-6', '12:00:00', '18:00:00', '2021-11-29', '2021-11-29', 6),
(7, 3, 5, 4, 'TR-7', '14:00:00', '18:00:00', '2021-11-29', '2021-11-29', 4),
(8, 2, 2, 3, 'TR-8', '05:00:00', '08:00:00', '2021-01-01', '2021-01-01', 3),
(9, 4, 4, 1, 'TR-9', '15:00:00', '21:00:00', '2020-12-02', '2020-12-02', 6),
(10, 3, 4, 1, 'TR-10', '08:00:00', '13:00:00', '2021-03-15', '2021-03-15', 5),
(11, 4, 2, 3, 'TR-11', '12:00:00', '16:00:00', '2021-03-15', '2021-03-15', 4),
(12, 3, 5, 1, 'TR-12', '12:00:00', '15:00:00', '2021-03-18', '2021-03-18', 3),
(13, 4, 4, 2, 'TR-13', '13:00:00', '17:00:00', '2021-03-18', '2021-03-18', 4),
(14, 2, 5, 3, 'TR-14', '10:00:00', '16:00:00', '2021-03-20', '2021-03-20', 6);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `user_id` int(10) NOT NULL,
  `fullname` varchar(40) NOT NULL,
  `username` varchar(20) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `address` text NOT NULL,
  `user_type` int(10) NOT NULL,
  `partner_type` varchar(20) DEFAULT '',
  `company_name` varchar(100) DEFAULT NULL,
  `date_created` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`user_id`, `fullname`, `username`, `email`, `password`, `address`, `user_type`, `partner_type`, `company_name`, `date_created`) VALUES
(1, 'Christian', 'Christian01', 'christian@gmail.com', 'Christian01', 'Office', 0, NULL, NULL, NULL),
(2, 'Farrel', 'Farrel01', 'farrel@gmail.com', 'Farrel01', 'Office', 0, NULL, NULL, NULL),
(3, 'Renald', 'Renald01', 'renald@gmail.com', 'Renald01', 'Office', 0, NULL, NULL, NULL),
(4, 'Thomas Budi', 'Thomas01', 'thomas01@gmail.com', 'Thomas01', 'Office', 0, NULL, NULL, NULL),
(5, 'User1', 'User01', 'user@gmail.com', 'User01', 'Jawa Barat', 0, NULL, NULL, '2021-11-22'),
(6, 'Partner1', 'Partner01', 'partner@gmail.com', 'Partner01', 'Jawa Tengah', 0, 'Bus', 'PT. Bus Indonesia', '2021-11-22'),
(13, 'Farrel', 'FarrelTampan', 'farrel1@gmail.com', 'Farrel', 'jawa barat', 0, NULL, NULL, '2021-11-22'),
(14, 'User02', 'User02', 'User02@gmail.com', 'User02', 'Jawa Tengah', 0, NULL, NULL, '2021-11-23'),
(15, 'Partner02', 'Partner02', 'partner02@gmail.com', 'Partner02', 'Jawa Barat	', 0, 'Hotel', 'Hotel A', '2021-11-23'),
(16, 'Partner03', 'Partner03', 'partner03@gmail.com', 'Partner03', 'Jawa Timur', 0, 'Flight', 'Air Asia', '2021-11-23'),
(17, 'User03', 'User03', 'user03@gmail.com', 'User03', 'Singapore', 0, NULL, NULL, '2021-11-24'),
(18, 'Saudaranya Farrel', 'SaudaraFarrel', 'saudarafarrel@gmail.com', 'farrel123', 'Jawa Utara', 0, NULL, NULL, '2021-11-25'),
(19, 'User05', 'User05', 'user05@gmail.com', 'User05', 'Jawa Tenggara', 0, NULL, NULL, '2021-11-25'),
(20, 'Partner05', 'Partner05', 'partner05', 'Partner05', 'Jawa Timur Laut', 0, 'Flight', 'Sriwijaya Air', '2021-11-25');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `airlines`
--
ALTER TABLE `airlines`
  ADD PRIMARY KEY (`airline_id`);

--
-- Indexes for table `airplanes`
--
ALTER TABLE `airplanes`
  ADD PRIMARY KEY (`airplane_id`),
  ADD KEY `airline_id` (`airline_id`);

--
-- Indexes for table `airports`
--
ALTER TABLE `airports`
  ADD PRIMARY KEY (`airport_id`);

--
-- Indexes for table `buscompanies`
--
ALTER TABLE `buscompanies`
  ADD PRIMARY KEY (`buscompany_id`);

--
-- Indexes for table `buses`
--
ALTER TABLE `buses`
  ADD PRIMARY KEY (`bus_id`),
  ADD KEY `buscompany_id` (`buscompany_id`);

--
-- Indexes for table `busstations`
--
ALTER TABLE `busstations`
  ADD PRIMARY KEY (`busstation_id`);

--
-- Indexes for table `bustrips`
--
ALTER TABLE `bustrips`
  ADD PRIMARY KEY (`bustrip_id`),
  ADD KEY `bus_id` (`bus_id`),
  ADD KEY `departure_busstation` (`departure_busstation`),
  ADD KEY `destination_busstation` (`destination_busstation`);

--
-- Indexes for table `flights`
--
ALTER TABLE `flights`
  ADD PRIMARY KEY (`flight_id`),
  ADD KEY `airplane_id` (`airplane_id`),
  ADD KEY `departure_airport` (`departure_airport`),
  ADD KEY `destination_airport` (`destination_airport`);

--
-- Indexes for table `hotels`
--
ALTER TABLE `hotels`
  ADD PRIMARY KEY (`hotel_id`);

--
-- Indexes for table `imagelist`
--
ALTER TABLE `imagelist`
  ADD PRIMARY KEY (`image_id`),
  ADD KEY `room_id` (`room_id`),
  ADD KEY `tour_id` (`tour_id`),
  ADD KEY `hotel_id` (`hotel_id`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`order_id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `flight_id` (`flight_id`),
  ADD KEY `traintrip_id` (`traintrip_id`),
  ADD KEY `bustrip_id` (`bustrip_id`),
  ADD KEY `hotel_id` (`room_id`),
  ADD KEY `tour_id` (`tour_id`);

--
-- Indexes for table `rooms`
--
ALTER TABLE `rooms`
  ADD PRIMARY KEY (`room_id`),
  ADD KEY `hotel_id` (`hotel_id`);

--
-- Indexes for table `roomstatus`
--
ALTER TABLE `roomstatus`
  ADD PRIMARY KEY (`roomstatus_id`),
  ADD KEY `room_id` (`room_id`);

--
-- Indexes for table `schedules`
--
ALTER TABLE `schedules`
  ADD PRIMARY KEY (`schedule_id`),
  ADD KEY `tour_id` (`tour_id`);

--
-- Indexes for table `seats`
--
ALTER TABLE `seats`
  ADD PRIMARY KEY (`seat_id`),
  ADD KEY `airplane_id` (`airplane_id`),
  ADD KEY `bus_id` (`bus_id`),
  ADD KEY `train_id` (`train_id`);

--
-- Indexes for table `stations`
--
ALTER TABLE `stations`
  ADD PRIMARY KEY (`station_id`);

--
-- Indexes for table `tours`
--
ALTER TABLE `tours`
  ADD PRIMARY KEY (`tour_id`);

--
-- Indexes for table `trains`
--
ALTER TABLE `trains`
  ADD PRIMARY KEY (`train_id`);

--
-- Indexes for table `traintrips`
--
ALTER TABLE `traintrips`
  ADD PRIMARY KEY (`traintrip_id`),
  ADD KEY `train_id` (`train_id`),
  ADD KEY `departure_station` (`departure_station`),
  ADD KEY `destination_station` (`destination_station`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`user_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `airlines`
--
ALTER TABLE `airlines`
  MODIFY `airline_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `airplanes`
--
ALTER TABLE `airplanes`
  MODIFY `airplane_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `airports`
--
ALTER TABLE `airports`
  MODIFY `airport_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `buscompanies`
--
ALTER TABLE `buscompanies`
  MODIFY `buscompany_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `buses`
--
ALTER TABLE `buses`
  MODIFY `bus_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `busstations`
--
ALTER TABLE `busstations`
  MODIFY `busstation_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `bustrips`
--
ALTER TABLE `bustrips`
  MODIFY `bustrip_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `flights`
--
ALTER TABLE `flights`
  MODIFY `flight_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- AUTO_INCREMENT for table `hotels`
--
ALTER TABLE `hotels`
  MODIFY `hotel_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `imagelist`
--
ALTER TABLE `imagelist`
  MODIFY `image_id` int(10) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `order_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT for table `rooms`
--
ALTER TABLE `rooms`
  MODIFY `room_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;

--
-- AUTO_INCREMENT for table `roomstatus`
--
ALTER TABLE `roomstatus`
  MODIFY `roomstatus_id` int(10) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `schedules`
--
ALTER TABLE `schedules`
  MODIFY `schedule_id` int(10) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `seats`
--
ALTER TABLE `seats`
  MODIFY `seat_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=50;

--
-- AUTO_INCREMENT for table `stations`
--
ALTER TABLE `stations`
  MODIFY `station_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `tours`
--
ALTER TABLE `tours`
  MODIFY `tour_id` int(10) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `trains`
--
ALTER TABLE `trains`
  MODIFY `train_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `traintrips`
--
ALTER TABLE `traintrips`
  MODIFY `traintrip_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `user_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=21;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `airplanes`
--
ALTER TABLE `airplanes`
  ADD CONSTRAINT `airplanes_ibfk_1` FOREIGN KEY (`airline_id`) REFERENCES `airlines` (`airline_id`);

--
-- Constraints for table `buses`
--
ALTER TABLE `buses`
  ADD CONSTRAINT `buses_ibfk_1` FOREIGN KEY (`buscompany_id`) REFERENCES `buscompanies` (`buscompany_id`);

--
-- Constraints for table `bustrips`
--
ALTER TABLE `bustrips`
  ADD CONSTRAINT `bustrips_ibfk_1` FOREIGN KEY (`bus_id`) REFERENCES `buses` (`bus_id`),
  ADD CONSTRAINT `bustrips_ibfk_2` FOREIGN KEY (`departure_busstation`) REFERENCES `busstations` (`busstation_id`),
  ADD CONSTRAINT `bustrips_ibfk_3` FOREIGN KEY (`destination_busstation`) REFERENCES `busstations` (`busstation_id`);

--
-- Constraints for table `flights`
--
ALTER TABLE `flights`
  ADD CONSTRAINT `flights_ibfk_1` FOREIGN KEY (`airplane_id`) REFERENCES `airplanes` (`airplane_id`),
  ADD CONSTRAINT `flights_ibfk_2` FOREIGN KEY (`departure_airport`) REFERENCES `airports` (`airport_id`),
  ADD CONSTRAINT `flights_ibfk_3` FOREIGN KEY (`destination_airport`) REFERENCES `airports` (`airport_id`);

--
-- Constraints for table `imagelist`
--
ALTER TABLE `imagelist`
  ADD CONSTRAINT `imagelist_ibfk_1` FOREIGN KEY (`room_id`) REFERENCES `rooms` (`room_id`),
  ADD CONSTRAINT `imagelist_ibfk_2` FOREIGN KEY (`tour_id`) REFERENCES `tours` (`tour_id`),
  ADD CONSTRAINT `imagelist_ibfk_3` FOREIGN KEY (`hotel_id`) REFERENCES `hotels` (`hotel_id`);

--
-- Constraints for table `orders`
--
ALTER TABLE `orders`
  ADD CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`),
  ADD CONSTRAINT `orders_ibfk_2` FOREIGN KEY (`flight_id`) REFERENCES `flights` (`flight_id`),
  ADD CONSTRAINT `orders_ibfk_3` FOREIGN KEY (`traintrip_id`) REFERENCES `traintrips` (`traintrip_id`),
  ADD CONSTRAINT `orders_ibfk_4` FOREIGN KEY (`bustrip_id`) REFERENCES `bustrips` (`bustrip_id`),
  ADD CONSTRAINT `orders_ibfk_5` FOREIGN KEY (`room_id`) REFERENCES `rooms` (`room_id`),
  ADD CONSTRAINT `orders_ibfk_6` FOREIGN KEY (`tour_id`) REFERENCES `tours` (`tour_id`);

--
-- Constraints for table `rooms`
--
ALTER TABLE `rooms`
  ADD CONSTRAINT `rooms_ibfk_2` FOREIGN KEY (`hotel_id`) REFERENCES `hotels` (`hotel_id`);

--
-- Constraints for table `roomstatus`
--
ALTER TABLE `roomstatus`
  ADD CONSTRAINT `roomstatus_ibfk_1` FOREIGN KEY (`room_id`) REFERENCES `rooms` (`room_id`);

--
-- Constraints for table `schedules`
--
ALTER TABLE `schedules`
  ADD CONSTRAINT `schedules_ibfk_1` FOREIGN KEY (`tour_id`) REFERENCES `tours` (`tour_id`);

--
-- Constraints for table `seats`
--
ALTER TABLE `seats`
  ADD CONSTRAINT `seats_ibfk_1` FOREIGN KEY (`airplane_id`) REFERENCES `airplanes` (`airplane_id`),
  ADD CONSTRAINT `seats_ibfk_2` FOREIGN KEY (`bus_id`) REFERENCES `buses` (`bus_id`),
  ADD CONSTRAINT `seats_ibfk_3` FOREIGN KEY (`train_id`) REFERENCES `trains` (`train_id`);

--
-- Constraints for table `traintrips`
--
ALTER TABLE `traintrips`
  ADD CONSTRAINT `traintrips_ibfk_1` FOREIGN KEY (`train_id`) REFERENCES `trains` (`train_id`),
  ADD CONSTRAINT `traintrips_ibfk_2` FOREIGN KEY (`departure_station`) REFERENCES `stations` (`station_id`),
  ADD CONSTRAINT `traintrips_ibfk_3` FOREIGN KEY (`destination_station`) REFERENCES `stations` (`station_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
