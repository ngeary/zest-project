-- Table schema
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
  PRIMARY KEY (`row_id`),
  KEY `iID` (`id`),
  KEY `iMemberID` (`member_id`),
  KEY `iFirstName` (`first_name`),
  KEY `iLastName` (`last_name`),
  KEY `iDOB` (`dob`),
  KEY `iCreatedTime` (`created_time`),
  KEY `iUpdatedTime` (`updated_time`)
)

-- Sample queries
SELECT row_id, id, member_id, first_name, last_name, dob, created_time FROM applications;

SELECT row_id, id, member_id, first_name, last_name, dob, created_time FROM applications WHERE member_id = 'A635';

SELECT row_id, id, member_id, first_name, last_name, dob, created_time FROM applications WHERE first_name = 'michael' AND last_name = 'scott' AND dob = '1965-08-08';

SELECT row_id, employment_data FROM applications WHERE created_time BETWEEN '2020-11-21' AND '2020-11-23';

SELECT row_id, id, member_id, first_name, last_name, dob, created_time FROM applications WHERE created_time > DATE(NOW()) - INTERVAL 1 WEEK;

SELECT row_id, id, member_id, first_name, last_name, dob, created_time FROM applications WHERE created_time > NOW() - INTERVAL 2 HOUR;

SELECT row_id, app_data FROM applications WHERE last_name = 'byrd';