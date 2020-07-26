#!/usr/bin/env bash
docker pull postgres:alpine
docker rmi -f petrjahoda/database:$1
docker build -t petrjahoda/database:$1 .
docker push petrjahoda/database:$1