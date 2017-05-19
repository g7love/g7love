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
-- Table structure for table `dynamic`
--

DROP TABLE IF EXISTS `dynamic`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `dynamic` (
  `id` int(8) NOT NULL AUTO_INCREMENT COMMENT '动态id',
  `userid` int(10) NOT NULL COMMENT '用户id',
  `content` varchar(600) NOT NULL DEFAULT '' COMMENT '动态内容',
  `pic1` char(36) NOT NULL DEFAULT '' COMMENT '图片1',
  `pic2` char(36) NOT NULL DEFAULT '' COMMENT '图片2',
  `pic3` char(36) NOT NULL DEFAULT '' COMMENT '图片3',
  `pic4` char(36) NOT NULL DEFAULT '' COMMENT '图片4',
  `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updatetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时间',
  `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除标记: 0,未删除 1,已删除',
  `reportNum` int(2) DEFAULT '0' COMMENT '举报次数',
  `praise` int(2) DEFAULT '0' COMMENT '点赞次数',
  `forwardingNum` int(2) DEFAULT '0' COMMENT '转发次数',
  PRIMARY KEY (`id`),
  KEY `userid` (`userid`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='发表动态表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dynamic`
--

LOCK TABLES `dynamic` WRITE;
/*!40000 ALTER TABLE `dynamic` DISABLE KEYS */;
INSERT INTO `dynamic` VALUES (1,100,'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Corrupti, illum voluptates consectetur consequatur ducimus. Necessitatibus, nobis consequatur hic eaque laborum laudantium. Adipisci, explicabo, asperiores molestias deleniti unde dolore enim quas.','/uploads/CuEKzGpWcAE5B8T.jpg','/uploads/CuKTefIWYAAw9Aj.jpg','/uploads/CuK6KgbWgAEhTaH.jpg','/uploads/CuK6KgbWgAEhTaH.jpg','2016-10-04 07:22:32','2016-10-08 13:10:07',0,10,115,30),(2,100,'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Corrupti, illum voluptates consectetur consequatur ducimus. Necessitatibus, nobis consequatur hic eaque laborum laudantium. Adipisci, explicabo, asperiores molestias deleniti unde dolore enim quas.','/uploads/CthYHkzXEAQw7qE.jpg','/uploads/CuKTefIWYAAw9Aj.jpg','/uploads/CuK6KgbWgAEhTaH.jpg','/uploads/CuK6KgbWgAEhTaH.jpg','2016-10-04 07:25:24','2016-10-08 13:12:05',0,8,97,0);
/*!40000 ALTER TABLE `dynamic` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-10-11 16:52:38
