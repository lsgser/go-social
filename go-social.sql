DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `surname` varchar(255) NOT NULL,
    `username` varchar(255) NOT NULL UNIQUE,
    `email` varchar(255) NOT NULL UNIQUE,
    `email_verified_at` timestamp DEFAULT NULL,
    `password` varchar(255) NOT NULL,
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
	UNIQUE KEY `users_username_unique` (`username`),
  	UNIQUE KEY `users_email_unique` (`email`)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `schools`;

CREATE TABLE `schools` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`school_name` varchar(255) NOT NULL UNIQUE,
	`school_icon` varchar(255) NOT NULL UNIQUE,
	`description` varchar(255) DEFAULT NULL,
	`created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE KEY `schools_school_name_unique` (`school_name`),
	UNIQUE KEY `schools_school_icon_unique` (`school_icon`)
) ENGINE=InnoDB;

/*
	students is a table that will store 
	users to a specific school
*/
DROP TABLE IF EXISTS `attends`;

CREATE TABLE `attends` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`school` bigint(20) unsigned NOT NULL,
	`student` bigint(20) unsigned NOT NULL,
	`created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	KEY `attends_school_id_index` (`school`),
	KEY `attends_student_id_index` (`student`),
	CONSTRAINT `attends_school_foreign` FOREIGN KEY (`school`) REFERENCES `schools` (`id`) ON DELETE CASCADE,
	CONSTRAINT `attends_student_foreign` FOREIGN KEY (`student`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `profiles`;

CREATE TABLE `profiles` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`user_id` bigint(20) unsigned NOT NULL,
	`school` varchar(255) DEFAULT NULL, 
	`profile_picture` varchar(255) DEFAULT NULL,
	`gender` tinyint(3) unsigned DEFAULT NULL,
	`birth_date` date DEFAULT NULL,
	`residence` varchar(255) DEFAULT NULL,
	`bio` varchar(255) DEFAULT NULL,
	`created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE KEY `profiles_user_id_unique` (`user_id`),
	KEY `profiles_user_id_index` (`user_id`),
	CONSTRAINT `profiles_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB;
