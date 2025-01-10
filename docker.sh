#!/bin/bash

POSTGRES_USER=todo_user
POSTGRES_PASSWORD=todo_password
POSTGRES_DB=tododb
POSTGRES_PORT=5435
HOST_DATA_DIR=/absolute/path/to/postgres_data

docker pull postgres:15-alpine

docker run -d \
  --name postgres_tododb \
  -e POSTGRES_USER=$POSTGRES_USER \
  -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
  -e POSTGRES_DB=$POSTGRES_DB \
  -p $POSTGRES_PORT:5432 \
  -v $HOST_DATA_DIR:/var/lib/postgresql/data \
  postgres:15-alpine
