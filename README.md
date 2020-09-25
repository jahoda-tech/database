[![developed_using](https://img.shields.io/badge/developed%20using-Jetbrains%20Goland-lightgrey)](https://www.jetbrains.com/go/)
<br/>
![GitHub](https://img.shields.io/github/license/petrjahoda/database)
[![GitHub last commit](https://img.shields.io/github/last-commit/petrjahoda/database)](https://github.com/petrjahoda/database/commits/master)
[![GitHub issues](https://img.shields.io/github/issues/petrjahoda/database)](https://github.com/petrjahoda/database/issues)
<br/>
![GitHub language count](https://img.shields.io/github/languages/count/petrjahoda/database)
![GitHub top language](https://img.shields.io/github/languages/top/petrjahoda/database)
![GitHub repo size](https://img.shields.io/github/repo-size/petrjahoda/database)
<br/>
[![Docker Pulls](https://img.shields.io/docker/pulls/petrjahoda/database)](https://hub.docker.com/r/petrjahoda/database)
[![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/petrjahoda/database?sort=date)](https://hub.docker.com/r/petrjahoda/database/tags)
<br/>
[![developed_using](https://img.shields.io/badge/database-PostgreSQL-red)](https://www.postgresql.org) [![developed_using](https://img.shields.io/badge/runtime-Docker-red)](https://www.docker.com)

# Database

## Description
Go module that holds database structure, used in every system service.
PostgreSQL 13 is used as a database service.

## Installation Information
Install under docker runtime using [this dockerfile image](https://github.com/petrjahoda/system/tree/master/latest) with this command: ```docker-compose up -d```

## Implementation Information
Check the software running with this command: ```docker stats```. <br/>
Database has to be running.
Database is running on port 5432, user postgres, password Zps05.....

### After installation fine tuning
- use https://pgtune.leopard.in.ua/#/
    - default: DB version 13
    - default: OS Type Linux
    - default: DB Type Mixed type of applications
    - default: number of connections: 10 per 1 workplace (for example: 100 per 10 workplaces)
- update configuration using ALTER SYSTEM tab
- restart database container        
        
## Developer Information
Use software only as a [part of a system](https://github.com/petrjahoda/system) using Docker runtime.<br/>
Do not run under linux, windows or mac on its own.


© 2020 Petr Jahoda
