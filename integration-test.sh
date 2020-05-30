#!/bin/bash

AVR_NAME=mock-avr
AVR_PORT=2323
NETWORK_NAME=avr_service_integration
SERVICE_NAME=avr-service
SERVICE_PORT=8080


#
# Setup
#
docker network create ${NETWORK_NAME}

docker run \
    -e "AVR_HOST=${AVR_NAME}" \
    -e "AVR_PORT=${AVR_PORT}" \
    -p "${SERVICE_PORT}:8080" \
    -d \
    --network ${NETWORK_NAME} \
    --network-alias ${SERVICE_NAME} \
    avr-service:latest

docker run \
    -e "PORT=${AVR_PORT}" \
    -p "${AVR_PORT}:${AVR_PORT}" \
    -d \
    --network ${NETWORK_NAME} \
    --network-alias ${AVR_NAME} \
    docker.pkg.github.com/jtoussaint/mock-denon-avr/mock-denon-avr:latest


#
# Run the tests
#
docker run \
    --network ${NETWORK_NAME} \
    -v "$(pwd)/postman-collections:/etc/newman" \
    -t postman/newman:ubuntu \
    run "mute_avr.json" \
    --environment="integration.postman_environment.json"

TEST_EXIT_CODE=${?}


#
# Clean up
#
docker stop $(docker network inspect -f '{{ range $key, $value := .Containers }}{{printf "%s\n" .Name}}{{ end }}' ${NETWORK_NAME})
docker network rm ${NETWORK_NAME}


# Check to see if the tests passed
if [[ ${TEST_EXIT_CODE} -ne 0 ]]
then
    echo "Tests failed"
    exit 1
fi
