#!/usr/bin/env bash
cd ..
go get -u all
cd postgres

docker pull postgres:13-alpine
docker rmi -f petrjahoda/database:latest
docker build -t petrjahoda/database:latest .
docker push petrjahoda/database:latest

docker rmi -f petrjahoda/database:2020.3.3
docker build -t petrjahoda/database:2020.3.3 .
docker push petrjahoda/database:2020.3.3
