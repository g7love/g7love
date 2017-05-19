-- MySQL dump 10.13  Distrib 5.7.15, for Linux (i686)
--
-- Host: localhost    Database: socialnetworking
-- ------------------------------------------------------
-- Server version	5.7.15

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `dynamiclog`
--

DROP TABLE IF EXISTS `dynamiclog`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `dynamiclog` (
  `id` int(8) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `dynamicId` int(7) NOT NULL COMMENT '内容id',
  `userid` int(10) NOT NULL COMMENT '用户id',
  `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除标记: 0,未删除 1,已删除',
  `reportNum` int(2) DEFAULT '0' COMMENT '举报次数',
  `praise` int(2) DEFAULT '0' COMMENT '点赞次数',
  `forwardingNum` int(2) DEFAULT '0' COMMENT '转发次数',
  PRIMARY KEY (`dynamicId`,`userid`),
  KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8 COMMENT='点赞，转发，举报日志记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dynamiclog`
--

LOCK TABLES `dynamiclog` WRITE;
/*!40000 ALTER TABLE `dynamiclog` DISABLE KEYS */;
INSERT INTO `dynamiclog` VALUES (43,1,100,0,1,1,0),(44,2,100,0,1,1,0);
/*!40000 ALTER TABLE `dynamiclog` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-10-11 16:52:37
