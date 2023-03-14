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

## [2023.1.3.14] - 2023-03-14

### Added

- ProductType struct
- ProductType and Location for Product
- CanChange for StockOrderRecord

## [2023.1.3.9] - 2023-03-09

### Changed

- StockRecord: StockId changed StockInId and StockOutId
- StockRecordItem: Location changed LocationInId and LocationOutId
- StockStateRecord: Removed CompanyId

## [2023.1.3.7] - 2023-03-07

### Added

- Closed to stock record

## [2023.1.3.6] - 2023-03-06

### Added

- Unique indexes for xstock tables

## [2023.1.2.17] - 2023-02-17

### Changed

- DateTimeStart and DateTimeEnd for StockOrderRecord
- CompanyTypeId changed to int from string

## [2023.1.2.7] - 2023-02-07

### Added

- updated to latest go and libraries
- database structure for xStock warehouse management system

## [2023.1.1.16] - 2023-01-16

### Added

- PowerFactor to Workplace table

## [2022.4.2.9] - 2022-11-09

### Added

- Checklist and ChecklistRecord tables
- ImageUrl to checklist table

## [2022.4.1.31] - 2022-10-31

### Changed

- workplaceport color changed to sql.nullstring

## [2022.4.1.27] - 2022-10-27

### Added

- Consumption for OrderRecord and DowntimeRecord

## [2022.4.1.26] - 2022-10-26

### Added

- ProductionConsumption and DowntimeConsumption to SummaryRecords and ShiftSummaryRecords

## [2022.3.3.12] - 2022-09-12

### Added

- ImageUrl to Layout table

## [2022.3.3.7] - 2022-09-07

### Added

- Layout table

## [2022.3.2.1] - 2022-08-01

### Added

- Unit to workplace

## [2022.3.1.20] - 2022-07-20

### Added

- Two new indexes to analog, digital, special and serial data

## [2022.3.1.14] - 2022-07-14

### Added

- ShiftSummaryRecord table

## [2022.2.3.21] - 2022-06-21

### Removed

- OperationOrderRecord

### Added

- OrderId for Operation

### Changed

- order and operation has unique barcode, not name

## [2022.2.3.20] - 2022-06-20

### Added

- OperationOrderRecord
- Note column for every possible table

## [2022.2.3.16] - 2022-06-16

### Changed

- Index for operation

### Added

- Special port table

## [2022.2.3.2] - 2022-06-02

### Removed

- ProductID for Part table

## [2022.2.2.20] - 2022-05-20

### Removed

- bookmarks table

## [2022.2.1.28] - 2022-04-28

### Removed

- creating database from postgres

## [2022.1.3.17] - 2022-03-17

### Added

- Voltage to Workplace table

## [2022.1.3.16] - 2022-03-16

### Removed

- WorkplaceSectionId in Workplace table

### Changed

- Updated to go 1.18
- Updated readme.md

## [2022.1.3.2] - 2022-03-02

### Added

- WebUserSettings table

## [2022.1.2.28] - 2022-02-28

### Changed

- Postgres updated to 14.2

## [2022.1.1.24] - 2022-01-24

### Added

- WorkplaceSectionRecord structure

## [2022.1.1.10] - 2022-01-10

### Added

- Record structure

## [2022.1.1.6] - 2022-01-06

### Changed

- CacheRecord changed to SummaryRecord

## [2022.1.1.3] - 2022-01-03

### Changed

- Alarm.WorkplaceID changed to sql.NullInt32
- Package.OrderID changed to Package.ProductID

## [2021.4.3.6] - 2021-12-06

### Changed

- updated to latest go 1.17.4
- updated to latest go libraries

## [2021.4.2.1] - 2021-11-01

### Added

- cacheRecord table
- added unique index to cacheRecords
- added tags

## [2021.4.1.28] - 2021-10-28

### Changed

- workshift start and end changed to int
- added tags

## [2021.4.1.4] - 2021-10-04

### Added

- updated to latest modules
- updated to latest postgres docker image
- removed table for workplace cached data

## [2021.3.1.20] - 2021-07-20

### Changed

- alarmRecord.workplaceId changed to sql.NullInt32

## [2021.2.3.28] - 2021-06-28

### Changed

- downtimeRecord.OrderId changed to downtimeRecord.OrderRecordId

## [2021.2.3.23] - 2021-06-23

### Changed

- added again row for AlarmRecord and again for DowntimeRecord and new for OrderRecord as sqlnull.Int32

## [2021.2.3.15] - 2021-06-15

### Changed

- updated to latest go 1.16.5
- updated to latest go libraries
- updated to latest postgres

### Added

- WebUserRecord table

## [2021.2.2.13] - 2021-05-13

### Changed

- updated to latest go 1.16.4
- updated to latest go libraries
- updated to latest postgres

## [2021.2.2.3] - 2021-05-03

### Changed

- updated to latest go
- updated to latest go libraries
- updated to latest postgres

## [2021.1.1.7] - 2021-04-07

### Removed

- date_time_end removed again
- three columns for state_service from workplaces table

### Added

- added PageCount table
- better index for state_records table

## [2021.1.1.6] - 2021-04-06

### Added

- date_time_end added again to StateRecord

## [2021.1.3.30] - 2021-03-30

### Added

- date_time_end to StateRecord

## [2021.1.3.23] - 2021-03-23

### Removed

- workplaceworkshift index

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