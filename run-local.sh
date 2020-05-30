#!/bin/bash

docker run \
    -e "AVR_HOST=192.168.86.32" \
    -e "AVR_PORT=23" \
    -e "PORT=8080" \
    -p 8080:8080 \
    -d \
    avr-service:latest
