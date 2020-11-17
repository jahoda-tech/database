#!/usr/bin/env bash
cd ..
go get -u all
cd postgres

docker pull postgres:alpine
docker rmi -f petrjahoda/database:latest
docker build -t petrjahoda/database:latest .
docker push petrjahoda/database:latest

docker rmi -f petrjahoda/database:2020.4.2
docker build -t petrjahoda/database:2020.4.2 .
docker push petrjahoda/database:2020.4.2
