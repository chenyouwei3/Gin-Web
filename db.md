数据库表
//operation_logs

CREATE TABLE `operation_logs` (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
`username` longtext,
`ip` longtext,
`method` longtext,
`query` longtext,
`path` longtext,
`status` bigint(20) DEFAULT NULL,
`start_time` datetime(3) DEFAULT NULL,
`time_cost` bigint(20) DEFAULT NULL,
`user_agent` longtext,
`errors` longtext,
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1766836196429807617 DEFAULT CHARSET=utf8mb4;

// apis

CREATE TABLE `apis` (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
`name` varchar(20) NOT NULL,
`url` varchar(20) NOT NULL,
`method` varchar(10) NOT NULL,
`desc` varchar(144) DEFAULT NULL,
PRIMARY KEY (`id`),
KEY `idx_apis_id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1759805366079598593 DEFAULT CHARSET=utf8mb4;

// roles

CREATE TABLE `roles` (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
`name` varchar(20) NOT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1759814479131975681 DEFAULT CHARSET=utf8mb4;

//users

CREATE TABLE `users` (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
`name` varchar(20) NOT NULL,
`account` varchar(20) NOT NULL,
`password` varchar(60) NOT NULL,
`role_id` bigint(20) NOT NULL,
`salt` varchar(20) NOT NULL,
`sex` varchar(5) NOT NULL,
PRIMARY KEY (`id`),
KEY `fk_roles_user` (`role_id`),
CONSTRAINT `fk_roles_user` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1762420553685151745 DEFAULT CHARSET=utf8mb4;

//role_apis

CREATE TABLE `role_apis` (
`api_id` bigint(20) NOT NULL,
`role_id` bigint(20) NOT NULL,
PRIMARY KEY (`api_id`,`role_id`),
KEY `fk_role_apis_role` (`role_id`),
CONSTRAINT `fk_role_apis_api` FOREIGN KEY (`api_id`) REFERENCES `apis` (`id`),
CONSTRAINT `fk_role_apis_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
