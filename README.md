# AVR Control

A simple API for controlling an Audio Video Receiver (AVR), specifically a Denon S750H.  This project consist of  three
components:

* **[avr-svc](avr-svc)** - A service that provides a REST Api in front of the AVR.
* **[mock-avr](mock-avr)** - A service that mocks the behavior of the AVR for integration tests.
* **[postman-collections](postman-collections)** - Integration tests built and run with postman/newman

The behavior of the Api can be found [here](avr.apib).

## Getting Started

After cloning this repo, run the `build-avr-svc.sh` script to build and push an image of the service locally.

Additionally, the mock service can be built and pushed to your local registry by running the `build-mock-avr.sh` script.

### Running unit tests

Tests for the service can be run by running `go test` in the `avr-svc` directory.  Alternativly tests are run as part of the docker build in `build-avr-svc.sh`.

### Running the service

The service can be run locally in disconnected mode by running `run-local.sh` after building the image.

| Env Var  |      Description      |
|----------|:--------------|
| AVR_HOST | The hostname / ip address of the AVR. |
| AVR_PORT | The port that the AVR listens on (23) |

### Running integration tests

Integration tests can be built by running the `run-integration-tests.sh` script.  This will build a new docker network, deploy the `avr-svc` and `mock-avr` to it, then run the postman collections in the `postman-collections` folder via newman.
