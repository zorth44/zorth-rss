可以订阅动漫花园的RSS信息到你的数据库

```sql
CREATE TABLE `rss_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `link` text NOT NULL,
  `description` text,
  `author` varchar(255) DEFAULT NULL,
  `pub_date` datetime NOT NULL,
  `enclosure_url` text,
  `enclosure_length` varchar(255) DEFAULT NULL,
  `enclosure_type` varchar(255) DEFAULT NULL,
  `guid` varchar(255) DEFAULT NULL,
  `category` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `guid` (`guid`)
) ENGINE=InnoDB AUTO_INCREMENT=501 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```

每隔一段时间都会访问动漫花园的RSS信息，并更新到数据库中。

后续的更新计划:
支持获取动漫花园的不在rss中的信息到数据库中