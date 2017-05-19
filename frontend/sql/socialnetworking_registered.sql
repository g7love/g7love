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
-- Table structure for table `registered`
--

DROP TABLE IF EXISTS `registered`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `registered` (
  `id` int(8) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `userid` int(10) NOT NULL COMMENT '用户id',
  `nickname` char(36) NOT NULL DEFAULT '' COMMENT '昵称',
  `email` char(36) NOT NULL COMMENT '邮箱',
  `password` char(36) NOT NULL DEFAULT '0' COMMENT '密码',
  `provinces` int(2) NOT NULL COMMENT '省份',
  `school` int(4) NOT NULL COMMENT '学校',
  `birthday` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生日',
  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别: 0,男 1,女',
  `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updatetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时间',
  `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除标记: 0,未删除 1,已删除',
  `motto` char(150) DEFAULT '' COMMENT '个性签名',
  `headPortrait` char(150) DEFAULT '/uploads/dF5n16wL_bigger.jpg' COMMENT '头像',
  `backgroundImage` char(150) DEFAULT '/uploads/600x200.jpg' COMMENT '背景条',
  `thumb` int(5) NOT NULL DEFAULT '0' COMMENT '主页点赞',
  PRIMARY KEY (`id`),
  UNIQUE KEY `userid` (`userid`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8 COMMENT='注册表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `registered`
--

LOCK TABLES `registered` WRITE;
/*!40000 ALTER TABLE `registered` DISABLE KEYS */;
INSERT INTO `registered` VALUES (20,100,'lizhengxiang','1065482100@qq.com','0',13,15,'1992-05-11 16:00:00',0,'2016-09-20 06:53:35','2016-09-20 06:53:35',0,'Online entrepreneur, fitness enthusiast, fashion aficionado, and tech lover. - That\'s me.','/uploads/cMhf51BU_400x400.jpg','/uploads/1500x500.jpg',9),(25,101,'lizhengxiang','2581913653@qq.com','0',13,15,'1992-05-11 16:00:00',0,'2016-09-20 07:58:53','2016-09-20 07:58:53',0,'','/uploads/dF5n16wL_bigger.jpg','/uploads/600x200.jpg',1),(26,1609000395,'111','1111','0',13,14,'1979-12-31 16:00:00',0,'2016-09-20 09:54:55','2016-09-20 09:54:55',0,'','/uploads/dF5n16wL_bigger.jpg','/uploads/600x200.jpg',0),(27,1609000493,'222','106543@qq.com','0',13,14,'1979-12-31 16:00:00',0,'2016-09-20 09:56:33','2016-09-20 09:56:33',0,'','/uploads/dF5n16wL_bigger.jpg','/uploads/600x200.jpg',0),(28,1609000552,'烦烦烦','33@qq.omm','0',10,11,'1979-12-31 16:00:00',0,'2016-09-20 10:15:52','2016-09-20 10:15:52',0,'','/uploads/dF5n16wL_bigger.jpg','/uploads/600x200.jpg',0),(29,1609000619,'111','222@qq.com','0',10,11,'1979-12-31 16:00:00',0,'2016-09-20 10:16:59','2016-09-20 10:16:59',0,'','/uploads/dF5n16wL_bigger.jpg','/uploads/600x200.jpg',0),(30,1609000715,'ssss','106@qq.com','0',10,11,'1979-12-31 16:00:00',0,'2016-09-20 10:20:15','2016-09-20 10:20:15',0,'','/uploads/dF5n16wL_bigger.jpg','/uploads/600x200.jpg',0),(31,1609000877,'www','www@qq.cm','0',13,14,'1979-12-31 16:00:00',0,'2016-09-20 12:19:37','2016-09-20 12:19:37',0,'','/uploads/dF5n16wL_bigger.jpg','/uploads/600x200.jpg',0),(32,1609000994,'www','www@qqq.cm','0',13,14,'1979-12-31 16:00:00',0,'2016-09-20 12:19:54','2016-09-20 12:19:54',0,'','/uploads/dF5n16wL_bigger.jpg','/uploads/600x200.jpg',0),(33,1609001040,'速冻发的地方好发的d','xx@qq.com','0',10,11,'1979-12-31 16:00:00',0,'2016-09-20 13:15:40','2016-09-20 13:15:40',0,'','/uploads/dF5n16wL_bigger.jpg','/uploads/600x200.jpg',0),(34,1609001169,'sss','ssss@qq.com','0',10,11,'1979-12-31 16:00:00',0,'2016-09-20 13:27:49','2016-09-20 13:27:49',0,'','/uploads/dF5n16wL_bigger.jpg','/uploads/600x200.jpg',0),(35,1609001277,'sss','sssss@qq.com','0',10,11,'1979-12-31 16:00:00',0,'2016-09-20 13:27:57','2016-09-20 13:27:57',0,'','/uploads/dF5n16wL_bigger.jpg','/uploads/600x200.jpg',0),(36,1609001396,'sss','sfdffssss@qq.com','0',10,11,'1979-12-31 16:00:00',0,'2016-09-20 13:28:16','2016-09-20 13:28:16',0,'','/uploads/dF5n16wL_bigger.jpg','/uploads/600x200.jpg',0),(37,1609001485,'dsds','ss@qq.com','0',13,14,'1979-12-31 16:00:00',0,'2016-09-21 01:44:45','2016-09-21 01:44:45',0,'','/uploads/dF5n16wL_bigger.jpg','/uploads/600x200.jpg',0),(38,1609001547,'dsds','s33s@qq.com','0',13,14,'1979-12-31 16:00:00',0,'2016-09-21 01:47:27','2016-09-21 01:47:27',0,'','/uploads/dF5n16wL_bigger.jpg','/uploads/600x200.jpg',0);
/*!40000 ALTER TABLE `registered` ENABLE KEYS */;
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
