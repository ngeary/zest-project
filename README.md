# zest-project

## Running the code

Execute `./run.sh` to run the code. All json files in the `data` directory will be processed automatically when the program is running.

## Packages

### anonymizer

Provides random values that can be used to anonymize data.

### db

Provides support for adding data to a MySQL database.

### processor

Contains all of the main code to read, parse, process, and write data.

## Data File Directories

### data

Contains the main data files to be ingested. Add json files here to process them.

### replacement_data

Contains replacement data for names and addresses for the purpose of anonymization.

### anon_data

Contains data that has been processed and anonymized.

## Database Schema

```
CREATE TABLE applications (
  row_id varchar(55) NOT NULL,
  id varchar(55) DEFAULT NULL,
  member_id varchar(55) DEFAULT NULL,
  first_name varchar(55) DEFAULT NULL,
  last_name varchar(55) DEFAULT NULL,
  dob varchar(55) DEFAULT NULL,
  app_data json DEFAULT NULL,
  employment_data json DEFAULT NULL,
  created_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (row_id),
  KEY iID (id),
  KEY iMemberID (member_id),
  KEY iFirstName (first_name),
  KEY iLastName (last_name),
  KEY iDOB (dob),
  KEY iCreatedTime (created_time),
  KEY iUpdatedTime (updated_time)
)
```
