-- phpMyAdmin SQL Dump
-- version 4.1.12
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Generation Time: Jun 12, 2022 at 04:04 AM
-- Server version: 5.6.16
-- PHP Version: 5.5.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Database: `pos_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `tbl_account`
--

CREATE TABLE IF NOT EXISTS `tbl_account` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `company_name` varchar(256) NOT NULL,
  `phone_number` varchar(12) NOT NULL,
  `email` varchar(256) NOT NULL,
  `address` varchar(500) NOT NULL,
  `business_type` tinyint(4) NOT NULL,
  `main_business_type` tinyint(4) DEFAULT NULL,
  `core_business_type` varchar(200) NOT NULL,
  `outlets_number` smallint(6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `phone_number` (`phone_number`),
  KEY `email` (`email`),
  KEY `company_name` (`company_name`),
  KEY `business_type` (`business_type`),
  KEY `core_business_type` (`core_business_type`),
  KEY `main_business_type` (`main_business_type`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=112 ;
