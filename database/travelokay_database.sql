-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 09, 2022 at 06:09 AM
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
  `airline_name` varchar(100) NOT NULL,
  `airline_contact` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `airlines`
--

INSERT INTO `airlines` (`airline_id`, `airline_name`, `airline_contact`) VALUES
(1, 'Air Asia', 'airasia@gmail.com'),
(2, 'Citilink', 'citilink@gmail.com'),
(3, 'Garuda Air', 'garudaair@gmail.com');

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
(1, 1, 'Boeing 456'),
(2, 1, 'Boeing 737'),
(3, 2, 'Boeing 457'),
(4, 2, 'Boeing 456'),
(5, 3, 'Airbus 841'),
(6, 3, 'Airbus 736');

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
(1, 'CGK', 'Bandara Internasional Soekarno-Hatta', 'tangerang', 'indonesia'),
(2, 'DPS', 'Bandara Internasional Ngurah Rai', 'denpasar', 'indonesia'),
(3, 'SUB', 'Bandara Internasional Juanda', 'sidoarjo', 'indonesia'),
(4, 'JPH', 'Bandara Haneda', 'tokyo', 'japan');

-- --------------------------------------------------------

--
-- Table structure for table `buscompanies`
--

CREATE TABLE `buscompanies` (
  `buscompany_id` int(10) NOT NULL,
  `buscompany_name` varchar(100) NOT NULL,
  `buscompany_contact` varchar(100) NOT NULL,
  `buscompany_address` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `buscompanies`
--

INSERT INTO `buscompanies` (`buscompany_id`, `buscompany_name`, `buscompany_contact`, `buscompany_address`) VALUES
(1, 'PT Bus Nusantara', 'nusantara@gmail.com', 'jawa tengah');

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
(1, 1, 'Mercedes a11'),
(2, 1, 'Mercedes b33');

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
(1, 'BA', 'Stasiun Semarang', 'semarang'),
(2, 'BB', 'Stasiun Bandung', 'semarang');

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
  `departure_time` datetime NOT NULL,
  `arrival_time` datetime NOT NULL,
  `travel_time` int(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `bustrips`
--

INSERT INTO `bustrips` (`bustrip_id`, `bus_id`, `departure_busstation`, `destination_busstation`, `bustrip_number`, `departure_time`, `arrival_time`, `travel_time`) VALUES
(1, 1, 1, 2, 'BT001', '2022-04-09 18:00:00', '2022-04-10 02:00:00', 8);

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
  `departure_time` datetime NOT NULL,
  `arrival_time` datetime NOT NULL,
  `travel_time` int(2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `flights`
--

INSERT INTO `flights` (`flight_id`, `airplane_id`, `departure_airport`, `destination_airport`, `flight_type`, `flight_number`, `departure_time`, `arrival_time`, `travel_time`) VALUES
(1, 1, 1, 2, 'domestic', 'F1111', '2022-04-09 10:00:00', '2022-04-09 12:00:00', 2),
(2, 2, 1, 3, 'domestic', 'F1112', '2022-04-10 10:00:00', '2022-04-10 13:00:00', 3),
(3, 3, 3, 4, 'international', 'F1113', '2022-04-11 14:00:00', '2022-04-11 18:00:00', 4),
(4, 4, 1, 4, 'international', 'F1114', '2022-04-13 15:00:00', '2022-04-13 19:00:00', 4);

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
(1, 'Hotel Indonesia', 5, '4.8', 1000, 'pool, park', 'jln. bundaran hi', 'jakarta', 'indoesia');

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `order_id` int(10) NOT NULL,
  `user_id` int(10) NOT NULL,
  `seat_id` int(10) DEFAULT NULL,
  `room_id` int(10) DEFAULT NULL,
  `tourschedule_id` int(10) DEFAULT NULL,
  `order_date` datetime NOT NULL DEFAULT current_timestamp(),
  `order_status` varchar(10) NOT NULL DEFAULT 'ordered',
  `transaction_type` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`order_id`, `user_id`, `seat_id`, `room_id`, `tourschedule_id`, `order_date`, `order_status`, `transaction_type`) VALUES
(1, 2, 1, NULL, NULL, '2022-04-09 11:04:16', 'ordered', 'BCA'),
(2, 2, NULL, 1, NULL, '2022-04-09 11:05:36', 'ordered', 'BNI'),
(3, 2, NULL, NULL, 1, '2022-04-09 11:06:29', 'ordered', 'BCA');

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
(1, 1, 'Suite Room', 'suite', 1000000, 'ac, shower, breakfast', 2, 1, NULL, NULL),
(2, 1, 'VIP room', 'vip', 3000000, 'ac, shower, breakfast', 1, 0, NULL, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `seats`
--

CREATE TABLE `seats` (
  `seat_id` int(10) NOT NULL,
  `flight_id` int(10) DEFAULT NULL,
  `traintrip_id` int(10) DEFAULT NULL,
  `bustrip_id` int(10) DEFAULT NULL,
  `seat_type` varchar(10) NOT NULL,
  `seat_name` varchar(5) NOT NULL,
  `seat_status` tinyint(1) NOT NULL DEFAULT 0,
  `baggage_capacity` int(3) NOT NULL,
  `seat_price` int(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `seats`
--

INSERT INTO `seats` (`seat_id`, `flight_id`, `traintrip_id`, `bustrip_id`, `seat_type`, `seat_name`, `seat_status`, `baggage_capacity`, `seat_price`) VALUES
(1, 1, NULL, NULL, 'bussiness', 'A1', 1, 10, 1000000),
(2, 1, NULL, NULL, 'bussiness', 'A2', 0, 10, 1000000),
(3, 1, NULL, NULL, 'bussiness', 'B1', 0, 10, 1000000),
(4, 1, NULL, NULL, 'bussiness', 'B2', 0, 10, 1000000),
(5, 1, NULL, NULL, 'economy', 'C1', 0, 8, 600000),
(6, 1, NULL, NULL, 'economy', 'C2', 0, 8, 600000),
(7, 1, NULL, NULL, 'economy', 'D1', 0, 8, 600000),
(8, 1, NULL, NULL, 'economy', 'D2', 0, 8, 600000),
(17, NULL, NULL, 1, 'executive', 'A1', 0, 10, 300000),
(18, NULL, NULL, 1, 'executive', 'A2', 0, 10, 300000),
(19, NULL, NULL, 1, 'executive', 'B1', 0, 10, 300000),
(20, NULL, NULL, 1, 'executive', 'B2', 0, 10, 300000),
(21, NULL, NULL, 1, 'economy', 'C1', 0, 8, 150000),
(22, NULL, NULL, 1, 'economy', 'C2', 0, 8, 150000),
(23, NULL, NULL, 1, 'economy', 'D1', 0, 8, 150000),
(24, NULL, NULL, 1, 'economy', 'D2', 0, 8, 150000);

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
(1, 'GMR', 'stasiun gambir', 'jakarta'),
(2, 'SMT', 'stasiun tawang', 'semarang'),
(3, 'BD', 'stasiun bandung', 'bandung');

-- --------------------------------------------------------

--
-- Table structure for table `tours`
--

CREATE TABLE `tours` (
  `tour_id` int(10) NOT NULL,
  `tour_name` varchar(100) NOT NULL,
  `tour_rating` decimal(2,1) NOT NULL DEFAULT 0.0,
  `tour_review` int(10) NOT NULL,
  `tour_desc` text NOT NULL,
  `tour_facility` text NOT NULL,
  `tour_address` text NOT NULL,
  `tour_city` varchar(50) NOT NULL,
  `tour_province` varchar(50) NOT NULL,
  `tour_country` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `tours`
--

INSERT INTO `tours` (`tour_id`, `tour_name`, `tour_rating`, `tour_review`, `tour_desc`, `tour_facility`, `tour_address`, `tour_city`, `tour_province`, `tour_country`) VALUES
(1, 'ragunan zoo', '4.6', 1040, 'the place of variant animals from around the world', 'park, restaurant, children ground', 'jln. ragunan 1', 'jakarta selatan', 'jakarta', 'indonesia');

-- --------------------------------------------------------

--
-- Table structure for table `tourschedules`
--

CREATE TABLE `tourschedules` (
  `schedule_id` int(10) NOT NULL,
  `tour_id` int(10) NOT NULL,
  `schedule_day` int(1) NOT NULL,
  `opentime` time NOT NULL,
  `closetime` time NOT NULL,
  `price` int(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `tourschedules`
--

INSERT INTO `tourschedules` (`schedule_id`, `tour_id`, `schedule_day`, `opentime`, `closetime`, `price`) VALUES
(1, 1, 0, '09:00:00', '18:00:00', 150000),
(2, 1, 1, '10:00:00', '16:00:00', 100000),
(3, 1, 2, '10:00:00', '16:00:00', 100000),
(4, 1, 3, '10:00:00', '16:00:00', 100000),
(5, 1, 4, '10:00:00', '16:00:00', 100000),
(6, 1, 5, '10:00:00', '16:00:00', 100000),
(7, 1, 6, '09:00:00', '18:00:00', 150000);

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
(1, 'Train A'),
(2, 'Train B'),
(3, 'Train C');

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
  `departure_time` datetime NOT NULL,
  `arrival_time` datetime NOT NULL,
  `travel_time` int(2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `traintrips`
--

INSERT INTO `traintrips` (`traintrip_id`, `train_id`, `departure_station`, `destination_station`, `trainTrip_number`, `departure_time`, `arrival_time`, `travel_time`) VALUES
(1, 1, 1, 2, 'TT001', '2022-04-09 10:00:00', '2022-04-09 18:00:00', 8);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `user_id` int(10) NOT NULL,
  `fullname` varchar(100) NOT NULL,
  `username` varchar(100) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `address` text NOT NULL,
  `user_type` int(1) NOT NULL,
  `partner_type` varchar(20) DEFAULT NULL,
  `company_name` varchar(255) DEFAULT NULL,
  `date_created` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`user_id`, `fullname`, `username`, `email`, `password`, `address`, `user_type`, `partner_type`, `company_name`, `date_created`) VALUES
(1, 'admin1', 'admin1', 'admin1@gmail.com', 'Admin1', 'Jakarta', 0, NULL, NULL, '2022-04-09 08:57:05'),
(2, 'user1', 'user1', 'user1@gmail.com', 'User1', 'Jawa Barat', 1, NULL, NULL, '2022-04-09 08:57:05'),
(3, 'flight partner1', 'flightpartner1', 'flightpartner1@gmail.com', 'FlightPartner1', 'Jawa Tengah', 2, 'flight', 'Air Asia', '2022-04-09 09:04:17'),
(4, 'train partner1', 'trainpartner1', 'trainpartner1@gmail.com', 'TrainPartner1', 'Jawa Timur', 2, 'train', 'PT KAI', '2022-04-09 09:04:17'),
(5, 'bus partner1', 'buspartner1', 'buspartner1@gmail.com', 'BusPartner1', 'Banten', 2, 'bus', 'PT Bus Indonesia', '2022-04-09 09:04:17'),
(6, 'hotel partner1', 'hotelpartner1', 'hotelpartner1@gmail.com', 'HotelPartner1', 'Bali', 2, 'hotel', 'Bali International Hotel', '2022-04-09 09:04:17'),
(7, 'tour partner1', 'tourpartner1', 'tourpartner1@gmail.com', 'TourPartner1', 'Bali', 2, 'tour', 'Sanggar Tari Kecak', '2022-04-09 09:04:17');

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
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`order_id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `hotel_id` (`room_id`),
  ADD KEY `tour_id` (`tourschedule_id`),
  ADD KEY `seat_id` (`seat_id`);

--
-- Indexes for table `rooms`
--
ALTER TABLE `rooms`
  ADD PRIMARY KEY (`room_id`),
  ADD KEY `hotel_id` (`hotel_id`);

--
-- Indexes for table `seats`
--
ALTER TABLE `seats`
  ADD PRIMARY KEY (`seat_id`),
  ADD KEY `airplane_id` (`flight_id`),
  ADD KEY `bus_id` (`bustrip_id`),
  ADD KEY `train_id` (`traintrip_id`);

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
-- Indexes for table `tourschedules`
--
ALTER TABLE `tourschedules`
  ADD PRIMARY KEY (`schedule_id`),
  ADD KEY `tour_id` (`tour_id`);

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
  MODIFY `airline_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `airplanes`
--
ALTER TABLE `airplanes`
  MODIFY `airplane_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `airports`
--
ALTER TABLE `airports`
  MODIFY `airport_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `buscompanies`
--
ALTER TABLE `buscompanies`
  MODIFY `buscompany_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `buses`
--
ALTER TABLE `buses`
  MODIFY `bus_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `busstations`
--
ALTER TABLE `busstations`
  MODIFY `busstation_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `bustrips`
--
ALTER TABLE `bustrips`
  MODIFY `bustrip_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `flights`
--
ALTER TABLE `flights`
  MODIFY `flight_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `hotels`
--
ALTER TABLE `hotels`
  MODIFY `hotel_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `order_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `rooms`
--
ALTER TABLE `rooms`
  MODIFY `room_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `seats`
--
ALTER TABLE `seats`
  MODIFY `seat_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;

--
-- AUTO_INCREMENT for table `stations`
--
ALTER TABLE `stations`
  MODIFY `station_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `tours`
--
ALTER TABLE `tours`
  MODIFY `tour_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `tourschedules`
--
ALTER TABLE `tourschedules`
  MODIFY `schedule_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `trains`
--
ALTER TABLE `trains`
  MODIFY `train_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `traintrips`
--
ALTER TABLE `traintrips`
  MODIFY `traintrip_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `user_id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

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
-- Constraints for table `orders`
--
ALTER TABLE `orders`
  ADD CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`),
  ADD CONSTRAINT `orders_ibfk_5` FOREIGN KEY (`room_id`) REFERENCES `rooms` (`room_id`),
  ADD CONSTRAINT `orders_ibfk_6` FOREIGN KEY (`tourschedule_id`) REFERENCES `tourschedules` (`schedule_id`),
  ADD CONSTRAINT `orders_ibfk_7` FOREIGN KEY (`seat_id`) REFERENCES `seats` (`seat_id`);

--
-- Constraints for table `rooms`
--
ALTER TABLE `rooms`
  ADD CONSTRAINT `rooms_ibfk_2` FOREIGN KEY (`hotel_id`) REFERENCES `hotels` (`hotel_id`);

--
-- Constraints for table `seats`
--
ALTER TABLE `seats`
  ADD CONSTRAINT `seats_ibfk_1` FOREIGN KEY (`flight_id`) REFERENCES `flights` (`flight_id`),
  ADD CONSTRAINT `seats_ibfk_2` FOREIGN KEY (`bustrip_id`) REFERENCES `bustrips` (`bustrip_id`),
  ADD CONSTRAINT `seats_ibfk_3` FOREIGN KEY (`traintrip_id`) REFERENCES `traintrips` (`traintrip_id`);

--
-- Constraints for table `tourschedules`
--
ALTER TABLE `tourschedules`
  ADD CONSTRAINT `tourschedules_ibfk_1` FOREIGN KEY (`tour_id`) REFERENCES `tours` (`tour_id`);

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
