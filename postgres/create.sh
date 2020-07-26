#!/usr/bin/env bash
docker pull postgres:alpine
docker rmi -f petrjahoda/database:latest
docker build -t petrjahoda/database:latest .
docker push petrjahoda/database:latest

docker rmi -f petrjahoda/database:2020.3.1
docker build -t petrjahoda/database:2020.3.1 .
docker push petrjahoda/database:2020.3.1
