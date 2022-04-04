#!/bin/bash

if [ "$(docker ps -q -f name=inventory-backend-mongo-db)" ]; then
  docker rm -f inventory-backend-mongo-db
fi