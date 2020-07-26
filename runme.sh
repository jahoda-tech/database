#!/usr/bin/env bash
docker run --name zapsi_database_mariadb --restart always -p 3306:3306 -d petrjahoda/zapsi_database_mariadb:latest
docker run --name zapsi_database_postgres --restart always -p 5432:5432 -d petrjahoda/zapsi_database_postgres:latest
docker run --name zapsi_database_postgres_updated --restart always -p 5433:5432 -d petrjahoda/zapsi_database_postgres:latest
docker run --name zapsi_database_timescale --restart always -p 5434:5432 -d petrjahoda/zapsi_database_timescale:latest
docker run --name zapsi_database_timescale_updated --restart always -p 5435:5432 -d petrjahoda/zapsi_database_timescale:latest