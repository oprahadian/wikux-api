-- MySQL dump 10.13  Distrib 8.0.16, for Linux (x86_64)
--
-- Host: localhost    Database: wikufest_db
-- ------------------------------------------------------
-- Server version	8.0.16

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8mb4 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `course`
--

DROP TABLE IF EXISTS `course`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `course` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `is_private_class` tinyint(1) DEFAULT NULL,
  `course_name` varchar(255) DEFAULT NULL,
  `description` text,
  `additional_link` text,
  `is_active` tinyint(1) DEFAULT NULL,
  `reg_date` datetime DEFAULT NULL,
  `upd_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `course`
--

LOCK TABLES `course` WRITE;
/*!40000 ALTER TABLE `course` DISABLE KEYS */;
INSERT INTO `course` VALUES (1,0,'THE COURSE NAME','THE COURSE DESCRIPTION','https://wikufest.org/',1,'2019-06-18 07:20:12',NULL),(2,1,'THE PRIVATE COURSE NAME','THE PRIVATE COURSE DESCRIPTION','https://wikufest.org/',1,'2019-06-18 07:20:44',NULL);
/*!40000 ALTER TABLE `course` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `course_instructor`
--

DROP TABLE IF EXISTS `course_instructor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `course_instructor` (
  `course_id` int(11) DEFAULT NULL,
  `instructor_id` int(11) DEFAULT NULL,
  `reg_date` datetime DEFAULT NULL,
  KEY `course_id` (`course_id`),
  KEY `instructor_id` (`instructor_id`),
  CONSTRAINT `course_instructor_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`),
  CONSTRAINT `course_instructor_ibfk_2` FOREIGN KEY (`instructor_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `course_instructor`
--

LOCK TABLES `course_instructor` WRITE;
/*!40000 ALTER TABLE `course_instructor` DISABLE KEYS */;
INSERT INTO `course_instructor` VALUES (1,1,'2019-06-19 12:48:40'),(2,1,'2019-06-19 12:49:03');
/*!40000 ALTER TABLE `course_instructor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `course_session`
--

DROP TABLE IF EXISTS `course_session`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `course_session` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `course_id` int(11) DEFAULT NULL,
  `room_name` varchar(255) DEFAULT NULL,
  `room_location` varchar(255) DEFAULT NULL,
  `quota` int(11) DEFAULT NULL,
  `available_quota` int(11) DEFAULT NULL,
  `start_time` datetime DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  `reg_date` datetime DEFAULT NULL,
  `upd_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `course_id` (`course_id`),
  CONSTRAINT `course_session_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `course_session`
--

LOCK TABLES `course_session` WRITE;
/*!40000 ALTER TABLE `course_session` DISABLE KEYS */;
INSERT INTO `course_session` VALUES (1,1,'ROOM FOR THE COURSE NAME','INDONESIA',-1,-1,'2019-06-18 07:23:18','2099-12-31 00:00:00',1,'2019-06-18 07:23:18',NULL),(2,2,'ROOM FOR THE RPIVATE COURSE NAME','INDONESIA',-1,-1,'2019-06-18 07:23:34','2099-12-31 00:00:00',1,'2019-06-18 07:23:34',NULL);
/*!40000 ALTER TABLE `course_session` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `course_session_enrollment`
--

DROP TABLE IF EXISTS `course_session_enrollment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `course_session_enrollment` (
  `course_session_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `reg_date` datetime DEFAULT NULL,
  KEY `course_session_id` (`course_session_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `course_session_enrollment_ibfk_1` FOREIGN KEY (`course_session_id`) REFERENCES `course_session` (`id`),
  CONSTRAINT `course_session_enrollment_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `course_session_enrollment`
--

LOCK TABLES `course_session_enrollment` WRITE;
/*!40000 ALTER TABLE `course_session_enrollment` DISABLE KEYS */;
INSERT INTO `course_session_enrollment` VALUES (1,1,'2019-06-18 07:25:02'),(2,1,'2019-06-18 07:25:07');
/*!40000 ALTER TABLE `course_session_enrollment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sport`
--

DROP TABLE IF EXISTS `sport`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `sport` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sport`
--

LOCK TABLES `sport` WRITE;
/*!40000 ALTER TABLE `sport` DISABLE KEYS */;
INSERT INTO `sport` VALUES (1,'Football'),(2,'Volleyball'),(3,'Basketball'),(4,'Badminton'),(5,'ESports');
/*!40000 ALTER TABLE `sport` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sport_team`
--

DROP TABLE IF EXISTS `sport_team`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `sport_team` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `team_name` varchar(255) DEFAULT NULL,
  `person_in_charge_id` int(11) DEFAULT NULL,
  `sport_id` int(11) DEFAULT NULL,
  `reg_date` datetime DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `person_in_charge_id` (`person_in_charge_id`),
  KEY `sport_id` (`sport_id`),
  CONSTRAINT `sport_team_ibfk_1` FOREIGN KEY (`person_in_charge_id`) REFERENCES `user` (`id`),
  CONSTRAINT `sport_team_ibfk_2` FOREIGN KEY (`sport_id`) REFERENCES `sport` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sport_team`
--

LOCK TABLES `sport_team` WRITE;
/*!40000 ALTER TABLE `sport_team` DISABLE KEYS */;
INSERT INTO `sport_team` VALUES (1,'Team Merah',1,1,'2019-06-28 09:28:29',1),(2,'Team Biru',1,1,'2019-06-28 09:28:41',1);
/*!40000 ALTER TABLE `sport_team` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sport_team_match`
--

DROP TABLE IF EXISTS `sport_team_match`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `sport_team_match` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sport_team_home_id` int(11) DEFAULT NULL,
  `sport_team_away_id` int(11) DEFAULT NULL,
  `start_date` datetime DEFAULT NULL,
  `end_date` datetime DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  `reg_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sport_team_home_id` (`sport_team_home_id`),
  KEY `sport_team_away_id` (`sport_team_away_id`),
  CONSTRAINT `sport_team_match_ibfk_1` FOREIGN KEY (`sport_team_home_id`) REFERENCES `sport_team` (`id`),
  CONSTRAINT `sport_team_match_ibfk_2` FOREIGN KEY (`sport_team_away_id`) REFERENCES `sport_team` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sport_team_match`
--

LOCK TABLES `sport_team_match` WRITE;
/*!40000 ALTER TABLE `sport_team_match` DISABLE KEYS */;
INSERT INTO `sport_team_match` VALUES (1,1,2,'2019-07-01 07:00:00','2019-07-01 09:00:00',1,'2019-06-28 09:42:46'),(2,2,1,'2019-07-05 07:00:00','2019-07-05 09:00:00',1,'2019-06-28 09:43:08');
/*!40000 ALTER TABLE `sport_team_match` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sport_team_member`
--

DROP TABLE IF EXISTS `sport_team_member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `sport_team_member` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `position_name` varchar(255) DEFAULT NULL,
  `team_member_id` int(11) DEFAULT NULL,
  `sport_team_id` int(11) DEFAULT NULL,
  `reg_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sport_team_id` (`sport_team_id`),
  KEY `team_member_id` (`team_member_id`),
  CONSTRAINT `sport_team_member_ibfk_1` FOREIGN KEY (`sport_team_id`) REFERENCES `sport_team` (`id`),
  CONSTRAINT `sport_team_member_ibfk_2` FOREIGN KEY (`team_member_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sport_team_member`
--

LOCK TABLES `sport_team_member` WRITE;
/*!40000 ALTER TABLE `sport_team_member` DISABLE KEYS */;
INSERT INTO `sport_team_member` VALUES (1,'Captain',1,1,'2019-06-28 09:52:59'),(2,'Captain',1,2,'2019-06-28 09:53:09');
/*!40000 ALTER TABLE `sport_team_member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sport_team_member_score`
--

DROP TABLE IF EXISTS `sport_team_member_score`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `sport_team_member_score` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sport_team_match_id` int(11) DEFAULT NULL,
  `sport_team_member_id` int(11) DEFAULT NULL,
  `score` int(11) DEFAULT NULL,
  `reg_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sport_team_member_id` (`sport_team_member_id`),
  KEY `sport_team_match_id` (`sport_team_match_id`),
  CONSTRAINT `sport_team_member_score_ibfk_1` FOREIGN KEY (`sport_team_member_id`) REFERENCES `sport_team_member` (`id`),
  CONSTRAINT `sport_team_member_score_ibfk_2` FOREIGN KEY (`sport_team_match_id`) REFERENCES `sport_team_match` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sport_team_member_score`
--

LOCK TABLES `sport_team_member_score` WRITE;
/*!40000 ALTER TABLE `sport_team_member_score` DISABLE KEYS */;
INSERT INTO `sport_team_member_score` VALUES (1,1,1,1,'2019-06-29 02:52:45'),(2,1,1,1,'2019-06-29 02:52:46'),(3,1,1,1,'2019-06-29 02:52:47');
/*!40000 ALTER TABLE `sport_team_member_score` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `password` varchar(255) DEFAULT NULL,
  `fullname` varchar(255) DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `reg_date` datetime DEFAULT NULL,
  `upd_date` datetime DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_email_IDX` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'971842e5e72d51ffabdc35640f76d3466d0c3620','OKTA PURNAMA RAHADIAN','oprahadian@gmail.com','2019-06-17 03:33:59','2019-06-20 06:39:35',1),(9,'5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8','','new.oprahadian@gmail.com','2019-06-26 08:55:17',NULL,1);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_coupon`
--

DROP TABLE IF EXISTS `user_coupon`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `user_coupon` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `coupon_name` varchar(50) DEFAULT NULL,
  `coupon_code` varchar(100) DEFAULT NULL,
  `is_redeemed` tinyint(1) DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  `reg_date` datetime DEFAULT NULL,
  `upd_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `user_coupon_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_coupon`
--

LOCK TABLES `user_coupon` WRITE;
/*!40000 ALTER TABLE `user_coupon` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_coupon` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_detail`
--

DROP TABLE IF EXISTS `user_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `user_detail` (
  `user_id` int(11) DEFAULT NULL,
  `class_name` varchar(255) DEFAULT NULL,
  KEY `user_id` (`user_id`),
  CONSTRAINT `user_detail_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_detail`
--

LOCK TABLES `user_detail` WRITE;
/*!40000 ALTER TABLE `user_detail` DISABLE KEYS */;
INSERT INTO `user_detail` VALUES (1,'NON CLASS');
/*!40000 ALTER TABLE `user_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_permission`
--

DROP TABLE IF EXISTS `user_permission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `user_permission` (
  `user_id` int(11) DEFAULT NULL,
  `permission` varchar(255) DEFAULT NULL,
  KEY `user_id` (`user_id`),
  CONSTRAINT `user_permission_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_permission`
--

LOCK TABLES `user_permission` WRITE;
/*!40000 ALTER TABLE `user_permission` DISABLE KEYS */;
INSERT INTO `user_permission` VALUES (1,'ADMIN');
/*!40000 ALTER TABLE `user_permission` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-06-29  9:28:35
