# Database Changelog

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/).

Please note, that this project, while following numbering syntax, it DOES NOT adhere
to [Semantic Versioning](http://semver.org/spec/v2.0.0.html) rules.

## Types of changes

* ```Added``` for new features.
* ```Changed``` for changes in existing functionality.
* ```Deprecated``` for soon-to-be removed features.
* ```Removed``` for now removed features.
* ```Fixed``` for any bug fixes.
* ```Security``` in case of vulnerabilities.

## [2021.1.3.22] - 2021-03-22

### Changed

- workplace port table changed

## [2021.1.3.18] - 2021-03-18

### Changed

- updated to latest postgres
- updated to latest libraries

## [2021.1.2.21] - 2021-02-21

### Changed

- updated to latest postgres 13.2
- reduced table columns
- changed database name
- changed default database password

## [2021.1.3.11] - 2021-03-11

### Changed

- barcode changed to string everywhere

## [2020.4.3.14] - 2020-12-14

### Changed

- updated to latest postgres
- updated to latest libraries

### Removed

- removed nano from dockerfile
- removed timezone from dockerfile

## [2020.4.2.17] - 2020-11-17

### Changed

- updated to latest postgres
- updated to latest libraries
- added locale table
- minor changes

## [2020.4.1.26] - 2020-10-26

### Changed

- updated to latest postgres

## [2020.3.3.25] - 2020-09-25

### Changed

- updated readme.md

## [2020.3.3.24] - 2020-09-24

### Changed

- docker image updated to postgres 13

## [2020.3.2.22] - 2020-08-22

### Added

- automatic go get -u all when creating docker image

## [2020.3.2.5] - 2020-08-05

### Changed

- removed Duration from Alarm

## [2020.3.2.4] - 2020-08-04

### Fixed

- fixed proper indexes for all tables

## [2020.3.1.30] - 2020-07-30

### Removed

- orderId and UserId from DowntimeRecord

## [2020.3.1.26] - 2020-07-26

### Changed

- changed repo name to "database"
- changed create.sh

## [2020.3.1.26] - 2020-07-26

### Removed

- removed ActualDataDateTime and ActualData from DevicePort

## [2020.3.1.22] - 2020-07-22

### Removed

- mariadb image
- sql server image
- timescale image

### Changed

- change to gorm v2
- updated tables structure

## [2020.2.3.27] - 2020-06-27

### Changed

- updated postgres image

### Added

- added timescale image

## [2020.2.4.2] - 2020-04-02

### Added

- add image column to workpplace

## [2020.1.3.5] - 2020-03-05

### Added

- workplace_id to DownTimeRecord

## [2020.1.2.29] - 2020-02-29

### Removed

- MySQL database

## [2020.1.2.28] - 2020-02-28

### Added

- MySQL database, latest version from dockerhub
- Automatic pulling latest official database image when creating zapsi database image
- New table SystemRecords for system information
- New table DeviceWorkplaceRecords for pairing devices and workplaces
- Creating MySQL database
- Creating SQLServer database

### Changed

- database name changed to zapsi3
- PostgreSQL database, upgraded to latest version from dockerhub
- MariaDB database, upgraded to latest version from dockerhub
- SQLServer database, upgraded to latest version from dockerhub
- password for mails is saved encrypted in database

### Removed

- WorkplaceId from devices table

### Fixed

- proper creating mssql database, NO ACTION instead of RESTRICT
- proper creating all databases (instead of using gorm.model own columns used)

## [2020.1.1.1] - 2020-01-01

### Added

- PostgreSQL database
- MariaDB database
- SQLServer database