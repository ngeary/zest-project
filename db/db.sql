CREATE TABLE `applications` (
  `row_id` varchar(55) NOT NULL,
  `id` varchar(55) DEFAULT NULL,
  `member_id` varchar(55) DEFAULT NULL,
  `first_name` varchar(55) DEFAULT NULL,
  `last_name` varchar(55) DEFAULT NULL,
  `dob` varchar(55) DEFAULT NULL,
  `app_data` json DEFAULT NULL,
  `employment_data` json DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`row_id`)
)